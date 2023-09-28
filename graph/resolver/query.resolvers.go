package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/kitadesign/gqlgen-test/graph"
)

// Status is the resolver for the status field.
func (r *queryResolver) Status(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }