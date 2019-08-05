package resolver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"upkeep/generated"
	"upkeep/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	users []*models.User
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	user := &models.User{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Name:  input.Name,
		Email: input.Email,
	}
	r.users = append(r.users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	_, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.users, nil
}

//type userResolver struct{ *Resolver }
//
//func (r *userResolver) User(ctx context.Context, obj *models.User) (*models.User, error) {
//	return &models.User{ID: obj.ID, Name: "name " + obj.Name, Email: "email " + obj.Email}, nil
//}

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