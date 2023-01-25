package api

import (
	"context"
	"github.com/google/uuid"
	"github.com/permitio/permit-golang/models"
	"github.com/permitio/permit-golang/openapi"
	"github.com/permitio/permit-golang/pkg/config"
	"github.com/permitio/permit-golang/pkg/errors"
	"go.uber.org/zap"
)

type Roles struct {
	permitBaseApi
}

func NewRolesApi(client *openapi.APIClient, config *config.PermitConfig) *Roles {
	return &Roles{
		permitBaseApi{
			client: client,
			config: config,
			logger: config.Logger,
		},
	}
}

type RoleRead struct {
	models.RoleRead
}

// List all roles in the current environment.
// Usage Example:
// `roles, err := PermitClient.Api.Roles.List(ctx,1, 10)`
func (r *Roles) List(ctx context.Context, page int, perPage int) ([]models.RoleRead, error) {
	perPageLimit := int32(DefaultPerPageLimit)
	if !isPaginationInLimit(int32(page), int32(perPage), perPageLimit) {
		err := errors.NewPermitPaginationError()
		r.logger.Error("error listing roles - max per page: "+string(perPageLimit), zap.Error(err))
		return nil, err
	}
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return nil, err
	}
	roles, httpRes, err := r.client.RolesApi.ListRoles(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment()).Page(int32(page)).PerPage(int32(perPage)).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error listing roles", zap.Error(err))
		return nil, err
	}
	return roles, nil
}

// Get a role by key.
// Usage Example:
// `role, err := PermitClient.Api.Roles.Get(ctx, "role-key")`
func (r *Roles) Get(ctx context.Context, roleKey string) (*models.RoleRead, error) {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return nil, err
	}
	role, httpRes, err := r.client.RolesApi.GetRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment(), roleKey).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error getting role: "+roleKey, zap.Error(err))
		return nil, err
	}

	return role, nil
}

func (r *Roles) GetByKey(ctx context.Context, roleKey string) (*models.RoleRead, error) {
	return r.Get(ctx, roleKey)
}

func (r *Roles) GetById(ctx context.Context, roleKey uuid.UUID) (*models.RoleRead, error) {
	return r.Get(ctx, roleKey.String())
}

func (r *Roles) Create(ctx context.Context, roleCreate models.RoleCreate) (*models.RoleRead, error) {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return nil, err
	}
	role, httpRes, err := r.client.RolesApi.CreateRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment()).RoleCreate(roleCreate).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error creating role: "+roleCreate.GetKey(), zap.Error(err))
		return nil, err
	}
	return role, nil
}

func (r *Roles) Update(ctx context.Context, roleKey string, roleUpdate models.RoleUpdate) (*models.RoleRead, error) {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return nil, err
	}
	role, httpRes, err := r.client.RolesApi.UpdateRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment(), roleKey).RoleUpdate(roleUpdate).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error updating role: "+roleKey, zap.Error(err))
		return nil, err
	}
	return role, nil
}

func (r *Roles) Delete(ctx context.Context, roleKey string) error {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return err
	}
	httpRes, err := r.client.RolesApi.DeleteRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment(), roleKey).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error deleting role: "+roleKey, zap.Error(err))
		return err
	}
	return nil
}

func (r *Roles) AssignPermissions(ctx context.Context, roleKey string, permissions []string) error {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return err
	}
	addRolePermissions := *models.NewAddRolePermissions(permissions)
	_, httpRes, err := r.client.RolesApi.AssignPermissionsToRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment(), roleKey).AddRolePermissions(addRolePermissions).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error assigning these permissions: '"+listToString(permissions)+"' to role: "+roleKey, zap.Error(err))
		return err
	}
	return nil
}

func (r *Roles) RemovePermissions(ctx context.Context, roleKey string, permissions []string) error {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return err
	}
	removeRolePermissions := *models.NewRemoveRolePermissions(permissions)
	_, httpRes, err := r.client.RolesApi.RemovePermissionsFromRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment(), roleKey).RemoveRolePermissions(removeRolePermissions).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error removing these permissions: '"+listToString(permissions)+"' from role: "+roleKey, zap.Error(err))
		return err
	}
	return nil
}

func (r *Roles) AddParentRole(ctx context.Context, roleKey string, parentRoleKey string) error {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return err
	}
	_, httpRes, err := r.client.RolesApi.AddParentRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment(), roleKey, parentRoleKey).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error adding parent role: '"+parentRoleKey+"' to role: "+roleKey, zap.Error(err))
		return err
	}
	return nil
}

func (r *Roles) RemoveParentRole(ctx context.Context, roleKey string, parentRoleKey string) error {
	err := r.lazyLoadContext(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		return err
	}
	_, httpRes, err := r.client.RolesApi.RemoveParentRole(ctx, r.config.Context.GetProject(), r.config.Context.GetEnvironment(), roleKey, parentRoleKey).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		r.logger.Error("error removing parent role: '"+parentRoleKey+"' from role: "+roleKey, zap.Error(err))
		return err
	}
	return nil
}
