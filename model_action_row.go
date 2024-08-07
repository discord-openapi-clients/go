/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ActionRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionRow{}

// ActionRow struct for ActionRow
type ActionRow struct {
	Type MessageComponentTypes `json:"type"`
	Components []ActionRowComponentsInner `json:"components"`
}

type _ActionRow ActionRow

// NewActionRow instantiates a new ActionRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionRow(type_ MessageComponentTypes, components []ActionRowComponentsInner) *ActionRow {
	this := ActionRow{}
	this.Type = type_
	this.Components = components
	return &this
}

// NewActionRowWithDefaults instantiates a new ActionRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionRowWithDefaults() *ActionRow {
	this := ActionRow{}
	return &this
}

// GetType returns the Type field value
func (o *ActionRow) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionRow) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActionRow) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetComponents returns the Components field value
func (o *ActionRow) GetComponents() []ActionRowComponentsInner {
	if o == nil {
		var ret []ActionRowComponentsInner
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *ActionRow) GetComponentsOk() ([]ActionRowComponentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *ActionRow) SetComponents(v []ActionRowComponentsInner) {
	o.Components = v
}

func (o ActionRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["components"] = o.Components
	return toSerialize, nil
}

func (o *ActionRow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"components",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varActionRow := _ActionRow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionRow)

	if err != nil {
		return err
	}

	*o = ActionRow(varActionRow)

	return err
}

type NullableActionRow struct {
	value *ActionRow
	isSet bool
}

func (v NullableActionRow) Get() *ActionRow {
	return v.value
}

func (v *NullableActionRow) Set(val *ActionRow) {
	v.value = val
	v.isSet = true
}

func (v NullableActionRow) IsSet() bool {
	return v.isSet
}

func (v *NullableActionRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionRow(val *ActionRow) *NullableActionRow {
	return &NullableActionRow{value: val, isSet: true}
}

func (v NullableActionRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


