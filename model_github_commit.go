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

// checks if the GithubCommit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubCommit{}

// GithubCommit struct for GithubCommit
type GithubCommit struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Message string `json:"message"`
	Author GithubAuthor `json:"author"`
}

type _GithubCommit GithubCommit

// NewGithubCommit instantiates a new GithubCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubCommit(id string, url string, message string, author GithubAuthor) *GithubCommit {
	this := GithubCommit{}
	this.Id = id
	this.Url = url
	this.Message = message
	this.Author = author
	return &this
}

// NewGithubCommitWithDefaults instantiates a new GithubCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubCommitWithDefaults() *GithubCommit {
	this := GithubCommit{}
	return &this
}

// GetId returns the Id field value
func (o *GithubCommit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GithubCommit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GithubCommit) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *GithubCommit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GithubCommit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GithubCommit) SetUrl(v string) {
	o.Url = v
}

// GetMessage returns the Message field value
func (o *GithubCommit) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GithubCommit) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GithubCommit) SetMessage(v string) {
	o.Message = v
}

// GetAuthor returns the Author field value
func (o *GithubCommit) GetAuthor() GithubAuthor {
	if o == nil {
		var ret GithubAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *GithubCommit) GetAuthorOk() (*GithubAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *GithubCommit) SetAuthor(v GithubAuthor) {
	o.Author = v
}

func (o GithubCommit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubCommit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["message"] = o.Message
	toSerialize["author"] = o.Author
	return toSerialize, nil
}

func (o *GithubCommit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"message",
		"author",
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

	varGithubCommit := _GithubCommit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubCommit)

	if err != nil {
		return err
	}

	*o = GithubCommit(varGithubCommit)

	return err
}

type NullableGithubCommit struct {
	value *GithubCommit
	isSet bool
}

func (v NullableGithubCommit) Get() *GithubCommit {
	return v.value
}

func (v *NullableGithubCommit) Set(val *GithubCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubCommit(val *GithubCommit) *NullableGithubCommit {
	return &NullableGithubCommit{value: val, isSet: true}
}

func (v NullableGithubCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


