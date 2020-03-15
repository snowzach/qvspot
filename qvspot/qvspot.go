package qvspot

import (
	"context"
)

type QVStore interface {
	VendorInsert(ctx context.Context, vendor *Vendor) error
	VendorGetById(ctx context.Context, vendorId string) (*Vendor, error)
	VendorDeleteById(ctx context.Context, vendorId string) error

	VendorLocationInsert(ctx context.Context, vendorLocation *VendorLocation) error
	VendorLocationGetById(ctx context.Context, vendorLocationId string) (*VendorLocation, error)
	VendorLocationDeleteById(ctx context.Context, vendorLocationId string) error

	ProductLocationInsert(ctx context.Context, productLocation *ProductLocation) error
	ProductLocationGetById(ctx context.Context, productLocationId string) (*ProductLocation, error)
	ProductLocationDeleteById(ctx context.Context, productLocationId string) error

	ProductInsert(ctx context.Context, product *Product) error
	ProductGetById(ctx context.Context, productId string) (*Product, error)
	ProductDeleteById(ctx context.Context, productId string) error

	ProductSearch(ctx context.Context, search *ProductSearchRequest) (*ProductSearchResponse, error)
}
