package postgres

import (
	"context"
	"database/sql"

	"github.com/snowzach/qvspot/qvspot"
	"github.com/snowzach/qvspot/store"
)

// LocationSave saves the location
func (c *Client) LocationSave(ctx context.Context, location *qvspot.Location) error {

	// Generate an ID if needed
	if location.Id == "" {
		location.Id = c.newID()
	}

	_, err := c.db.ExecContext(ctx, `
		INSERT INTO location (id, created, updated, name, description, position)
		VALUES($1, NOW(), NOW(), $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
		SET 
		updated = NOW(),
		name = $2,
		description = $3,
		position = $4
	`, location.Id, location.Name, location.Description, location.Position)
	if err != nil {
		return err
	}
	return nil

}

// LocationGetByID returns the the location by id
func (c *Client) LocationGetById(ctx context.Context, id string) (*qvspot.Location, error) {

	location := new(qvspot.Location)
	err := c.db.GetContext(ctx, location, `SELECT * FROM location WHERE id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return location, nil

}

// LocationDeleteById deletes a location by id
func (c *Client) LocationDeleteById(ctx context.Context, id string) error {

	_, err := c.db.ExecContext(ctx, `DELETE FROM location WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil

}
