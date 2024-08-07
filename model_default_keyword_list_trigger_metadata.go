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

// checks if the DefaultKeywordListTriggerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultKeywordListTriggerMetadata{}

// DefaultKeywordListTriggerMetadata struct for DefaultKeywordListTriggerMetadata
type DefaultKeywordListTriggerMetadata struct {
	AllowList []string `json:"allow_list,omitempty"`
	Presets []AutomodKeywordPresetType `json:"presets,omitempty"`
}

// NewDefaultKeywordListTriggerMetadata instantiates a new DefaultKeywordListTriggerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultKeywordListTriggerMetadata() *DefaultKeywordListTriggerMetadata {
	this := DefaultKeywordListTriggerMetadata{}
	return &this
}

// NewDefaultKeywordListTriggerMetadataWithDefaults instantiates a new DefaultKeywordListTriggerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultKeywordListTriggerMetadataWithDefaults() *DefaultKeywordListTriggerMetadata {
	this := DefaultKeywordListTriggerMetadata{}
	return &this
}

// GetAllowList returns the AllowList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListTriggerMetadata) GetAllowList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListTriggerMetadata) GetAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowList) {
		return nil, false
	}
	return o.AllowList, true
}

// HasAllowList returns a boolean if a field has been set.
func (o *DefaultKeywordListTriggerMetadata) HasAllowList() bool {
	if o != nil && !IsNil(o.AllowList) {
		return true
	}

	return false
}

// SetAllowList gets a reference to the given []string and assigns it to the AllowList field.
func (o *DefaultKeywordListTriggerMetadata) SetAllowList(v []string) {
	o.AllowList = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListTriggerMetadata) GetPresets() []AutomodKeywordPresetType {
	if o == nil {
		var ret []AutomodKeywordPresetType
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListTriggerMetadata) GetPresetsOk() ([]AutomodKeywordPresetType, bool) {
	if o == nil || IsNil(o.Presets) {
		return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *DefaultKeywordListTriggerMetadata) HasPresets() bool {
	if o != nil && !IsNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []AutomodKeywordPresetType and assigns it to the Presets field.
func (o *DefaultKeywordListTriggerMetadata) SetPresets(v []AutomodKeywordPresetType) {
	o.Presets = v
}

func (o DefaultKeywordListTriggerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultKeywordListTriggerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowList != nil {
		toSerialize["allow_list"] = o.AllowList
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	return toSerialize, nil
}

type NullableDefaultKeywordListTriggerMetadata struct {
	value *DefaultKeywordListTriggerMetadata
	isSet bool
}

func (v NullableDefaultKeywordListTriggerMetadata) Get() *DefaultKeywordListTriggerMetadata {
	return v.value
}

func (v *NullableDefaultKeywordListTriggerMetadata) Set(val *DefaultKeywordListTriggerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordListTriggerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordListTriggerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordListTriggerMetadata(val *DefaultKeywordListTriggerMetadata) *NullableDefaultKeywordListTriggerMetadata {
	return &NullableDefaultKeywordListTriggerMetadata{value: val, isSet: true}
}

func (v NullableDefaultKeywordListTriggerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordListTriggerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


