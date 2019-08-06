package resolver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"upkeep/generated"
	"upkeep/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	users []*models.User
}

func (resolver *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{resolver}
}
func (resolver *Resolver) Query() generated.QueryResolver {
	return &queryResolver{resolver}
}

type mutationResolver struct{ *Resolver }

func (resolver *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	user := &models.User{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Name:  input.Name,
		Email: input.Email,
	}
	resolver.users = append(resolver.users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }

func (resolver *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	_, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return resolver.users, nil
}
func (resolver *queryResolver) User(ctx context.Context, email *string) ([]*models.User, error) {
	_, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var result []*models.User
	for _, v := range resolver.users {
		if strings.EqualFold(v.Email, *email) {
			result = append(result, v)
		}
	}
	return result, nil

}

// Define a function to recover the gin.Context from the context.Context struct
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}