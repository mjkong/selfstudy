package main

import(
    "os"
    "fmt"
    "net/http"
    "gopkg.in/alecthomas/kingpin.v2"
)

var (
  app = kingpin.New("mj", "kingpin test")

  start     = app.Command("start", "Start the orderer node").Default()
  version   = app.Command("version", "Show version information")
  benchmark = app.Command("benchmark", "Run orderer in benchmark mode")

)


func main(){

  fullCmd := kingpin.MustParse(app.Parse(os.Args[1:]))
             fmt.Println(fullCmd)

             // "version" command
  if fullCmd == version.FullCommand() {
    fmt.Println("1.0")
    return
  }

  if fullCmd == start.FullCommand(){
    http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello World"))
    })

    http.ListenAndServe(":5000", nil)
  }
}
