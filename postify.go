package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	remoteUri := os.Getenv("POSTIFY_REMOTE_URL")
	if remoteUri == "" {
		panic("Please set the env 'POSTIFY_REMOTE_URL' with the target remote host")
	}
	remote, err := url.Parse(remoteUri)
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				log.Println("POSTifying request to POST from GET", r.URL)
				r.Method = http.MethodPost
			}
			log.Printf("%s: %s /n", r.Method, r.URL)
			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
