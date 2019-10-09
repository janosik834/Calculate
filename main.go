package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"

)
type Numbers struct {
  First string `json:"First Number"`
  Second string `json:"Second Number"`}

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,  "<h1>Calculate </h1>"+
        "<form action=\"/calculate\" >"+
        "First Number: <input type=\"text\" name=\"FirstNumber\"></input><br>"+
        "Second Number: <input type=\"text\" name=\"SecondNumber\"></input><br>"+
        "<input type=\"submit\" value=\"Factorial\">"+
        "</form>")

}

func allNumbers(w http.ResponseWriter, r *http.Request) {
  var twoNumbers = Numbers{First: "3", Second: "43"}
  fmt.Fprintf(w, "Endpoint Hit: All Articles Endpoint")
  json.NewEncoder(w).Encode(twoNumbers)
}
func handleRequests(){
  http.HandleFunc("/", homePage)
  http.HandleFunc("/calculate", allNumbers)
  log.Fatal(http.ListenAndServe(":8487", nil))
  }

  func main() {
handleRequests()

  }
