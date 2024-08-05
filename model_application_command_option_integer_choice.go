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

// checks if the ApplicationCommandOptionIntegerChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandOptionIntegerChoice{}

// ApplicationCommandOptionIntegerChoice struct for ApplicationCommandOptionIntegerChoice
type ApplicationCommandOptionIntegerChoice struct {
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Value int64 `json:"value"`
}

type _ApplicationCommandOptionIntegerChoice ApplicationCommandOptionIntegerChoice

// NewApplicationCommandOptionIntegerChoice instantiates a new ApplicationCommandOptionIntegerChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandOptionIntegerChoice(name string, value int64) *ApplicationCommandOptionIntegerChoice {
	this := ApplicationCommandOptionIntegerChoice{}
	this.Name = name
	this.Value = value
	return &this
}

// NewApplicationCommandOptionIntegerChoiceWithDefaults instantiates a new ApplicationCommandOptionIntegerChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandOptionIntegerChoiceWithDefaults() *ApplicationCommandOptionIntegerChoice {
	this := ApplicationCommandOptionIntegerChoice{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCommandOptionIntegerChoice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandOptionIntegerChoice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandOptionIntegerChoice) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandOptionIntegerChoice) GetNameLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandOptionIntegerChoice) GetNameLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return nil, false
	}
	return &o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandOptionIntegerChoice) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandOptionIntegerChoice) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}

// GetValue returns the Value field value
func (o *ApplicationCommandOptionIntegerChoice) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandOptionIntegerChoice) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ApplicationCommandOptionIntegerChoice) SetValue(v int64) {
	o.Value = v
}

func (o ApplicationCommandOptionIntegerChoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandOptionIntegerChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.NameLocalizations != nil {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ApplicationCommandOptionIntegerChoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varApplicationCommandOptionIntegerChoice := _ApplicationCommandOptionIntegerChoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandOptionIntegerChoice)

	if err != nil {
		return err
	}

	*o = ApplicationCommandOptionIntegerChoice(varApplicationCommandOptionIntegerChoice)

	return err
}

type NullableApplicationCommandOptionIntegerChoice struct {
	value *ApplicationCommandOptionIntegerChoice
	isSet bool
}

func (v NullableApplicationCommandOptionIntegerChoice) Get() *ApplicationCommandOptionIntegerChoice {
	return v.value
}

func (v *NullableApplicationCommandOptionIntegerChoice) Set(val *ApplicationCommandOptionIntegerChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandOptionIntegerChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandOptionIntegerChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandOptionIntegerChoice(val *ApplicationCommandOptionIntegerChoice) *NullableApplicationCommandOptionIntegerChoice {
	return &NullableApplicationCommandOptionIntegerChoice{value: val, isSet: true}
}

func (v NullableApplicationCommandOptionIntegerChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandOptionIntegerChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


