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

// checks if the GuildScheduledEventStatuses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildScheduledEventStatuses{}

// GuildScheduledEventStatuses struct for GuildScheduledEventStatuses
type GuildScheduledEventStatuses struct {
}

// NewGuildScheduledEventStatuses instantiates a new GuildScheduledEventStatuses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildScheduledEventStatuses() *GuildScheduledEventStatuses {
	this := GuildScheduledEventStatuses{}
	return &this
}

// NewGuildScheduledEventStatusesWithDefaults instantiates a new GuildScheduledEventStatuses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildScheduledEventStatusesWithDefaults() *GuildScheduledEventStatuses {
	this := GuildScheduledEventStatuses{}
	return &this
}

func (o GuildScheduledEventStatuses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildScheduledEventStatuses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableGuildScheduledEventStatuses struct {
	value *GuildScheduledEventStatuses
	isSet bool
}

func (v NullableGuildScheduledEventStatuses) Get() *GuildScheduledEventStatuses {
	return v.value
}

func (v *NullableGuildScheduledEventStatuses) Set(val *GuildScheduledEventStatuses) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildScheduledEventStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildScheduledEventStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildScheduledEventStatuses(val *GuildScheduledEventStatuses) *NullableGuildScheduledEventStatuses {
	return &NullableGuildScheduledEventStatuses{value: val, isSet: true}
}

func (v NullableGuildScheduledEventStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildScheduledEventStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


