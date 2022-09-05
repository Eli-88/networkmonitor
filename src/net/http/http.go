package http

import (
	"io"
	"io/ioutil"
	"net/http"
	"networkmonitor/logger"
)

// interface compliance
var _ Server = &httpServer{}

type Method string
type Target string

const (
	GET  Method = "GET"
	POST Method = "POST"
)

type httpServer struct {
	mux  *http.ServeMux
	addr string
}

func MakeServer(addr string) Server {
	return &httpServer{
		mux:  http.NewServeMux(),
		addr: addr,
	}
}

func (h *httpServer) RegisterHttpHandler(methods []Method, target Target, handler RequestHandler) {
	h.mux.HandleFunc(string(target), func(w http.ResponseWriter, r *http.Request) {
		for _, v := range methods {
			if v == Method(r.Method) {
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					logger.Error(err)
					goto Fail
				}
				response, err := handler.OnHttpRequest(body)
				if err != nil {
					logger.Error(err)
					goto Fail
				}
				io.WriteString(w, response)
				goto End
			}
		}
	Fail:
		io.WriteString(w, "invalid request")
	End:
		return
	})
}

func (h *httpServer) Run() error {
	return http.ListenAndServe(h.addr, h.mux)
}
