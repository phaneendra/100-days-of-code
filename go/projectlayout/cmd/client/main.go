package main

import (
	"fmt"
	"github.com/phaneendra/go-eggs/go/projectlayout"
	"github.com/phaneendra/go-eggs/go/projectlayout/services/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", projectlayout.Config())
	fmt.Println("Client:", clientlib.Hello())
}
