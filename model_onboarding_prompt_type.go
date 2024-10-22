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

// checks if the OnboardingPromptType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnboardingPromptType{}

// OnboardingPromptType struct for OnboardingPromptType
type OnboardingPromptType struct {
}

// NewOnboardingPromptType instantiates a new OnboardingPromptType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnboardingPromptType() *OnboardingPromptType {
	this := OnboardingPromptType{}
	return &this
}

// NewOnboardingPromptTypeWithDefaults instantiates a new OnboardingPromptType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnboardingPromptTypeWithDefaults() *OnboardingPromptType {
	this := OnboardingPromptType{}
	return &this
}

func (o OnboardingPromptType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnboardingPromptType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableOnboardingPromptType struct {
	value *OnboardingPromptType
	isSet bool
}

func (v NullableOnboardingPromptType) Get() *OnboardingPromptType {
	return v.value
}

func (v *NullableOnboardingPromptType) Set(val *OnboardingPromptType) {
	v.value = val
	v.isSet = true
}

func (v NullableOnboardingPromptType) IsSet() bool {
	return v.isSet
}

func (v *NullableOnboardingPromptType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnboardingPromptType(val *OnboardingPromptType) *NullableOnboardingPromptType {
	return &NullableOnboardingPromptType{value: val, isSet: true}
}

func (v NullableOnboardingPromptType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnboardingPromptType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


