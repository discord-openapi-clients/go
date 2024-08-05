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

// checks if the InteractionApplicationCommandAutocompleteCallbackIntegerData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractionApplicationCommandAutocompleteCallbackIntegerData{}

// InteractionApplicationCommandAutocompleteCallbackIntegerData struct for InteractionApplicationCommandAutocompleteCallbackIntegerData
type InteractionApplicationCommandAutocompleteCallbackIntegerData struct {
	Choices []InteractionApplicationCommandAutocompleteCallbackIntegerDataChoicesInner `json:"choices,omitempty"`
}

// NewInteractionApplicationCommandAutocompleteCallbackIntegerData instantiates a new InteractionApplicationCommandAutocompleteCallbackIntegerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractionApplicationCommandAutocompleteCallbackIntegerData() *InteractionApplicationCommandAutocompleteCallbackIntegerData {
	this := InteractionApplicationCommandAutocompleteCallbackIntegerData{}
	return &this
}

// NewInteractionApplicationCommandAutocompleteCallbackIntegerDataWithDefaults instantiates a new InteractionApplicationCommandAutocompleteCallbackIntegerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractionApplicationCommandAutocompleteCallbackIntegerDataWithDefaults() *InteractionApplicationCommandAutocompleteCallbackIntegerData {
	this := InteractionApplicationCommandAutocompleteCallbackIntegerData{}
	return &this
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractionApplicationCommandAutocompleteCallbackIntegerData) GetChoices() []InteractionApplicationCommandAutocompleteCallbackIntegerDataChoicesInner {
	if o == nil {
		var ret []InteractionApplicationCommandAutocompleteCallbackIntegerDataChoicesInner
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractionApplicationCommandAutocompleteCallbackIntegerData) GetChoicesOk() ([]InteractionApplicationCommandAutocompleteCallbackIntegerDataChoicesInner, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *InteractionApplicationCommandAutocompleteCallbackIntegerData) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []InteractionApplicationCommandAutocompleteCallbackIntegerDataChoicesInner and assigns it to the Choices field.
func (o *InteractionApplicationCommandAutocompleteCallbackIntegerData) SetChoices(v []InteractionApplicationCommandAutocompleteCallbackIntegerDataChoicesInner) {
	o.Choices = v
}

func (o InteractionApplicationCommandAutocompleteCallbackIntegerData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractionApplicationCommandAutocompleteCallbackIntegerData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	return toSerialize, nil
}

type NullableInteractionApplicationCommandAutocompleteCallbackIntegerData struct {
	value *InteractionApplicationCommandAutocompleteCallbackIntegerData
	isSet bool
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackIntegerData) Get() *InteractionApplicationCommandAutocompleteCallbackIntegerData {
	return v.value
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackIntegerData) Set(val *InteractionApplicationCommandAutocompleteCallbackIntegerData) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackIntegerData) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackIntegerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractionApplicationCommandAutocompleteCallbackIntegerData(val *InteractionApplicationCommandAutocompleteCallbackIntegerData) *NullableInteractionApplicationCommandAutocompleteCallbackIntegerData {
	return &NullableInteractionApplicationCommandAutocompleteCallbackIntegerData{value: val, isSet: true}
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackIntegerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackIntegerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


