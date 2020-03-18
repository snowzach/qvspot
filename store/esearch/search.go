package esearch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"

	"github.com/snowzach/qvspot/qvspot"
)

// Search searches for products or items
func (e *esearch) Search(ctx context.Context, search *qvspot.SearchRequest) (*qvspot.SearchResponse, error) {

	query := elastic.NewBoolQuery().Filter(elastic.NewTermQuery("type", TypeItem))

	if search.Position != nil {
		query.Filter(elastic.NewGeoDistanceQuery("location.position").Point(search.Position.Lat, search.Position.Lon).Distance(search.Distance))
	}

	// Field search
	if search.Search != "" {
		query.Must(elastic.NewMultiMatchQuery(search.Search, "id", "product.id", "product.name", "product.desc"))
	}

	// Item Attributes
	for attr, values := range search.Attr {
		list := make([]interface{}, len(values.List), len(values.List))
		for i, value := range values.List {
			list[i] = value
		}
		query.Filter(elastic.NewTermsQuery(attr, list...))
	}

	// Item Numeric Attributes
	for attr, values := range search.AttrNum {
		query.Filter(elastic.NewRangeQuery(attr).Gte(values.Min).Lte(values.Max))
	}

	if search.Product != nil {

		if search.Product.Id != "" {
			query.Filter(elastic.NewTermQuery("product.id", search.Product.Id))
		}

		// Product Attributes
		for attr, values := range search.Product.Attr {
			list := make([]interface{}, len(values.List), len(values.List))
			for i, value := range values.List {
				list[i] = value
			}
			query.Filter(elastic.NewTermsQuery("product."+attr, list...))
		}

		// Product Numeric Attributes
		for attr, values := range search.Product.AttrNum {
			query.Filter(elastic.NewRangeQuery("product." + attr).Gte(values.Min).Lte(values.Max))
		}

	}

	if search.ByProduct {

		res, err := e.client.Search().
			Index(e.indexName(IndexAll, IndexVendor)).
			Query(query).
			Aggregation("product_id", elastic.NewTermsAggregation().Field("product_id").SubAggregation("product", elastic.NewTopHitsAggregation().From(int(search.Offset)).Size(int(search.Limit)))).
			Size(0).
			Do(ctx)
		if err != nil {
			return nil, fmt.Errorf("Could not get search results: %v", err)
		}

		results := make([]*qvspot.SearchResult, 0)

		if products, ok := res.Aggregations.Terms("product_id"); ok {
			for _, productMatch := range products.Buckets {
				item := new(qvspot.Item)
				if hits, ok := productMatch.Aggregations.TopHits("product"); ok {
					if err = json.Unmarshal(hits.Hits.Hits[0].Source, &item); err != nil {
						return nil, fmt.Errorf("Could not parse hit: %s", err)
					}
					results = append(results, &qvspot.SearchResult{
						Product:   item.Product,
						ItemCount: productMatch.DocCount,
					})
				}
			}
		}

		return &qvspot.SearchResponse{
			Results: results,
		}, nil

	}

	res, err := e.client.Search().
		Index(e.indexName(IndexAll, IndexVendor)).
		Query(query).
		Size(int(search.Limit)).From(int(search.Offset)).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not get search results: %v", err)
	}

	results := make([]*qvspot.SearchResult, len(res.Hits.Hits), len(res.Hits.Hits))

	for i, hit := range res.Hits.Hits {
		item := new(qvspot.Item)
		err := json.Unmarshal(hit.Source, &item)
		if err != nil {
			return nil, fmt.Errorf("Could not parse item: %s", err)
		}
		results[i] = &qvspot.SearchResult{
			Item: item,
		}
	}

	return &qvspot.SearchResponse{
		Results: results,
	}, nil

}
