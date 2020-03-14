package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
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


    fmt.Println("==========================================")
    dump, err := httputil.DumpRequest(r, true)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "%q", dump)
    fmt.Printf("%q", dump)
}

func main() {
    finish := make(chan bool)
    http.HandleFunc("/", handler)
    go func() {
      http.ListenAndServe(":8080", nil)
    }()
    go func() {
      http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
    }()
    <-finish
}
