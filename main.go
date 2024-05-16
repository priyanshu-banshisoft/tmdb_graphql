package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  apiKey := os.Getenv("NEATFLIX_API_KEY")
  if apiKey == "" {
	log.Fatal("TMDB_API_KEY environment variable is required")
}

h := handler.New(&handler.Config{
	Schema: &graphql.Schema{},
	Pretty: true,
})

http.Handle("/graphql", h)

log.Println("Server is running on port 8080")
log.Fatal(http.ListenAndServe(":8080", nil))

}
