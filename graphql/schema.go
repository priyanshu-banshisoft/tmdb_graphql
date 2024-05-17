package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/priyanshu-banshisoft/tmdb_graphql/tmdb"
)



var Schema graphql.Schema

func init() {
    var queryType = graphql.NewObject(
        graphql.ObjectConfig{
            Name: "Query",
            Fields: graphql.Fields{
                "popularMovies": &graphql.Field{
                    Type: graphql.NewList(movieType),
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        return client.FetchPopularMovies()
                    },
                },
                "movie": &graphql.Field{
                    Type: movieType,
                    Args: graphql.FieldConfigArgument{
                        "id": &graphql.ArgumentConfig{
                            Type: graphql.Int,
                        },
                    },
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        apiKey := p.Context.Value("apiKey").(string)
                        client := tmdb.NewClient(apiKey)
                        id, ok := p.Args["id"].(int)
                        if ok {
                            return client.FetchMovieDetails(id)
                        }
                        return nil, nil
                    },
                },
            },
        },
    )

    var err error
    Schema, err = graphql.NewSchema(
        graphql.SchemaConfig{
            Query: queryType,
        },
    )
    if err != nil {
        panic(err)
    }
}