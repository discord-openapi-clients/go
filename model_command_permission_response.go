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

// checks if the CommandPermissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandPermissionResponse{}

// CommandPermissionResponse struct for CommandPermissionResponse
type CommandPermissionResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type ApplicationCommandPermissionType `json:"type"`
	Permission bool `json:"permission"`
}

type _CommandPermissionResponse CommandPermissionResponse

// NewCommandPermissionResponse instantiates a new CommandPermissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandPermissionResponse(id string, type_ ApplicationCommandPermissionType, permission bool) *CommandPermissionResponse {
	this := CommandPermissionResponse{}
	this.Id = id
	this.Type = type_
	this.Permission = permission
	return &this
}

// NewCommandPermissionResponseWithDefaults instantiates a new CommandPermissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandPermissionResponseWithDefaults() *CommandPermissionResponse {
	this := CommandPermissionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CommandPermissionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommandPermissionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommandPermissionResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *CommandPermissionResponse) GetType() ApplicationCommandPermissionType {
	if o == nil {
		var ret ApplicationCommandPermissionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CommandPermissionResponse) GetTypeOk() (*ApplicationCommandPermissionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CommandPermissionResponse) SetType(v ApplicationCommandPermissionType) {
	o.Type = v
}

// GetPermission returns the Permission field value
func (o *CommandPermissionResponse) GetPermission() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *CommandPermissionResponse) GetPermissionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *CommandPermissionResponse) SetPermission(v bool) {
	o.Permission = v
}

func (o CommandPermissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandPermissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["permission"] = o.Permission
	return toSerialize, nil
}

func (o *CommandPermissionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"permission",
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

	varCommandPermissionResponse := _CommandPermissionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommandPermissionResponse)

	if err != nil {
		return err
	}

	*o = CommandPermissionResponse(varCommandPermissionResponse)

	return err
}

type NullableCommandPermissionResponse struct {
	value *CommandPermissionResponse
	isSet bool
}

func (v NullableCommandPermissionResponse) Get() *CommandPermissionResponse {
	return v.value
}

func (v *NullableCommandPermissionResponse) Set(val *CommandPermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandPermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandPermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandPermissionResponse(val *CommandPermissionResponse) *NullableCommandPermissionResponse {
	return &NullableCommandPermissionResponse{value: val, isSet: true}
}

func (v NullableCommandPermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandPermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


