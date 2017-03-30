package hydraportal

import (
	"Hydra/hydraconfigurator"
	"Hydra/hydradblayer"
	"Hydra/hydraweb/hydrarestapi"
	"html/template"
	"log"
	"net/http"
)

var hydraWebTemplate *template.Template

func Run() error {
	var err error
	hydraWebTemplate, err = template.ParseFiles("./hydraweb/hydraportal/cover/Crew/crew.html", "./hydraweb/hydraportal/cover/about/about.html")
	if err != nil {
		return err
	}
	conf := struct {
		Filespath string `json:"filespath"`
	}{}
	err = hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, &conf, "./hydraweb/portalconfig.json")
	if err != nil {
		return err
	}

	hydrarestapi.InitializeAPIHandlers()
	log.Println(conf.Filespath)
	fs := http.FileServer(http.Dir(conf.Filespath))
	http.Handle("/", fs)
	http.HandleFunc("/Crew/", crewhandler)
	http.HandleFunc("/about/", abouthandler)
	err = http.ListenAndServe(":8061", nil)
	return err
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
