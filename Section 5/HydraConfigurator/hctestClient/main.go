package main

import (
	"Hydra/HydraConfigurator"
	"fmt"
)

type ConfS struct {
	TS      string  `name:"testString"`
	TB      bool    `name:"testBool"`
	TF      float64 `name:"testFloat"`
	TestInt int
}

func main() {
	configstruct := new(ConfS)
	HydraConfigurator.GetConfiguration(HydraConfigurator.CUSTOM, configstruct, "configfile.conf")
	fmt.Println(*configstruct)

	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println(5 * configstruct.TestInt)

	fmt.Println(configstruct.TS)
}
