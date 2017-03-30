package main

import (
	"Hydra/hlogger"
	"Hydra/hydrachat"
	"Hydra/hydraweb/hydraportal"
	"flag"
	"strings"
	//"Hydra/hydradblayer/passwordvault"
	//"crypto/md5"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")
	operation := flag.String("o", "w", "Operation: w for web \n c for chat")
	flag.Parse()
	switch strings.ToLower(*operation) {
	case "c":
		err := hydrachat.Run()
		if err != nil {
			logger.Println("Could not run hydra chat", err)
		}
	case "w":
		err := hydraportal.Run()
		if err != nil {
			logger.Println("could not run hydra web portal", err)
		}
	}

	/*
		code to populate some password in the pass vault
		db,err := passwordvault.ConnectPasswordVault()
		if err!= nil{
			return
		}
		minapss := md5.Sum([]byte("minaspass"))
		jimpass := md5.Sum([]byte("jimspass"))
		caropass := md5.Sum([]byte("carospass"))
		passwordvault.AddBytesToVault(db,"Mina",minapss[:])
		passwordvault.AddBytesToVault(db,"Jim",jimpass[:])
		passwordvault.AddBytesToVault(db,"Caro",caropass[:])
		db.Close()
	*/

}
