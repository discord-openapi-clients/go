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

// checks if the BaseCreateMessageCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseCreateMessageCreateRequest{}

// BaseCreateMessageCreateRequest struct for BaseCreateMessageCreateRequest
type BaseCreateMessageCreateRequest struct {
	Content NullableString `json:"content,omitempty"`
	Embeds []RichEmbed `json:"embeds,omitempty"`
	AllowedMentions NullableMessageAllowedMentionsRequest `json:"allowed_mentions,omitempty"`
	StickerIds []string `json:"sticker_ids,omitempty"`
	Components []ActionRow `json:"components,omitempty"`
	Flags NullableInt32 `json:"flags,omitempty"`
	Attachments []MessageAttachmentRequest `json:"attachments,omitempty"`
	Poll NullablePollCreateRequest `json:"poll,omitempty"`
}

// NewBaseCreateMessageCreateRequest instantiates a new BaseCreateMessageCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCreateMessageCreateRequest() *BaseCreateMessageCreateRequest {
	this := BaseCreateMessageCreateRequest{}
	return &this
}

// NewBaseCreateMessageCreateRequestWithDefaults instantiates a new BaseCreateMessageCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCreateMessageCreateRequestWithDefaults() *BaseCreateMessageCreateRequest {
	this := BaseCreateMessageCreateRequest{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *BaseCreateMessageCreateRequest) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *BaseCreateMessageCreateRequest) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *BaseCreateMessageCreateRequest) UnsetContent() {
	o.Content.Unset()
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetEmbeds() []RichEmbed {
	if o == nil {
		var ret []RichEmbed
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetEmbedsOk() ([]RichEmbed, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []RichEmbed and assigns it to the Embeds field.
func (o *BaseCreateMessageCreateRequest) SetEmbeds(v []RichEmbed) {
	o.Embeds = v
}

// GetAllowedMentions returns the AllowedMentions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetAllowedMentions() MessageAllowedMentionsRequest {
	if o == nil || IsNil(o.AllowedMentions.Get()) {
		var ret MessageAllowedMentionsRequest
		return ret
	}
	return *o.AllowedMentions.Get()
}

// GetAllowedMentionsOk returns a tuple with the AllowedMentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetAllowedMentionsOk() (*MessageAllowedMentionsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedMentions.Get(), o.AllowedMentions.IsSet()
}

// HasAllowedMentions returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasAllowedMentions() bool {
	if o != nil && o.AllowedMentions.IsSet() {
		return true
	}

	return false
}

// SetAllowedMentions gets a reference to the given NullableMessageAllowedMentionsRequest and assigns it to the AllowedMentions field.
func (o *BaseCreateMessageCreateRequest) SetAllowedMentions(v MessageAllowedMentionsRequest) {
	o.AllowedMentions.Set(&v)
}
// SetAllowedMentionsNil sets the value for AllowedMentions to be an explicit nil
func (o *BaseCreateMessageCreateRequest) SetAllowedMentionsNil() {
	o.AllowedMentions.Set(nil)
}

// UnsetAllowedMentions ensures that no value is present for AllowedMentions, not even an explicit nil
func (o *BaseCreateMessageCreateRequest) UnsetAllowedMentions() {
	o.AllowedMentions.Unset()
}

// GetStickerIds returns the StickerIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetStickerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.StickerIds
}

// GetStickerIdsOk returns a tuple with the StickerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetStickerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StickerIds) {
		return nil, false
	}
	return o.StickerIds, true
}

// HasStickerIds returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasStickerIds() bool {
	if o != nil && !IsNil(o.StickerIds) {
		return true
	}

	return false
}

// SetStickerIds gets a reference to the given []string and assigns it to the StickerIds field.
func (o *BaseCreateMessageCreateRequest) SetStickerIds(v []string) {
	o.StickerIds = v
}

// GetComponents returns the Components field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetComponents() []ActionRow {
	if o == nil {
		var ret []ActionRow
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetComponentsOk() ([]ActionRow, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ActionRow and assigns it to the Components field.
func (o *BaseCreateMessageCreateRequest) SetComponents(v []ActionRow) {
	o.Components = v
}

// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *BaseCreateMessageCreateRequest) SetFlags(v int32) {
	o.Flags.Set(&v)
}
// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *BaseCreateMessageCreateRequest) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *BaseCreateMessageCreateRequest) UnsetFlags() {
	o.Flags.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetAttachments() []MessageAttachmentRequest {
	if o == nil {
		var ret []MessageAttachmentRequest
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetAttachmentsOk() ([]MessageAttachmentRequest, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []MessageAttachmentRequest and assigns it to the Attachments field.
func (o *BaseCreateMessageCreateRequest) SetAttachments(v []MessageAttachmentRequest) {
	o.Attachments = v
}

// GetPoll returns the Poll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseCreateMessageCreateRequest) GetPoll() PollCreateRequest {
	if o == nil || IsNil(o.Poll.Get()) {
		var ret PollCreateRequest
		return ret
	}
	return *o.Poll.Get()
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseCreateMessageCreateRequest) GetPollOk() (*PollCreateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Poll.Get(), o.Poll.IsSet()
}

// HasPoll returns a boolean if a field has been set.
func (o *BaseCreateMessageCreateRequest) HasPoll() bool {
	if o != nil && o.Poll.IsSet() {
		return true
	}

	return false
}

// SetPoll gets a reference to the given NullablePollCreateRequest and assigns it to the Poll field.
func (o *BaseCreateMessageCreateRequest) SetPoll(v PollCreateRequest) {
	o.Poll.Set(&v)
}
// SetPollNil sets the value for Poll to be an explicit nil
func (o *BaseCreateMessageCreateRequest) SetPollNil() {
	o.Poll.Set(nil)
}

// UnsetPoll ensures that no value is present for Poll, not even an explicit nil
func (o *BaseCreateMessageCreateRequest) UnsetPoll() {
	o.Poll.Unset()
}

func (o BaseCreateMessageCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseCreateMessageCreateRequest) ToMap() (map[string]interface{}, error) {
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
	if o.StickerIds != nil {
		toSerialize["sticker_ids"] = o.StickerIds
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Poll.IsSet() {
		toSerialize["poll"] = o.Poll.Get()
	}
	return toSerialize, nil
}

type NullableBaseCreateMessageCreateRequest struct {
	value *BaseCreateMessageCreateRequest
	isSet bool
}

func (v NullableBaseCreateMessageCreateRequest) Get() *BaseCreateMessageCreateRequest {
	return v.value
}

func (v *NullableBaseCreateMessageCreateRequest) Set(val *BaseCreateMessageCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCreateMessageCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCreateMessageCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCreateMessageCreateRequest(val *BaseCreateMessageCreateRequest) *NullableBaseCreateMessageCreateRequest {
	return &NullableBaseCreateMessageCreateRequest{value: val, isSet: true}
}

func (v NullableBaseCreateMessageCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCreateMessageCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


