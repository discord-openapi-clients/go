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

// checks if the DefaultKeywordRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultKeywordRuleResponse{}

// DefaultKeywordRuleResponse struct for DefaultKeywordRuleResponse
type DefaultKeywordRuleResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	CreatorId string `json:"creator_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	EventType AutomodEventType `json:"event_type"`
	Actions []DefaultKeywordRuleResponseActionsInner `json:"actions"`
	TriggerType AutomodTriggerType `json:"trigger_type"`
	Enabled NullableBool `json:"enabled,omitempty"`
	ExemptRoles []string `json:"exempt_roles,omitempty"`
	ExemptChannels []string `json:"exempt_channels,omitempty"`
	TriggerMetadata DefaultKeywordListTriggerMetadataResponse `json:"trigger_metadata"`
}

type _DefaultKeywordRuleResponse DefaultKeywordRuleResponse

// NewDefaultKeywordRuleResponse instantiates a new DefaultKeywordRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultKeywordRuleResponse(id string, guildId string, creatorId string, name string, eventType AutomodEventType, actions []DefaultKeywordRuleResponseActionsInner, triggerType AutomodTriggerType, triggerMetadata DefaultKeywordListTriggerMetadataResponse) *DefaultKeywordRuleResponse {
	this := DefaultKeywordRuleResponse{}
	this.Id = id
	this.GuildId = guildId
	this.CreatorId = creatorId
	this.Name = name
	this.EventType = eventType
	this.Actions = actions
	this.TriggerType = triggerType
	this.TriggerMetadata = triggerMetadata
	return &this
}

// NewDefaultKeywordRuleResponseWithDefaults instantiates a new DefaultKeywordRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultKeywordRuleResponseWithDefaults() *DefaultKeywordRuleResponse {
	this := DefaultKeywordRuleResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DefaultKeywordRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DefaultKeywordRuleResponse) SetId(v string) {
	o.Id = v
}

// GetGuildId returns the GuildId field value
func (o *DefaultKeywordRuleResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *DefaultKeywordRuleResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetCreatorId returns the CreatorId field value
func (o *DefaultKeywordRuleResponse) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *DefaultKeywordRuleResponse) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetName returns the Name field value
func (o *DefaultKeywordRuleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DefaultKeywordRuleResponse) SetName(v string) {
	o.Name = v
}

// GetEventType returns the EventType field value
func (o *DefaultKeywordRuleResponse) GetEventType() AutomodEventType {
	if o == nil {
		var ret AutomodEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetEventTypeOk() (*AutomodEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *DefaultKeywordRuleResponse) SetEventType(v AutomodEventType) {
	o.EventType = v
}

// GetActions returns the Actions field value
func (o *DefaultKeywordRuleResponse) GetActions() []DefaultKeywordRuleResponseActionsInner {
	if o == nil {
		var ret []DefaultKeywordRuleResponseActionsInner
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetActionsOk() ([]DefaultKeywordRuleResponseActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *DefaultKeywordRuleResponse) SetActions(v []DefaultKeywordRuleResponseActionsInner) {
	o.Actions = v
}

// GetTriggerType returns the TriggerType field value
func (o *DefaultKeywordRuleResponse) GetTriggerType() AutomodTriggerType {
	if o == nil {
		var ret AutomodTriggerType
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetTriggerTypeOk() (*AutomodTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *DefaultKeywordRuleResponse) SetTriggerType(v AutomodTriggerType) {
	o.TriggerType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordRuleResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordRuleResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *DefaultKeywordRuleResponse) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *DefaultKeywordRuleResponse) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *DefaultKeywordRuleResponse) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *DefaultKeywordRuleResponse) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetExemptRoles returns the ExemptRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordRuleResponse) GetExemptRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptRoles
}

// GetExemptRolesOk returns a tuple with the ExemptRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordRuleResponse) GetExemptRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptRoles) {
		return nil, false
	}
	return o.ExemptRoles, true
}

// HasExemptRoles returns a boolean if a field has been set.
func (o *DefaultKeywordRuleResponse) HasExemptRoles() bool {
	if o != nil && !IsNil(o.ExemptRoles) {
		return true
	}

	return false
}

// SetExemptRoles gets a reference to the given []string and assigns it to the ExemptRoles field.
func (o *DefaultKeywordRuleResponse) SetExemptRoles(v []string) {
	o.ExemptRoles = v
}

// GetExemptChannels returns the ExemptChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultKeywordRuleResponse) GetExemptChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptChannels
}

// GetExemptChannelsOk returns a tuple with the ExemptChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultKeywordRuleResponse) GetExemptChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptChannels) {
		return nil, false
	}
	return o.ExemptChannels, true
}

// HasExemptChannels returns a boolean if a field has been set.
func (o *DefaultKeywordRuleResponse) HasExemptChannels() bool {
	if o != nil && !IsNil(o.ExemptChannels) {
		return true
	}

	return false
}

// SetExemptChannels gets a reference to the given []string and assigns it to the ExemptChannels field.
func (o *DefaultKeywordRuleResponse) SetExemptChannels(v []string) {
	o.ExemptChannels = v
}

// GetTriggerMetadata returns the TriggerMetadata field value
func (o *DefaultKeywordRuleResponse) GetTriggerMetadata() DefaultKeywordListTriggerMetadataResponse {
	if o == nil {
		var ret DefaultKeywordListTriggerMetadataResponse
		return ret
	}

	return o.TriggerMetadata
}

// GetTriggerMetadataOk returns a tuple with the TriggerMetadata field value
// and a boolean to check if the value has been set.
func (o *DefaultKeywordRuleResponse) GetTriggerMetadataOk() (*DefaultKeywordListTriggerMetadataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerMetadata, true
}

// SetTriggerMetadata sets field value
func (o *DefaultKeywordRuleResponse) SetTriggerMetadata(v DefaultKeywordListTriggerMetadataResponse) {
	o.TriggerMetadata = v
}

func (o DefaultKeywordRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultKeywordRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["guild_id"] = o.GuildId
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["name"] = o.Name
	toSerialize["event_type"] = o.EventType
	toSerialize["actions"] = o.Actions
	toSerialize["trigger_type"] = o.TriggerType
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.ExemptRoles != nil {
		toSerialize["exempt_roles"] = o.ExemptRoles
	}
	if o.ExemptChannels != nil {
		toSerialize["exempt_channels"] = o.ExemptChannels
	}
	toSerialize["trigger_metadata"] = o.TriggerMetadata
	return toSerialize, nil
}

func (o *DefaultKeywordRuleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"guild_id",
		"creator_id",
		"name",
		"event_type",
		"actions",
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

	varDefaultKeywordRuleResponse := _DefaultKeywordRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDefaultKeywordRuleResponse)

	if err != nil {
		return err
	}

	*o = DefaultKeywordRuleResponse(varDefaultKeywordRuleResponse)

	return err
}

type NullableDefaultKeywordRuleResponse struct {
	value *DefaultKeywordRuleResponse
	isSet bool
}

func (v NullableDefaultKeywordRuleResponse) Get() *DefaultKeywordRuleResponse {
	return v.value
}

func (v *NullableDefaultKeywordRuleResponse) Set(val *DefaultKeywordRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordRuleResponse(val *DefaultKeywordRuleResponse) *NullableDefaultKeywordRuleResponse {
	return &NullableDefaultKeywordRuleResponse{value: val, isSet: true}
}

func (v NullableDefaultKeywordRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


