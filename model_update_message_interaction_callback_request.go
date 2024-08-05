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

// checks if the UpdateMessageInteractionCallbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMessageInteractionCallbackRequest{}

// UpdateMessageInteractionCallbackRequest struct for UpdateMessageInteractionCallbackRequest
type UpdateMessageInteractionCallbackRequest struct {
	Type InteractionCallbackTypes `json:"type"`
	Data NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial `json:"data,omitempty"`
}

type _UpdateMessageInteractionCallbackRequest UpdateMessageInteractionCallbackRequest

// NewUpdateMessageInteractionCallbackRequest instantiates a new UpdateMessageInteractionCallbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMessageInteractionCallbackRequest(type_ InteractionCallbackTypes) *UpdateMessageInteractionCallbackRequest {
	this := UpdateMessageInteractionCallbackRequest{}
	this.Type = type_
	return &this
}

// NewUpdateMessageInteractionCallbackRequestWithDefaults instantiates a new UpdateMessageInteractionCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMessageInteractionCallbackRequestWithDefaults() *UpdateMessageInteractionCallbackRequest {
	this := UpdateMessageInteractionCallbackRequest{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateMessageInteractionCallbackRequest) GetType() InteractionCallbackTypes {
	if o == nil {
		var ret InteractionCallbackTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateMessageInteractionCallbackRequest) GetTypeOk() (*InteractionCallbackTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateMessageInteractionCallbackRequest) SetType(v InteractionCallbackTypes) {
	o.Type = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMessageInteractionCallbackRequest) GetData() IncomingWebhookUpdateForInteractionCallbackRequestPartial {
	if o == nil || IsNil(o.Data.Get()) {
		var ret IncomingWebhookUpdateForInteractionCallbackRequestPartial
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMessageInteractionCallbackRequest) GetDataOk() (*IncomingWebhookUpdateForInteractionCallbackRequestPartial, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *UpdateMessageInteractionCallbackRequest) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial and assigns it to the Data field.
func (o *UpdateMessageInteractionCallbackRequest) SetData(v IncomingWebhookUpdateForInteractionCallbackRequestPartial) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *UpdateMessageInteractionCallbackRequest) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *UpdateMessageInteractionCallbackRequest) UnsetData() {
	o.Data.Unset()
}

func (o UpdateMessageInteractionCallbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMessageInteractionCallbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	return toSerialize, nil
}

func (o *UpdateMessageInteractionCallbackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varUpdateMessageInteractionCallbackRequest := _UpdateMessageInteractionCallbackRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMessageInteractionCallbackRequest)

	if err != nil {
		return err
	}

	*o = UpdateMessageInteractionCallbackRequest(varUpdateMessageInteractionCallbackRequest)

	return err
}

type NullableUpdateMessageInteractionCallbackRequest struct {
	value *UpdateMessageInteractionCallbackRequest
	isSet bool
}

func (v NullableUpdateMessageInteractionCallbackRequest) Get() *UpdateMessageInteractionCallbackRequest {
	return v.value
}

func (v *NullableUpdateMessageInteractionCallbackRequest) Set(val *UpdateMessageInteractionCallbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMessageInteractionCallbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMessageInteractionCallbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMessageInteractionCallbackRequest(val *UpdateMessageInteractionCallbackRequest) *NullableUpdateMessageInteractionCallbackRequest {
	return &NullableUpdateMessageInteractionCallbackRequest{value: val, isSet: true}
}

func (v NullableUpdateMessageInteractionCallbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMessageInteractionCallbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


