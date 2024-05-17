package main

import (
	"context"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/priyanshu-banshisoft/tmdb_graphql/graphql"
)

func main() {
    h := handler.New(&handler.Config{
        Schema: &graphql.Schema,
        Pretty: true,
    })

    http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("API_KEY")
        if apiKey == "" {
            http.Error(w, "Authorization header is required", http.StatusUnauthorized)
            return
        }

        ctx := r.Context()
        ctx = context.WithValue(ctx, "apiKey", apiKey)
        r = r.WithContext(ctx)
        h.ContextHandler(ctx, w, r)
    })

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
