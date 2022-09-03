package transport

import (
	"io"
	"io/ioutil"
	"net/http"
	"networkmonitor/core/logger"
)

// interface compliance
var _ HttpServer = &httpServer{}

type HttpMethod string
type HttpTarget string

const (
	HTTP_GET  HttpMethod = "GET"
	HTTP_POST HttpMethod = "POST"
)

type httpServer struct {
	mux  *http.ServeMux
	addr string
}

func MakeHttpServer(addr string) HttpServer {
	return &httpServer{
		mux:  http.NewServeMux(),
		addr: addr,
	}
}

func (h *httpServer) RegisterHttpHandler(methods []HttpMethod, target HttpTarget, handler HttpRequestHandler) {
	h.mux.HandleFunc(string(target), func(w http.ResponseWriter, r *http.Request) {
		for _, v := range methods {
			if v == HttpMethod(r.Method) {
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

func (h *httpServer) Run() {
	http.ListenAndServe(h.addr, h.mux)
}
