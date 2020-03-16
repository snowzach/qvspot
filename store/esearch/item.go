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
	IdPrefixItem = "item:"
	TypeItem     = "item"
)

// Includes the type field
type ESItem struct {
	*qvspot.Item
	Type string `json:"type"`
}

// ItemInsert inserts/replaces a item
func (e *esearch) ItemInsert(ctx context.Context, item *qvspot.Item) error {

	request := elastic.NewBulkIndexRequest().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixItem + item.Id).
		Doc(&ESItem{Item: item, Type: TypeItem})

	request.Source() // Convert to json so we can modify b
	e.bulk.Add(request)

	return nil

}

// ItemGetById returns a item by id
func (e *esearch) ItemGetById(ctx context.Context, itemId string) (*qvspot.Item, error) {

	res, err := e.client.Get().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixItem + itemId).
		Do(ctx)
	if elastic.IsNotFound(err) {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("Could not get item: %v", err)
	}

	// Unmarshal the block
	item := new(qvspot.Item)
	err = json.Unmarshal(res.Source, item)
	return item, err

}

// ItemDeleteById removes a item by id
func (e *esearch) ItemDeleteById(ctx context.Context, itemId string) error {

	_, err := e.client.Delete().
		Index(e.indexName(IndexAll, IndexVendor)).
		Id(IdPrefixItem + itemId).
		Refresh("true").
		Do(ctx)
	return err
}
