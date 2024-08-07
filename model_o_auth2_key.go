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

// checks if the OAuth2Key type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Key{}

// OAuth2Key struct for OAuth2Key
type OAuth2Key struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N string `json:"n"`
	E string `json:"e"`
	Alg string `json:"alg"`
}

type _OAuth2Key OAuth2Key

// NewOAuth2Key instantiates a new OAuth2Key object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Key(kty string, use string, kid string, n string, e string, alg string) *OAuth2Key {
	this := OAuth2Key{}
	this.Kty = kty
	this.Use = use
	this.Kid = kid
	this.N = n
	this.E = e
	this.Alg = alg
	return &this
}

// NewOAuth2KeyWithDefaults instantiates a new OAuth2Key object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2KeyWithDefaults() *OAuth2Key {
	this := OAuth2Key{}
	return &this
}

// GetKty returns the Kty field value
func (o *OAuth2Key) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *OAuth2Key) SetKty(v string) {
	o.Kty = v
}

// GetUse returns the Use field value
func (o *OAuth2Key) GetUse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Use
}

// GetUseOk returns a tuple with the Use field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Use, true
}

// SetUse sets field value
func (o *OAuth2Key) SetUse(v string) {
	o.Use = v
}

// GetKid returns the Kid field value
func (o *OAuth2Key) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *OAuth2Key) SetKid(v string) {
	o.Kid = v
}

// GetN returns the N field value
func (o *OAuth2Key) GetN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N, true
}

// SetN sets field value
func (o *OAuth2Key) SetN(v string) {
	o.N = v
}

// GetE returns the E field value
func (o *OAuth2Key) GetE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.E, true
}

// SetE sets field value
func (o *OAuth2Key) SetE(v string) {
	o.E = v
}

// GetAlg returns the Alg field value
func (o *OAuth2Key) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *OAuth2Key) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *OAuth2Key) SetAlg(v string) {
	o.Alg = v
}

func (o OAuth2Key) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Key) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kty"] = o.Kty
	toSerialize["use"] = o.Use
	toSerialize["kid"] = o.Kid
	toSerialize["n"] = o.N
	toSerialize["e"] = o.E
	toSerialize["alg"] = o.Alg
	return toSerialize, nil
}

func (o *OAuth2Key) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kty",
		"use",
		"kid",
		"n",
		"e",
		"alg",
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

	varOAuth2Key := _OAuth2Key{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuth2Key)

	if err != nil {
		return err
	}

	*o = OAuth2Key(varOAuth2Key)

	return err
}

type NullableOAuth2Key struct {
	value *OAuth2Key
	isSet bool
}

func (v NullableOAuth2Key) Get() *OAuth2Key {
	return v.value
}

func (v *NullableOAuth2Key) Set(val *OAuth2Key) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Key) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Key) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Key(val *OAuth2Key) *NullableOAuth2Key {
	return &NullableOAuth2Key{value: val, isSet: true}
}

func (v NullableOAuth2Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Key) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


