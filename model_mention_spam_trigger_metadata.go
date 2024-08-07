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

// checks if the MentionSpamTriggerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionSpamTriggerMetadata{}

// MentionSpamTriggerMetadata struct for MentionSpamTriggerMetadata
type MentionSpamTriggerMetadata struct {
	MentionTotalLimit int32 `json:"mention_total_limit"`
	MentionRaidProtectionEnabled NullableBool `json:"mention_raid_protection_enabled,omitempty"`
}

type _MentionSpamTriggerMetadata MentionSpamTriggerMetadata

// NewMentionSpamTriggerMetadata instantiates a new MentionSpamTriggerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionSpamTriggerMetadata(mentionTotalLimit int32) *MentionSpamTriggerMetadata {
	this := MentionSpamTriggerMetadata{}
	this.MentionTotalLimit = mentionTotalLimit
	return &this
}

// NewMentionSpamTriggerMetadataWithDefaults instantiates a new MentionSpamTriggerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionSpamTriggerMetadataWithDefaults() *MentionSpamTriggerMetadata {
	this := MentionSpamTriggerMetadata{}
	return &this
}

// GetMentionTotalLimit returns the MentionTotalLimit field value
func (o *MentionSpamTriggerMetadata) GetMentionTotalLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MentionTotalLimit
}

// GetMentionTotalLimitOk returns a tuple with the MentionTotalLimit field value
// and a boolean to check if the value has been set.
func (o *MentionSpamTriggerMetadata) GetMentionTotalLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MentionTotalLimit, true
}

// SetMentionTotalLimit sets field value
func (o *MentionSpamTriggerMetadata) SetMentionTotalLimit(v int32) {
	o.MentionTotalLimit = v
}

// GetMentionRaidProtectionEnabled returns the MentionRaidProtectionEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionSpamTriggerMetadata) GetMentionRaidProtectionEnabled() bool {
	if o == nil || IsNil(o.MentionRaidProtectionEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.MentionRaidProtectionEnabled.Get()
}

// GetMentionRaidProtectionEnabledOk returns a tuple with the MentionRaidProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionSpamTriggerMetadata) GetMentionRaidProtectionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionRaidProtectionEnabled.Get(), o.MentionRaidProtectionEnabled.IsSet()
}

// HasMentionRaidProtectionEnabled returns a boolean if a field has been set.
func (o *MentionSpamTriggerMetadata) HasMentionRaidProtectionEnabled() bool {
	if o != nil && o.MentionRaidProtectionEnabled.IsSet() {
		return true
	}

	return false
}

// SetMentionRaidProtectionEnabled gets a reference to the given NullableBool and assigns it to the MentionRaidProtectionEnabled field.
func (o *MentionSpamTriggerMetadata) SetMentionRaidProtectionEnabled(v bool) {
	o.MentionRaidProtectionEnabled.Set(&v)
}
// SetMentionRaidProtectionEnabledNil sets the value for MentionRaidProtectionEnabled to be an explicit nil
func (o *MentionSpamTriggerMetadata) SetMentionRaidProtectionEnabledNil() {
	o.MentionRaidProtectionEnabled.Set(nil)
}

// UnsetMentionRaidProtectionEnabled ensures that no value is present for MentionRaidProtectionEnabled, not even an explicit nil
func (o *MentionSpamTriggerMetadata) UnsetMentionRaidProtectionEnabled() {
	o.MentionRaidProtectionEnabled.Unset()
}

func (o MentionSpamTriggerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionSpamTriggerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mention_total_limit"] = o.MentionTotalLimit
	if o.MentionRaidProtectionEnabled.IsSet() {
		toSerialize["mention_raid_protection_enabled"] = o.MentionRaidProtectionEnabled.Get()
	}
	return toSerialize, nil
}

func (o *MentionSpamTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mention_total_limit",
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

	varMentionSpamTriggerMetadata := _MentionSpamTriggerMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMentionSpamTriggerMetadata)

	if err != nil {
		return err
	}

	*o = MentionSpamTriggerMetadata(varMentionSpamTriggerMetadata)

	return err
}

type NullableMentionSpamTriggerMetadata struct {
	value *MentionSpamTriggerMetadata
	isSet bool
}

func (v NullableMentionSpamTriggerMetadata) Get() *MentionSpamTriggerMetadata {
	return v.value
}

func (v *NullableMentionSpamTriggerMetadata) Set(val *MentionSpamTriggerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionSpamTriggerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionSpamTriggerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionSpamTriggerMetadata(val *MentionSpamTriggerMetadata) *NullableMentionSpamTriggerMetadata {
	return &NullableMentionSpamTriggerMetadata{value: val, isSet: true}
}

func (v NullableMentionSpamTriggerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionSpamTriggerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


