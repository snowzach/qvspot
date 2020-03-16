package qvspot

import (
	"context"
)

type QVStore interface {
	VendorInsert(ctx context.Context, vendor *Vendor) error
	VendorGetById(ctx context.Context, vendorId string) (*Vendor, error)
	VendorDeleteById(ctx context.Context, vendorId string) error

	LocationInsert(ctx context.Context, location *Location) error
	LocationGetById(ctx context.Context, locationId string) (*Location, error)
	LocationDeleteById(ctx context.Context, locationId string) error

	ItemInsert(ctx context.Context, item *Item) error
	ItemGetById(ctx context.Context, itemId string) (*Item, error)
	ItemDeleteById(ctx context.Context, itemId string) error

	ProductInsert(ctx context.Context, product *Product) error
	ProductGetById(ctx context.Context, productId string) (*Product, error)
	ProductDeleteById(ctx context.Context, productId string) error

	ProductSearch(ctx context.Context, search *ProductSearchRequest) (*ProductSearchResponse, error)
}
