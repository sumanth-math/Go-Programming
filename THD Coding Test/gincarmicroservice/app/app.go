package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	// read port number from environmental variable
	portNumber = os.Getenv("PORT_NUMBER")
	router     = gin.Default()
)

func StartApplication() {
	mapRoutes()
	if portNumber == "" {
		portNumber = ":9001"
	} else {
		portNumber = fmt.Sprintf(":%s", portNumber)
	}
	log.Fatal(router.Run(portNumber))
}
