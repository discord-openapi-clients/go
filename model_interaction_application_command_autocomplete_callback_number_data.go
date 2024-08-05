/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InteractionApplicationCommandAutocompleteCallbackNumberData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractionApplicationCommandAutocompleteCallbackNumberData{}

// InteractionApplicationCommandAutocompleteCallbackNumberData struct for InteractionApplicationCommandAutocompleteCallbackNumberData
type InteractionApplicationCommandAutocompleteCallbackNumberData struct {
	Choices []InteractionApplicationCommandAutocompleteCallbackNumberDataChoicesInner `json:"choices,omitempty"`
}

// NewInteractionApplicationCommandAutocompleteCallbackNumberData instantiates a new InteractionApplicationCommandAutocompleteCallbackNumberData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractionApplicationCommandAutocompleteCallbackNumberData() *InteractionApplicationCommandAutocompleteCallbackNumberData {
	this := InteractionApplicationCommandAutocompleteCallbackNumberData{}
	return &this
}

// NewInteractionApplicationCommandAutocompleteCallbackNumberDataWithDefaults instantiates a new InteractionApplicationCommandAutocompleteCallbackNumberData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractionApplicationCommandAutocompleteCallbackNumberDataWithDefaults() *InteractionApplicationCommandAutocompleteCallbackNumberData {
	this := InteractionApplicationCommandAutocompleteCallbackNumberData{}
	return &this
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) GetChoices() []InteractionApplicationCommandAutocompleteCallbackNumberDataChoicesInner {
	if o == nil {
		var ret []InteractionApplicationCommandAutocompleteCallbackNumberDataChoicesInner
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) GetChoicesOk() ([]InteractionApplicationCommandAutocompleteCallbackNumberDataChoicesInner, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []InteractionApplicationCommandAutocompleteCallbackNumberDataChoicesInner and assigns it to the Choices field.
func (o *InteractionApplicationCommandAutocompleteCallbackNumberData) SetChoices(v []InteractionApplicationCommandAutocompleteCallbackNumberDataChoicesInner) {
	o.Choices = v
}

func (o InteractionApplicationCommandAutocompleteCallbackNumberData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractionApplicationCommandAutocompleteCallbackNumberData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	return toSerialize, nil
}

type NullableInteractionApplicationCommandAutocompleteCallbackNumberData struct {
	value *InteractionApplicationCommandAutocompleteCallbackNumberData
	isSet bool
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackNumberData) Get() *InteractionApplicationCommandAutocompleteCallbackNumberData {
	return v.value
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackNumberData) Set(val *InteractionApplicationCommandAutocompleteCallbackNumberData) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackNumberData) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackNumberData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractionApplicationCommandAutocompleteCallbackNumberData(val *InteractionApplicationCommandAutocompleteCallbackNumberData) *NullableInteractionApplicationCommandAutocompleteCallbackNumberData {
	return &NullableInteractionApplicationCommandAutocompleteCallbackNumberData{value: val, isSet: true}
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackNumberData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackNumberData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

