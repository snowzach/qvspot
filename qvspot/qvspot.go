package qvspot

type QVStore interface {
	VendorInsert(vendor *Vendor) error
	VendorGetById(vendorId string) (*Vendor, error)
	VendorDeleteById(vendorId string) error

	VendorLocationInsert(vendorLocation *VendorLocation) error
	VendorLocationGetById(vendorLocationId string) (*VendorLocation, error)
	VendorLocationDeleteById(vendorLocationId string) error

	ProductLocationInsert(productLocation *ProductLocation) error
	ProductLocationGetById(productLocationId string) (*ProductLocation, error)
	ProductLocationDeleteById(productLocationId string) error

	ProductInsert(product *Product) error
	ProductGetById(productId string) (*Product, error)
	ProductDeleteById(productId string) error
}
