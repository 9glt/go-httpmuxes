# go-httpmuxes
Add same urls to several http.NewServeMux


```go
package main

import (
    httpmuxes "github.com/9glt/go-httpmuxes"
    "net/http"
    "log"
)
func main() {
    httpmuxes.Register("https", http.NewServeMux())
    httpmuxes.Register("http", http.NewServeMux())
    // ....

    httpmuxes.HandleFuncAll("/http-https", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("hi"))
    })

    httpmuxes.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hi http"))
    }, []string{"http"})

    httpmuxes.HandleFunc("/https", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hi https"))
    }, []string{"https"})

    // ...

    go func() { log.Fatal(http.ListenAndServe(":80", httpmuxes.Get("http"))}) 
    log.Fatal(http.ListenAndServe(":443", httpmuxes.Get("https")))
}
```