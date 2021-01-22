package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cristian-moreno-ruiz/go-wallet/controllers"

	"github.com/go-shadow/moment"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Your favorite wallet!")
	fmt.Println("Endpoint Hit: /")
}

func summary(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is today's currency values:\n")
	fmt.Fprintf(w, fmt.Sprint(rates[moment.New().Format("YYYY-MM-DD")].print()))
	fmt.Println("Endpoint Hit: /summary")
}

func crypto(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://api.exchangerate.host/latest?base=BTC", nil)

	// TODO: Better way of handling errrors??
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(request)

	// TODO: Better way of handling errrors??
	if err != nil {
		fmt.Println(err)
		return
	}

	var result map[string]entry
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Fprintf(w, "This is today's BTC price:\n")
	fmt.Fprintf(w, fmt.Sprint(result["rates"].print()))
	fmt.Println("Endpoint Hit: /summary")
}

func history(w http.ResponseWriter, r *http.Request) {
	days := strings.TrimPrefix(r.URL.Path, "/history/")
	daysInt, _ := strconv.Atoi(days)
	updateRates(daysInt)
	fmt.Fprintf(w, "Last "+days+" days' currency values:\n")
	printHistory(w, daysInt)

	fmt.Println("Endpoint Hit: /history")
}

func printHistory(w http.ResponseWriter, days int) {
	to := moment.New()
	day := moment.New().Subtract("d", days)
	for day.Diff(to, "days") <= 0 {
		fmt.Fprintf(w, "\n"+day.Format("YYYY-MM-DD")+":")
		fmt.Fprintf(w, rates[day.Format("YYYY-MM-DD")].print())
		day.Add("days", 1)
	}
}

func updateRates(days int) {
	to := moment.New()
	from := moment.New().Subtract("d", days)

	// Check if the stored dates contain the requested dates, return if it is the case
	if (ratesTo != nil && to.Diff(ratesTo, "days") <= 0) && (ratesFrom != nil && from.Diff(ratesFrom, "days") >= 0) {
		return
	}

	client := http.Client{}
	request, err := http.NewRequest("GET", "https://api.exchangerate.host/timeseries?start_date="+from.Format("YYYY-MM-DD")+"&end_date="+to.Format("YYYY-MM-DD")+"&base=USD", nil)

	// TODO: Better way of handling errrors??
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(request)

	// TODO: Better way of handling errrors??
	if err != nil {
		fmt.Println(err)
		return
	}

	var result map[string]map[string]entry
	json.NewDecoder(resp.Body).Decode(&result)

	rates = result["rates"]
	ratesTo = to
	ratesFrom = from
	fmt.Println("Values updated from, to", ratesFrom.Format("YYYY-MM-DD"), ratesTo.Format("YYYY-MM-DD"))
}

func handleRequests() {
	http.HandleFunc("/", root)
	http.HandleFunc("/summary", summary)
	http.HandleFunc("/crypto", crypto)
	http.HandleFunc("/history/", history)
	// TODO: Deprecate this route
	http.HandleFunc("/taxes/", controllers.Handle)
	http.HandleFunc("/wallet/upload", controllers.Import)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

type entry struct {
	USD float32
	EUR float32
	CRC float32
}

func (entry entry) print() string {
	return "\nUSD: " + fmt.Sprint(entry.USD) + "\n" + "EUR: " + fmt.Sprint(entry.EUR) + "\n" + "CRC: " + fmt.Sprint(entry.CRC) + "\n"
}

var ratesFrom, ratesTo *moment.Moment
var rates map[string]entry

func main() {
	updateRates(1)
	handleRequests()
}
