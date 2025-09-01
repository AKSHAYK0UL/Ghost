package ghost

import (
	"fmt"
	"log"
	"net/http"

	"ghost_0.1/context"
)

type HandlerFunc func(*context.Context)

type Ghost struct {
	routes map[string]HandlerFunc
}

// constructor
func New() *Ghost {
	return &Ghost{routes: make(map[string]HandlerFunc)}
}

// Get route
func (g *Ghost) GET(path string, handler HandlerFunc) {

	g.routes[http.MethodGet+" "+path] = handler
}

// POST route
func (g *Ghost) POST(path string, handler HandlerFunc) {
	g.routes[http.MethodPost+" "+path] = handler
}

// PUT route
func (g *Ghost) PUT(path string, handler HandlerFunc) {
	g.routes[http.MethodPut+" "+path] = handler
}

// DELETE route
func (g *Ghost) DELETE(path string, handler HandlerFunc) {
	g.routes[http.MethodDelete+" "+path] = handler
}

// handles all requests
func (g *Ghost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &context.Context{
		Request: r,
		Writer:  w,
		Params:  make(map[string]string),
	}

	// Find handler
	key := r.Method + " " + r.URL.Path
	if handler, exists := g.routes[key]; exists {
		handler(ctx)
	} else {
		ctx.Error(404, "Route not found")
	}
}

// Start server
func (g *Ghost) Start(port string) {
	fmt.Printf("Server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, g))
}
