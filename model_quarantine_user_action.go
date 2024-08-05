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

// checks if the QuarantineUserAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuarantineUserAction{}

// QuarantineUserAction struct for QuarantineUserAction
type QuarantineUserAction struct {
	Type AutomodActionType `json:"type"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _QuarantineUserAction QuarantineUserAction

// NewQuarantineUserAction instantiates a new QuarantineUserAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuarantineUserAction(type_ AutomodActionType) *QuarantineUserAction {
	this := QuarantineUserAction{}
	this.Type = type_
	return &this
}

// NewQuarantineUserActionWithDefaults instantiates a new QuarantineUserAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuarantineUserActionWithDefaults() *QuarantineUserAction {
	this := QuarantineUserAction{}
	return &this
}

// GetType returns the Type field value
func (o *QuarantineUserAction) GetType() AutomodActionType {
	if o == nil {
		var ret AutomodActionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *QuarantineUserAction) GetTypeOk() (*AutomodActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *QuarantineUserAction) SetType(v AutomodActionType) {
	o.Type = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *QuarantineUserAction) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuarantineUserAction) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *QuarantineUserAction) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *QuarantineUserAction) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o QuarantineUserAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuarantineUserAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *QuarantineUserAction) UnmarshalJSON(data []byte) (err error) {
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

	varQuarantineUserAction := _QuarantineUserAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuarantineUserAction)

	if err != nil {
		return err
	}

	*o = QuarantineUserAction(varQuarantineUserAction)

	return err
}

type NullableQuarantineUserAction struct {
	value *QuarantineUserAction
	isSet bool
}

func (v NullableQuarantineUserAction) Get() *QuarantineUserAction {
	return v.value
}

func (v *NullableQuarantineUserAction) Set(val *QuarantineUserAction) {
	v.value = val
	v.isSet = true
}

func (v NullableQuarantineUserAction) IsSet() bool {
	return v.isSet
}

func (v *NullableQuarantineUserAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuarantineUserAction(val *QuarantineUserAction) *NullableQuarantineUserAction {
	return &NullableQuarantineUserAction{value: val, isSet: true}
}

func (v NullableQuarantineUserAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuarantineUserAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


