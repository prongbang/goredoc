# goredoc

[![GoDoc](https://godoc.org/github.com/prongbang/goredoc?status.svg)](https://godoc.org/github.com/prongbang/goredoc)
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/goredoc?_=1)](https://goreportcard.com/report/github.com/prongbang/goredoc?_=1)

Go [ReDoc](https://github.com/Redocly/redoc) is an embedded OpenAPI/Swagger documentation for Go, with handler implementations for: `net/http`, `fiber`, `gin`, and `echo`.

```shell
go get github.com/prongbang/goredoc
```

## How to use

```go
doc := goredoc.New(goredoc.Config{
    Title:   "API Documentation",
    SpecURL: "https://petstore.swagger.io/v2/swagger.json",
})
```

- net/http

```go
import (
    "github.com/prongbang/goredoc"
    "net/http"
)

http.HandleFunc("/docs", doc.Handler())
log.Fatal(http.ListenAndServe(":8080", nil))
```

- fiber

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/adaptor"
    "github.com/prongbang/goredoc"
)

f := fiber.New()
f.Get("/docs", adaptor.HTTPHandlerFunc(doc.Handler()))
log.Fatal(f.Listen(":8080"))
```

- gin

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/prongbang/goredoc"
)

g := gin.New()
g.GET("/docs", gin.WrapF(doc.Handler()))
log.Fatal(g.Run(":8080"))
```

- echo

```go
import (
    "github.com/labstack/echo/v4"
    "github.com/prongbang/goredoc"
)

e := echo.New()
e.GET("/docs", echo.WrapHandler(doc.Handler()))
e.Logger.Fatal(e.Start(":8080"))
```

## Listen

[http://localhost:8080/docs](http://localhost:8080/docs)