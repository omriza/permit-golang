package permit

import (
	"context"
	"github.com/permitio/permit-golang/pkg/api"
	config "github.com/permitio/permit-golang/pkg/config"
	"github.com/permitio/permit-golang/pkg/enforcement"
	"github.com/permitio/permit-golang/pkg/models"
)

type Client struct {
	config      config.PermitConfig
	Api         *api.PermitApiClient
	Elements    *api.Elements
	enforcement *enforcement.PermitEnforcer
}

var New = NewPermit

func NewPermit(config config.PermitConfig) *Client {
	ctx := context.Background()
	apiClient := api.NewPermitApiClient(ctx, &config)
	enforcerClient := enforcement.NewPermitEnforcerClient(&config)
	return &Client{
		config:      config,
		Api:         apiClient,
		Elements:    apiClient.Elements,
		enforcement: enforcerClient,
	}
}

func (c *Client) SyncUser(ctx context.Context, user models.UserCreate) (*models.UserRead, error) {
	return c.Api.Users.SyncUser(ctx, user)
}
func (c *Client) Check(user enforcement.User, action enforcement.Action, resource enforcement.Resource) (bool, error) {
	return c.enforcement.Check(user, action, resource)
}

type PermitInterface interface {
	Check(user enforcement.User, action enforcement.Action, resource enforcement.Resource) (bool, error)
	SyncUser(ctx context.Context, user models.UserCreate) (*models.UserRead, error)
}
