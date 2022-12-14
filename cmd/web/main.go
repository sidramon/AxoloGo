package main

import (
	"fmt"
	"net/http"

	"github.com/sidramon/AxoloGo/Config"
	"github.com/sidramon/AxoloGo/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err);
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":3000"

	handlers.CreateTemplates(&appConfig)

	/* 
		<- Routes 
	*/

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)
	
	/*
		->
	*/

	fmt.Println("(http://localhost:3000) - Server started on port", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}