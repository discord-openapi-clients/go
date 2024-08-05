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

// checks if the SpamLinkRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpamLinkRuleResponse{}

// SpamLinkRuleResponse struct for SpamLinkRuleResponse
type SpamLinkRuleResponse struct {
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
	TriggerMetadata map[string]interface{} `json:"trigger_metadata"`
}

type _SpamLinkRuleResponse SpamLinkRuleResponse

// NewSpamLinkRuleResponse instantiates a new SpamLinkRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpamLinkRuleResponse(id string, guildId string, creatorId string, name string, eventType AutomodEventType, actions []DefaultKeywordRuleResponseActionsInner, triggerType AutomodTriggerType, triggerMetadata map[string]interface{}) *SpamLinkRuleResponse {
	this := SpamLinkRuleResponse{}
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

// NewSpamLinkRuleResponseWithDefaults instantiates a new SpamLinkRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpamLinkRuleResponseWithDefaults() *SpamLinkRuleResponse {
	this := SpamLinkRuleResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SpamLinkRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SpamLinkRuleResponse) SetId(v string) {
	o.Id = v
}

// GetGuildId returns the GuildId field value
func (o *SpamLinkRuleResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *SpamLinkRuleResponse) SetGuildId(v string) {
	o.GuildId = v
}

// GetCreatorId returns the CreatorId field value
func (o *SpamLinkRuleResponse) GetCreatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetCreatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *SpamLinkRuleResponse) SetCreatorId(v string) {
	o.CreatorId = v
}

// GetName returns the Name field value
func (o *SpamLinkRuleResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpamLinkRuleResponse) SetName(v string) {
	o.Name = v
}

// GetEventType returns the EventType field value
func (o *SpamLinkRuleResponse) GetEventType() AutomodEventType {
	if o == nil {
		var ret AutomodEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetEventTypeOk() (*AutomodEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *SpamLinkRuleResponse) SetEventType(v AutomodEventType) {
	o.EventType = v
}

// GetActions returns the Actions field value
func (o *SpamLinkRuleResponse) GetActions() []DefaultKeywordRuleResponseActionsInner {
	if o == nil {
		var ret []DefaultKeywordRuleResponseActionsInner
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetActionsOk() ([]DefaultKeywordRuleResponseActionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *SpamLinkRuleResponse) SetActions(v []DefaultKeywordRuleResponseActionsInner) {
	o.Actions = v
}

// GetTriggerType returns the TriggerType field value
func (o *SpamLinkRuleResponse) GetTriggerType() AutomodTriggerType {
	if o == nil {
		var ret AutomodTriggerType
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetTriggerTypeOk() (*AutomodTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *SpamLinkRuleResponse) SetTriggerType(v AutomodTriggerType) {
	o.TriggerType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpamLinkRuleResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpamLinkRuleResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *SpamLinkRuleResponse) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *SpamLinkRuleResponse) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *SpamLinkRuleResponse) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *SpamLinkRuleResponse) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetExemptRoles returns the ExemptRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpamLinkRuleResponse) GetExemptRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptRoles
}

// GetExemptRolesOk returns a tuple with the ExemptRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpamLinkRuleResponse) GetExemptRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptRoles) {
		return nil, false
	}
	return o.ExemptRoles, true
}

// HasExemptRoles returns a boolean if a field has been set.
func (o *SpamLinkRuleResponse) HasExemptRoles() bool {
	if o != nil && !IsNil(o.ExemptRoles) {
		return true
	}

	return false
}

// SetExemptRoles gets a reference to the given []string and assigns it to the ExemptRoles field.
func (o *SpamLinkRuleResponse) SetExemptRoles(v []string) {
	o.ExemptRoles = v
}

// GetExemptChannels returns the ExemptChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpamLinkRuleResponse) GetExemptChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExemptChannels
}

// GetExemptChannelsOk returns a tuple with the ExemptChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpamLinkRuleResponse) GetExemptChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExemptChannels) {
		return nil, false
	}
	return o.ExemptChannels, true
}

// HasExemptChannels returns a boolean if a field has been set.
func (o *SpamLinkRuleResponse) HasExemptChannels() bool {
	if o != nil && !IsNil(o.ExemptChannels) {
		return true
	}

	return false
}

// SetExemptChannels gets a reference to the given []string and assigns it to the ExemptChannels field.
func (o *SpamLinkRuleResponse) SetExemptChannels(v []string) {
	o.ExemptChannels = v
}

// GetTriggerMetadata returns the TriggerMetadata field value
func (o *SpamLinkRuleResponse) GetTriggerMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.TriggerMetadata
}

// GetTriggerMetadataOk returns a tuple with the TriggerMetadata field value
// and a boolean to check if the value has been set.
func (o *SpamLinkRuleResponse) GetTriggerMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.TriggerMetadata, true
}

// SetTriggerMetadata sets field value
func (o *SpamLinkRuleResponse) SetTriggerMetadata(v map[string]interface{}) {
	o.TriggerMetadata = v
}

func (o SpamLinkRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpamLinkRuleResponse) ToMap() (map[string]interface{}, error) {
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

func (o *SpamLinkRuleResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSpamLinkRuleResponse := _SpamLinkRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpamLinkRuleResponse)

	if err != nil {
		return err
	}

	*o = SpamLinkRuleResponse(varSpamLinkRuleResponse)

	return err
}

type NullableSpamLinkRuleResponse struct {
	value *SpamLinkRuleResponse
	isSet bool
}

func (v NullableSpamLinkRuleResponse) Get() *SpamLinkRuleResponse {
	return v.value
}

func (v *NullableSpamLinkRuleResponse) Set(val *SpamLinkRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpamLinkRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpamLinkRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpamLinkRuleResponse(val *SpamLinkRuleResponse) *NullableSpamLinkRuleResponse {
	return &NullableSpamLinkRuleResponse{value: val, isSet: true}
}

func (v NullableSpamLinkRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpamLinkRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


