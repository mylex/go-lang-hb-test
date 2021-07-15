package server

import (
	"fmt"
	"log"
	"net/http"
)

func Init(port int) {
	portPrefix := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(portPrefix, Router())

	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}
