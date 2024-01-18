package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"main.go/graph"
)

func main() {
	app := fiber.New()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	app.Post("/query", func(c *fiber.Ctx) error {
		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			srv.ServeHTTP(w, r)
		})
		fasthttpadaptor.NewFastHTTPHandlerFunc(h)(c.Context())
		return nil
	})

	log.Fatal(app.Listen(":3030"))
}
