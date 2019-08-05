package main


/**
 * go run github.com/99designs/gqlgen
 * go run server/server.go
 *
 * Resolvers isn't auto generated after gqlgen.
 */

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
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
	// Disable Console Coloer, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Logger())

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
