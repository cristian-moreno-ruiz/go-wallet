package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Your favorite wallet!")
	fmt.Println("Endpoint Hit: /")
}

func summary(w http.ResponseWriter, r *http.Request) {
	// TODO: Next step is to fetch Current data from an API
	fmt.Fprintf(w, "This is today's currency values:\n")
	fmt.Fprintf(w, fmt.Sprint(print(entries[len(entries)-1])))
	fmt.Println("Endpoint Hit: /summary")
}

func history(w http.ResponseWriter, r *http.Request) {
	// TODO: Next step is to fetch Current data from an API
	fmt.Fprintf(w, "Last "+fmt.Sprint(len(entries))+" days' currency values:\n")

	for i := 0; i < len(entries); i++ {
		fmt.Fprintf(w, fmt.Sprint(print(entries[len(entries)-i-1])))
	}
	fmt.Println("Endpoint Hit: /summary")
}

func handleRequests() {
	http.HandleFunc("/", root)
	http.HandleFunc("/summary", summary)
	http.HandleFunc("/history", history)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

type entry struct {
	USD float32
	EUR float32
	CRC float32
}

func print(this entry) string {
	return "\nUSD: " + fmt.Sprint(this.USD) + "\n" + "EUR: " + fmt.Sprint(this.EUR) + "\n" + "CRC: " + fmt.Sprint(this.CRC) + "\n"
}

var entries []entry

func main() {
	// This is an initialization with static data
	entries = []entry{
		entry{1.0, 0.82, 613.27},
		entry{1.0, 0.83, 611.59},
	}
	handleRequests()
}
