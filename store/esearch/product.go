package esearch

import (
	"github.com/olivere/elastic/v7"

	"github.com/snowzach/qvspot/qvspot"
)

// ProductInsert replaces a product
func (e *esearch) ProductInsert(product *qvspot.Product) error {

	request := elastic.NewBulkIndexRequest().
		Index(e.indexName(IndexTypeProduct)).
		Id(product.Id).
		Doc(product)

	request.Source() // Convert to json so we can modify b
	e.bulk.Add(request)

	return nil

}

// ProductDeleteById removes a product by ProductId
func (e *esearch) ProductDeleteById(productId string) error {

	_, err := e.client.Delete().
		Index(e.indexName(IndexTypeProduct)).
		Id(productId).
		Refresh("true").
		Do(e.ctx)
	return err
}
