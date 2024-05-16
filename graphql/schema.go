package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/priyanshu-banshisoft/tmdb_graphql/tmdb"
)



func NewSchema(apiKey string) graphql.Schema {
    client := tmdb.NewClient(apiKey)

    var queryType = graphql.NewObject(
        graphql.ObjectConfig{
            Name: "Query",
            Fields: graphql.Fields{
                "popularMovies": &graphql.Field{
                    Type: graphql.NewList(movieType),
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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

    schema, _ := graphql.NewSchema(
        graphql.SchemaConfig{
            Query: queryType,
        },
    )

    return schema
}