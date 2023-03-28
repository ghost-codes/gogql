package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"math/rand"
	"fmt"

	"github.com/ghost-codes/gogql/graph/model"
)

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
    video := model.Video{
        ID: fmt.Sprintf("T%d",rand.Int()),
        Title: input.Title,
        URL: input.URL,
        Auther: &model.User{ID: input.UserID, Name: "user",},
    }
    name :="Hope"
    r.Store.Queries.CreateUser(ctx,&name)
    return &video,nil
        
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
    return nil,nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
