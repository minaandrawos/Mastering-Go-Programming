package hydraportal

import (
	"Hydra/hydraconfigurator"
	"Hydra/hydradblayer"
	"Hydra/hydraweb/hydrarestapi"
	"bufio"
	"html/template"
	"log"
	"net"
	"net/http"
	"sync"

	"Hydra/hydradblayer/passwordvault"
	"bytes"
	"crypto/md5"

	"golang.org/x/net/websocket"
)

var hydraWebTemplate *template.Template
var historylog struct {
	logs []string
	sync.RWMutex
}

func Run() error {
	var err error

	conf := struct {
		Filespath string   `json:"filespath"`
		Templates []string `json:"templates"`
	}{}
	err = hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, &conf, "./hydraweb/portalconfig.json")
	if err != nil {
		return err
	}

	hydraWebTemplate, err = template.ParseFiles(conf.Templates...)
	if err != nil {
		return err
	}

	hydrarestapi.InitializeAPIHandlers()
	log.Println(conf.Filespath)
	fs := http.FileServer(http.Dir(conf.Filespath))
	http.Handle("/", fs)
	http.HandleFunc("/Crew/", crewhandler)
	http.HandleFunc("/about/", abouthandler)
	http.HandleFunc("/Chat/", chathandler)
	http.Handle("/ChatRoom/", websocket.Handler(chatWS))
	go func() {
		err = http.ListenAndServeTLS(":8062", "cert.pem", "key.pem", nil)
		log.Println(err)
	}()
	err = http.ListenAndServe(":8061", nil)
	return err
}

func chathandler(w http.ResponseWriter, r *http.Request) {
	nameStruct := struct{ Name string }{}
	r.ParseForm()
	if len(r.Form) == 0 {
		if cookie, err := r.Cookie("username"); err != nil {
			hydraWebTemplate.ExecuteTemplate(w, "login.html", nil)
			return
		} else {
			nameStruct.Name = cookie.Value
			hydraWebTemplate.ExecuteTemplate(w, "chat.html", nameStruct)
			return
		}
	}

	if r.Method == "POST" {

		user := r.Form["username"][0]
		pass := r.Form["password"][0]
		if !verifyPassword(user, pass) {
			hydraWebTemplate.ExecuteTemplate(w, "login.html", nil)
			return
		}
		nameStruct.Name = user
		if _, ok := r.Form["rememberme"]; ok {
			cookie := http.Cookie{Name: "username", Value: user}
			http.SetCookie(w, &cookie)
		}
	}
	hydraWebTemplate.ExecuteTemplate(w, "chat.html", nameStruct)
}

func verifyPassword(username, pass string) bool {
	db, err := passwordvault.ConnectPasswordVault()
	if err != nil {
		return false
	}
	defer db.Close()
	data, err := passwordvault.GetPasswordBytes(db, username)
	if err != nil {
		return false
	}
	hashedPass := md5.Sum([]byte(pass))
	return bytes.Equal(hashedPass[:], data)
}

func chatWS(ws *websocket.Conn) {
	conn, err := net.Dial("tcp", "127.0.0.1:2100")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	historylog.RLock()
	for _, log := range historylog.logs {
		err := websocket.Message.Send(ws, log)
		if err != nil {
			return
		}
	}
	historylog.RUnlock()

	if len(historylog.logs) > 20 {
		historylog.Lock()
		historylog.logs = historylog.logs[1:]
		historylog.Unlock()
	}

	go func() {
		//Use scanner to receive chat messages
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			err := websocket.Message.Send(ws, message)
			if err != nil {
				return
			}
		}
	}()

	for {
		// receive text frame
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			return
		}
		_, err = conn.Write([]byte(message))
		if err == nil {
			historylog.Lock()
			historylog.logs = append(historylog.logs, message)
			historylog.Unlock()
		}
	}
}

func crewhandler(w http.ResponseWriter, r *http.Request) {
	dblayer, err := hydradblayer.ConnectDatabase("mysql", "gouser:gouser@/Hydra")
	if err != nil {
		return
	}
	all, err := dblayer.AllMembers()
	if err != nil {
		return
	}
	err = hydraWebTemplate.ExecuteTemplate(w, "crew.html", all)
	if err != nil {
		log.Println(err)
	}
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	about := struct {
		Msg string `json:"message"`
	}{}
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, &about, "./hydraweb/about.json")
	if err != nil {
		return
	}
	err = hydraWebTemplate.ExecuteTemplate(w, "about.html", about)
	if err != nil {
		log.Println(err)
	}
}
