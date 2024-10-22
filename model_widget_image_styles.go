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

// checks if the WidgetImageStyles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetImageStyles{}

// WidgetImageStyles struct for WidgetImageStyles
type WidgetImageStyles struct {
}

// NewWidgetImageStyles instantiates a new WidgetImageStyles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetImageStyles() *WidgetImageStyles {
	this := WidgetImageStyles{}
	return &this
}

// NewWidgetImageStylesWithDefaults instantiates a new WidgetImageStyles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetImageStylesWithDefaults() *WidgetImageStyles {
	this := WidgetImageStyles{}
	return &this
}

func (o WidgetImageStyles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetImageStyles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableWidgetImageStyles struct {
	value *WidgetImageStyles
	isSet bool
}

func (v NullableWidgetImageStyles) Get() *WidgetImageStyles {
	return v.value
}

func (v *NullableWidgetImageStyles) Set(val *WidgetImageStyles) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetImageStyles) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetImageStyles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetImageStyles(val *WidgetImageStyles) *NullableWidgetImageStyles {
	return &NullableWidgetImageStyles{value: val, isSet: true}
}

func (v NullableWidgetImageStyles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetImageStyles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


