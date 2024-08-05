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

// checks if the UpdateGuildRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGuildRoleRequest{}

// UpdateGuildRoleRequest struct for UpdateGuildRoleRequest
type UpdateGuildRoleRequest struct {
	Name NullableString `json:"name,omitempty"`
	Permissions NullableInt32 `json:"permissions,omitempty"`
	Color NullableInt32 `json:"color,omitempty"`
	Hoist NullableBool `json:"hoist,omitempty"`
	Mentionable NullableBool `json:"mentionable,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	UnicodeEmoji NullableString `json:"unicode_emoji,omitempty"`
}

// NewUpdateGuildRoleRequest instantiates a new UpdateGuildRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGuildRoleRequest() *UpdateGuildRoleRequest {
	this := UpdateGuildRoleRequest{}
	return &this
}

// NewUpdateGuildRoleRequestWithDefaults instantiates a new UpdateGuildRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGuildRoleRequestWithDefaults() *UpdateGuildRoleRequest {
	this := UpdateGuildRoleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildRoleRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGuildRoleRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateGuildRoleRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateGuildRoleRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateGuildRoleRequest) UnsetName() {
	o.Name.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildRoleRequest) GetPermissions() int32 {
	if o == nil || IsNil(o.Permissions.Get()) {
		var ret int32
		return ret
	}
	return *o.Permissions.Get()
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildRoleRequest) GetPermissionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions.Get(), o.Permissions.IsSet()
}

// HasPermissions returns a boolean if a field has been set.
func (o *UpdateGuildRoleRequest) HasPermissions() bool {
	if o != nil && o.Permissions.IsSet() {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NullableInt32 and assigns it to the Permissions field.
func (o *UpdateGuildRoleRequest) SetPermissions(v int32) {
	o.Permissions.Set(&v)
}
// SetPermissionsNil sets the value for Permissions to be an explicit nil
func (o *UpdateGuildRoleRequest) SetPermissionsNil() {
	o.Permissions.Set(nil)
}

// UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
func (o *UpdateGuildRoleRequest) UnsetPermissions() {
	o.Permissions.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildRoleRequest) GetColor() int32 {
	if o == nil || IsNil(o.Color.Get()) {
		var ret int32
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildRoleRequest) GetColorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *UpdateGuildRoleRequest) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableInt32 and assigns it to the Color field.
func (o *UpdateGuildRoleRequest) SetColor(v int32) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *UpdateGuildRoleRequest) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *UpdateGuildRoleRequest) UnsetColor() {
	o.Color.Unset()
}

// GetHoist returns the Hoist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildRoleRequest) GetHoist() bool {
	if o == nil || IsNil(o.Hoist.Get()) {
		var ret bool
		return ret
	}
	return *o.Hoist.Get()
}

// GetHoistOk returns a tuple with the Hoist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildRoleRequest) GetHoistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hoist.Get(), o.Hoist.IsSet()
}

// HasHoist returns a boolean if a field has been set.
func (o *UpdateGuildRoleRequest) HasHoist() bool {
	if o != nil && o.Hoist.IsSet() {
		return true
	}

	return false
}

// SetHoist gets a reference to the given NullableBool and assigns it to the Hoist field.
func (o *UpdateGuildRoleRequest) SetHoist(v bool) {
	o.Hoist.Set(&v)
}
// SetHoistNil sets the value for Hoist to be an explicit nil
func (o *UpdateGuildRoleRequest) SetHoistNil() {
	o.Hoist.Set(nil)
}

// UnsetHoist ensures that no value is present for Hoist, not even an explicit nil
func (o *UpdateGuildRoleRequest) UnsetHoist() {
	o.Hoist.Unset()
}

// GetMentionable returns the Mentionable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildRoleRequest) GetMentionable() bool {
	if o == nil || IsNil(o.Mentionable.Get()) {
		var ret bool
		return ret
	}
	return *o.Mentionable.Get()
}

// GetMentionableOk returns a tuple with the Mentionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildRoleRequest) GetMentionableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mentionable.Get(), o.Mentionable.IsSet()
}

// HasMentionable returns a boolean if a field has been set.
func (o *UpdateGuildRoleRequest) HasMentionable() bool {
	if o != nil && o.Mentionable.IsSet() {
		return true
	}

	return false
}

// SetMentionable gets a reference to the given NullableBool and assigns it to the Mentionable field.
func (o *UpdateGuildRoleRequest) SetMentionable(v bool) {
	o.Mentionable.Set(&v)
}
// SetMentionableNil sets the value for Mentionable to be an explicit nil
func (o *UpdateGuildRoleRequest) SetMentionableNil() {
	o.Mentionable.Set(nil)
}

// UnsetMentionable ensures that no value is present for Mentionable, not even an explicit nil
func (o *UpdateGuildRoleRequest) UnsetMentionable() {
	o.Mentionable.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildRoleRequest) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildRoleRequest) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *UpdateGuildRoleRequest) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *UpdateGuildRoleRequest) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *UpdateGuildRoleRequest) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *UpdateGuildRoleRequest) UnsetIcon() {
	o.Icon.Unset()
}

// GetUnicodeEmoji returns the UnicodeEmoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildRoleRequest) GetUnicodeEmoji() string {
	if o == nil || IsNil(o.UnicodeEmoji.Get()) {
		var ret string
		return ret
	}
	return *o.UnicodeEmoji.Get()
}

// GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildRoleRequest) GetUnicodeEmojiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnicodeEmoji.Get(), o.UnicodeEmoji.IsSet()
}

// HasUnicodeEmoji returns a boolean if a field has been set.
func (o *UpdateGuildRoleRequest) HasUnicodeEmoji() bool {
	if o != nil && o.UnicodeEmoji.IsSet() {
		return true
	}

	return false
}

// SetUnicodeEmoji gets a reference to the given NullableString and assigns it to the UnicodeEmoji field.
func (o *UpdateGuildRoleRequest) SetUnicodeEmoji(v string) {
	o.UnicodeEmoji.Set(&v)
}
// SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil
func (o *UpdateGuildRoleRequest) SetUnicodeEmojiNil() {
	o.UnicodeEmoji.Set(nil)
}

// UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil
func (o *UpdateGuildRoleRequest) UnsetUnicodeEmoji() {
	o.UnicodeEmoji.Unset()
}

func (o UpdateGuildRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGuildRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Permissions.IsSet() {
		toSerialize["permissions"] = o.Permissions.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Hoist.IsSet() {
		toSerialize["hoist"] = o.Hoist.Get()
	}
	if o.Mentionable.IsSet() {
		toSerialize["mentionable"] = o.Mentionable.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.UnicodeEmoji.IsSet() {
		toSerialize["unicode_emoji"] = o.UnicodeEmoji.Get()
	}
	return toSerialize, nil
}

type NullableUpdateGuildRoleRequest struct {
	value *UpdateGuildRoleRequest
	isSet bool
}

func (v NullableUpdateGuildRoleRequest) Get() *UpdateGuildRoleRequest {
	return v.value
}

func (v *NullableUpdateGuildRoleRequest) Set(val *UpdateGuildRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildRoleRequest(val *UpdateGuildRoleRequest) *NullableUpdateGuildRoleRequest {
	return &NullableUpdateGuildRoleRequest{value: val, isSet: true}
}

func (v NullableUpdateGuildRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


