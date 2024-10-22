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

// checks if the GuildScheduledEventPrivacyLevels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildScheduledEventPrivacyLevels{}

// GuildScheduledEventPrivacyLevels struct for GuildScheduledEventPrivacyLevels
type GuildScheduledEventPrivacyLevels struct {
}

// NewGuildScheduledEventPrivacyLevels instantiates a new GuildScheduledEventPrivacyLevels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildScheduledEventPrivacyLevels() *GuildScheduledEventPrivacyLevels {
	this := GuildScheduledEventPrivacyLevels{}
	return &this
}

// NewGuildScheduledEventPrivacyLevelsWithDefaults instantiates a new GuildScheduledEventPrivacyLevels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildScheduledEventPrivacyLevelsWithDefaults() *GuildScheduledEventPrivacyLevels {
	this := GuildScheduledEventPrivacyLevels{}
	return &this
}

func (o GuildScheduledEventPrivacyLevels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildScheduledEventPrivacyLevels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableGuildScheduledEventPrivacyLevels struct {
	value *GuildScheduledEventPrivacyLevels
	isSet bool
}

func (v NullableGuildScheduledEventPrivacyLevels) Get() *GuildScheduledEventPrivacyLevels {
	return v.value
}

func (v *NullableGuildScheduledEventPrivacyLevels) Set(val *GuildScheduledEventPrivacyLevels) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildScheduledEventPrivacyLevels) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildScheduledEventPrivacyLevels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildScheduledEventPrivacyLevels(val *GuildScheduledEventPrivacyLevels) *NullableGuildScheduledEventPrivacyLevels {
	return &NullableGuildScheduledEventPrivacyLevels{value: val, isSet: true}
}

func (v NullableGuildScheduledEventPrivacyLevels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildScheduledEventPrivacyLevels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


