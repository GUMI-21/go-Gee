package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	Roots    map[string]*node
	Handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Roots:    make(map[string]*node),
		Handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.Roots[method]
	if !ok {
		r.Roots[method] = &node{}
	}
	r.Roots[method].insert(pattern, parts, 0)
	r.Handlers[key] = handler
}

func (r *Router) GetRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.Roots[method]

	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *Router) Handle(c *Context) {
	n, params := r.GetRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.Handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND\n")
	}
}
