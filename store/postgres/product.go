package postgres

import (
	"context"
	"database/sql"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// ProductSave saves the product
func (c *Client) ProductSave(ctx context.Context, product *qvspot.Product) error {

	// Generate an ID if needed
	if product.Id == "" {
		product.Id = c.newID()
	}

	_, err := c.db.ExecContext(ctx, `
		INSERT INTO product (id, created, updated, vendor_id, name, description, pic_url, attr, attr_num)
		VALUES($1, NOW(), NOW(), $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE
		SET 
		updated = NOW(),
		vendor_id = $2,
		name = $3,
		description = $4,
		pic_url = $5,
		attr = $6,
		attr_num = $7
	`, product.Id, product.VendorId, product.Name, product.Description, product.PicUrl, product.Attr, product.AttrNum)
	if err != nil {
		return err
	}
	return nil

}

// ProductGetByID returns the the product by id
func (c *Client) ProductGetById(ctx context.Context, id string) (*qvspot.Product, error) {

	product := new(qvspot.Product)
	err := c.db.GetContext(ctx, product, `SELECT * FROM product WHERE id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return product, nil

}

// ProductDeleteById deletes a product by id
func (c *Client) ProductDeleteById(ctx context.Context, id string) error {

	_, err := c.db.ExecContext(ctx, `DELETE FROM product WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil

}
