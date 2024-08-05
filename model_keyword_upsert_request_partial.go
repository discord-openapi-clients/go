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

// checks if the KeywordUpsertRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordUpsertRequestPartial{}

// KeywordUpsertRequestPartial struct for KeywordUpsertRequestPartial
type KeywordUpsertRequestPartial struct {
	Name *string `json:"name,omitempty"`
	EventType *AutomodEventType `json:"event_type,omitempty"`
	Actions []DefaultKeywordListUpsertRequestActionsInner `json:"actions,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	ExemptRoles []string `json:"exempt_roles,omitempty"`
	ExemptChannels []string `json:"exempt_channels,omitempty"`
	TriggerType *AutomodTriggerType `json:"trigger_type,omitempty"`
	TriggerMetadata NullableKeywordTriggerMetadata `json:"trigger_metadata,omitempty"`
}

// NewKeywordUpsertRequestPartial instantiates a new KeywordUpsertRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordUpsertRequestPartial() *KeywordUpsertRequestPartial {
	this := KeywordUpsertRequestPartial{}
	return &this
}

// NewKeywordUpsertRequestPartialWithDefaults instantiates a new KeywordUpsertRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordUpsertRequestPartialWithDefaults() *KeywordUpsertRequestPartial {
	this := KeywordUpsertRequestPartial{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeywordUpsertRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordUpsertRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeywordUpsertRequestPartial) SetName(v string) {
	o.Name = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *KeywordUpsertRequestPartial) GetEventType() AutomodEventType {
	if o == nil || IsNil(o.EventType) {
		var ret AutomodEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordUpsertRequestPartial) GetEventTypeOk() (*AutomodEventType, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given AutomodEventType and assigns it to the EventType field.
func (o *KeywordUpsertRequestPartial) SetEventType(v AutomodEventType) {
	o.EventType = &v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordUpsertRequestPartial) GetActions() []DefaultKeywordListUpsertRequestActionsInner {
	if o == nil {
		var ret []DefaultKeywordListUpsertRequestActionsInner
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordUpsertRequestPartial) GetActionsOk() ([]DefaultKeywordListUpsertRequestActionsInner, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []DefaultKeywordListUpsertRequestActionsInner and assigns it to the Actions field.
func (o *KeywordUpsertRequestPartial) SetActions(v []DefaultKeywordListUpsertRequestActionsInner) {
	o.Actions = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordUpsertRequestPartial) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordUpsertRequestPartial) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *KeywordUpsertRequestPartial) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *KeywordUpsertRequestPartial) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *KeywordUpsertRequestPartial) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetExemptRoles returns the ExemptRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordUpsertRequestPartial) GetExemptRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptRoles
}

// GetExemptRolesOk returns a tuple with the ExemptRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordUpsertRequestPartial) GetExemptRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptRoles) {
		return nil, false
	}
	return o.ExemptRoles, true
}

// HasExemptRoles returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasExemptRoles() bool {
	if o != nil && !IsNil(o.ExemptRoles) {
		return true
	}

	return false
}

// SetExemptRoles gets a reference to the given []string and assigns it to the ExemptRoles field.
func (o *KeywordUpsertRequestPartial) SetExemptRoles(v []string) {
	o.ExemptRoles = v
}

// GetExemptChannels returns the ExemptChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordUpsertRequestPartial) GetExemptChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptChannels
}

// GetExemptChannelsOk returns a tuple with the ExemptChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordUpsertRequestPartial) GetExemptChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptChannels) {
		return nil, false
	}
	return o.ExemptChannels, true
}

// HasExemptChannels returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasExemptChannels() bool {
	if o != nil && !IsNil(o.ExemptChannels) {
		return true
	}

	return false
}

// SetExemptChannels gets a reference to the given []string and assigns it to the ExemptChannels field.
func (o *KeywordUpsertRequestPartial) SetExemptChannels(v []string) {
	o.ExemptChannels = v
}

// GetTriggerType returns the TriggerType field value if set, zero value otherwise.
func (o *KeywordUpsertRequestPartial) GetTriggerType() AutomodTriggerType {
	if o == nil || IsNil(o.TriggerType) {
		var ret AutomodTriggerType
		return ret
	}
	return *o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordUpsertRequestPartial) GetTriggerTypeOk() (*AutomodTriggerType, bool) {
	if o == nil || IsNil(o.TriggerType) {
		return nil, false
	}
	return o.TriggerType, true
}

// HasTriggerType returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasTriggerType() bool {
	if o != nil && !IsNil(o.TriggerType) {
		return true
	}

	return false
}

// SetTriggerType gets a reference to the given AutomodTriggerType and assigns it to the TriggerType field.
func (o *KeywordUpsertRequestPartial) SetTriggerType(v AutomodTriggerType) {
	o.TriggerType = &v
}

// GetTriggerMetadata returns the TriggerMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeywordUpsertRequestPartial) GetTriggerMetadata() KeywordTriggerMetadata {
	if o == nil || IsNil(o.TriggerMetadata.Get()) {
		var ret KeywordTriggerMetadata
		return ret
	}
	return *o.TriggerMetadata.Get()
}

// GetTriggerMetadataOk returns a tuple with the TriggerMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeywordUpsertRequestPartial) GetTriggerMetadataOk() (*KeywordTriggerMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.TriggerMetadata.Get(), o.TriggerMetadata.IsSet()
}

// HasTriggerMetadata returns a boolean if a field has been set.
func (o *KeywordUpsertRequestPartial) HasTriggerMetadata() bool {
	if o != nil && o.TriggerMetadata.IsSet() {
		return true
	}

	return false
}

// SetTriggerMetadata gets a reference to the given NullableKeywordTriggerMetadata and assigns it to the TriggerMetadata field.
func (o *KeywordUpsertRequestPartial) SetTriggerMetadata(v KeywordTriggerMetadata) {
	o.TriggerMetadata.Set(&v)
}
// SetTriggerMetadataNil sets the value for TriggerMetadata to be an explicit nil
func (o *KeywordUpsertRequestPartial) SetTriggerMetadataNil() {
	o.TriggerMetadata.Set(nil)
}

// UnsetTriggerMetadata ensures that no value is present for TriggerMetadata, not even an explicit nil
func (o *KeywordUpsertRequestPartial) UnsetTriggerMetadata() {
	o.TriggerMetadata.Unset()
}

func (o KeywordUpsertRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeywordUpsertRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.ExemptRoles != nil {
		toSerialize["exempt_roles"] = o.ExemptRoles
	}
	if o.ExemptChannels != nil {
		toSerialize["exempt_channels"] = o.ExemptChannels
	}
	if !IsNil(o.TriggerType) {
		toSerialize["trigger_type"] = o.TriggerType
	}
	if o.TriggerMetadata.IsSet() {
		toSerialize["trigger_metadata"] = o.TriggerMetadata.Get()
	}
	return toSerialize, nil
}

type NullableKeywordUpsertRequestPartial struct {
	value *KeywordUpsertRequestPartial
	isSet bool
}

func (v NullableKeywordUpsertRequestPartial) Get() *KeywordUpsertRequestPartial {
	return v.value
}

func (v *NullableKeywordUpsertRequestPartial) Set(val *KeywordUpsertRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordUpsertRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordUpsertRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordUpsertRequestPartial(val *KeywordUpsertRequestPartial) *NullableKeywordUpsertRequestPartial {
	return &NullableKeywordUpsertRequestPartial{value: val, isSet: true}
}

func (v NullableKeywordUpsertRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeywordUpsertRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


