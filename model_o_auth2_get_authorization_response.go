/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the OAuth2GetAuthorizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2GetAuthorizationResponse{}

// OAuth2GetAuthorizationResponse struct for OAuth2GetAuthorizationResponse
type OAuth2GetAuthorizationResponse struct {
	Application ApplicationResponse `json:"application"`
	Expires time.Time `json:"expires"`
	Scopes []OAuth2Scopes `json:"scopes"`
	User NullableUserResponse `json:"user,omitempty"`
}

type _OAuth2GetAuthorizationResponse OAuth2GetAuthorizationResponse

// NewOAuth2GetAuthorizationResponse instantiates a new OAuth2GetAuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2GetAuthorizationResponse(application ApplicationResponse, expires time.Time, scopes []OAuth2Scopes) *OAuth2GetAuthorizationResponse {
	this := OAuth2GetAuthorizationResponse{}
	this.Application = application
	this.Expires = expires
	this.Scopes = scopes
	return &this
}

// NewOAuth2GetAuthorizationResponseWithDefaults instantiates a new OAuth2GetAuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2GetAuthorizationResponseWithDefaults() *OAuth2GetAuthorizationResponse {
	this := OAuth2GetAuthorizationResponse{}
	return &this
}

// GetApplication returns the Application field value
func (o *OAuth2GetAuthorizationResponse) GetApplication() ApplicationResponse {
	if o == nil {
		var ret ApplicationResponse
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *OAuth2GetAuthorizationResponse) GetApplicationOk() (*ApplicationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *OAuth2GetAuthorizationResponse) SetApplication(v ApplicationResponse) {
	o.Application = v
}

// GetExpires returns the Expires field value
func (o *OAuth2GetAuthorizationResponse) GetExpires() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *OAuth2GetAuthorizationResponse) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *OAuth2GetAuthorizationResponse) SetExpires(v time.Time) {
	o.Expires = v
}

// GetScopes returns the Scopes field value
func (o *OAuth2GetAuthorizationResponse) GetScopes() []OAuth2Scopes {
	if o == nil {
		var ret []OAuth2Scopes
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *OAuth2GetAuthorizationResponse) GetScopesOk() ([]OAuth2Scopes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *OAuth2GetAuthorizationResponse) SetScopes(v []OAuth2Scopes) {
	o.Scopes = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuth2GetAuthorizationResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuth2GetAuthorizationResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *OAuth2GetAuthorizationResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *OAuth2GetAuthorizationResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *OAuth2GetAuthorizationResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *OAuth2GetAuthorizationResponse) UnsetUser() {
	o.User.Unset()
}

func (o OAuth2GetAuthorizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2GetAuthorizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application"] = o.Application
	toSerialize["expires"] = o.Expires
	toSerialize["scopes"] = o.Scopes
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

func (o *OAuth2GetAuthorizationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"application",
		"expires",
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

	varOAuth2GetAuthorizationResponse := _OAuth2GetAuthorizationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2GetAuthorizationResponse)

	if err != nil {
		return err
	}

	*o = OAuth2GetAuthorizationResponse(varOAuth2GetAuthorizationResponse)

	return err
}

type NullableOAuth2GetAuthorizationResponse struct {
	value *OAuth2GetAuthorizationResponse
	isSet bool
}

func (v NullableOAuth2GetAuthorizationResponse) Get() *OAuth2GetAuthorizationResponse {
	return v.value
}

func (v *NullableOAuth2GetAuthorizationResponse) Set(val *OAuth2GetAuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2GetAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2GetAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2GetAuthorizationResponse(val *OAuth2GetAuthorizationResponse) *NullableOAuth2GetAuthorizationResponse {
	return &NullableOAuth2GetAuthorizationResponse{value: val, isSet: true}
}

func (v NullableOAuth2GetAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2GetAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


