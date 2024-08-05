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

// checks if the Button type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Button{}

// Button struct for Button
type Button struct {
	Type MessageComponentTypes `json:"type"`
	CustomId NullableString `json:"custom_id,omitempty"`
	Style ButtonStyleTypes `json:"style"`
	Label NullableString `json:"label,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
	Emoji NullableEmoji `json:"emoji,omitempty"`
	Url NullableString `json:"url,omitempty"`
	SkuId *string `json:"sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _Button Button

// NewButton instantiates a new Button object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewButton(type_ MessageComponentTypes, style ButtonStyleTypes) *Button {
	this := Button{}
	this.Type = type_
	this.Style = style
	return &this
}

// NewButtonWithDefaults instantiates a new Button object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewButtonWithDefaults() *Button {
	this := Button{}
	return &this
}

// GetType returns the Type field value
func (o *Button) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Button) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Button) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetCustomId returns the CustomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Button) GetCustomId() string {
	if o == nil || IsNil(o.CustomId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomId.Get()
}

// GetCustomIdOk returns a tuple with the CustomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Button) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomId.Get(), o.CustomId.IsSet()
}

// HasCustomId returns a boolean if a field has been set.
func (o *Button) HasCustomId() bool {
	if o != nil && o.CustomId.IsSet() {
		return true
	}

	return false
}

// SetCustomId gets a reference to the given NullableString and assigns it to the CustomId field.
func (o *Button) SetCustomId(v string) {
	o.CustomId.Set(&v)
}
// SetCustomIdNil sets the value for CustomId to be an explicit nil
func (o *Button) SetCustomIdNil() {
	o.CustomId.Set(nil)
}

// UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
func (o *Button) UnsetCustomId() {
	o.CustomId.Unset()
}

// GetStyle returns the Style field value
func (o *Button) GetStyle() ButtonStyleTypes {
	if o == nil {
		var ret ButtonStyleTypes
		return ret
	}

	return o.Style
}

// GetStyleOk returns a tuple with the Style field value
// and a boolean to check if the value has been set.
func (o *Button) GetStyleOk() (*ButtonStyleTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Style, true
}

// SetStyle sets field value
func (o *Button) SetStyle(v ButtonStyleTypes) {
	o.Style = v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Button) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Button) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *Button) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *Button) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *Button) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *Button) UnsetLabel() {
	o.Label.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Button) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Button) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *Button) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *Button) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}
// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *Button) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *Button) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Button) GetEmoji() Emoji {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret Emoji
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Button) GetEmojiOk() (*Emoji, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *Button) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableEmoji and assigns it to the Emoji field.
func (o *Button) SetEmoji(v Emoji) {
	o.Emoji.Set(&v)
}
// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *Button) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *Button) UnsetEmoji() {
	o.Emoji.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Button) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Button) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Button) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Button) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Button) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Button) UnsetUrl() {
	o.Url.Unset()
}

// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *Button) GetSkuId() string {
	if o == nil || IsNil(o.SkuId) {
		var ret string
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Button) GetSkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkuId) {
		return nil, false
	}
	return o.SkuId, true
}

// HasSkuId returns a boolean if a field has been set.
func (o *Button) HasSkuId() bool {
	if o != nil && !IsNil(o.SkuId) {
		return true
	}

	return false
}

// SetSkuId gets a reference to the given string and assigns it to the SkuId field.
func (o *Button) SetSkuId(v string) {
	o.SkuId = &v
}

func (o Button) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Button) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.CustomId.IsSet() {
		toSerialize["custom_id"] = o.CustomId.Get()
	}
	toSerialize["style"] = o.Style
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.SkuId) {
		toSerialize["sku_id"] = o.SkuId
	}
	return toSerialize, nil
}

func (o *Button) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"style",
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

	varButton := _Button{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varButton)

	if err != nil {
		return err
	}

	*o = Button(varButton)

	return err
}

type NullableButton struct {
	value *Button
	isSet bool
}

func (v NullableButton) Get() *Button {
	return v.value
}

func (v *NullableButton) Set(val *Button) {
	v.value = val
	v.isSet = true
}

func (v NullableButton) IsSet() bool {
	return v.isSet
}

func (v *NullableButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButton(val *Button) *NullableButton {
	return &NullableButton{value: val, isSet: true}
}

func (v NullableButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


