package main

import (
    "fmt"
    "net/http"
    "sort"
)

func handler(w http.ResponseWriter, r *http.Request) {

    // adding debug header to test (strong/weak) ETags in combination with NGINX
    w.Header().Set("ETag", "HelloWorld")

    var requestKeys []string
    for k := range r.Header {
        requestKeys = append(requestKeys, k)
    }
    sort.Strings(requestKeys)

    fmt.Println("==========================================")
    fmt.Fprintln(w, "<b>request.RequestURI:</b>", r.RequestURI, "</br>")
    fmt.Println("<b>request.RequestURI:</b>", r.RequestURI, "</br>")
    fmt.Fprintln(w, "<b>request.RemoteAddr:</b>", r.RemoteAddr, "</br>")
    fmt.Println("<b>request.RemoteAddr:</b>", r.RemoteAddr, "</br>")
    fmt.Fprintln(w, "<b>request.TLS:</b>", r.TLS, "</br>")
    fmt.Println("<b>request.TLS:</b>", r.TLS, "</br>")


    fmt.Fprintln(w, "<b>Request Headers:</b></br>")
    fmt.Println("<b>Request Headers:</b></br>")
    for _, k := range requestKeys {
        fmt.Fprintln(w, k, ":", r.Header[k], "</br>")
        fmt.Println(k, ":", r.Header[k], "</br>")
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
