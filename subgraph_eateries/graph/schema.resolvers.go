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
	fmt.Printf("INSIDE EATERIES FOR LOCATION: %v\n", obj)
	fmt.Printf("obj: %v\n", obj)
	fmt.Println("AFTER obj:")
	fmt.Printf("OBJ.LATITUDE: %v\n", *obj.Latitude)
	fmt.Printf("OBJ.LONGITUDE: %v\n", *obj.Longitude)
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
		fmt.Printf("Name: %s, Longitude: %f, Latitude: %f, Price: %s, Distance: %f, Rating: %f, ID: %s\n", b.Name, b.Coordinates.Longitude, b.Coordinates.Latitude, b.Price, b.Distance, b.Rating, b.ID)
		rating := b.Rating
		eateries = append(eateries, &model.Eatery{
			ID:     b.ID,
			Name:   b.Name,
			Rating: &rating,
		})
	}

	return eateries, nil
}

// Eatery is the resolver for the eatery field.
func (r *queryResolver) Eatery(ctx context.Context, id string) (*model.Eatery, error) {
	fmt.Printf("INSIDE EATERY: %v\n", id)
	result, err := r.YelpClient.SearchBusinessDetails(id)
	if err != nil {
		log.Printf("Eatery ERROR: failed to search business details, %s", err.Error())
		return nil, err
	}
	fmt.Printf("result: %v\n", result)
	fmt.Printf("id: %v\n", id)
	return &model.Eatery{
		ID:     id,
		Name:   result.Name,
		Rating: &result.Rating,
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
		fmt.Printf("Name: %s, Longitude: %f, Latitude: %f, Price: %s, Distance: %f, Rating: %f, ID: %s\n", b.Name, b.Coordinates.Longitude, b.Coordinates.Latitude, b.Price, b.Distance, b.Rating, b.ID)
		rating := b.Rating
		eateries = append(eateries, &model.Eatery{
			ID:     b.ID,
			Name:   b.Name,
			Rating: &rating,
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
