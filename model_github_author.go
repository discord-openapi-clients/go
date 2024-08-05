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

// checks if the GithubAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubAuthor{}

// GithubAuthor struct for GithubAuthor
type GithubAuthor struct {
	Username NullableString `json:"username,omitempty"`
	Name string `json:"name"`
}

type _GithubAuthor GithubAuthor

// NewGithubAuthor instantiates a new GithubAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubAuthor(name string) *GithubAuthor {
	this := GithubAuthor{}
	this.Name = name
	return &this
}

// NewGithubAuthorWithDefaults instantiates a new GithubAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubAuthorWithDefaults() *GithubAuthor {
	this := GithubAuthor{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubAuthor) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubAuthor) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *GithubAuthor) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *GithubAuthor) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *GithubAuthor) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *GithubAuthor) UnsetUsername() {
	o.Username.Unset()
}

// GetName returns the Name field value
func (o *GithubAuthor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GithubAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GithubAuthor) SetName(v string) {
	o.Name = v
}

func (o GithubAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GithubAuthor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGithubAuthor := _GithubAuthor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubAuthor)

	if err != nil {
		return err
	}

	*o = GithubAuthor(varGithubAuthor)

	return err
}

type NullableGithubAuthor struct {
	value *GithubAuthor
	isSet bool
}

func (v NullableGithubAuthor) Get() *GithubAuthor {
	return v.value
}

func (v *NullableGithubAuthor) Set(val *GithubAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubAuthor(val *GithubAuthor) *NullableGithubAuthor {
	return &NullableGithubAuthor{value: val, isSet: true}
}

func (v NullableGithubAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

