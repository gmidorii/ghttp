package ghttp

const (
	HTTP  = "http"
	HTTPS = "https"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

type GHttp struct {
	Protocol string
	Domain   string
	Path     string
	Method   string
	Query    map[string]string
}

func NewGHttp(protocol, domain, path, method string) *GHttp {
	return &GHttp{
		Protocol: protocol,
		Domain:   domain,
		Path:     path,
		Method:   method,
	}
}
