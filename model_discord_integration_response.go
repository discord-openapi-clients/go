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

// checks if the DiscordIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscordIntegrationResponse{}

// DiscordIntegrationResponse struct for DiscordIntegrationResponse
type DiscordIntegrationResponse struct {
	Type IntegrationTypes `json:"type"`
	Name NullableString `json:"name,omitempty"`
	Account NullableAccountResponse `json:"account,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Application IntegrationApplicationResponse `json:"application"`
	Scopes []OAuth2Scopes `json:"scopes"`
	User NullableUserResponse `json:"user,omitempty"`
}

type _DiscordIntegrationResponse DiscordIntegrationResponse

// NewDiscordIntegrationResponse instantiates a new DiscordIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscordIntegrationResponse(type_ IntegrationTypes, id string, application IntegrationApplicationResponse, scopes []OAuth2Scopes) *DiscordIntegrationResponse {
	this := DiscordIntegrationResponse{}
	this.Type = type_
	this.Id = id
	this.Application = application
	this.Scopes = scopes
	return &this
}

// NewDiscordIntegrationResponseWithDefaults instantiates a new DiscordIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscordIntegrationResponseWithDefaults() *DiscordIntegrationResponse {
	this := DiscordIntegrationResponse{}
	return &this
}

// GetType returns the Type field value
func (o *DiscordIntegrationResponse) GetType() IntegrationTypes {
	if o == nil {
		var ret IntegrationTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiscordIntegrationResponse) SetType(v IntegrationTypes) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DiscordIntegrationResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DiscordIntegrationResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetName() {
	o.Name.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetAccount() AccountResponse {
	if o == nil || IsNil(o.Account.Get()) {
		var ret AccountResponse
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableAccountResponse and assigns it to the Account field.
func (o *DiscordIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *DiscordIntegrationResponse) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetAccount() {
	o.Account.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *DiscordIntegrationResponse) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *DiscordIntegrationResponse) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetId returns the Id field value
func (o *DiscordIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscordIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetApplication returns the Application field value
func (o *DiscordIntegrationResponse) GetApplication() IntegrationApplicationResponse {
	if o == nil {
		var ret IntegrationApplicationResponse
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetApplicationOk() (*IntegrationApplicationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *DiscordIntegrationResponse) SetApplication(v IntegrationApplicationResponse) {
	o.Application = v
}

// GetScopes returns the Scopes field value
func (o *DiscordIntegrationResponse) GetScopes() []OAuth2Scopes {
	if o == nil {
		var ret []OAuth2Scopes
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *DiscordIntegrationResponse) GetScopesOk() ([]OAuth2Scopes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *DiscordIntegrationResponse) SetScopes(v []OAuth2Scopes) {
	o.Scopes = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscordIntegrationResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscordIntegrationResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *DiscordIntegrationResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *DiscordIntegrationResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *DiscordIntegrationResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *DiscordIntegrationResponse) UnsetUser() {
	o.User.Unset()
}

func (o DiscordIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscordIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["application"] = o.Application
	toSerialize["scopes"] = o.Scopes
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *DiscordIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"application",
		"scopes",
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

	varDiscordIntegrationResponse := _DiscordIntegrationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscordIntegrationResponse)

	if err != nil {
		return err
	}

	*o = DiscordIntegrationResponse(varDiscordIntegrationResponse)

	return err
}

type NullableDiscordIntegrationResponse struct {
	value *DiscordIntegrationResponse
	isSet bool
}

func (v NullableDiscordIntegrationResponse) Get() *DiscordIntegrationResponse {
	return v.value
}

func (v *NullableDiscordIntegrationResponse) Set(val *DiscordIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscordIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscordIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscordIntegrationResponse(val *DiscordIntegrationResponse) *NullableDiscordIntegrationResponse {
	return &NullableDiscordIntegrationResponse{value: val, isSet: true}
}

func (v NullableDiscordIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscordIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


