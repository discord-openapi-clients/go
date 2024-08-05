/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateGuildEmojiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGuildEmojiRequest{}

// UpdateGuildEmojiRequest struct for UpdateGuildEmojiRequest
type UpdateGuildEmojiRequest struct {
	Name *string `json:"name,omitempty"`
	Roles []GetEntitlementsSkuIdsParameterOneOfInner `json:"roles,omitempty"`
}

// NewUpdateGuildEmojiRequest instantiates a new UpdateGuildEmojiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGuildEmojiRequest() *UpdateGuildEmojiRequest {
	this := UpdateGuildEmojiRequest{}
	return &this
}

// NewUpdateGuildEmojiRequestWithDefaults instantiates a new UpdateGuildEmojiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGuildEmojiRequestWithDefaults() *UpdateGuildEmojiRequest {
	this := UpdateGuildEmojiRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGuildEmojiRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildEmojiRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGuildEmojiRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGuildEmojiRequest) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildEmojiRequest) GetRoles() []GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil {
		var ret []GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildEmojiRequest) GetRolesOk() ([]GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateGuildEmojiRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the Roles field.
func (o *UpdateGuildEmojiRequest) SetRoles(v []GetEntitlementsSkuIdsParameterOneOfInner) {
	o.Roles = v
}

func (o UpdateGuildEmojiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGuildEmojiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableUpdateGuildEmojiRequest struct {
	value *UpdateGuildEmojiRequest
	isSet bool
}

func (v NullableUpdateGuildEmojiRequest) Get() *UpdateGuildEmojiRequest {
	return v.value
}

func (v *NullableUpdateGuildEmojiRequest) Set(val *UpdateGuildEmojiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildEmojiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildEmojiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildEmojiRequest(val *UpdateGuildEmojiRequest) *NullableUpdateGuildEmojiRequest {
	return &NullableUpdateGuildEmojiRequest{value: val, isSet: true}
}

func (v NullableUpdateGuildEmojiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildEmojiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

