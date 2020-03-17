package postgres

import (
	"context"
	"database/sql"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// VendorSave saves the vendor
func (c *Client) VendorSave(ctx context.Context, vendor *qvspot.Vendor) error {

	// Generate an ID if needed
	if vendor.Id == "" {
		vendor.Id = c.newID()
	}

	_, err := c.db.ExecContext(ctx, `
		INSERT INTO vendor (id, created, updated, name, description)
		VALUES($1, NOW(), NOW(), $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET 
		updated = NOW(),
		name = $2,
		description = $3
	`, vendor.Id, vendor.Name, vendor.Description)
	if err != nil {
		return err
	}
	return nil

}

// VendorGetByID returns the the vendor by id
func (c *Client) VendorGetById(ctx context.Context, id string) (*qvspot.Vendor, error) {

	vendor := new(qvspot.Vendor)
	err := c.db.GetContext(ctx, vendor, `SELECT * FROM vendor WHERE id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return vendor, nil

}

// VendorDeleteById deletes a vendor by id
func (c *Client) VendorDeleteById(ctx context.Context, id string) error {

	_, err := c.db.ExecContext(ctx, `DELETE FROM vendor WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil

}
