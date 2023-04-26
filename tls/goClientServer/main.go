package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	clientMode := os.Args[1]

	if strings.Compare(clientMode, "client") == 0 {
		// producer(props, topic)
		client := setClient()

		// make a GET request to the server using HTTPS
		resp, err := client.Get("https://localhost:443")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// print the response body
		fmt.Println(string(body))

	} else if strings.Compare(clientMode, "server") == 0 {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, World!")
		})

		// if err := http.ListenAndServe(":80", nil); err != nil {
		err := http.ListenAndServeTLS(":443", "/etc/tls/server.pem", "/etc/tls/server-key.pem", nil)
		// err := http.ListenAndServeTLS(":443", "./certifs/server.pem", "./certifs/server-key.pem", nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	} else {
		fmt.Printf("Invalid option")
	}
}

func setClient() *http.Client {
	cert, err := tls.LoadX509KeyPair("./certifs/client.pem", "./certifs/client-key.pem")
	if err != nil {
		panic(err)
	}

	// read the CA certificate file
	caCert, err := ioutil.ReadFile("./certifs/ca.pem")
	if err != nil {
		panic(err)
	}

	// create a new CA certificate pool and add the CA certificate to it
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// create a new HTTP client with a custom TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs:      caCertPool,
			},
		},
	}

	return client
}
