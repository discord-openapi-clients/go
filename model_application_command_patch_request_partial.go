/*
Discord HTTP API (Preview)

Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.

API version: 10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApplicationCommandPatchRequestPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandPatchRequestPartial{}

// ApplicationCommandPatchRequestPartial struct for ApplicationCommandPatchRequestPartial
type ApplicationCommandPatchRequestPartial struct {
	Name *string `json:"name,omitempty"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Options []ApplicationCommandCreateRequestOptionsInner `json:"options,omitempty"`
	DefaultMemberPermissions NullableInt32 `json:"default_member_permissions,omitempty"`
	DmPermission NullableBool `json:"dm_permission,omitempty"`
}

// NewApplicationCommandPatchRequestPartial instantiates a new ApplicationCommandPatchRequestPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandPatchRequestPartial() *ApplicationCommandPatchRequestPartial {
	this := ApplicationCommandPatchRequestPartial{}
	return &this
}

// NewApplicationCommandPatchRequestPartialWithDefaults instantiates a new ApplicationCommandPatchRequestPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandPatchRequestPartialWithDefaults() *ApplicationCommandPatchRequestPartial {
	this := ApplicationCommandPatchRequestPartial{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationCommandPatchRequestPartial) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandPatchRequestPartial) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationCommandPatchRequestPartial) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationCommandPatchRequestPartial) SetName(v string) {
	o.Name = &v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandPatchRequestPartial) GetNameLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandPatchRequestPartial) GetNameLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return nil, false
	}
	return &o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandPatchRequestPartial) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandPatchRequestPartial) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandPatchRequestPartial) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandPatchRequestPartial) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationCommandPatchRequestPartial) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApplicationCommandPatchRequestPartial) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApplicationCommandPatchRequestPartial) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApplicationCommandPatchRequestPartial) UnsetDescription() {
	o.Description.Unset()
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandPatchRequestPartial) GetDescriptionLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandPatchRequestPartial) GetDescriptionLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return nil, false
	}
	return &o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandPatchRequestPartial) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandPatchRequestPartial) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandPatchRequestPartial) GetOptions() []ApplicationCommandCreateRequestOptionsInner {
	if o == nil {
		var ret []ApplicationCommandCreateRequestOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandPatchRequestPartial) GetOptionsOk() ([]ApplicationCommandCreateRequestOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApplicationCommandPatchRequestPartial) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ApplicationCommandCreateRequestOptionsInner and assigns it to the Options field.
func (o *ApplicationCommandPatchRequestPartial) SetOptions(v []ApplicationCommandCreateRequestOptionsInner) {
	o.Options = v
}

// GetDefaultMemberPermissions returns the DefaultMemberPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandPatchRequestPartial) GetDefaultMemberPermissions() int32 {
	if o == nil || IsNil(o.DefaultMemberPermissions.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultMemberPermissions.Get()
}

// GetDefaultMemberPermissionsOk returns a tuple with the DefaultMemberPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandPatchRequestPartial) GetDefaultMemberPermissionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultMemberPermissions.Get(), o.DefaultMemberPermissions.IsSet()
}

// HasDefaultMemberPermissions returns a boolean if a field has been set.
func (o *ApplicationCommandPatchRequestPartial) HasDefaultMemberPermissions() bool {
	if o != nil && o.DefaultMemberPermissions.IsSet() {
		return true
	}

	return false
}

// SetDefaultMemberPermissions gets a reference to the given NullableInt32 and assigns it to the DefaultMemberPermissions field.
func (o *ApplicationCommandPatchRequestPartial) SetDefaultMemberPermissions(v int32) {
	o.DefaultMemberPermissions.Set(&v)
}
// SetDefaultMemberPermissionsNil sets the value for DefaultMemberPermissions to be an explicit nil
func (o *ApplicationCommandPatchRequestPartial) SetDefaultMemberPermissionsNil() {
	o.DefaultMemberPermissions.Set(nil)
}

// UnsetDefaultMemberPermissions ensures that no value is present for DefaultMemberPermissions, not even an explicit nil
func (o *ApplicationCommandPatchRequestPartial) UnsetDefaultMemberPermissions() {
	o.DefaultMemberPermissions.Unset()
}

// GetDmPermission returns the DmPermission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandPatchRequestPartial) GetDmPermission() bool {
	if o == nil || IsNil(o.DmPermission.Get()) {
		var ret bool
		return ret
	}
	return *o.DmPermission.Get()
}

// GetDmPermissionOk returns a tuple with the DmPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandPatchRequestPartial) GetDmPermissionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DmPermission.Get(), o.DmPermission.IsSet()
}

// HasDmPermission returns a boolean if a field has been set.
func (o *ApplicationCommandPatchRequestPartial) HasDmPermission() bool {
	if o != nil && o.DmPermission.IsSet() {
		return true
	}

	return false
}

// SetDmPermission gets a reference to the given NullableBool and assigns it to the DmPermission field.
func (o *ApplicationCommandPatchRequestPartial) SetDmPermission(v bool) {
	o.DmPermission.Set(&v)
}
// SetDmPermissionNil sets the value for DmPermission to be an explicit nil
func (o *ApplicationCommandPatchRequestPartial) SetDmPermissionNil() {
	o.DmPermission.Set(nil)
}

// UnsetDmPermission ensures that no value is present for DmPermission, not even an explicit nil
func (o *ApplicationCommandPatchRequestPartial) UnsetDmPermission() {
	o.DmPermission.Unset()
}

func (o ApplicationCommandPatchRequestPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandPatchRequestPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.NameLocalizations != nil {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DescriptionLocalizations != nil {
		toSerialize["description_localizations"] = o.DescriptionLocalizations
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.DefaultMemberPermissions.IsSet() {
		toSerialize["default_member_permissions"] = o.DefaultMemberPermissions.Get()
	}
	if o.DmPermission.IsSet() {
		toSerialize["dm_permission"] = o.DmPermission.Get()
	}
	return toSerialize, nil
}

type NullableApplicationCommandPatchRequestPartial struct {
	value *ApplicationCommandPatchRequestPartial
	isSet bool
}

func (v NullableApplicationCommandPatchRequestPartial) Get() *ApplicationCommandPatchRequestPartial {
	return v.value
}

func (v *NullableApplicationCommandPatchRequestPartial) Set(val *ApplicationCommandPatchRequestPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandPatchRequestPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandPatchRequestPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandPatchRequestPartial(val *ApplicationCommandPatchRequestPartial) *NullableApplicationCommandPatchRequestPartial {
	return &NullableApplicationCommandPatchRequestPartial{value: val, isSet: true}
}

func (v NullableApplicationCommandPatchRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandPatchRequestPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

