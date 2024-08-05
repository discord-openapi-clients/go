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

// checks if the UpdateGuildWidgetSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGuildWidgetSettingsRequest{}

// UpdateGuildWidgetSettingsRequest struct for UpdateGuildWidgetSettingsRequest
type UpdateGuildWidgetSettingsRequest struct {
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Enabled NullableBool `json:"enabled,omitempty"`
}

// NewUpdateGuildWidgetSettingsRequest instantiates a new UpdateGuildWidgetSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGuildWidgetSettingsRequest() *UpdateGuildWidgetSettingsRequest {
	this := UpdateGuildWidgetSettingsRequest{}
	return &this
}

// NewUpdateGuildWidgetSettingsRequestWithDefaults instantiates a new UpdateGuildWidgetSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGuildWidgetSettingsRequestWithDefaults() *UpdateGuildWidgetSettingsRequest {
	this := UpdateGuildWidgetSettingsRequest{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *UpdateGuildWidgetSettingsRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGuildWidgetSettingsRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *UpdateGuildWidgetSettingsRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *UpdateGuildWidgetSettingsRequest) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateGuildWidgetSettingsRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateGuildWidgetSettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateGuildWidgetSettingsRequest) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *UpdateGuildWidgetSettingsRequest) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *UpdateGuildWidgetSettingsRequest) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *UpdateGuildWidgetSettingsRequest) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o UpdateGuildWidgetSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGuildWidgetSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableUpdateGuildWidgetSettingsRequest struct {
	value *UpdateGuildWidgetSettingsRequest
	isSet bool
}

func (v NullableUpdateGuildWidgetSettingsRequest) Get() *UpdateGuildWidgetSettingsRequest {
	return v.value
}

func (v *NullableUpdateGuildWidgetSettingsRequest) Set(val *UpdateGuildWidgetSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGuildWidgetSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGuildWidgetSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGuildWidgetSettingsRequest(val *UpdateGuildWidgetSettingsRequest) *NullableUpdateGuildWidgetSettingsRequest {
	return &NullableUpdateGuildWidgetSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateGuildWidgetSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGuildWidgetSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


