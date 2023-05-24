package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request Path: %s\n", r.URL)
		/*
			"Fprintf" also specifies destination or where we need the output to display.
			Here, "w" means we are going to be displaying the output in the browser.
		*/
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	/*
		Note: In tutorial it said this "http.ListenAndServe(":80", nil)".
		But, this will keep on letting windows firewall pop up whenever we run `go run .`.
		We could setup new windows firewall inbound rules, but why mess with your computer, when you can just change one line of code.
	*/
	// 💡 Note: On Production, don't forget to change it back to ":3000" or something else.
	http.ListenAndServe("127.0.0.1:8080", nil)
}
