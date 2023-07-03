/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"time"
)

// checks if the RoleAssignmentDetailedRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleAssignmentDetailedRead{}

// RoleAssignmentDetailedRead struct for RoleAssignmentDetailedRead
type RoleAssignmentDetailedRead struct {
	// Unique id of the role assignment
	Id string `json:"id"`
	Role RoleAssignmentRole `json:"role"`
	User RoleAssignmentUser `json:"user"`
	Tenant RoleAssignmentTenant `json:"tenant"`
	// Unique id of the organization that the role assignment belongs to.
	OrganizationId string `json:"organization_id"`
	// Unique id of the project that the role assignment belongs to.
	ProjectId string `json:"project_id"`
	// Unique id of the environment that the role assignment belongs to.
	EnvironmentId string `json:"environment_id"`
	// Date and time when the role assignment was created (ISO_8601 format).
	CreatedAt time.Time `json:"created_at"`
}

// NewRoleAssignmentDetailedRead instantiates a new RoleAssignmentDetailedRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentDetailedRead(id string, role RoleAssignmentRole, user RoleAssignmentUser, tenant RoleAssignmentTenant, organizationId string, projectId string, environmentId string, createdAt time.Time) *RoleAssignmentDetailedRead {
	this := RoleAssignmentDetailedRead{}
	this.Id = id
	this.Role = role
	this.User = user
	this.Tenant = tenant
	this.OrganizationId = organizationId
	this.ProjectId = projectId
	this.EnvironmentId = environmentId
	this.CreatedAt = createdAt
	return &this
}

// NewRoleAssignmentDetailedReadWithDefaults instantiates a new RoleAssignmentDetailedRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentDetailedReadWithDefaults() *RoleAssignmentDetailedRead {
	this := RoleAssignmentDetailedRead{}
	return &this
}

// GetId returns the Id field value
func (o *RoleAssignmentDetailedRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleAssignmentDetailedRead) SetId(v string) {
	o.Id = v
}

// GetRole returns the Role field value
func (o *RoleAssignmentDetailedRead) GetRole() RoleAssignmentRole {
	if o == nil {
		var ret RoleAssignmentRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetRoleOk() (*RoleAssignmentRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RoleAssignmentDetailedRead) SetRole(v RoleAssignmentRole) {
	o.Role = v
}

// GetUser returns the User field value
func (o *RoleAssignmentDetailedRead) GetUser() RoleAssignmentUser {
	if o == nil {
		var ret RoleAssignmentUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetUserOk() (*RoleAssignmentUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RoleAssignmentDetailedRead) SetUser(v RoleAssignmentUser) {
	o.User = v
}

// GetTenant returns the Tenant field value
func (o *RoleAssignmentDetailedRead) GetTenant() RoleAssignmentTenant {
	if o == nil {
		var ret RoleAssignmentTenant
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetTenantOk() (*RoleAssignmentTenant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *RoleAssignmentDetailedRead) SetTenant(v RoleAssignmentTenant) {
	o.Tenant = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *RoleAssignmentDetailedRead) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *RoleAssignmentDetailedRead) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProjectId returns the ProjectId field value
func (o *RoleAssignmentDetailedRead) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RoleAssignmentDetailedRead) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *RoleAssignmentDetailedRead) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *RoleAssignmentDetailedRead) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RoleAssignmentDetailedRead) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RoleAssignmentDetailedRead) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RoleAssignmentDetailedRead) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o RoleAssignmentDetailedRead) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleAssignmentDetailedRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["role"] = o.Role
	toSerialize["user"] = o.User
	toSerialize["tenant"] = o.Tenant
	toSerialize["organization_id"] = o.OrganizationId
	toSerialize["project_id"] = o.ProjectId
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableRoleAssignmentDetailedRead struct {
	value *RoleAssignmentDetailedRead
	isSet bool
}

func (v NullableRoleAssignmentDetailedRead) Get() *RoleAssignmentDetailedRead {
	return v.value
}

func (v *NullableRoleAssignmentDetailedRead) Set(val *RoleAssignmentDetailedRead) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentDetailedRead) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentDetailedRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentDetailedRead(val *RoleAssignmentDetailedRead) *NullableRoleAssignmentDetailedRead {
	return &NullableRoleAssignmentDetailedRead{value: val, isSet: true}
}

func (v NullableRoleAssignmentDetailedRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentDetailedRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
