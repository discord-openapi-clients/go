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

// checks if the UpdateMessageInteractionCallbackResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMessageInteractionCallbackResponse{}

// UpdateMessageInteractionCallbackResponse struct for UpdateMessageInteractionCallbackResponse
type UpdateMessageInteractionCallbackResponse struct {
	Type InteractionCallbackTypes `json:"type"`
	Message MessageResponse `json:"message"`
}

type _UpdateMessageInteractionCallbackResponse UpdateMessageInteractionCallbackResponse

// NewUpdateMessageInteractionCallbackResponse instantiates a new UpdateMessageInteractionCallbackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMessageInteractionCallbackResponse(type_ InteractionCallbackTypes, message MessageResponse) *UpdateMessageInteractionCallbackResponse {
	this := UpdateMessageInteractionCallbackResponse{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewUpdateMessageInteractionCallbackResponseWithDefaults instantiates a new UpdateMessageInteractionCallbackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMessageInteractionCallbackResponseWithDefaults() *UpdateMessageInteractionCallbackResponse {
	this := UpdateMessageInteractionCallbackResponse{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateMessageInteractionCallbackResponse) GetType() InteractionCallbackTypes {
	if o == nil {
		var ret InteractionCallbackTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateMessageInteractionCallbackResponse) GetTypeOk() (*InteractionCallbackTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateMessageInteractionCallbackResponse) SetType(v InteractionCallbackTypes) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *UpdateMessageInteractionCallbackResponse) GetMessage() MessageResponse {
	if o == nil {
		var ret MessageResponse
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UpdateMessageInteractionCallbackResponse) GetMessageOk() (*MessageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UpdateMessageInteractionCallbackResponse) SetMessage(v MessageResponse) {
	o.Message = v
}

func (o UpdateMessageInteractionCallbackResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMessageInteractionCallbackResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *UpdateMessageInteractionCallbackResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"message",
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

	varUpdateMessageInteractionCallbackResponse := _UpdateMessageInteractionCallbackResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMessageInteractionCallbackResponse)

	if err != nil {
		return err
	}

	*o = UpdateMessageInteractionCallbackResponse(varUpdateMessageInteractionCallbackResponse)

	return err
}

type NullableUpdateMessageInteractionCallbackResponse struct {
	value *UpdateMessageInteractionCallbackResponse
	isSet bool
}

func (v NullableUpdateMessageInteractionCallbackResponse) Get() *UpdateMessageInteractionCallbackResponse {
	return v.value
}

func (v *NullableUpdateMessageInteractionCallbackResponse) Set(val *UpdateMessageInteractionCallbackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMessageInteractionCallbackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMessageInteractionCallbackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMessageInteractionCallbackResponse(val *UpdateMessageInteractionCallbackResponse) *NullableUpdateMessageInteractionCallbackResponse {
	return &NullableUpdateMessageInteractionCallbackResponse{value: val, isSet: true}
}

func (v NullableUpdateMessageInteractionCallbackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMessageInteractionCallbackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

