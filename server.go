package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	fmt.Println("Listening on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://type.fit/api/quotes")
	if err != nil {
		// Equivalenet to Println with a call to os.Exit that follows
		log.Fatalln(err)
	}
	// The defer keyword will make the resp.Body.Close() not run until the surrounding function returns
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	type Quote struct {
		Text   string `json:"text"`
		Author string `json:"author"`
	}

	var quotes []Quote
	json.Unmarshal([]byte(string(body)), &quotes)

	if err != nil {
		fmt.Println("error:", err)
	}
	tpl, err := template.ParseFiles("static/tmpl/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(w, quotes[rand.Intn(1641)])
	if err != nil {
		panic(err)
	}

	// fmt.Fprintf(w, quotes[rand.Intn(1641)].Text)
}
