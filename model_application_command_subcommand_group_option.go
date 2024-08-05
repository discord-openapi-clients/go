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

// checks if the ApplicationCommandSubcommandGroupOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandSubcommandGroupOption{}

// ApplicationCommandSubcommandGroupOption struct for ApplicationCommandSubcommandGroupOption
type ApplicationCommandSubcommandGroupOption struct {
	Type ApplicationCommandOptionType `json:"type"`
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Required NullableBool `json:"required,omitempty"`
	Options []ApplicationCommandSubcommandOption `json:"options,omitempty"`
}

type _ApplicationCommandSubcommandGroupOption ApplicationCommandSubcommandGroupOption

// NewApplicationCommandSubcommandGroupOption instantiates a new ApplicationCommandSubcommandGroupOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandSubcommandGroupOption(type_ ApplicationCommandOptionType, name string, description string) *ApplicationCommandSubcommandGroupOption {
	this := ApplicationCommandSubcommandGroupOption{}
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationCommandSubcommandGroupOptionWithDefaults instantiates a new ApplicationCommandSubcommandGroupOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandSubcommandGroupOptionWithDefaults() *ApplicationCommandSubcommandGroupOption {
	this := ApplicationCommandSubcommandGroupOption{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationCommandSubcommandGroupOption) GetType() ApplicationCommandOptionType {
	if o == nil {
		var ret ApplicationCommandOptionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandSubcommandGroupOption) GetTypeOk() (*ApplicationCommandOptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandSubcommandGroupOption) SetType(v ApplicationCommandOptionType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ApplicationCommandSubcommandGroupOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandSubcommandGroupOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandSubcommandGroupOption) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandGroupOption) GetNameLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandGroupOption) GetNameLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return nil, false
	}
	return &o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandGroupOption) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandSubcommandGroupOption) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}

// GetDescription returns the Description field value
func (o *ApplicationCommandSubcommandGroupOption) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandSubcommandGroupOption) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationCommandSubcommandGroupOption) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandGroupOption) GetDescriptionLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandGroupOption) GetDescriptionLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return nil, false
	}
	return &o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandGroupOption) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandSubcommandGroupOption) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandGroupOption) GetRequired() bool {
	if o == nil || IsNil(o.Required.Get()) {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandGroupOption) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandGroupOption) HasRequired() bool {
	if o != nil && o.Required.IsSet() {
		return true
	}

	return false
}

// SetRequired gets a reference to the given NullableBool and assigns it to the Required field.
func (o *ApplicationCommandSubcommandGroupOption) SetRequired(v bool) {
	o.Required.Set(&v)
}
// SetRequiredNil sets the value for Required to be an explicit nil
func (o *ApplicationCommandSubcommandGroupOption) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil
func (o *ApplicationCommandSubcommandGroupOption) UnsetRequired() {
	o.Required.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandSubcommandGroupOption) GetOptions() []ApplicationCommandSubcommandOption {
	if o == nil {
		var ret []ApplicationCommandSubcommandOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandSubcommandGroupOption) GetOptionsOk() ([]ApplicationCommandSubcommandOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApplicationCommandSubcommandGroupOption) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ApplicationCommandSubcommandOption and assigns it to the Options field.
func (o *ApplicationCommandSubcommandGroupOption) SetOptions(v []ApplicationCommandSubcommandOption) {
	o.Options = v
}

func (o ApplicationCommandSubcommandGroupOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandSubcommandGroupOption) ToMap() (map[string]interface{}, error) {
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

func (o *ApplicationCommandSubcommandGroupOption) UnmarshalJSON(data []byte) (err error) {
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

	varApplicationCommandSubcommandGroupOption := _ApplicationCommandSubcommandGroupOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandSubcommandGroupOption)

	if err != nil {
		return err
	}

	*o = ApplicationCommandSubcommandGroupOption(varApplicationCommandSubcommandGroupOption)

	return err
}

type NullableApplicationCommandSubcommandGroupOption struct {
	value *ApplicationCommandSubcommandGroupOption
	isSet bool
}

func (v NullableApplicationCommandSubcommandGroupOption) Get() *ApplicationCommandSubcommandGroupOption {
	return v.value
}

func (v *NullableApplicationCommandSubcommandGroupOption) Set(val *ApplicationCommandSubcommandGroupOption) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandSubcommandGroupOption) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandSubcommandGroupOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandSubcommandGroupOption(val *ApplicationCommandSubcommandGroupOption) *NullableApplicationCommandSubcommandGroupOption {
	return &NullableApplicationCommandSubcommandGroupOption{value: val, isSet: true}
}

func (v NullableApplicationCommandSubcommandGroupOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandSubcommandGroupOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


