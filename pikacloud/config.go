package pikacloud

import (
	"log"

	"github.com/pikacloud/gopikacloud"
)

// Config for Pikacloud client
type Config struct {
	Token string
}

// Client returns a new client for accessing Pikacloud
func (c *Config) Client() (*gopikacloud.Client, error) {
	client := gopikacloud.NewClient(c.Token)
	log.Printf("[INFO] Pikacloud Client configured")
	return client, nil
}
