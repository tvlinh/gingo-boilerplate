package main

import (
	"log"
	"tvlinh/gingo-boilerplate/cmd/api"
)

func main() {
	err := api.Run()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
