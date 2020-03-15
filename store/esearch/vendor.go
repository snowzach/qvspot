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
	IdPrefixVendor = "vendor:"
	TypeVendor     = "vendor"
)

// Includes the type field
type ESVendor struct {
	*qvspot.Vendor
	Type string `json:"type"`
}

// VendorInsert inserts/replaces a vendor
func (e *esearch) VendorInsert(ctx context.Context, vendor *qvspot.Vendor) error {

	request := elastic.NewBulkIndexRequest().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixVendor + vendor.Id).
		Doc(&ESVendor{Vendor: vendor, Type: TypeVendor})

	request.Source() // Convert to json so we can modify b
	e.bulk.Add(request)

	return nil

}

// VendorGetById returns a vendor by id
func (e *esearch) VendorGetById(ctx context.Context, vendorId string) (*qvspot.Vendor, error) {

	res, err := e.client.Get().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixVendor + vendorId).
		Do(ctx)
	if elastic.IsNotFound(err) {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("Could not get vendor: %v", err)
	}

	// Unmarshal the block
	vendor := new(qvspot.Vendor)
	err = json.Unmarshal(res.Source, vendor)
	return vendor, err

}

// VendorDeleteById removes a vendor by id
func (e *esearch) VendorDeleteById(ctx context.Context, vendorId string) error {

	_, err := e.client.Delete().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixVendor + vendorId).
		Refresh("true").
		Do(ctx)
	return err
}
