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

type Users struct {
	permitBaseApi
}

func NewUsersApi(client *openapi.APIClient, config *config.PermitConfig) *Users {
	return &Users{
		permitBaseApi{
			client: client,
			config: config,
			logger: config.Logger,
		},
	}
}

// List the users from your context's environment.
// Usage Example:
//  `users, err := PermitClient.Api.Users.List(ctx, 1, 10)`
func (u *Users) List(ctx context.Context, page int, perPage int) ([]models.UserRead, error) {
	perPageLimit := int32(DefaultPerPageLimit)
	if !isPaginationInLimit(int32(page), int32(perPage), perPageLimit) {
		err := errors.NewPermitPaginationError()
		u.logger.Error("error listing users - max per page: "+string(perPageLimit), zap.Error(err))
		return nil, err
	}
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	users, httpRes, err := u.client.UsersApi.ListUsers(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment()).Page(int32(page)).PerPage(int32(perPage)).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error listing users", zap.Error(err))
		return nil, err
	}
	return users.GetData(), nil
}

// Get a user from your context's environment.
// Usage Example:
//  `user, err := PermitClient.Api.Users.Get(ctx, "user-key")`
func (u *Users) Get(ctx context.Context, userKey string) (*models.UserRead, error) {
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	user, httpRes, err := u.client.UsersApi.GetUser(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment(), userKey).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error getting user: "+userKey, zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *Users) GetByKey(ctx context.Context, userKey string) (*models.UserRead, error) {
	return u.Get(ctx, userKey)
}

func (u *Users) GetById(ctx context.Context, userId uuid.UUID) (*models.UserRead, error) {
	return u.Get(ctx, userId.String())
}

func (u *Users) Create(ctx context.Context, userCreate models.UserCreate) (*models.UserRead, error) {
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	user, httpRes, err := u.client.UsersApi.CreateUser(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment()).UserCreate(userCreate).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error creating user:"+userCreate.GetKey(), zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *Users) Update(ctx context.Context, userKey string, userUpdate models.UserUpdate) (*models.UserRead, error) {
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	user, httpRes, err := u.client.UsersApi.UpdateUser(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment(), userKey).UserUpdate(userUpdate).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error updating user:"+userKey, zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *Users) Delete(ctx context.Context, userKey string) error {
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return err
	}
	httpRes, err := u.client.UsersApi.DeleteUser(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment(), userKey).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error deleting user:"+userKey, zap.Error(err))
		return err
	}
	return nil
}

func (u *Users) AssignRole(ctx context.Context, userKey string, roleKey string, tenantKey string) (*models.RoleAssignmentRead, error) {
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	userRoleCreate := *models.NewUserRoleCreate(roleKey, tenantKey)
	roleAssignmentRead, httpRes, err := u.client.UsersApi.AssignRoleToUser(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment(), userKey).UserRoleCreate(userRoleCreate).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error assigning role:"+roleKey+" to user:"+userKey, zap.Error(err))
		return nil, err
	}
	return roleAssignmentRead, nil
}

func (u *Users) UnassignRole(ctx context.Context, userKey string, roleKey string, tenantKey string) (*models.UserRead, error) {
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	UserRoleRemove := *models.NewUserRoleRemove(roleKey, tenantKey)
	user, httpRes, err := u.client.UsersApi.UnassignRoleFromUser(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment(), userKey).UserRoleRemove(UserRoleRemove).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error unassigning role:"+roleKey+" from user:"+userKey, zap.Error(err))
		return nil, err
	}
	return user, nil
}
func (u *Users) GetAssignedRoles(ctx context.Context, userKey string, tenantKey string, page int, perPage int) ([]models.RoleAssignmentRead, error) {
	perPageLimit := int32(DefaultPerPageLimit)
	if !isPaginationInLimit(int32(page), int32(perPage), perPageLimit) {
		err := errors.NewPermitPaginationError()
		u.logger.Error("error listing users - max per page: "+string(perPageLimit), zap.Error(err))
		return nil, err
	}
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	roleAssignments, httpRes, err := u.client.RoleAssignmentsApi.ListRoleAssignments(ctx, u.config.Context.GetProject(), u.config.Context.GetEnvironment()).
		User(userKey).
		Tenant(tenantKey).
		Page(int32(page)).PerPage(int32(perPage)).Execute()
	err = errors.HttpErrorHandle(err, httpRes)
	if err != nil {
		u.logger.Error("error listing roles for user:"+userKey, zap.Error(err))
		return nil, err
	}
	return roleAssignments, nil
}

func (u *Users) SyncUser(ctx context.Context, user models.UserCreate) (*models.UserRead, error) {
	err := u.lazyLoadContext(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	existUser, err := u.Get(ctx, user.GetKey())
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	if existUser != nil {
		u.logger.Info("User already exists, updating it...", zap.String("user", user.GetKey()))
		userUpdate := models.NewUserUpdate()
		if email := user.GetEmail(); email != "" {
			userUpdate.SetEmail(user.GetEmail())
		}
		if firstName := user.GetFirstName(); firstName != "" {
			userUpdate.SetFirstName(user.GetFirstName())
		}
		if lastName := user.GetLastName(); lastName != "" {
			userUpdate.SetLastName(user.GetLastName())
		}
		userUpdate.SetAttributes(user.GetAttributes())
		userRead, err := u.Update(ctx, user.GetKey(), *userUpdate)
		if err != nil {
			u.logger.Error("error updating user: "+user.GetKey(), zap.Error(err))
			return nil, err
		}
		return userRead, nil
	}
	u.logger.Info("User does not exist, creating it...", zap.String("user", user.GetKey()))
	userRead, err := u.Create(ctx, user)
	if err != nil {
		u.logger.Error("error creating user: "+user.GetKey(), zap.Error(err))
		return nil, err
	}
	return userRead, err
}
