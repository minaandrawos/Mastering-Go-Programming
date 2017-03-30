package hydrarestapi

import (
	"Hydra/hydraconfigurator"
	"log"
	"net/http"
)

type DBlayerconfig struct {
	DB string `json:"database"`
	Conn string `json:"connectionstring"`
}

func RunAPI() {
	conf := new(DBlayerconfig)
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "./hydraweb/apiconfig.json")
	if err != nil {
		log.Fatal("Error decoding JSON",err)
	}
	h := newhydraCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		log.Fatal("Error connecting to db ", err)
	}
	http.Handle("/hydracrew/", h)
	http.ListenAndServe(":8061", nil)
}
