package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Today's amazing weather!")

		resp, err := http.Get("https://quotes.rest/qod?language=en")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))

	})

	fmt.Println("Listening on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
