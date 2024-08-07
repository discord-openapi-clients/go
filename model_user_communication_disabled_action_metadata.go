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

// checks if the UserCommunicationDisabledActionMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCommunicationDisabledActionMetadata{}

// UserCommunicationDisabledActionMetadata struct for UserCommunicationDisabledActionMetadata
type UserCommunicationDisabledActionMetadata struct {
	DurationSeconds NullableInt32 `json:"duration_seconds,omitempty"`
}

// NewUserCommunicationDisabledActionMetadata instantiates a new UserCommunicationDisabledActionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCommunicationDisabledActionMetadata() *UserCommunicationDisabledActionMetadata {
	this := UserCommunicationDisabledActionMetadata{}
	return &this
}

// NewUserCommunicationDisabledActionMetadataWithDefaults instantiates a new UserCommunicationDisabledActionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCommunicationDisabledActionMetadataWithDefaults() *UserCommunicationDisabledActionMetadata {
	this := UserCommunicationDisabledActionMetadata{}
	return &this
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCommunicationDisabledActionMetadata) GetDurationSeconds() int32 {
	if o == nil || IsNil(o.DurationSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.DurationSeconds.Get()
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCommunicationDisabledActionMetadata) GetDurationSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationSeconds.Get(), o.DurationSeconds.IsSet()
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *UserCommunicationDisabledActionMetadata) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given NullableInt32 and assigns it to the DurationSeconds field.
func (o *UserCommunicationDisabledActionMetadata) SetDurationSeconds(v int32) {
	o.DurationSeconds.Set(&v)
}
// SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil
func (o *UserCommunicationDisabledActionMetadata) SetDurationSecondsNil() {
	o.DurationSeconds.Set(nil)
}

// UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
func (o *UserCommunicationDisabledActionMetadata) UnsetDurationSeconds() {
	o.DurationSeconds.Unset()
}

func (o UserCommunicationDisabledActionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCommunicationDisabledActionMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DurationSeconds.IsSet() {
		toSerialize["duration_seconds"] = o.DurationSeconds.Get()
	}
	return toSerialize, nil
}

type NullableUserCommunicationDisabledActionMetadata struct {
	value *UserCommunicationDisabledActionMetadata
	isSet bool
}

func (v NullableUserCommunicationDisabledActionMetadata) Get() *UserCommunicationDisabledActionMetadata {
	return v.value
}

func (v *NullableUserCommunicationDisabledActionMetadata) Set(val *UserCommunicationDisabledActionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCommunicationDisabledActionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCommunicationDisabledActionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCommunicationDisabledActionMetadata(val *UserCommunicationDisabledActionMetadata) *NullableUserCommunicationDisabledActionMetadata {
	return &NullableUserCommunicationDisabledActionMetadata{value: val, isSet: true}
}

func (v NullableUserCommunicationDisabledActionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCommunicationDisabledActionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


