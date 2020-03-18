package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/phamstack/godek/graph/generated"
	"github.com/phamstack/godek/helpers"
	"github.com/phamstack/godek/models"
)

func (r *mutationResolver) LoginGoogle(ctx context.Context, token string, name string, email string, avatar string) (*models.Auth, error) {
	user, err := r.Services.User.ByEmail(email)

	if err != models.ErrNotFound && err != nil {
		return nil, err
	}

	if err == models.ErrNotFound {
		userCount := r.Services.User.Count()
		username := helpers.GenerateUsername(email, userCount)

		newUser := &models.User{
			Name:     name,
			Email:    email,
			Username: username,
			Avatar:   avatar,
		}
		r.Services.User.Create(newUser)

		authToken := r.Services.User.GenerateAuthToken(newUser)

		return &models.Auth{
			User:  newUser,
			Token: authToken,
		}, nil
	}

	authToken := r.Services.User.GenerateAuthToken(user)
	return &models.Auth{
		User:  user,
		Token: authToken,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (*models.User, error) {
	fmt.Println(ctx)
	return nil, nil
}

func (r *mutationResolver) LogoutAll(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
