package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"subgraph_eateries/graph/generated"
	"subgraph_eateries/graph/model"

	"github.com/jmatth11/yfusion"
)

// EateriesForLocation is the resolver for the eateriesForLocation field.
func (r *locationResolver) EateriesForLocation(ctx context.Context, obj *model.Location) ([]*model.Eatery, error) {
	bs := &yfusion.BusinessSearchParams{}
	bs.SetLatitude(*obj.Latitude)
	bs.SetLongitude(*obj.Longitude)
	bs.SetTerm("food")
	result, err := r.YelpClient.SearchBusiness(bs)
	if err != nil {
		log.Printf("EateriesForLocation ERROR: failed to search business, %s", err.Error())
		return nil, err
	}

	eateries := []*model.Eatery{}
	for _, b := range result.Businesses {
		rating := b.Rating
		distance := b.Distance
		reviewCount := b.ReviewCount
		url := b.URL
		eateries = append(eateries, &model.Eatery{
			ID:          b.ID,
			Name:        b.Name,
			Rating:      &rating,
			Distance:    &distance,
			ReviewCount: &reviewCount,
			URL:         &url,
		})
	}

	return eateries, nil
}

// Eatery is the resolver for the eatery field.
func (r *queryResolver) Eatery(ctx context.Context, id string) (*model.Eatery, error) {
	result, err := r.YelpClient.SearchBusinessDetails(id)
	if err != nil {
		log.Printf("Eatery ERROR: failed to search business details, %s", err.Error())
		return nil, err
	}

	return &model.Eatery{
		ID:          id,
		Name:        result.Name,
		Rating:      &result.Rating,
		Distance:    &result.Distance,
		ReviewCount: &result.ReviewCount,
		URL:         &result.URL,
		Location: &model.Location{
			ID:        fmt.Sprintf("%f%s%f", result.Coordinates.Latitude, ",", result.Coordinates.Longitude),
			Latitude:  &result.Coordinates.Latitude,
			Longitude: &result.Coordinates.Longitude,
		},
	}, nil
}

// EateriesForCity is the resolver for the eateriesForCity field.
func (r *queryResolver) EateriesForCity(ctx context.Context, city string) ([]*model.Eatery, error) {
	bs := &yfusion.BusinessSearchParams{}
	bs.SetLocation(city)
	bs.SetTerm("food")
	result, err := r.YelpClient.SearchBusiness(bs)
	if err != nil {
		log.Printf("EateriesForCity ERROR: failed to search business, %s", err.Error())
		return nil, err
	}

	eateries := []*model.Eatery{}
	for _, b := range result.Businesses {
		rating := b.Rating
		distance := b.Distance
		reviewCount := b.ReviewCount
		url := b.URL
		eateries = append(eateries, &model.Eatery{
			ID:          b.ID,
			Name:        b.Name,
			Rating:      &rating,
			Distance:    &distance,
			ReviewCount: &reviewCount,
			URL:         &url,
		})
	}

	return eateries, nil
}

// Location returns generated.LocationResolver implementation.
func (r *Resolver) Location() generated.LocationResolver { return &locationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type locationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
