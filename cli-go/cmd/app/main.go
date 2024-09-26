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

	fmt.Printf("Operations: %+v", opts)
}
