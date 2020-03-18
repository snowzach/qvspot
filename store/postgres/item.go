package postgres

import (
	"context"
	"database/sql"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// ItemSave saves the item
func (c *Client) ItemSave(ctx context.Context, item *qvspot.Item) error {

	// Generate an ID if needed
	if item.Id == "" {
		item.Id = c.newID()
	}

	_, err := c.db.ExecContext(ctx, `
		INSERT INTO item (id, created, updated, vendor_id, product_id, location_id, stock, price, unit, start_time, end_time, attr, attr_num)
		VALUES($1, NOW(), NOW(), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (id) DO UPDATE
		SET 
		updated = NOW(),
		vendor_id = $2,
		product_id = $3,
		location_id = $4,
		stock = $5,
		price = $6,
		unit = $7,
		start_time = $8,
		end_time = $9,
		attr = $10,
		attr_num = $11
	`, item.Id, item.VendorId, item.ProductId, item.LocationId, item.Stock, item.Price, item.Unit, item.StartTime, item.EndTime, item.Attr, item.AttrNum)
	if err != nil {
		return err
	}
	return nil

}

// ItemGetByID returns the the item by id
func (c *Client) ItemGetById(ctx context.Context, id string) (*qvspot.Item, error) {

	item := new(qvspot.Item)
	err := c.db.GetContext(ctx, item, `SELECT * FROM item WHERE id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return item, nil

}

// ItemDeleteById deletes a item by id
func (c *Client) ItemDeleteById(ctx context.Context, id string) error {

	_, err := c.db.ExecContext(ctx, `DELETE FROM item WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil

}
