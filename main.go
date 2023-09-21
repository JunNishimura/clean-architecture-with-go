package main

import (
	"log"

	"github.com/JunNishimura/clean-architecture-with-go/driver"
)

func main() {
	if err := driver.Run(); err != nil {
		log.Fatal(err)
	}
}
