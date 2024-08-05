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

// checks if the PartialGuildSubscriptionIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialGuildSubscriptionIntegrationResponse{}

// PartialGuildSubscriptionIntegrationResponse struct for PartialGuildSubscriptionIntegrationResponse
type PartialGuildSubscriptionIntegrationResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Type IntegrationTypes `json:"type"`
	Name NullableString `json:"name,omitempty"`
	Account NullableAccountResponse `json:"account,omitempty"`
}

type _PartialGuildSubscriptionIntegrationResponse PartialGuildSubscriptionIntegrationResponse

// NewPartialGuildSubscriptionIntegrationResponse instantiates a new PartialGuildSubscriptionIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialGuildSubscriptionIntegrationResponse(id string, type_ IntegrationTypes) *PartialGuildSubscriptionIntegrationResponse {
	this := PartialGuildSubscriptionIntegrationResponse{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewPartialGuildSubscriptionIntegrationResponseWithDefaults instantiates a new PartialGuildSubscriptionIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialGuildSubscriptionIntegrationResponseWithDefaults() *PartialGuildSubscriptionIntegrationResponse {
	this := PartialGuildSubscriptionIntegrationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PartialGuildSubscriptionIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartialGuildSubscriptionIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartialGuildSubscriptionIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *PartialGuildSubscriptionIntegrationResponse) GetType() IntegrationTypes {
	if o == nil {
		var ret IntegrationTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PartialGuildSubscriptionIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PartialGuildSubscriptionIntegrationResponse) SetType(v IntegrationTypes) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialGuildSubscriptionIntegrationResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGuildSubscriptionIntegrationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PartialGuildSubscriptionIntegrationResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PartialGuildSubscriptionIntegrationResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) UnsetName() {
	o.Name.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialGuildSubscriptionIntegrationResponse) GetAccount() AccountResponse {
	if o == nil || IsNil(o.Account.Get()) {
		var ret AccountResponse
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGuildSubscriptionIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *PartialGuildSubscriptionIntegrationResponse) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableAccountResponse and assigns it to the Account field.
func (o *PartialGuildSubscriptionIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *PartialGuildSubscriptionIntegrationResponse) UnsetAccount() {
	o.Account.Unset()
}

func (o PartialGuildSubscriptionIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialGuildSubscriptionIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	return toSerialize, nil
}

func (o *PartialGuildSubscriptionIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPartialGuildSubscriptionIntegrationResponse := _PartialGuildSubscriptionIntegrationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPartialGuildSubscriptionIntegrationResponse)

	if err != nil {
		return err
	}

	*o = PartialGuildSubscriptionIntegrationResponse(varPartialGuildSubscriptionIntegrationResponse)

	return err
}

type NullablePartialGuildSubscriptionIntegrationResponse struct {
	value *PartialGuildSubscriptionIntegrationResponse
	isSet bool
}

func (v NullablePartialGuildSubscriptionIntegrationResponse) Get() *PartialGuildSubscriptionIntegrationResponse {
	return v.value
}

func (v *NullablePartialGuildSubscriptionIntegrationResponse) Set(val *PartialGuildSubscriptionIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialGuildSubscriptionIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialGuildSubscriptionIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialGuildSubscriptionIntegrationResponse(val *PartialGuildSubscriptionIntegrationResponse) *NullablePartialGuildSubscriptionIntegrationResponse {
	return &NullablePartialGuildSubscriptionIntegrationResponse{value: val, isSet: true}
}

func (v NullablePartialGuildSubscriptionIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialGuildSubscriptionIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


