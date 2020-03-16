package esearch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

const (
	IdPrefixLocation = "location:"
	TypeLocation     = "location"
)

// Includes the type field
type ESLocation struct {
	*qvspot.Location
	Type string `json:"type"`
}

// LocationInsert inserts/replaces a location
func (e *esearch) LocationInsert(ctx context.Context, location *qvspot.Location) error {

	request := elastic.NewBulkIndexRequest().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixLocation + location.Id).
		Doc(&ESLocation{Location: location, Type: TypeLocation})

	request.Source() // Convert to json so we can modify b
	e.bulk.Add(request)

	return nil

}

// LocationGetById returns a location by id
func (e *esearch) LocationGetById(ctx context.Context, locationId string) (*qvspot.Location, error) {

	res, err := e.client.Get().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixLocation + locationId).
		Do(ctx)
	if elastic.IsNotFound(err) {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("Could not get location: %v", err)
	}

	// Unmarshal the block
	location := new(qvspot.Location)
	err = json.Unmarshal(res.Source, location)
	return location, err

}

// LocationDeleteById removes a location by id
func (e *esearch) LocationDeleteById(ctx context.Context, locationId string) error {

	_, err := e.client.Delete().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixLocation + locationId).
		Refresh("true").
		Do(ctx)
	return err
}
