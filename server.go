package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	NASA_API_KEY := os.Getenv("NASA_API_KEY")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Today's amazing weather!")

		resp, err := http.Get("https://api.nasa.gov/insight_weather/?api_key=" + NASA_API_KEY + "&feedtype=json&ver=1.0")
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
