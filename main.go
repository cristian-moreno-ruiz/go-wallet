package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-shadow/moment"
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
	days := strings.TrimPrefix(r.URL.Path, "/history/")
	fmt.Println("Requesting days:", days)

	daysInt, _ := strconv.Atoi(days)
	updateRates(daysInt)
	fmt.Fprintf(w, "Last "+days+" days' currency values:\n")

	for i := 0; i < len(entries); i++ {
		fmt.Fprintf(w, fmt.Sprint(print(entries[len(entries)-i-1])))
	}
	fmt.Println("Endpoint Hit: /history")
}

func updateRates(days int) {
	to := moment.New()
	from := moment.New().Subtract("d", days)
	fmt.Println("Values requested from, to", from.Format("YYYY-MM-DD"), to.Format("YYYY-MM-DD"))

	// TODO: Consider if I need to fetch or already have the data
	fmt.Println("compare TO", to.Diff(&ratesTo, "days"))

	// TODO: Do request with first and last days

	// TODO: Pick the values of interest from the response

	// TODO: Store them, including recording which days I have
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://api.exchangerate.host/2020-04-04", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	// var result map[string]interface{}
	var result map[string]entry
	json.NewDecoder(resp.Body).Decode(&result)

	//TODO:
	// lastUpdated = result["date"]
	rates := result["rates"]

	fmt.Println(rates.USD)
}

func handleRequests() {
	http.HandleFunc("/", root)
	http.HandleFunc("/summary", summary)
	http.HandleFunc("/history/", history)
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
var ratesFrom, ratesTo moment.Moment //= moment.New(), moment.New()

func main() {
	// This is an initialization with static data
	entries = []entry{
		entry{1.0, 0.82, 613.27},
		entry{1.0, 0.83, 611.59},
	}
	handleRequests()
}
