package api

import (
	"context"
	"encoding/json"
)

// Favorite is a Favorite
type Favorite struct {
	ID           string `json:"id,omitempty"`
	Created      *Time  `json:"created,omitempty"`
	ForeignKey   string `json:"foreign_key,omitempty"`
	ForeignModel string `json:"foreign_model,omitempty"`
	Modified     *Time  `json:"modified,omitempty"`
}

// CreateFavorite Creates a new Passbolt Favorite for the given Resource ID
func (c *Client) CreateFavorite(ctx context.Context, resourceID string) (*Favorite, error) {
	msg, err := c.DoCustomRequest(ctx, "POST", "/favorites/resource/"+resourceID+".json", "v2", nil, nil)
	if err != nil {
		return nil, err
	}

	var favorite Favorite
	err = json.Unmarshal(msg.Body, &favorite)
	if err != nil {
		return nil, err
	}
	return &favorite, nil
}

// DeleteFavorite Deletes a Passbolt Favorite
func (c *Client) DeleteFavorite(ctx context.Context, favoriteID string) error {
	_, err := c.DoCustomRequest(ctx, "DELETE", "/favorites/"+favoriteID+".json", "v2", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
