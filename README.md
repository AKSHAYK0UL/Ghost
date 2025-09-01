# Ghost - Minimal HTTP Framework in Go

Ghost is a lightweight HTTP framework built in Go. It provides a simple way to define routes, handle requests, and send responses without relying on third-party frameworks.

---

## 🚀 Features

* Minimal and lightweight design
* Routing support (`GET`, `POST`, `PUT`, `DELETE`)
* Custom request/response context
* Easy JSON and plain text responses
* Query parameters and headers support
* Error handling helper

---

## 📂 Project Structure

```
Ghost/
├── context/        # Custom Context (request/response handling)
│   └── context.go
├── ghost/          # Framework core (router, server)
│   └── ghost.go
└── main.go         # Example usage
```

---

## 🛠️ Installation

Clone the repo:

```bash
git clone https://github.com/AKSHAYK0UL/Ghost.git
cd Ghost
```

Run the server:

```bash
go run main.go
```

---

## 📖 Example

```go
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

	// GET route
	mux.GET("/", func(ctx *context.Context) {
		ctx.String(http.StatusOK, data)
	})

	// POST route
	mux.POST("/add", func(ctx *context.Context) {
		var body string
		ctx.BindJSON(&body)
		fmt.Println("Body:", body)
		data = body
		ctx.String(http.StatusAccepted, "Message Accepted")
	})

	// Start server
	mux.Start("8080")
}
```

---

## 📬 API Example

### GET Request

```bash
curl http://localhost:8080/
```

### POST Request

```bash
curl -X POST http://localhost:8080/add \
     -H "Content-Type: application/json" \
     -d '"Hello Ghost"'
```

Response:

```
Message Accepted
```

---

## 📌 TODO

* Add route parameters (`/add/:id`) extraction
* Add middleware support

---
