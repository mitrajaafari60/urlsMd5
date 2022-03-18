package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"urlsMd5/internal/controller"
	"urlsMd5/internal/service"
	"urlsMd5/pkg/flagArgs"
)

var Parallel int

func main() {
	flag.IntVar(&Parallel, "parallel", 10, "Number of parallel process")
	flag.Parse()

	index := flagArgs.FlagsCount(os.Args, "parallel")
	urls := os.Args[index:]
	inputData := make(chan string, len(urls))

	for _, data := range urls {
		inputData <- data
	}

	for p := 0; p < Parallel; p++ {
		go service.Worker(inputData)
	}

	http.HandleFunc("/md5urls", controller.CalculateUrlMd5)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
