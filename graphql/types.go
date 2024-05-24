package graphql

import "github.com/graphql-go/graphql"

var movieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movie",
		Fields: graphql.Fields{
			"backdrop_path": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"original_title": &graphql.Field{
				Type: graphql.String,
			},
			"overview": &graphql.Field{
				Type: graphql.String,
			},
			"poster_path": &graphql.Field{
				Type: graphql.String,
			},
			"media_type": &graphql.Field{
				Type: graphql.String,
			},
			"adult": &graphql.Field{
				Type: graphql.Boolean,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"original_language": &graphql.Field{
				Type: graphql.String,
			},
			"genre_ids": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"genres" : &graphql.Field{
				Type: graphql.NewList(genres),
			},
			"popularity": &graphql.Field{
				Type: graphql.Float,
			},
			"release_date": &graphql.Field{
				Type: graphql.String,
			},
			"video": &graphql.Field{
				Type: graphql.Boolean,
			},
			"vote_average": &graphql.Field{
				Type: graphql.Float,
			},
			"vote_count": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var genres = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Genres",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var credit = graphql.NewObject(
	graphql.ObjectConfig{
		Name : "Credits",
		Fields: graphql.Fields{
			"id": &graphql.Field{ 
				Type: graphql.Int,
			},
			"cast": &graphql.Field{
				Type: graphql.NewList(cast),
			},
		},
	},
)

var cast = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Cast",
		Fields: graphql.Fields{
			"id": &graphql.Field{ 
				Type: graphql.Int,
			},
			"profile_path": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
