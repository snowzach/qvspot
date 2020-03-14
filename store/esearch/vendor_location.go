package esearch

import (
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

const (
	IdPrefixVendorLocation = "vendor_location:"
	TypeVendorLocation     = "vendorLocation"
)

// Includes the type field
type ESVendorLocation struct {
	*qvspot.VendorLocation
	Type string `json:"type"`
}

// VendorLocationInsert inserts/replaces a vendorLocation
func (e *esearch) VendorLocationInsert(vendorLocation *qvspot.VendorLocation) error {

	request := elastic.NewBulkIndexRequest().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixVendorLocation + vendorLocation.Id).
		Doc(&ESVendorLocation{VendorLocation: vendorLocation, Type: TypeVendorLocation})

	request.Source() // Convert to json so we can modify b
	e.bulk.Add(request)

	return nil

}

// VendorLocationGetById returns a vendorLocation by id
func (e *esearch) VendorLocationGetById(vendorLocationId string) (*qvspot.VendorLocation, error) {

	res, err := e.client.Get().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixVendorLocation + vendorLocationId).
		Do(e.ctx)
	if elastic.IsNotFound(err) {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("Could not get vendorLocation: %v", err)
	}

	// Unmarshal the block
	vendorLocation := new(qvspot.VendorLocation)
	err = json.Unmarshal(res.Source, vendorLocation)
	return vendorLocation, err

}

// VendorLocationDeleteById removes a vendorLocation by id
func (e *esearch) VendorLocationDeleteById(vendorLocationId string) error {

	_, err := e.client.Delete().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixVendorLocation + vendorLocationId).
		Refresh("true").
		Do(e.ctx)
	return err
}
