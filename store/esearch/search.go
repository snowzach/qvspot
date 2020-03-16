package esearch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"

	"github.com/snowzach/qvspot/qvspot"
)

// VendorInsert inserts/replaces a vendor
func (e *esearch) ProductSearch(ctx context.Context, search *qvspot.ProductSearchRequest) (*qvspot.ProductSearchResponse, error) {

	query := elastic.NewBoolQuery().Filter(elastic.NewTermQuery("type", TypeItem))

	// Field search
	if search.Search != "" {
		query.Must(elastic.NewMultiMatchQuery(search.Search, "id", "product.id", "product.name", "product.desc"))
	}

	// Attributes
	for attr, values := range search.Attr {
		list := make([]interface{}, len(values.List), len(values.List))
		for i, value := range values.List {
			list[i] = value
		}
		query.Filter(elastic.NewTermsQuery(attr, list...))
	}

	// Numeric Attributes
	for attr, values := range search.AttrNum {
		query.Filter(elastic.NewRangeQuery(attr).Gte(values.Min).Lte(values.Max))
	}

	res, err := e.client.Search().
		Index(e.indexName(IndexAll, IndexVendor)).
		Query(query).
		From(int(search.Offset)).Size(int(search.Limit)).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not get search results: %v", err)
	}

	results := make([]*qvspot.ProductResult, len(res.Hits.Hits), len(res.Hits.Hits))

	for i, hit := range res.Hits.Hits {
		product := new(qvspot.Product)
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			return nil, fmt.Errorf("Could not parse Block: %s", err)
		}
		results[i] = &qvspot.ProductResult{
			Product: product,
		}
	}

	return &qvspot.ProductSearchResponse{
		Results: results,
	}, nil
}
