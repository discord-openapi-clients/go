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

// checks if the MessageReferenceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReferenceResponse{}

// MessageReferenceResponse struct for MessageReferenceResponse
type MessageReferenceResponse struct {
	Type NullableMessageReferenceType `json:"type,omitempty"`
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	MessageId *string `json:"message_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _MessageReferenceResponse MessageReferenceResponse

// NewMessageReferenceResponse instantiates a new MessageReferenceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReferenceResponse(channelId string) *MessageReferenceResponse {
	this := MessageReferenceResponse{}
	this.ChannelId = channelId
	return &this
}

// NewMessageReferenceResponseWithDefaults instantiates a new MessageReferenceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReferenceResponseWithDefaults() *MessageReferenceResponse {
	this := MessageReferenceResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageReferenceResponse) GetType() MessageReferenceType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret MessageReferenceType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageReferenceResponse) GetTypeOk() (*MessageReferenceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MessageReferenceResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableMessageReferenceType and assigns it to the Type field.
func (o *MessageReferenceResponse) SetType(v MessageReferenceType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *MessageReferenceResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *MessageReferenceResponse) UnsetType() {
	o.Type.Unset()
}

// GetChannelId returns the ChannelId field value
func (o *MessageReferenceResponse) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *MessageReferenceResponse) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *MessageReferenceResponse) SetChannelId(v string) {
	o.ChannelId = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *MessageReferenceResponse) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReferenceResponse) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *MessageReferenceResponse) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *MessageReferenceResponse) SetMessageId(v string) {
	o.MessageId = &v
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *MessageReferenceResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReferenceResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *MessageReferenceResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *MessageReferenceResponse) SetGuildId(v string) {
	o.GuildId = &v
}

func (o MessageReferenceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReferenceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	toSerialize["channel_id"] = o.ChannelId
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	return toSerialize, nil
}

func (o *MessageReferenceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel_id",
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

	varMessageReferenceResponse := _MessageReferenceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageReferenceResponse)

	if err != nil {
		return err
	}

	*o = MessageReferenceResponse(varMessageReferenceResponse)

	return err
}

type NullableMessageReferenceResponse struct {
	value *MessageReferenceResponse
	isSet bool
}

func (v NullableMessageReferenceResponse) Get() *MessageReferenceResponse {
	return v.value
}

func (v *NullableMessageReferenceResponse) Set(val *MessageReferenceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReferenceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReferenceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReferenceResponse(val *MessageReferenceResponse) *NullableMessageReferenceResponse {
	return &NullableMessageReferenceResponse{value: val, isSet: true}
}

func (v NullableMessageReferenceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReferenceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

