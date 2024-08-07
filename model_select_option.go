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

// checks if the SelectOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectOption{}

// SelectOption struct for SelectOption
type SelectOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Description NullableString `json:"description,omitempty"`
	Emoji NullableEmoji `json:"emoji,omitempty"`
	Default NullableBool `json:"default,omitempty"`
}

type _SelectOption SelectOption

// NewSelectOption instantiates a new SelectOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectOption(label string, value string) *SelectOption {
	this := SelectOption{}
	this.Label = label
	this.Value = value
	return &this
}

// NewSelectOptionWithDefaults instantiates a new SelectOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectOptionWithDefaults() *SelectOption {
	this := SelectOption{}
	return &this
}

// GetLabel returns the Label field value
func (o *SelectOption) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SelectOption) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SelectOption) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *SelectOption) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SelectOption) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SelectOption) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectOption) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectOption) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SelectOption) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SelectOption) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SelectOption) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SelectOption) UnsetDescription() {
	o.Description.Unset()
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectOption) GetEmoji() Emoji {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret Emoji
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectOption) GetEmojiOk() (*Emoji, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *SelectOption) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableEmoji and assigns it to the Emoji field.
func (o *SelectOption) SetEmoji(v Emoji) {
	o.Emoji.Set(&v)
}
// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *SelectOption) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *SelectOption) UnsetEmoji() {
	o.Emoji.Unset()
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectOption) GetDefault() bool {
	if o == nil || IsNil(o.Default.Get()) {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectOption) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *SelectOption) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableBool and assigns it to the Default field.
func (o *SelectOption) SetDefault(v bool) {
	o.Default.Set(&v)
}
// SetDefaultNil sets the value for Default to be an explicit nil
func (o *SelectOption) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *SelectOption) UnsetDefault() {
	o.Default.Unset()
}

func (o SelectOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	return toSerialize, nil
}

func (o *SelectOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"value",
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

	varSelectOption := _SelectOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSelectOption)

	if err != nil {
		return err
	}

	*o = SelectOption(varSelectOption)

	return err
}

type NullableSelectOption struct {
	value *SelectOption
	isSet bool
}

func (v NullableSelectOption) Get() *SelectOption {
	return v.value
}

func (v *NullableSelectOption) Set(val *SelectOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectOption(val *SelectOption) *NullableSelectOption {
	return &NullableSelectOption{value: val, isSet: true}
}

func (v NullableSelectOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


