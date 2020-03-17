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
		query.Filter(elastic.NewTermsQuery("product."+attr, list...))
	}

	// Numeric Attributes
	for attr, values := range search.AttrNum {
		query.Filter(elastic.NewRangeQuery("product." + attr).Gte(values.Min).Lte(values.Max))
	}

	res, err := e.client.Search().
		Index(e.indexName(IndexAll, IndexVendor)).
		Query(query).
		Aggregation("product_id", elastic.NewTermsAggregation().Field("product_id").SubAggregation("product", elastic.NewTopHitsAggregation().Size(1))).
		Size(0).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not get search results: %v", err)
	}

	results := make([]*qvspot.ProductResult, len(res.Hits.Hits), len(res.Hits.Hits))

	if products, ok := res.Aggregations.Terms("product_id"); ok {
		for _, productMatch := range products.Buckets {
			item := new(qvspot.Item)
			if hits, ok := productMatch.Aggregations.TopHits("product"); ok {
				if err = json.Unmarshal(hits.Hits.Hits[0].Source, &item); err != nil {
					return nil, fmt.Errorf("Could not parse hit: %s", err)
				}
				results = append(results, &qvspot.ProductResult{
					Product: item.Product,
					Matches: productMatch.DocCount,
				})
			}
		}
	}

	return &qvspot.ProductSearchResponse{
		Results: results,
	}, nil
}
