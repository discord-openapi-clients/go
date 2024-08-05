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

// checks if the InteractionApplicationCommandAutocompleteCallbackStringData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InteractionApplicationCommandAutocompleteCallbackStringData{}

// InteractionApplicationCommandAutocompleteCallbackStringData struct for InteractionApplicationCommandAutocompleteCallbackStringData
type InteractionApplicationCommandAutocompleteCallbackStringData struct {
	Choices []InteractionApplicationCommandAutocompleteCallbackStringDataChoicesInner `json:"choices,omitempty"`
}

// NewInteractionApplicationCommandAutocompleteCallbackStringData instantiates a new InteractionApplicationCommandAutocompleteCallbackStringData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractionApplicationCommandAutocompleteCallbackStringData() *InteractionApplicationCommandAutocompleteCallbackStringData {
	this := InteractionApplicationCommandAutocompleteCallbackStringData{}
	return &this
}

// NewInteractionApplicationCommandAutocompleteCallbackStringDataWithDefaults instantiates a new InteractionApplicationCommandAutocompleteCallbackStringData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractionApplicationCommandAutocompleteCallbackStringDataWithDefaults() *InteractionApplicationCommandAutocompleteCallbackStringData {
	this := InteractionApplicationCommandAutocompleteCallbackStringData{}
	return &this
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractionApplicationCommandAutocompleteCallbackStringData) GetChoices() []InteractionApplicationCommandAutocompleteCallbackStringDataChoicesInner {
	if o == nil {
		var ret []InteractionApplicationCommandAutocompleteCallbackStringDataChoicesInner
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractionApplicationCommandAutocompleteCallbackStringData) GetChoicesOk() ([]InteractionApplicationCommandAutocompleteCallbackStringDataChoicesInner, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *InteractionApplicationCommandAutocompleteCallbackStringData) HasChoices() bool {
	if o != nil && !IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []InteractionApplicationCommandAutocompleteCallbackStringDataChoicesInner and assigns it to the Choices field.
func (o *InteractionApplicationCommandAutocompleteCallbackStringData) SetChoices(v []InteractionApplicationCommandAutocompleteCallbackStringDataChoicesInner) {
	o.Choices = v
}

func (o InteractionApplicationCommandAutocompleteCallbackStringData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InteractionApplicationCommandAutocompleteCallbackStringData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	return toSerialize, nil
}

type NullableInteractionApplicationCommandAutocompleteCallbackStringData struct {
	value *InteractionApplicationCommandAutocompleteCallbackStringData
	isSet bool
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackStringData) Get() *InteractionApplicationCommandAutocompleteCallbackStringData {
	return v.value
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackStringData) Set(val *InteractionApplicationCommandAutocompleteCallbackStringData) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackStringData) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackStringData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractionApplicationCommandAutocompleteCallbackStringData(val *InteractionApplicationCommandAutocompleteCallbackStringData) *NullableInteractionApplicationCommandAutocompleteCallbackStringData {
	return &NullableInteractionApplicationCommandAutocompleteCallbackStringData{value: val, isSet: true}
}

func (v NullableInteractionApplicationCommandAutocompleteCallbackStringData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractionApplicationCommandAutocompleteCallbackStringData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


