package http

type RequestHandler interface {
	OnHttpRequest([]byte) (string, error)
}

type Server interface {
	RegisterHttpHandler([]Method, Target, RequestHandler)
	Run()
}
