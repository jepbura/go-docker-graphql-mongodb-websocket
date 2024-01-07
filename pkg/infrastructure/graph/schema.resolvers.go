package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/infrastructure/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if r.Usecase.UserUsecaseInterface == nil {
		return nil, fmt.Errorf("not implemented: CreateUser - CreateUser")
	}

	user, err := r.Usecase.UserUsecaseInterface.SaveUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	if r.Usecase.UserUsecaseInterface == nil {
		return "", fmt.Errorf("not implemented: DeleteUser - DeleteUser")
	}

	deleteUserId, err := r.Usecase.UserUsecaseInterface.DeleteUser(ctx, id)
	if err != nil {
		return "", err
	}

	return deleteUserId, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	if r.Usecase.UserUsecaseInterface == nil {
		return nil, fmt.Errorf("not implemented: Users - users")
	}

	users, err := r.Usecase.UserUsecaseInterface.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	if r.Usecase.UserUsecaseInterface == nil {
		return nil, fmt.Errorf("not implemented: User - user")
	}

	user, err := r.Usecase.UserUsecaseInterface.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
