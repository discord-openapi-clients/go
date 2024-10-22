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

// checks if the StageInstancesPrivacyLevels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StageInstancesPrivacyLevels{}

// StageInstancesPrivacyLevels struct for StageInstancesPrivacyLevels
type StageInstancesPrivacyLevels struct {
}

// NewStageInstancesPrivacyLevels instantiates a new StageInstancesPrivacyLevels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageInstancesPrivacyLevels() *StageInstancesPrivacyLevels {
	this := StageInstancesPrivacyLevels{}
	return &this
}

// NewStageInstancesPrivacyLevelsWithDefaults instantiates a new StageInstancesPrivacyLevels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageInstancesPrivacyLevelsWithDefaults() *StageInstancesPrivacyLevels {
	this := StageInstancesPrivacyLevels{}
	return &this
}

func (o StageInstancesPrivacyLevels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StageInstancesPrivacyLevels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableStageInstancesPrivacyLevels struct {
	value *StageInstancesPrivacyLevels
	isSet bool
}

func (v NullableStageInstancesPrivacyLevels) Get() *StageInstancesPrivacyLevels {
	return v.value
}

func (v *NullableStageInstancesPrivacyLevels) Set(val *StageInstancesPrivacyLevels) {
	v.value = val
	v.isSet = true
}

func (v NullableStageInstancesPrivacyLevels) IsSet() bool {
	return v.isSet
}

func (v *NullableStageInstancesPrivacyLevels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageInstancesPrivacyLevels(val *StageInstancesPrivacyLevels) *NullableStageInstancesPrivacyLevels {
	return &NullableStageInstancesPrivacyLevels{value: val, isSet: true}
}

func (v NullableStageInstancesPrivacyLevels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageInstancesPrivacyLevels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


