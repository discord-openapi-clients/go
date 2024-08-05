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

// checks if the DefaultKeywordListUpsertRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultKeywordListUpsertRequest{}

// DefaultKeywordListUpsertRequest struct for DefaultKeywordListUpsertRequest
type DefaultKeywordListUpsertRequest struct {
	Name string `json:"name"`
	EventType AutomodEventType `json:"event_type"`
	Actions []DefaultKeywordListUpsertRequestActionsInner `json:"actions,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	ExemptRoles []string `json:"exempt_roles,omitempty"`
	ExemptChannels []string `json:"exempt_channels,omitempty"`
	TriggerType AutomodTriggerType `json:"trigger_type"`
	TriggerMetadata DefaultKeywordListTriggerMetadata `json:"trigger_metadata"`
}

type _DefaultKeywordListUpsertRequest DefaultKeywordListUpsertRequest

// NewDefaultKeywordListUpsertRequest instantiates a new DefaultKeywordListUpsertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultKeywordListUpsertRequest(name string, eventType AutomodEventType, triggerType AutomodTriggerType, triggerMetadata DefaultKeywordListTriggerMetadata) *DefaultKeywordListUpsertRequest {
	this := DefaultKeywordListUpsertRequest{}
	this.Name = name
	this.EventType = eventType
	this.TriggerType = triggerType
	this.TriggerMetadata = triggerMetadata
	return &this
}

// NewDefaultKeywordListUpsertRequestWithDefaults instantiates a new DefaultKeywordListUpsertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultKeywordListUpsertRequestWithDefaults() *DefaultKeywordListUpsertRequest {
	this := DefaultKeywordListUpsertRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DefaultKeywordListUpsertRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DefaultKeywordListUpsertRequest) SetName(v string) {
	o.Name = v
}

// GetEventType returns the EventType field value
func (o *DefaultKeywordListUpsertRequest) GetEventType() AutomodEventType {
	if o == nil {
		var ret AutomodEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequest) GetEventTypeOk() (*AutomodEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *DefaultKeywordListUpsertRequest) SetEventType(v AutomodEventType) {
	o.EventType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequest) GetActions() []DefaultKeywordListUpsertRequestActionsInner {
	if o == nil {
		var ret []DefaultKeywordListUpsertRequestActionsInner
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequest) GetActionsOk() ([]DefaultKeywordListUpsertRequestActionsInner, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequest) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []DefaultKeywordListUpsertRequestActionsInner and assigns it to the Actions field.
func (o *DefaultKeywordListUpsertRequest) SetActions(v []DefaultKeywordListUpsertRequestActionsInner) {
	o.Actions = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequest) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *DefaultKeywordListUpsertRequest) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *DefaultKeywordListUpsertRequest) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *DefaultKeywordListUpsertRequest) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetExemptRoles returns the ExemptRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequest) GetExemptRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptRoles
}

// GetExemptRolesOk returns a tuple with the ExemptRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequest) GetExemptRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptRoles) {
		return nil, false
	}
	return o.ExemptRoles, true
}

// HasExemptRoles returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequest) HasExemptRoles() bool {
	if o != nil && !IsNil(o.ExemptRoles) {
		return true
	}

	return false
}

// SetExemptRoles gets a reference to the given []string and assigns it to the ExemptRoles field.
func (o *DefaultKeywordListUpsertRequest) SetExemptRoles(v []string) {
	o.ExemptRoles = v
}

// GetExemptChannels returns the ExemptChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordListUpsertRequest) GetExemptChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptChannels
}

// GetExemptChannelsOk returns a tuple with the ExemptChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordListUpsertRequest) GetExemptChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptChannels) {
		return nil, false
	}
	return o.ExemptChannels, true
}

// HasExemptChannels returns a boolean if a field has been set.
func (o *DefaultKeywordListUpsertRequest) HasExemptChannels() bool {
	if o != nil && !IsNil(o.ExemptChannels) {
		return true
	}

	return false
}

// SetExemptChannels gets a reference to the given []string and assigns it to the ExemptChannels field.
func (o *DefaultKeywordListUpsertRequest) SetExemptChannels(v []string) {
	o.ExemptChannels = v
}

// GetTriggerType returns the TriggerType field value
func (o *DefaultKeywordListUpsertRequest) GetTriggerType() AutomodTriggerType {
	if o == nil {
		var ret AutomodTriggerType
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequest) GetTriggerTypeOk() (*AutomodTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *DefaultKeywordListUpsertRequest) SetTriggerType(v AutomodTriggerType) {
	o.TriggerType = v
}

// GetTriggerMetadata returns the TriggerMetadata field value
func (o *DefaultKeywordListUpsertRequest) GetTriggerMetadata() DefaultKeywordListTriggerMetadata {
	if o == nil {
		var ret DefaultKeywordListTriggerMetadata
		return ret
	}

	return o.TriggerMetadata
}

// GetTriggerMetadataOk returns a tuple with the TriggerMetadata field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordListUpsertRequest) GetTriggerMetadataOk() (*DefaultKeywordListTriggerMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerMetadata, true
}

// SetTriggerMetadata sets field value
func (o *DefaultKeywordListUpsertRequest) SetTriggerMetadata(v DefaultKeywordListTriggerMetadata) {
	o.TriggerMetadata = v
}

func (o DefaultKeywordListUpsertRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultKeywordListUpsertRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["event_type"] = o.EventType
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
	toSerialize["trigger_type"] = o.TriggerType
	toSerialize["trigger_metadata"] = o.TriggerMetadata
	return toSerialize, nil
}

func (o *DefaultKeywordListUpsertRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"event_type",
		"trigger_type",
		"trigger_metadata",
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

	varDefaultKeywordListUpsertRequest := _DefaultKeywordListUpsertRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDefaultKeywordListUpsertRequest)

	if err != nil {
		return err
	}

	*o = DefaultKeywordListUpsertRequest(varDefaultKeywordListUpsertRequest)

	return err
}

type NullableDefaultKeywordListUpsertRequest struct {
	value *DefaultKeywordListUpsertRequest
	isSet bool
}

func (v NullableDefaultKeywordListUpsertRequest) Get() *DefaultKeywordListUpsertRequest {
	return v.value
}

func (v *NullableDefaultKeywordListUpsertRequest) Set(val *DefaultKeywordListUpsertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordListUpsertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordListUpsertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordListUpsertRequest(val *DefaultKeywordListUpsertRequest) *NullableDefaultKeywordListUpsertRequest {
	return &NullableDefaultKeywordListUpsertRequest{value: val, isSet: true}
}

func (v NullableDefaultKeywordListUpsertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordListUpsertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

