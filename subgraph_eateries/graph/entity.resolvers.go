package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"subgraph_eateries/graph/generated"
	"subgraph_eateries/graph/model"
)

// FindLocationByIDAndLatitudeAndLongitude is the resolver for the findLocationByIDAndLatitudeAndLongitude field.
func (r *entityResolver) FindLocationByIDAndLatitudeAndLongitude(ctx context.Context, id string, latitude *float64, longitude *float64) (*model.Location, error) {
	return &model.Location{
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
