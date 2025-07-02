package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	//Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	//Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//Create a custom Server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	//Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}
	//HTTP 1.1 Server without TLS
	// http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Could not start server", err)
	// }

}

//Lec 155
// First, the specified package and its dependencies are downloaded from the internet. So, it's not just the HTTP2 package that will be downloaded. If it is dependent on any other package, that package will also be downloaded.

//PEM - Privacy Enhanced Mail - Commonly used for stoing and transmitting cryptographic keys, certficates and other data. 64 encoded, making it ASCII text, as it is easier to read and transport in text-based protocols. PEM files have specific footers and headers to identify the type of content they hold. 
//DER - Distinguished Encoding Rules, it is a binary encoding for x509 certificates. 
