package esearch

import (
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

const (
	IdPrefixProductLocation = "product_location:"
	TypeProductLocation     = "productLocation"
)

// Includes the type field
type ESProductLocation struct {
	*qvspot.ProductLocation
	Type string `json:"type"`
}

// ProductLocationInsert inserts/replaces a productLocation
func (e *esearch) ProductLocationInsert(productLocation *qvspot.ProductLocation) error {

	request := elastic.NewBulkIndexRequest().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixProductLocation + productLocation.Id).
		Doc(&ESProductLocation{ProductLocation: productLocation, Type: TypeProductLocation})

	request.Source() // Convert to json so we can modify b
	e.bulk.Add(request)

	return nil

}

// ProductLocationGetById returns a productLocation by id
func (e *esearch) ProductLocationGetById(productLocationId string) (*qvspot.ProductLocation, error) {

	res, err := e.client.Get().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixProductLocation + productLocationId).
		Do(e.ctx)
	if elastic.IsNotFound(err) {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("Could not get productLocation: %v", err)
	}

	// Unmarshal the block
	productLocation := new(qvspot.ProductLocation)
	err = json.Unmarshal(res.Source, productLocation)
	return productLocation, err

}

// ProductLocationDeleteById removes a productLocation by id
func (e *esearch) ProductLocationDeleteById(productLocationId string) error {

	_, err := e.client.Delete().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixProductLocation + productLocationId).
		Refresh("true").
		Do(e.ctx)
	return err
}
