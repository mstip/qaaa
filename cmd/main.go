package main

import (
	"log"

	"github.com/mstip/qaaa/pkg/store"
	"github.com/mstip/qaaa/pkg/web"
)

func main() {
	log.Println("qaaa-server is running")
	store := store.NewStore()
	ws, err := web.NewWebServer(store)
	if err != nil {
		log.Fatal(err)
	}
	ws.RunAndServe()
}
