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

// checks if the MessageMentionChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageMentionChannelResponse{}

// MessageMentionChannelResponse struct for MessageMentionChannelResponse
type MessageMentionChannelResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Type ChannelTypes `json:"type"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _MessageMentionChannelResponse MessageMentionChannelResponse

// NewMessageMentionChannelResponse instantiates a new MessageMentionChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageMentionChannelResponse(id string, name string, type_ ChannelTypes, guildId string) *MessageMentionChannelResponse {
	this := MessageMentionChannelResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.GuildId = guildId
	return &this
}

// NewMessageMentionChannelResponseWithDefaults instantiates a new MessageMentionChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageMentionChannelResponseWithDefaults() *MessageMentionChannelResponse {
	this := MessageMentionChannelResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MessageMentionChannelResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageMentionChannelResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageMentionChannelResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MessageMentionChannelResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessageMentionChannelResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessageMentionChannelResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MessageMentionChannelResponse) GetType() ChannelTypes {
	if o == nil {
		var ret ChannelTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageMentionChannelResponse) GetTypeOk() (*ChannelTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageMentionChannelResponse) SetType(v ChannelTypes) {
	o.Type = v
}

// GetGuildId returns the GuildId field value
func (o *MessageMentionChannelResponse) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *MessageMentionChannelResponse) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *MessageMentionChannelResponse) SetGuildId(v string) {
	o.GuildId = v
}

func (o MessageMentionChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageMentionChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["guild_id"] = o.GuildId
	return toSerialize, nil
}

func (o *MessageMentionChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"guild_id",
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

	varMessageMentionChannelResponse := _MessageMentionChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageMentionChannelResponse)

	if err != nil {
		return err
	}

	*o = MessageMentionChannelResponse(varMessageMentionChannelResponse)

	return err
}

type NullableMessageMentionChannelResponse struct {
	value *MessageMentionChannelResponse
	isSet bool
}

func (v NullableMessageMentionChannelResponse) Get() *MessageMentionChannelResponse {
	return v.value
}

func (v *NullableMessageMentionChannelResponse) Set(val *MessageMentionChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageMentionChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageMentionChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageMentionChannelResponse(val *MessageMentionChannelResponse) *NullableMessageMentionChannelResponse {
	return &NullableMessageMentionChannelResponse{value: val, isSet: true}
}

func (v NullableMessageMentionChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageMentionChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


