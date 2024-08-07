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

// checks if the MessageStickerItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageStickerItemResponse{}

// MessageStickerItemResponse struct for MessageStickerItemResponse
type MessageStickerItemResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	FormatType StickerFormatTypes `json:"format_type"`
}

type _MessageStickerItemResponse MessageStickerItemResponse

// NewMessageStickerItemResponse instantiates a new MessageStickerItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageStickerItemResponse(id string, name string, formatType StickerFormatTypes) *MessageStickerItemResponse {
	this := MessageStickerItemResponse{}
	this.Id = id
	this.Name = name
	this.FormatType = formatType
	return &this
}

// NewMessageStickerItemResponseWithDefaults instantiates a new MessageStickerItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageStickerItemResponseWithDefaults() *MessageStickerItemResponse {
	this := MessageStickerItemResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MessageStickerItemResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageStickerItemResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageStickerItemResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MessageStickerItemResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessageStickerItemResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessageStickerItemResponse) SetName(v string) {
	o.Name = v
}

// GetFormatType returns the FormatType field value
func (o *MessageStickerItemResponse) GetFormatType() StickerFormatTypes {
	if o == nil {
		var ret StickerFormatTypes
		return ret
	}

	return o.FormatType
}

// GetFormatTypeOk returns a tuple with the FormatType field value
// and a boolean to check if the value has been set.
func (o *MessageStickerItemResponse) GetFormatTypeOk() (*StickerFormatTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormatType, true
}

// SetFormatType sets field value
func (o *MessageStickerItemResponse) SetFormatType(v StickerFormatTypes) {
	o.FormatType = v
}

func (o MessageStickerItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageStickerItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["format_type"] = o.FormatType
	return toSerialize, nil
}

func (o *MessageStickerItemResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"format_type",
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

	varMessageStickerItemResponse := _MessageStickerItemResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageStickerItemResponse)

	if err != nil {
		return err
	}

	*o = MessageStickerItemResponse(varMessageStickerItemResponse)

	return err
}

type NullableMessageStickerItemResponse struct {
	value *MessageStickerItemResponse
	isSet bool
}

func (v NullableMessageStickerItemResponse) Get() *MessageStickerItemResponse {
	return v.value
}

func (v *NullableMessageStickerItemResponse) Set(val *MessageStickerItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageStickerItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageStickerItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageStickerItemResponse(val *MessageStickerItemResponse) *NullableMessageStickerItemResponse {
	return &NullableMessageStickerItemResponse{value: val, isSet: true}
}

func (v NullableMessageStickerItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageStickerItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


