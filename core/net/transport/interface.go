package transport

type HttpRequestHandler interface {
	OnHttpRequest([]byte) (string, error)
}

type HttpServer interface {
	RegisterHttpHandler([]HttpMethod, HttpTarget, HttpRequestHandler)
	Run()
}
