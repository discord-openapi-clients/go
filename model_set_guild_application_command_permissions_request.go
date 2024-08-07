/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
)

// checks if the SetGuildApplicationCommandPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetGuildApplicationCommandPermissionsRequest{}

// SetGuildApplicationCommandPermissionsRequest struct for SetGuildApplicationCommandPermissionsRequest
type SetGuildApplicationCommandPermissionsRequest struct {
	Permissions []ApplicationCommandPermission `json:"permissions,omitempty"`
}

// NewSetGuildApplicationCommandPermissionsRequest instantiates a new SetGuildApplicationCommandPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetGuildApplicationCommandPermissionsRequest() *SetGuildApplicationCommandPermissionsRequest {
	this := SetGuildApplicationCommandPermissionsRequest{}
	return &this
}

// NewSetGuildApplicationCommandPermissionsRequestWithDefaults instantiates a new SetGuildApplicationCommandPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetGuildApplicationCommandPermissionsRequestWithDefaults() *SetGuildApplicationCommandPermissionsRequest {
	this := SetGuildApplicationCommandPermissionsRequest{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetGuildApplicationCommandPermissionsRequest) GetPermissions() []ApplicationCommandPermission {
	if o == nil {
		var ret []ApplicationCommandPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetGuildApplicationCommandPermissionsRequest) GetPermissionsOk() ([]ApplicationCommandPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *SetGuildApplicationCommandPermissionsRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []ApplicationCommandPermission and assigns it to the Permissions field.
func (o *SetGuildApplicationCommandPermissionsRequest) SetPermissions(v []ApplicationCommandPermission) {
	o.Permissions = v
}

func (o SetGuildApplicationCommandPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetGuildApplicationCommandPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableSetGuildApplicationCommandPermissionsRequest struct {
	value *SetGuildApplicationCommandPermissionsRequest
	isSet bool
}

func (v NullableSetGuildApplicationCommandPermissionsRequest) Get() *SetGuildApplicationCommandPermissionsRequest {
	return v.value
}

func (v *NullableSetGuildApplicationCommandPermissionsRequest) Set(val *SetGuildApplicationCommandPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetGuildApplicationCommandPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetGuildApplicationCommandPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetGuildApplicationCommandPermissionsRequest(val *SetGuildApplicationCommandPermissionsRequest) *NullableSetGuildApplicationCommandPermissionsRequest {
	return &NullableSetGuildApplicationCommandPermissionsRequest{value: val, isSet: true}
}

func (v NullableSetGuildApplicationCommandPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetGuildApplicationCommandPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


