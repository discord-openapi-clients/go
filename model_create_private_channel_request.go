/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discord

import (
	"encoding/json"
)

// checks if the CreatePrivateChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePrivateChannelRequest{}

// CreatePrivateChannelRequest struct for CreatePrivateChannelRequest
type CreatePrivateChannelRequest struct {
	RecipientId *string `json:"recipient_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	AccessTokens []string `json:"access_tokens,omitempty"`
	Nicks map[string]string `json:"nicks,omitempty"`
}

// NewCreatePrivateChannelRequest instantiates a new CreatePrivateChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePrivateChannelRequest() *CreatePrivateChannelRequest {
	this := CreatePrivateChannelRequest{}
	return &this
}

// NewCreatePrivateChannelRequestWithDefaults instantiates a new CreatePrivateChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePrivateChannelRequestWithDefaults() *CreatePrivateChannelRequest {
	this := CreatePrivateChannelRequest{}
	return &this
}

// GetRecipientId returns the RecipientId field value if set, zero value otherwise.
func (o *CreatePrivateChannelRequest) GetRecipientId() string {
	if o == nil || IsNil(o.RecipientId) {
		var ret string
		return ret
	}
	return *o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePrivateChannelRequest) GetRecipientIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientId) {
		return nil, false
	}
	return o.RecipientId, true
}

// HasRecipientId returns a boolean if a field has been set.
func (o *CreatePrivateChannelRequest) HasRecipientId() bool {
	if o != nil && !IsNil(o.RecipientId) {
		return true
	}

	return false
}

// SetRecipientId gets a reference to the given string and assigns it to the RecipientId field.
func (o *CreatePrivateChannelRequest) SetRecipientId(v string) {
	o.RecipientId = &v
}

// GetAccessTokens returns the AccessTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrivateChannelRequest) GetAccessTokens() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AccessTokens
}

// GetAccessTokensOk returns a tuple with the AccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrivateChannelRequest) GetAccessTokensOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokens) {
		return nil, false
	}
	return o.AccessTokens, true
}

// HasAccessTokens returns a boolean if a field has been set.
func (o *CreatePrivateChannelRequest) HasAccessTokens() bool {
	if o != nil && !IsNil(o.AccessTokens) {
		return true
	}

	return false
}

// SetAccessTokens gets a reference to the given []string and assigns it to the AccessTokens field.
func (o *CreatePrivateChannelRequest) SetAccessTokens(v []string) {
	o.AccessTokens = v
}

// GetNicks returns the Nicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePrivateChannelRequest) GetNicks() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Nicks
}

// GetNicksOk returns a tuple with the Nicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePrivateChannelRequest) GetNicksOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Nicks) {
		return nil, false
	}
	return &o.Nicks, true
}

// HasNicks returns a boolean if a field has been set.
func (o *CreatePrivateChannelRequest) HasNicks() bool {
	if o != nil && !IsNil(o.Nicks) {
		return true
	}

	return false
}

// SetNicks gets a reference to the given map[string]string and assigns it to the Nicks field.
func (o *CreatePrivateChannelRequest) SetNicks(v map[string]string) {
	o.Nicks = v
}

func (o CreatePrivateChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePrivateChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecipientId) {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if o.AccessTokens != nil {
		toSerialize["access_tokens"] = o.AccessTokens
	}
	if o.Nicks != nil {
		toSerialize["nicks"] = o.Nicks
	}
	return toSerialize, nil
}

type NullableCreatePrivateChannelRequest struct {
	value *CreatePrivateChannelRequest
	isSet bool
}

func (v NullableCreatePrivateChannelRequest) Get() *CreatePrivateChannelRequest {
	return v.value
}

func (v *NullableCreatePrivateChannelRequest) Set(val *CreatePrivateChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePrivateChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePrivateChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePrivateChannelRequest(val *CreatePrivateChannelRequest) *NullableCreatePrivateChannelRequest {
	return &NullableCreatePrivateChannelRequest{value: val, isSet: true}
}

func (v NullableCreatePrivateChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePrivateChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


