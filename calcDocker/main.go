package main

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
  "bytes"
  "io/ioutil"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "github.com/janosik834/Calculate/calculateP"

)

var port string = ":8989"
var url string = "http://localhost" + port +"/calculate"

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  if r.URL.Path != "/" {
          http.Error(w, "404 not found.", http.StatusNotFound)
          return
      }
      switch r.Method {
      case "GET":
           http.ServeFile(w, r, "index.html")
      case "POST":
          if err := r.ParseForm(); err != nil {
              fmt.Fprintf(w, "ParseForm() err: %v", err)
              return
          }
          a, err1 := strconv.Atoi(r.FormValue("a"))
          b, err2 := strconv.Atoi(r.FormValue("b"))
          if err1 != nil {
              fmt.Fprintf(w, "'First Number' is not a number and is changed to 0 \n")
          }
          if err2 != nil {
              fmt.Fprintf(w, "'Second Number' is not a number and is changed to 0\n ")
          }
          var n calculateP.Numbers = calculateP.Numbers{First: a, Second: b}
          jsonN, _ :=  json.Marshal(n)
          req, err := http.NewRequest("POST",url, bytes.NewBuffer(jsonN))
          if err != nil {
              panic(err)
          }
          req.Header.Set("Content-Type", "application/json")
          client := &http.Client{}
          resp, err := client.Do(req)
          if err != nil {
            panic(err)
          }
          defer resp.Body.Close()
          fmt.Fprintf(w, "response Status: %v \n", resp.Status)
          fmt.Fprintf(w, "response Headers: %v \n", resp.Header)
          body, _ := ioutil.ReadAll(resp.Body)
          var jsonRecive interface{}
          json.Unmarshal(body, &jsonRecive)
          fmt.Fprintf(w, "response Body: %v \n", jsonRecive)
      default:
          fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
      }
}

func calculate(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
  var twoNumbers calculateP.Numbers
  var twoNumbersRec calculateP.NumbersRec
  var bad = calculateP.Errorstruct{"Incorrect input"}
  if err := json.NewDecoder(r.Body).Decode(&twoNumbers); err != nil {
    fmt.Fprintf(w, "json.Decoder err: %v", err)
    send(w, http.StatusBadRequest, bad)
    return
  }
  if (twoNumbers.First < 0) || (twoNumbers.Second < 0) {
    send(w, http.StatusBadRequest, bad)
    return
  }
  c := make(chan int)
  go factorial(twoNumbers.First, c)
  go factorial(twoNumbers.Second, c)
  twoNumbersRec.Second, twoNumbersRec.First = <-c, <-c
  send(w, http.StatusOK, twoNumbersRec)
}
func handleRequests(){
  router := httprouter.New()
  router.GET("/", homePage)
  router.POST("/", homePage)
  router.POST("/calculate", calculate)
  log.Fatal(http.ListenAndServe(port, router))
  }

  func main() {
handleRequests()

  }
  // send returns a JSON encoded representation of `val`
// with status code `code`.
func send(w http.ResponseWriter, code int, val interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    return json.NewEncoder(w).Encode(val)
}

func factorial(n int, c chan int) {
  if n == 0 {
    c <- 1
  } else {
    muxx :=1;
    for i := 1; i < n; i++ {
      muxx = muxx * i
    }
    muxx = muxx * n
    c <-muxx
  }

}
