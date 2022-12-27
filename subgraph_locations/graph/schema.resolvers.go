package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"subgraph_locations/graph/generated"
	"subgraph_locations/graph/model"
)

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, id string) (*model.Location, error) {
	// hacky validation
	if id == "" {
		log.Println("empty id")
		return nil, nil
	}

	resp, err := r.RestyClient.R().
		SetQueryParams(map[string]string{
			"access_key": r.PositionstackAccessKey,
			"query":      id,
			"limit":      "1",
		}).
		SetHeader("Accept", "application/json").
		Get("http://api.positionstack.com/v1/reverse")

	if err != nil {
		log.Printf("Location ERROR: positionstack reverse geocoding failed, %s", err.Error())
		return nil, err
	}

	// hacky unmarshal rather than -> autogen structs
	jsonMap := make(map[string](interface{}))
	jsonErr := json.Unmarshal(resp.Body(), &jsonMap)
	if jsonErr != nil {
		log.Printf("ERROR: fail to unmarshal json, %s", err.Error())
		return nil, jsonErr
	}
	log.Printf("INFO: jsonMap:, %v", jsonMap)
	jsonData := jsonMap["data"].([]interface{})
	locationData, ok := jsonData[0].(map[string]interface{})
	if !ok {
		log.Printf("ERROR: failed to cast location data, %s", "Sorry sometimes we can't rely on positionstack: https://github.com/apilayer/positionstack/issues/7")
		return nil, errors.New("error: Sorry sometimes we can't rely on positionstack: https://github.com/apilayer/positionstack/issues/7")
	}

	latitude := locationData["latitude"].(float64)
	longitude := locationData["longitude"].(float64)
	label := locationData["label"].(string)
	locality := locationData["locality"].(string)
	county := locationData["county"].(string)
	continent := locationData["continent"].(string)

	return &model.Location{ID: id,
		Latitude:  &latitude,
		Longitude: &longitude,
		Label:     &label,
		Locality:  &locality,
		County:    &county,
		Continent: &continent,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
