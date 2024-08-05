/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the IncomingWebhookUpdateForInteractionCallbackRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomingWebhookUpdateForInteractionCallbackRequestPartial{}

// IncomingWebhookUpdateForInteractionCallbackRequestPartial struct for IncomingWebhookUpdateForInteractionCallbackRequestPartial
type IncomingWebhookUpdateForInteractionCallbackRequestPartial struct {
	Content NullableString `json:"content,omitempty"`
	Embeds []RichEmbed `json:"embeds,omitempty"`
	AllowedMentions NullableMessageAllowedMentionsRequest `json:"allowed_mentions,omitempty"`
	Components []ActionRow `json:"components,omitempty"`
	Attachments []MessageAttachmentRequest `json:"attachments,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
}

// NewIncomingWebhookUpdateForInteractionCallbackRequestPartial instantiates a new IncomingWebhookUpdateForInteractionCallbackRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomingWebhookUpdateForInteractionCallbackRequestPartial() *IncomingWebhookUpdateForInteractionCallbackRequestPartial {
	this := IncomingWebhookUpdateForInteractionCallbackRequestPartial{}
	return &this
}

// NewIncomingWebhookUpdateForInteractionCallbackRequestPartialWithDefaults instantiates a new IncomingWebhookUpdateForInteractionCallbackRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomingWebhookUpdateForInteractionCallbackRequestPartialWithDefaults() *IncomingWebhookUpdateForInteractionCallbackRequestPartial {
	this := IncomingWebhookUpdateForInteractionCallbackRequestPartial{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) UnsetContent() {
	o.Content.Unset()
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetEmbeds() []RichEmbed {
	if o == nil {
		var ret []RichEmbed
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetEmbedsOk() ([]RichEmbed, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []RichEmbed and assigns it to the Embeds field.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetEmbeds(v []RichEmbed) {
	o.Embeds = v
}

// GetAllowedMentions returns the AllowedMentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetAllowedMentions() MessageAllowedMentionsRequest {
	if o == nil || IsNil(o.AllowedMentions.Get()) {
		var ret MessageAllowedMentionsRequest
		return ret
	}
	return *o.AllowedMentions.Get()
}

// GetAllowedMentionsOk returns a tuple with the AllowedMentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedMentions.Get(), o.AllowedMentions.IsSet()
}

// HasAllowedMentions returns a boolean if a field has been set.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) HasAllowedMentions() bool {
	if o != nil && o.AllowedMentions.IsSet() {
		return true
	}

	return false
}

// SetAllowedMentions gets a reference to the given NullableMessageAllowedMentionsRequest and assigns it to the AllowedMentions field.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetAllowedMentions(v MessageAllowedMentionsRequest) {
	o.AllowedMentions.Set(&v)
}
// SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetAllowedMentionsNil() {
	o.AllowedMentions.Set(nil)
}

// UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) UnsetAllowedMentions() {
	o.AllowedMentions.Unset()
}

// GetComponents returns the Components field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetComponents() []ActionRow {
	if o == nil {
		var ret []ActionRow
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetComponentsOk() ([]ActionRow, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ActionRow and assigns it to the Components field.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetComponents(v []ActionRow) {
	o.Components = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetAttachments() []MessageAttachmentRequest {
	if o == nil {
		var ret []MessageAttachmentRequest
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetAttachmentsOk() ([]MessageAttachmentRequest, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []MessageAttachmentRequest and assigns it to the Attachments field.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetAttachments(v []MessageAttachmentRequest) {
	o.Attachments = v
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetFlags(v int32) {
	o.Flags.Set(&v)
}
// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *IncomingWebhookUpdateForInteractionCallbackRequestPartial) UnsetFlags() {
	o.Flags.Unset()
}

func (o IncomingWebhookUpdateForInteractionCallbackRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomingWebhookUpdateForInteractionCallbackRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.Embeds != nil {
		toSerialize["embeds"] = o.Embeds
	}
	if o.AllowedMentions.IsSet() {
		toSerialize["allowed_mentions"] = o.AllowedMentions.Get()
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	return toSerialize, nil
}

type NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial struct {
	value *IncomingWebhookUpdateForInteractionCallbackRequestPartial
	isSet bool
}

func (v NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial) Get() *IncomingWebhookUpdateForInteractionCallbackRequestPartial {
	return v.value
}

func (v *NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial) Set(val *IncomingWebhookUpdateForInteractionCallbackRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomingWebhookUpdateForInteractionCallbackRequestPartial(val *IncomingWebhookUpdateForInteractionCallbackRequestPartial) *NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial {
	return &NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial{value: val, isSet: true}
}

func (v NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

