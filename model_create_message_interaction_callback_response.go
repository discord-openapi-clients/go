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

// checks if the CreateMessageInteractionCallbackResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessageInteractionCallbackResponse{}

// CreateMessageInteractionCallbackResponse struct for CreateMessageInteractionCallbackResponse
type CreateMessageInteractionCallbackResponse struct {
	Type InteractionCallbackTypes `json:"type"`
	Message MessageResponse `json:"message"`
}

type _CreateMessageInteractionCallbackResponse CreateMessageInteractionCallbackResponse

// NewCreateMessageInteractionCallbackResponse instantiates a new CreateMessageInteractionCallbackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageInteractionCallbackResponse(type_ InteractionCallbackTypes, message MessageResponse) *CreateMessageInteractionCallbackResponse {
	this := CreateMessageInteractionCallbackResponse{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewCreateMessageInteractionCallbackResponseWithDefaults instantiates a new CreateMessageInteractionCallbackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageInteractionCallbackResponseWithDefaults() *CreateMessageInteractionCallbackResponse {
	this := CreateMessageInteractionCallbackResponse{}
	return &this
}

// GetType returns the Type field value
func (o *CreateMessageInteractionCallbackResponse) GetType() InteractionCallbackTypes {
	if o == nil {
		var ret InteractionCallbackTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateMessageInteractionCallbackResponse) GetTypeOk() (*InteractionCallbackTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateMessageInteractionCallbackResponse) SetType(v InteractionCallbackTypes) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *CreateMessageInteractionCallbackResponse) GetMessage() MessageResponse {
	if o == nil {
		var ret MessageResponse
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CreateMessageInteractionCallbackResponse) GetMessageOk() (*MessageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CreateMessageInteractionCallbackResponse) SetMessage(v MessageResponse) {
	o.Message = v
}

func (o CreateMessageInteractionCallbackResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageInteractionCallbackResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *CreateMessageInteractionCallbackResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateMessageInteractionCallbackResponse := _CreateMessageInteractionCallbackResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMessageInteractionCallbackResponse)

	if err != nil {
		return err
	}

	*o = CreateMessageInteractionCallbackResponse(varCreateMessageInteractionCallbackResponse)

	return err
}

type NullableCreateMessageInteractionCallbackResponse struct {
	value *CreateMessageInteractionCallbackResponse
	isSet bool
}

func (v NullableCreateMessageInteractionCallbackResponse) Get() *CreateMessageInteractionCallbackResponse {
	return v.value
}

func (v *NullableCreateMessageInteractionCallbackResponse) Set(val *CreateMessageInteractionCallbackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageInteractionCallbackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageInteractionCallbackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageInteractionCallbackResponse(val *CreateMessageInteractionCallbackResponse) *NullableCreateMessageInteractionCallbackResponse {
	return &NullableCreateMessageInteractionCallbackResponse{value: val, isSet: true}
}

func (v NullableCreateMessageInteractionCallbackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageInteractionCallbackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


