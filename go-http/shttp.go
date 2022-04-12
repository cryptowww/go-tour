package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

func main(){
	// http server
	fmt.Println("--starting a server")
	//log.Fatal(http.ListenAndServe(":8080",nil))
	s := &http.Server{
		Addr:			":8080",
		//Handler:		myHandler,
		ReadTimeout:	10 * time.Second,
		WriteTimeout:	10 * time.Second,
		MaxHeaderBytes:	1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
