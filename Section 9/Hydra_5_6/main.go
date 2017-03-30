package main

import (
	"Hydra/hlogger"
	"Hydra/hydrachat"
	//"Hydra/hydraweb/hydrarestapi"
	//"Hydra/hydraweb"
	"Hydra/hydraweb/hydraportal"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	err := hydrachat.Run()
	if err != nil {
		logger.Println("Could not run hydra chat", err)
	}

	//hydrarestapi.RunAPI()
	//hydraweb.Run()
	hydraportal.Run()
}
