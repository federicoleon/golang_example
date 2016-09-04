package main

import (
	"encoding/json"
	"fmt"
	"github.com/federicoleon/golang_example/sort" // External dependency
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/sort", sortSlice)
	fmt.Println("Webserver ready for traffic :-)")
	http.ListenAndServe(":8080", nil)
}

type SortExample struct {
	Original []int `json: "original"`
	Result   []int `json: "result"`
}

func sortSlice(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var e SortExample
	err := decoder.Decode(&e)
	if err != nil {
		io.WriteString(w, "unable to parse JSON")
		return
	}

	e.Result, err = sort.BubbleSort(e.Original)

	result, err := json.Marshal(e)
	if err != nil {
		io.WriteString(w, "unable to return a response")
		return
	}

	io.WriteString(w, string(result))
}
