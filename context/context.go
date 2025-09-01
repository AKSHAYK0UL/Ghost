package context

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// context for request and response
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Params  map[string]string
}

// JSON response
func (ctx *Context) JSON(code int, data any) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(code)
	json.NewEncoder(ctx.Writer).Encode(data)
}

// sends plain text response
func (ctx *Context) String(code int, text string) {
	ctx.Writer.Header().Set("Content-Type", "text/plain")
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Write([]byte(text))
}

// read JSON from request body
func (ctx *Context) BindJSON(v any) error {
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(v)
}

// query parameter
func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

// query parameter with default value
func (ctx *Context) QueryDefault(key, defaultValue string) string {
	if value := ctx.Query(key); value != "" {
		return value
	}
	return defaultValue
}

// query parameter as int
func (ctx *Context) QueryInt(key string) (int, error) {
	value := ctx.Query(key)
	if value == "" {
		return 0, fmt.Errorf("parameter %s not found", key)
	}
	return strconv.Atoi(value)
}

// request header
func (ctx *Context) GetHeader(key string) string {
	return ctx.Request.Header.Get(key)
}

// SetHeader
func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

// status code
func (ctx *Context) Status(code int) {
	ctx.Writer.WriteHeader(code)
}

// error response
func (ctx *Context) Error(code int, message string) {
	ctx.JSON(code, map[string]string{"error": message})
}
