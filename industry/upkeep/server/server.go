package main

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"log"
	"os"
	//"log"
	"github.com/gin-gonic/gin"
	//"net/http
	"upkeep/generated"
	"upkeep/resolver"
)

const defaultPort = "9000"
const defaultGraphqlEndpoint = "graphql"

// At the Resolver level, gqlgen gives you access to the context.Context object.
// One way to access the gin.Context is to add it to the context and retrieve it again.
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}


// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/" + defaultGraphqlEndpoint)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	// router := gin.Default() creates a default middleware.
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	//router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(GinContextToContextMiddleware())
	router.POST("/" + defaultGraphqlEndpoint, graphqlHandler())
	router.GET("/", playgroundHandler())

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground handling by Gin", port)
	router.Run(":" + port)
}

//func main() {
//	port := os.Getenv("PORT")
//	if port == "" {
//		port = defaultPort
//	}
//
//	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
//	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))
//
//	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
//	log.Fatal(http.ListenAndServe(":"+port, nil))
//}
