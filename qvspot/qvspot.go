package qvspot

type QVStore interface {
	VendorInsert(vendor *Vendor) error
	VendorGetById(vendorId string) (*Vendor, error)
	VendorDeleteById(vendorId string) error

	ProductInsert(product *Product) error
	ProductGetById(productId string) (*Product, error)
	ProductDeleteById(productId string) error
}
