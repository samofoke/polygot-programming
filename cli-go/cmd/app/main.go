package main

import (
	"fmt"
	"log"

	"github.com/samofoke/polygot/pkg/progfile"
)

func main() {
	opts, err := progfile.GetOps()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	mainConfig, err := progfile.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get the config %v", err)
	}

	fmt.Printf("Operations: %+v", opts)
	fmt.Printf("config: %+v", mainConfig)
}
