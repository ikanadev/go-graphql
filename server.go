package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/vkevv/go-graphql/src/config"
	"github.com/vkevv/go-graphql/src/db"
	"github.com/vkevv/go-graphql/src/graph"
	"github.com/vkevv/go-graphql/src/graph/generated"
	"github.com/vkevv/go-graphql/src/graph/model"
	"github.com/vkevv/go-graphql/src/middleware"
)

// go run github.com/99designs/gqlgen generate

func graphQLHandler(DB *pg.DB, conf config.Config) gin.HandlerFunc {
	config := generated.Config{Resolvers: graph.NewResolver(DB, conf)}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playGroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	conf := config.GetConfig()

	DB := db.New(&pg.Options{
		User:     conf.DB.Username,
		Password: conf.DB.Password,
		Database: conf.DB.Database,
		Addr:     conf.DB.Host + ":" + conf.DB.Port,
	})
	DB.AddQueryHook(db.Logger{})
	defer DB.Close()

	err := db.CreateTables(DB, &model.User{}, &model.Todo{})
	if err != nil {
		panic("Can't create tables" + err.Error())
	}

	router := gin.New()
	router.Use(middleware.AuthMiddleware(conf))
	router.POST("/query", graphQLHandler(DB, conf))
	router.GET("/", playGroundHandler())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.App.Port)
	router.Run(":" + conf.App.Port)
}
