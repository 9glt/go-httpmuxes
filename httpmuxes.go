package httpmuxes

import "net/http"

var (
	muxes = map[string]*http.ServeMux{}
)

func Register(name string, mux *http.ServeMux) {
	muxes[name] = mux
}

func Get(name string) *http.ServeMux {
	return muxes[name]
}

func HandleFuncAll(path string, fn http.HandlerFunc) {
	for _, mux := range muxes {
		mux.HandleFunc(path, fn)
	}
}

func HandleFunc(path string, fn http.HandlerFunc, muxs []string) {
	for _, mux := range muxs {
		muxes[mux].HandleFunc(path, fn)
	}
}
