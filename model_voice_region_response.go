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

// checks if the VoiceRegionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoiceRegionResponse{}

// VoiceRegionResponse struct for VoiceRegionResponse
type VoiceRegionResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Custom bool `json:"custom"`
	Deprecated bool `json:"deprecated"`
	Optimal bool `json:"optimal"`
}

type _VoiceRegionResponse VoiceRegionResponse

// NewVoiceRegionResponse instantiates a new VoiceRegionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceRegionResponse(id string, name string, custom bool, deprecated bool, optimal bool) *VoiceRegionResponse {
	this := VoiceRegionResponse{}
	this.Id = id
	this.Name = name
	this.Custom = custom
	this.Deprecated = deprecated
	this.Optimal = optimal
	return &this
}

// NewVoiceRegionResponseWithDefaults instantiates a new VoiceRegionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceRegionResponseWithDefaults() *VoiceRegionResponse {
	this := VoiceRegionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *VoiceRegionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VoiceRegionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VoiceRegionResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *VoiceRegionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VoiceRegionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VoiceRegionResponse) SetName(v string) {
	o.Name = v
}

// GetCustom returns the Custom field value
func (o *VoiceRegionResponse) GetCustom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *VoiceRegionResponse) GetCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *VoiceRegionResponse) SetCustom(v bool) {
	o.Custom = v
}

// GetDeprecated returns the Deprecated field value
func (o *VoiceRegionResponse) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *VoiceRegionResponse) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *VoiceRegionResponse) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetOptimal returns the Optimal field value
func (o *VoiceRegionResponse) GetOptimal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Optimal
}

// GetOptimalOk returns a tuple with the Optimal field value
// and a boolean to check if the value has been set.
func (o *VoiceRegionResponse) GetOptimalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Optimal, true
}

// SetOptimal sets field value
func (o *VoiceRegionResponse) SetOptimal(v bool) {
	o.Optimal = v
}

func (o VoiceRegionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceRegionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["custom"] = o.Custom
	toSerialize["deprecated"] = o.Deprecated
	toSerialize["optimal"] = o.Optimal
	return toSerialize, nil
}

func (o *VoiceRegionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"custom",
		"deprecated",
		"optimal",
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

	varVoiceRegionResponse := _VoiceRegionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVoiceRegionResponse)

	if err != nil {
		return err
	}

	*o = VoiceRegionResponse(varVoiceRegionResponse)

	return err
}

type NullableVoiceRegionResponse struct {
	value *VoiceRegionResponse
	isSet bool
}

func (v NullableVoiceRegionResponse) Get() *VoiceRegionResponse {
	return v.value
}

func (v *NullableVoiceRegionResponse) Set(val *VoiceRegionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceRegionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceRegionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceRegionResponse(val *VoiceRegionResponse) *NullableVoiceRegionResponse {
	return &NullableVoiceRegionResponse{value: val, isSet: true}
}

func (v NullableVoiceRegionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceRegionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


