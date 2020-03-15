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
	IdPrefixProduct = "product:"
	TypeProduct     = "product"
)

// Includes the type field
type ESProduct struct {
	*qvspot.Product
	Type string `json:"type"`
}

// ProductInsert inserts/replaces a product
func (e *esearch) ProductInsert(ctx context.Context, product *qvspot.Product) error {

	request := elastic.NewBulkIndexRequest().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixProduct + product.Id).
		Doc(&ESProduct{Product: product, Type: TypeProduct})

	request.Source() // Convert to json so we can modify b
	e.bulk.Add(request)

	return nil

}

// ProductGetById returns a product by id
func (e *esearch) ProductGetById(ctx context.Context, productId string) (*qvspot.Product, error) {

	res, err := e.client.Get().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixProduct + productId).
		Do(ctx)
	if elastic.IsNotFound(err) {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("Could not get product: %v", err)
	}

	// Unmarshal the block
	product := new(qvspot.Product)
	err = json.Unmarshal(res.Source, product)
	return product, err

}

// ProductDeleteById removes a product by id
func (e *esearch) ProductDeleteById(ctx context.Context, productId string) error {

	_, err := e.client.Delete().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixProduct + productId).
		Refresh("true").
		Do(ctx)
	return err
}
