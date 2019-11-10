package main

import (
	"fmt"
	"github.com/phaneendra/go-eggs/go/projectlayout"
	"github.com/phaneendra/go-eggs/go/projectlayout/internal/auth"
	"github.com/phaneendra/go-eggs/go/projectlayout/services/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", projectlayout.Config())
	fmt.Println("Server:", serverlib.Hello())
	fmt.Println("Auth:", auth.Auth())
}
