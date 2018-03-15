package main

import (
  "fmt"
  "net/http"
  "strings"
)

func main() {
  http.HandleFunc("/", fnHandler)
  http.HandleFunc("/interactive", interactiveHandler)
  http.ListenAndServe(":7000", nil)
}

func fnHandler(w http.ResponseWriter, r *http.Request) {
  txt := r.FormValue("text")
  options := strings.Split(txt, "|")
  question := options[0]
  options = options[1:]
  
  fmt.Println("Poll: " + question)
  fmt.Println("Options: " + strings.Join(options,","))
  
  poll := generatePoll(question, options)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(poll))
  w.WriteHeader(200)

}

func generatePoll(question string, options []string) string {

	appender := ""
	optionsVal := "["
	for i := 0; i < len(options); i++ {
		optionsVal += appender + `{
		        "name": "`+ options[i] +`",
		        "text": "`+ options[i] +`",
		        "type": "button",
		        "value": "`+ options[i] +`"
		    }`
		appender = ","
	}
	optionsVal += "]"

	return `{
		    "response_type": "in_channel",
		    "text": "`+ question +`",
		    "attachments": [
		        {
		            "fallback": "Oops.. For some reason you aren't able to start a poll.",
		            "callback_id": "meetie_poll",
		            "color": "#3AA3E3",
		            "attachment_type": "default",
		            "actions": `+ optionsVal +`
		        }
		    ]
		}`
}


func interactiveHandler(w http.ResponseWriter, r *http.Request) {

  txt := r.FormValue("payload")
  w.Write([]byte("Thanks for selecting " + txt + "!"))
  w.WriteHeader(200)

}


