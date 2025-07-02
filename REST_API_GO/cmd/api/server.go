package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type user struct {
	Name string `json: "name"`
	Age  string `json: "age"`
	City string `json: "city"`
}

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET Method on Teachers Route"))
			fmt.Println("Hello GET Method on Teachers Route")
		case http.MethodPost:

			//Parse form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form: ", http.StatusBadRequest)
				return
			}
			fmt.Println("Form: ", r.Form)

			//Prepare response data
			response := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0]
			}
			fmt.Println("Processed Response Map: ", response)

			//RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()
			fmt.Println("RAW Body: ", body)
			fmt.Println("RAW Body: ", string(body))

			//If you expect json data, then unmarshal it
			var userInstance user
			err = json.Unmarshal(body, &userInstance)
			if err != nil {
				return
			}
			fmt.Println(userInstance)

			w.Write([]byte("Hello POST Method on Teachers Route"))
			fmt.Println("Hello POST Method on Teachers Route")
		case http.MethodPut:
			w.Write([]byte("Hello PUT Method on Teachers Route"))
			fmt.Println("Hello PUT Method on Teachers Route")
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH Method on Teachers Route"))
			fmt.Println("Hello PATCH Method on Teachers Route")
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE Method on Teachers Route"))
			fmt.Println("Hello DELETE Method on Teachers Route")
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Execs Route"))
		fmt.Println("Hello Execs Route")
	})

	fmt.Println("Server is running on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server: ", err)
	}
}
