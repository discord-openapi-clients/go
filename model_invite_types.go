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

// checks if the InviteTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteTypes{}

// InviteTypes struct for InviteTypes
type InviteTypes struct {
}

// NewInviteTypes instantiates a new InviteTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteTypes() *InviteTypes {
	this := InviteTypes{}
	return &this
}

// NewInviteTypesWithDefaults instantiates a new InviteTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteTypesWithDefaults() *InviteTypes {
	this := InviteTypes{}
	return &this
}

func (o InviteTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableInviteTypes struct {
	value *InviteTypes
	isSet bool
}

func (v NullableInviteTypes) Get() *InviteTypes {
	return v.value
}

func (v *NullableInviteTypes) Set(val *InviteTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteTypes(val *InviteTypes) *NullableInviteTypes {
	return &NullableInviteTypes{value: val, isSet: true}
}

func (v NullableInviteTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


