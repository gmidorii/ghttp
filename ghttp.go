package ghttp

import "fmt"

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

// GHttp is main struct
type GHttp struct {
	Protocol string
	Domain   string
	Path     string
	Method   string
	Query    map[string]string
}

// NewGHttp create GHttp struct
func NewGHttp(protocol, domain, path, method string) *GHttp {
	return &GHttp{
		Protocol: protocol,
		Domain:   domain,
		Path:     path,
		Method:   method,
		Query:    make(map[string]string),
	}
}

// PutQuery is putting query paramaeter.
// same value overwrite new value
func (g *GHttp) PutQuery(key, v string) {
	g.Query[key] = v
}

func (g *GHttp) Exec() ([]byte, error) {
	switch g.Method {
	case GET:
		return execGet(*g)
	case POST:
	case DELETE:
	case PUT:
	default:
		return nil, fmt.Errorf("unexpected method: %v", g.Method)
	}
}
