package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"subgraph_eateries/graph/generated"
	"subgraph_eateries/graph/model"
)

// FindLocationByIDAndLatitudeAndLongitude is the resolver for the findLocationByIDAndLatitudeAndLongitude field.
func (r *entityResolver) FindLocationByIDAndLatitudeAndLongitude(ctx context.Context, id string, latitude *float64, longitude *float64) (*model.Location, error) {
	fmt.Printf("THIS IS ID: %v\n", id)
	fmt.Printf("THIS IS latitude: %v\n", latitude)
	fmt.Printf("THIS IS longitude: %v\n", longitude)
	return &model.Location{
		ID:        id,
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
	// panic(fmt.Errorf("not implemented: FindLocationByIDAndLatitudeAndLongitude - findLocationByIDAndLatitudeAndLongitude"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
