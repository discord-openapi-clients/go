/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ChannelSelect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelSelect{}

// ChannelSelect struct for ChannelSelect
type ChannelSelect struct {
	Type MessageComponentTypes `json:"type"`
	CustomId string `json:"custom_id"`
	Placeholder NullableString `json:"placeholder,omitempty"`
	MinValues NullableInt32 `json:"min_values,omitempty"`
	MaxValues NullableInt32 `json:"max_values,omitempty"`
	Disabled NullableBool `json:"disabled,omitempty"`
	DefaultValues []ChannelSelectDefaultValue `json:"default_values,omitempty"`
	ChannelTypes []ChannelTypes `json:"channel_types,omitempty"`
}

type _ChannelSelect ChannelSelect

// NewChannelSelect instantiates a new ChannelSelect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelSelect(type_ MessageComponentTypes, customId string) *ChannelSelect {
	this := ChannelSelect{}
	this.Type = type_
	this.CustomId = customId
	return &this
}

// NewChannelSelectWithDefaults instantiates a new ChannelSelect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelSelectWithDefaults() *ChannelSelect {
	this := ChannelSelect{}
	return &this
}

// GetType returns the Type field value
func (o *ChannelSelect) GetType() MessageComponentTypes {
	if o == nil {
		var ret MessageComponentTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChannelSelect) GetTypeOk() (*MessageComponentTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChannelSelect) SetType(v MessageComponentTypes) {
	o.Type = v
}

// GetCustomId returns the CustomId field value
func (o *ChannelSelect) GetCustomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomId
}

// GetCustomIdOk returns a tuple with the CustomId field value
// and a boolean to check if the value has been set.
func (o *ChannelSelect) GetCustomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomId, true
}

// SetCustomId sets field value
func (o *ChannelSelect) SetCustomId(v string) {
	o.CustomId = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelSelect) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder.Get()) {
		var ret string
		return ret
	}
	return *o.Placeholder.Get()
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelSelect) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placeholder.Get(), o.Placeholder.IsSet()
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *ChannelSelect) HasPlaceholder() bool {
	if o != nil && o.Placeholder.IsSet() {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given NullableString and assigns it to the Placeholder field.
func (o *ChannelSelect) SetPlaceholder(v string) {
	o.Placeholder.Set(&v)
}
// SetPlaceholderNil sets the value for Placeholder to be an explicit nil
func (o *ChannelSelect) SetPlaceholderNil() {
	o.Placeholder.Set(nil)
}

// UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
func (o *ChannelSelect) UnsetPlaceholder() {
	o.Placeholder.Unset()
}

// GetMinValues returns the MinValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelSelect) GetMinValues() int32 {
	if o == nil || IsNil(o.MinValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MinValues.Get()
}

// GetMinValuesOk returns a tuple with the MinValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelSelect) GetMinValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinValues.Get(), o.MinValues.IsSet()
}

// HasMinValues returns a boolean if a field has been set.
func (o *ChannelSelect) HasMinValues() bool {
	if o != nil && o.MinValues.IsSet() {
		return true
	}

	return false
}

// SetMinValues gets a reference to the given NullableInt32 and assigns it to the MinValues field.
func (o *ChannelSelect) SetMinValues(v int32) {
	o.MinValues.Set(&v)
}
// SetMinValuesNil sets the value for MinValues to be an explicit nil
func (o *ChannelSelect) SetMinValuesNil() {
	o.MinValues.Set(nil)
}

// UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
func (o *ChannelSelect) UnsetMinValues() {
	o.MinValues.Unset()
}

// GetMaxValues returns the MaxValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelSelect) GetMaxValues() int32 {
	if o == nil || IsNil(o.MaxValues.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValues.Get()
}

// GetMaxValuesOk returns a tuple with the MaxValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelSelect) GetMaxValuesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxValues.Get(), o.MaxValues.IsSet()
}

// HasMaxValues returns a boolean if a field has been set.
func (o *ChannelSelect) HasMaxValues() bool {
	if o != nil && o.MaxValues.IsSet() {
		return true
	}

	return false
}

// SetMaxValues gets a reference to the given NullableInt32 and assigns it to the MaxValues field.
func (o *ChannelSelect) SetMaxValues(v int32) {
	o.MaxValues.Set(&v)
}
// SetMaxValuesNil sets the value for MaxValues to be an explicit nil
func (o *ChannelSelect) SetMaxValuesNil() {
	o.MaxValues.Set(nil)
}

// UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
func (o *ChannelSelect) UnsetMaxValues() {
	o.MaxValues.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelSelect) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelSelect) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *ChannelSelect) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *ChannelSelect) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}
// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *ChannelSelect) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *ChannelSelect) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelSelect) GetDefaultValues() []ChannelSelectDefaultValue {
	if o == nil {
		var ret []ChannelSelectDefaultValue
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelSelect) GetDefaultValuesOk() ([]ChannelSelectDefaultValue, bool) {
	if o == nil || IsNil(o.DefaultValues) {
		return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *ChannelSelect) HasDefaultValues() bool {
	if o != nil && !IsNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given []ChannelSelectDefaultValue and assigns it to the DefaultValues field.
func (o *ChannelSelect) SetDefaultValues(v []ChannelSelectDefaultValue) {
	o.DefaultValues = v
}

// GetChannelTypes returns the ChannelTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelSelect) GetChannelTypes() []ChannelTypes {
	if o == nil {
		var ret []ChannelTypes
		return ret
	}
	return o.ChannelTypes
}

// GetChannelTypesOk returns a tuple with the ChannelTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelSelect) GetChannelTypesOk() ([]ChannelTypes, bool) {
	if o == nil || IsNil(o.ChannelTypes) {
		return nil, false
	}
	return o.ChannelTypes, true
}

// HasChannelTypes returns a boolean if a field has been set.
func (o *ChannelSelect) HasChannelTypes() bool {
	if o != nil && !IsNil(o.ChannelTypes) {
		return true
	}

	return false
}

// SetChannelTypes gets a reference to the given []ChannelTypes and assigns it to the ChannelTypes field.
func (o *ChannelSelect) SetChannelTypes(v []ChannelTypes) {
	o.ChannelTypes = v
}

func (o ChannelSelect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelSelect) ToMap() (map[string]interface{}, error) {
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
	if o.ChannelTypes != nil {
		toSerialize["channel_types"] = o.ChannelTypes
	}
	return toSerialize, nil
}

func (o *ChannelSelect) UnmarshalJSON(data []byte) (err error) {
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

	varChannelSelect := _ChannelSelect{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelSelect)

	if err != nil {
		return err
	}

	*o = ChannelSelect(varChannelSelect)

	return err
}

type NullableChannelSelect struct {
	value *ChannelSelect
	isSet bool
}

func (v NullableChannelSelect) Get() *ChannelSelect {
	return v.value
}

func (v *NullableChannelSelect) Set(val *ChannelSelect) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelSelect) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelSelect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelSelect(val *ChannelSelect) *NullableChannelSelect {
	return &NullableChannelSelect{value: val, isSet: true}
}

func (v NullableChannelSelect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelSelect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

