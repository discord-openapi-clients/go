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

// checks if the ApplicationCommandSubcommandOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandSubcommandOption{}

// ApplicationCommandSubcommandOption struct for ApplicationCommandSubcommandOption
type ApplicationCommandSubcommandOption struct {
	Type ApplicationCommandOptionType `json:"type"`
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Required NullableBool `json:"required,omitempty"`
	Options []ApplicationCommandSubcommandOptionOptionsInner `json:"options,omitempty"`
}

type _ApplicationCommandSubcommandOption ApplicationCommandSubcommandOption

// NewApplicationCommandSubcommandOption instantiates a new ApplicationCommandSubcommandOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandSubcommandOption(type_ ApplicationCommandOptionType, name string, description string) *ApplicationCommandSubcommandOption {
	this := ApplicationCommandSubcommandOption{}
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationCommandSubcommandOptionWithDefaults instantiates a new ApplicationCommandSubcommandOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandSubcommandOptionWithDefaults() *ApplicationCommandSubcommandOption {
	this := ApplicationCommandSubcommandOption{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationCommandSubcommandOption) GetType() ApplicationCommandOptionType {
	if o == nil {
		var ret ApplicationCommandOptionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandSubcommandOption) GetTypeOk() (*ApplicationCommandOptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandSubcommandOption) SetType(v ApplicationCommandOptionType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ApplicationCommandSubcommandOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandSubcommandOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandSubcommandOption) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandOption) GetNameLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandOption) GetNameLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return nil, false
	}
	return &o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandOption) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandSubcommandOption) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}

// GetDescription returns the Description field value
func (o *ApplicationCommandSubcommandOption) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandSubcommandOption) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationCommandSubcommandOption) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandOption) GetDescriptionLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandOption) GetDescriptionLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return nil, false
	}
	return &o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandOption) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandSubcommandOption) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandOption) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandOption) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandOption) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *ApplicationCommandSubcommandOption) SetRequired(v bool) {
	o.Required.Set(&v)
}
// SetRequiredNil sets the value for Required to be an explicit nil
func (o *ApplicationCommandSubcommandOption) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *ApplicationCommandSubcommandOption) UnsetRequired() {
	o.Required.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandOption) GetOptions() []ApplicationCommandSubcommandOptionOptionsInner {
	if o == nil {
		var ret []ApplicationCommandSubcommandOptionOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandOption) GetOptionsOk() ([]ApplicationCommandSubcommandOptionOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandOption) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ApplicationCommandSubcommandOptionOptionsInner and assigns it to the Options field.
func (o *ApplicationCommandSubcommandOption) SetOptions(v []ApplicationCommandSubcommandOptionOptionsInner) {
	o.Options = v
}

func (o ApplicationCommandSubcommandOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandSubcommandOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if o.NameLocalizations != nil {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	toSerialize["description"] = o.Description
	if o.DescriptionLocalizations != nil {
		toSerialize["description_localizations"] = o.DescriptionLocalizations
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *ApplicationCommandSubcommandOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"description",
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

	varApplicationCommandSubcommandOption := _ApplicationCommandSubcommandOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandSubcommandOption)

	if err != nil {
		return err
	}

	*o = ApplicationCommandSubcommandOption(varApplicationCommandSubcommandOption)

	return err
}

type NullableApplicationCommandSubcommandOption struct {
	value *ApplicationCommandSubcommandOption
	isSet bool
}

func (v NullableApplicationCommandSubcommandOption) Get() *ApplicationCommandSubcommandOption {
	return v.value
}

func (v *NullableApplicationCommandSubcommandOption) Set(val *ApplicationCommandSubcommandOption) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandSubcommandOption) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandSubcommandOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandSubcommandOption(val *ApplicationCommandSubcommandOption) *NullableApplicationCommandSubcommandOption {
	return &NullableApplicationCommandSubcommandOption{value: val, isSet: true}
}

func (v NullableApplicationCommandSubcommandOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandSubcommandOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


