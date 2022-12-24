package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"subgraph_eateries/graph/generated"
	"subgraph_eateries/graph/model"
)

// FindLocationByID is the resolver for the findLocationByID field.
func (r *entityResolver) FindLocationByID(ctx context.Context, id string) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: FindLocationByID - findLocationByID"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
