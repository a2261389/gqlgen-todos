// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/a2261389/gqlgen-todos/graph/generated"
	"github.com/a2261389/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	for index, val := range r.todos {
		if input.ID == val.ID {
			r.todos[index] = &model.Todo{
				ID:   val.ID,
				Text: input.Text,
				User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
			}
			return r.todos[index], nil
		}
	}
	return nil, errors.New("not found todo")
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input *model.DeleteTodo) (bool, error) {
	for index, val := range r.todos {
	fmt.Printf("%#v\n", val)
		if input.ID == val.ID {
			r.todos = append(r.todos[:index], r.todos[index+1:]...)
			return true, nil
		}
	}
	return false, nil
}


func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
