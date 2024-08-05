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

// checks if the MentionableSelect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionableSelect{}

// MentionableSelect struct for MentionableSelect
type MentionableSelect struct {
	Type MessageComponentTypes `json:"type"`
	CustomId string `json:"custom_id"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	MinValues NullableInt32 `json:"min_values,omitempty"`
	MaxValues NullableInt32 `json:"max_values,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
	DefaultValues []MentionableSelectDefaultValuesInner `json:"default_values,omitempty"`
}

type _MentionableSelect MentionableSelect

// NewMentionableSelect instantiates a new MentionableSelect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionableSelect(type_ MessageComponentTypes, customId string) *MentionableSelect {
	this := MentionableSelect{}
	this.Type = type_
	this.CustomId = customId
	return &this
}

// NewMentionableSelectWithDefaults instantiates a new MentionableSelect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionableSelectWithDefaults() *MentionableSelect {
	this := MentionableSelect{}
	return &this
}

// GetType returns the Type field value
func (o *MentionableSelect) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MentionableSelect) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MentionableSelect) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetCustomId returns the CustomId field value
func (o *MentionableSelect) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *MentionableSelect) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *MentionableSelect) SetCustomId(v string) {
	o.CustomId = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionableSelect) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionableSelect) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *MentionableSelect) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *MentionableSelect) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}
// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *MentionableSelect) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *MentionableSelect) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetMinValues returns the MinValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionableSelect) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MinValues.Get()
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionableSelect) GetMinValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinValues.Get(), o.MinValues.IsSet()
}

// HasMinValues returns a boolean if a field has been set.
func (o *MentionableSelect) HasMinValues() bool {
	if o != nil && o.MinValues.IsSet() {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given NullableInt32 and assigns it to the MinValues field.
func (o *MentionableSelect) SetMinValues(v int32) {
	o.MinValues.Set(&v)
}
// SetMinValuesNil sets the value for MinValues to be an explicit nil
func (o *MentionableSelect) SetMinValuesNil() {
	o.MinValues.Set(nil)
}

// UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
func (o *MentionableSelect) UnsetMinValues() {
	o.MinValues.Unset()
}

// GetMaxValues returns the MaxValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionableSelect) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValues.Get()
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionableSelect) GetMaxValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxValues.Get(), o.MaxValues.IsSet()
}

// HasMaxValues returns a boolean if a field has been set.
func (o *MentionableSelect) HasMaxValues() bool {
	if o != nil && o.MaxValues.IsSet() {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given NullableInt32 and assigns it to the MaxValues field.
func (o *MentionableSelect) SetMaxValues(v int32) {
	o.MaxValues.Set(&v)
}
// SetMaxValuesNil sets the value for MaxValues to be an explicit nil
func (o *MentionableSelect) SetMaxValuesNil() {
	o.MaxValues.Set(nil)
}

// UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
func (o *MentionableSelect) UnsetMaxValues() {
	o.MaxValues.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionableSelect) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionableSelect) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *MentionableSelect) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *MentionableSelect) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}
// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *MentionableSelect) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *MentionableSelect) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionableSelect) GetDefaultValues() []MentionableSelectDefaultValuesInner {
	if o == nil {
		var ret []MentionableSelectDefaultValuesInner
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionableSelect) GetDefaultValuesOk() ([]MentionableSelectDefaultValuesInner, bool) {
	if o == nil || IsNil(o.DefaultValues) {
		return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *MentionableSelect) HasDefaultValues() bool {
	if o != nil && !IsNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given []MentionableSelectDefaultValuesInner and assigns it to the DefaultValues field.
func (o *MentionableSelect) SetDefaultValues(v []MentionableSelectDefaultValuesInner) {
	o.DefaultValues = v
}

func (o MentionableSelect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionableSelect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["custom_id"] = o.CustomId
	if o.Placeholder.IsSet() {
		toSerialize["placeholder"] = o.Placeholder.Get()
	}
	if o.MinValues.IsSet() {
		toSerialize["min_values"] = o.MinValues.Get()
	}
	if o.MaxValues.IsSet() {
		toSerialize["max_values"] = o.MaxValues.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	if o.DefaultValues != nil {
		toSerialize["default_values"] = o.DefaultValues
	}
	return toSerialize, nil
}

func (o *MentionableSelect) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"custom_id",
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

	varMentionableSelect := _MentionableSelect{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMentionableSelect)

	if err != nil {
		return err
	}

	*o = MentionableSelect(varMentionableSelect)

	return err
}

type NullableMentionableSelect struct {
	value *MentionableSelect
	isSet bool
}

func (v NullableMentionableSelect) Get() *MentionableSelect {
	return v.value
}

func (v *NullableMentionableSelect) Set(val *MentionableSelect) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionableSelect) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionableSelect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionableSelect(val *MentionableSelect) *NullableMentionableSelect {
	return &NullableMentionableSelect{value: val, isSet: true}
}

func (v NullableMentionableSelect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionableSelect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


