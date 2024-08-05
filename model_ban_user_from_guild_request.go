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

// checks if the BanUserFromGuildRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BanUserFromGuildRequest{}

// BanUserFromGuildRequest struct for BanUserFromGuildRequest
type BanUserFromGuildRequest struct {
	DeleteMessageSeconds NullableInt32 `json:"delete_message_seconds,omitempty"`
	DeleteMessageDays NullableInt32 `json:"delete_message_days,omitempty"`
}

// NewBanUserFromGuildRequest instantiates a new BanUserFromGuildRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanUserFromGuildRequest() *BanUserFromGuildRequest {
	this := BanUserFromGuildRequest{}
	return &this
}

// NewBanUserFromGuildRequestWithDefaults instantiates a new BanUserFromGuildRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanUserFromGuildRequestWithDefaults() *BanUserFromGuildRequest {
	this := BanUserFromGuildRequest{}
	return &this
}

// GetDeleteMessageSeconds returns the DeleteMessageSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BanUserFromGuildRequest) GetDeleteMessageSeconds() int32 {
	if o == nil || IsNil(o.DeleteMessageSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.DeleteMessageSeconds.Get()
}

// GetDeleteMessageSecondsOk returns a tuple with the DeleteMessageSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BanUserFromGuildRequest) GetDeleteMessageSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteMessageSeconds.Get(), o.DeleteMessageSeconds.IsSet()
}

// HasDeleteMessageSeconds returns a boolean if a field has been set.
func (o *BanUserFromGuildRequest) HasDeleteMessageSeconds() bool {
	if o != nil && o.DeleteMessageSeconds.IsSet() {
		return true
	}

	return false
}

// SetDeleteMessageSeconds gets a reference to the given NullableInt32 and assigns it to the DeleteMessageSeconds field.
func (o *BanUserFromGuildRequest) SetDeleteMessageSeconds(v int32) {
	o.DeleteMessageSeconds.Set(&v)
}
// SetDeleteMessageSecondsNil sets the value for DeleteMessageSeconds to be an explicit nil
func (o *BanUserFromGuildRequest) SetDeleteMessageSecondsNil() {
	o.DeleteMessageSeconds.Set(nil)
}

// UnsetDeleteMessageSeconds ensures that no value is present for DeleteMessageSeconds, not even an explicit nil
func (o *BanUserFromGuildRequest) UnsetDeleteMessageSeconds() {
	o.DeleteMessageSeconds.Unset()
}

// GetDeleteMessageDays returns the DeleteMessageDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BanUserFromGuildRequest) GetDeleteMessageDays() int32 {
	if o == nil || IsNil(o.DeleteMessageDays.Get()) {
		var ret int32
		return ret
	}
	return *o.DeleteMessageDays.Get()
}

// GetDeleteMessageDaysOk returns a tuple with the DeleteMessageDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BanUserFromGuildRequest) GetDeleteMessageDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteMessageDays.Get(), o.DeleteMessageDays.IsSet()
}

// HasDeleteMessageDays returns a boolean if a field has been set.
func (o *BanUserFromGuildRequest) HasDeleteMessageDays() bool {
	if o != nil && o.DeleteMessageDays.IsSet() {
		return true
	}

	return false
}

// SetDeleteMessageDays gets a reference to the given NullableInt32 and assigns it to the DeleteMessageDays field.
func (o *BanUserFromGuildRequest) SetDeleteMessageDays(v int32) {
	o.DeleteMessageDays.Set(&v)
}
// SetDeleteMessageDaysNil sets the value for DeleteMessageDays to be an explicit nil
func (o *BanUserFromGuildRequest) SetDeleteMessageDaysNil() {
	o.DeleteMessageDays.Set(nil)
}

// UnsetDeleteMessageDays ensures that no value is present for DeleteMessageDays, not even an explicit nil
func (o *BanUserFromGuildRequest) UnsetDeleteMessageDays() {
	o.DeleteMessageDays.Unset()
}

func (o BanUserFromGuildRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BanUserFromGuildRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteMessageSeconds.IsSet() {
		toSerialize["delete_message_seconds"] = o.DeleteMessageSeconds.Get()
	}
	if o.DeleteMessageDays.IsSet() {
		toSerialize["delete_message_days"] = o.DeleteMessageDays.Get()
	}
	return toSerialize, nil
}

type NullableBanUserFromGuildRequest struct {
	value *BanUserFromGuildRequest
	isSet bool
}

func (v NullableBanUserFromGuildRequest) Get() *BanUserFromGuildRequest {
	return v.value
}

func (v *NullableBanUserFromGuildRequest) Set(val *BanUserFromGuildRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBanUserFromGuildRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBanUserFromGuildRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanUserFromGuildRequest(val *BanUserFromGuildRequest) *NullableBanUserFromGuildRequest {
	return &NullableBanUserFromGuildRequest{value: val, isSet: true}
}

func (v NullableBanUserFromGuildRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanUserFromGuildRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


