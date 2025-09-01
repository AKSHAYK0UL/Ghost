package main

import (
	"fmt"
	"net/http"

	"ghost_0.1/context"
	ghost "ghost_0.1/ghost"
)

func main() {
	fmt.Println("Ghost")
	data := "None"

	mux := ghost.New()

	mux.GET("/", func(ctx *context.Context) {
		ctx.String(http.StatusOK, data)
	})

	mux.POST("/add/:id", func(ctx *context.Context) {
		var body string
		ctx.BindJSON(&body)
		fmt.Println("Body:", body)
		data = body
		ctx.String(http.StatusAccepted, "Message Accepted")
	})

	mux.Start("8080")
}
