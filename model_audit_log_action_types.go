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

// checks if the AuditLogActionTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLogActionTypes{}

// AuditLogActionTypes struct for AuditLogActionTypes
type AuditLogActionTypes struct {
}

// NewAuditLogActionTypes instantiates a new AuditLogActionTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLogActionTypes() *AuditLogActionTypes {
	this := AuditLogActionTypes{}
	return &this
}

// NewAuditLogActionTypesWithDefaults instantiates a new AuditLogActionTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogActionTypesWithDefaults() *AuditLogActionTypes {
	this := AuditLogActionTypes{}
	return &this
}

func (o AuditLogActionTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditLogActionTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAuditLogActionTypes struct {
	value *AuditLogActionTypes
	isSet bool
}

func (v NullableAuditLogActionTypes) Get() *AuditLogActionTypes {
	return v.value
}

func (v *NullableAuditLogActionTypes) Set(val *AuditLogActionTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLogActionTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLogActionTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLogActionTypes(val *AuditLogActionTypes) *NullableAuditLogActionTypes {
	return &NullableAuditLogActionTypes{value: val, isSet: true}
}

func (v NullableAuditLogActionTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLogActionTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


