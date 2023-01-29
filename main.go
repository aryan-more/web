package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

const deploy = false

func main() {
	e := echo.New()
	e.Static("/", "/static")

	if deploy {
		log.Fatal(e.StartAutoTLS("aryanmore.us.to"))
	} else {
		log.Fatal(e.Start(":80"))
	}

}
