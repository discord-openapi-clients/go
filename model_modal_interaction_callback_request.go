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

// checks if the ModalInteractionCallbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModalInteractionCallbackRequest{}

// ModalInteractionCallbackRequest struct for ModalInteractionCallbackRequest
type ModalInteractionCallbackRequest struct {
	Type InteractionCallbackTypes `json:"type"`
	Data ModalInteractionCallbackData `json:"data"`
}

type _ModalInteractionCallbackRequest ModalInteractionCallbackRequest

// NewModalInteractionCallbackRequest instantiates a new ModalInteractionCallbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModalInteractionCallbackRequest(type_ InteractionCallbackTypes, data ModalInteractionCallbackData) *ModalInteractionCallbackRequest {
	this := ModalInteractionCallbackRequest{}
	this.Type = type_
	this.Data = data
	return &this
}

// NewModalInteractionCallbackRequestWithDefaults instantiates a new ModalInteractionCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModalInteractionCallbackRequestWithDefaults() *ModalInteractionCallbackRequest {
	this := ModalInteractionCallbackRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ModalInteractionCallbackRequest) GetType() InteractionCallbackTypes {
	if o == nil {
		var ret InteractionCallbackTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModalInteractionCallbackRequest) GetTypeOk() (*InteractionCallbackTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ModalInteractionCallbackRequest) SetType(v InteractionCallbackTypes) {
	o.Type = v
}

// GetData returns the Data field value
func (o *ModalInteractionCallbackRequest) GetData() ModalInteractionCallbackData {
	if o == nil {
		var ret ModalInteractionCallbackData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ModalInteractionCallbackRequest) GetDataOk() (*ModalInteractionCallbackData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ModalInteractionCallbackRequest) SetData(v ModalInteractionCallbackData) {
	o.Data = v
}

func (o ModalInteractionCallbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModalInteractionCallbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ModalInteractionCallbackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"data",
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

	varModalInteractionCallbackRequest := _ModalInteractionCallbackRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModalInteractionCallbackRequest)

	if err != nil {
		return err
	}

	*o = ModalInteractionCallbackRequest(varModalInteractionCallbackRequest)

	return err
}

type NullableModalInteractionCallbackRequest struct {
	value *ModalInteractionCallbackRequest
	isSet bool
}

func (v NullableModalInteractionCallbackRequest) Get() *ModalInteractionCallbackRequest {
	return v.value
}

func (v *NullableModalInteractionCallbackRequest) Set(val *ModalInteractionCallbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModalInteractionCallbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModalInteractionCallbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModalInteractionCallbackRequest(val *ModalInteractionCallbackRequest) *NullableModalInteractionCallbackRequest {
	return &NullableModalInteractionCallbackRequest{value: val, isSet: true}
}

func (v NullableModalInteractionCallbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModalInteractionCallbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


