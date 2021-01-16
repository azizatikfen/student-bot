package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
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
		joinZoomMeeting(me.ID, me.Password)
		fmt.Println(me)
		fmt.Println(string(body))
	})
	http.ListenAndServe(":8080", nil)
}

func joinZoomMeeting(id, passwd string) {
	cmd := exec.Command(`C:\Users\yusuf\AppData\Roaming\Zoom\bin\Zoom.exe`, `"--url=zoommtg://zoom.us/join?action=join&confno=`+id+`&pwd=`+passwd+`"`)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
