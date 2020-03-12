package qvspot

type ProductStore interface {
	ProductInsert(product *Product) error
	ProductDeleteById(productId string) error
}
