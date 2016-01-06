package main

import (
	"flag"
	"fmt"
	"github.com/mathieupassenaud/leds/api"
	"github.com/mathieupassenaud/leds/backend"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Launch services")
	renderer.Init()

	r := api.CreateRouter()
	httpTimeout := flag.Int("http-timeout", 10, "timeout on http connections (in seconds)")
	httpAddr := flag.String("http-addr", ":8080", "http address to listen to (use 'none' to disable)")

	if *httpAddr != "none" {
		// Serve in a goroutine since ListenAndServe is blocking and we may also serve https below
		go func() {
			fmt.Printf("Listening to http requests on '%s'", *httpAddr)
			server := http.Server{
				Addr:        *httpAddr,
				Handler:     r,
				ReadTimeout: time.Duration(*httpTimeout) * time.Second,
			}

			if err := server.ListenAndServe(); err != nil {
				log.Fatal(err)
			}
		}()
	}

	select {}

}
