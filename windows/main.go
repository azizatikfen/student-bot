package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Meeting s
type Meeting struct {
	ID       string
	Password string
}

func main() {
	http.HandleFunc("/meeting", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer r.Body.Close()
		var me Meeting
		err = json.Unmarshal(body, &me)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(me)
		fmt.Println(string(body))
	})
	http.ListenAndServe(":8080", nil)
}
