package main



import (

    "fmt"

    "net/http"

)



func helloWorld(w http.ResponseWriter, r *http.Request){

    fmt.Fprintf(w, "Hello World v2")

}



func main() {
    fmt.Print("o deploy rodou")
    http.HandleFunc("/", helloWorld)

    http.ListenAndServe(":6111", nil)

}
