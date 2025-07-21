package http

import (
	"fmt"
	"net/http"
	"strings"
)

type RouterGroup struct {
	prefix string
	mux    *http.ServeMux
}

func NewRouterGroup (prefix string, mux *http.ServeMux) *RouterGroup {
	return &RouterGroup{prefix ,mux}
}

func (g *RouterGroup) HandleFunc(pattern string, handlerFunc http.HandlerFunc) {
	parts := strings.SplitN(pattern, " ", 2)

	if len(parts) > 2{
		panic("invalid methodAndRoute format; expected 'METHOD /route'")
	}

	method := parts[0]
	route := parts[1]

	fullRoute := g.prefix + strings.TrimPrefix(route, "/")

	fullPattern := method + " " + fullRoute

	fmt.Printf("Registered route at -> %s\n", fullPattern)

	g.mux.HandleFunc(fullPattern, handlerFunc)
}