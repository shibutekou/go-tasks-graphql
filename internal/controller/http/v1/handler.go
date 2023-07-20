package v1

import (
	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/vgekko/go-tasks-graphql/graph"
	"github.com/vgekko/go-tasks-graphql/internal/usecase"
	"golang.org/x/exp/slog"
)

func NewHandler(handler *gin.Engine, useCase *usecase.Usecase, log *slog.Logger) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())

	v1 := handler.Group("/v1")
	{
		v1.POST("/graphql", graphqlHandler(useCase, log))

		newTaskRoutes(v1, useCase.Task, log)
	}

}

func graphqlHandler(uc *usecase.Usecase, log *slog.Logger) gin.HandlerFunc {
	h := graphqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Uc: uc, Log: log}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
