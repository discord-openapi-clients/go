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

// checks if the ApplicationCommandUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandUpdateRequest{}

// ApplicationCommandUpdateRequest struct for ApplicationCommandUpdateRequest
type ApplicationCommandUpdateRequest struct {
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	Options []ApplicationCommandCreateRequestOptionsInner `json:"options,omitempty"`
	DefaultMemberPermissions NullableInt32 `json:"default_member_permissions,omitempty"`
	DmPermission NullableBool `json:"dm_permission,omitempty"`
	Contexts []InteractionContextType `json:"contexts,omitempty"`
	IntegrationTypes []ApplicationIntegrationType `json:"integration_types,omitempty"`
	Type NullableApplicationCommandType `json:"type,omitempty"`
	Id *string `json:"id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _ApplicationCommandUpdateRequest ApplicationCommandUpdateRequest

// NewApplicationCommandUpdateRequest instantiates a new ApplicationCommandUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandUpdateRequest(name string) *ApplicationCommandUpdateRequest {
	this := ApplicationCommandUpdateRequest{}
	this.Name = name
	return &this
}

// NewApplicationCommandUpdateRequestWithDefaults instantiates a new ApplicationCommandUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandUpdateRequestWithDefaults() *ApplicationCommandUpdateRequest {
	this := ApplicationCommandUpdateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCommandUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandUpdateRequest) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetNameLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetNameLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return nil, false
	}
	return &o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandUpdateRequest) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ApplicationCommandUpdateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ApplicationCommandUpdateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ApplicationCommandUpdateRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetDescriptionLocalizations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetDescriptionLocalizationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return nil, false
	}
	return &o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandUpdateRequest) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetOptions() []ApplicationCommandCreateRequestOptionsInner {
	if o == nil {
		var ret []ApplicationCommandCreateRequestOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetOptionsOk() ([]ApplicationCommandCreateRequestOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ApplicationCommandCreateRequestOptionsInner and assigns it to the Options field.
func (o *ApplicationCommandUpdateRequest) SetOptions(v []ApplicationCommandCreateRequestOptionsInner) {
	o.Options = v
}

// GetDefaultMemberPermissions returns the DefaultMemberPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetDefaultMemberPermissions() int32 {
	if o == nil || IsNil(o.DefaultMemberPermissions.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultMemberPermissions.Get()
}

// GetDefaultMemberPermissionsOk returns a tuple with the DefaultMemberPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetDefaultMemberPermissionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultMemberPermissions.Get(), o.DefaultMemberPermissions.IsSet()
}

// HasDefaultMemberPermissions returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasDefaultMemberPermissions() bool {
	if o != nil && o.DefaultMemberPermissions.IsSet() {
		return true
	}

	return false
}

// SetDefaultMemberPermissions gets a reference to the given NullableInt32 and assigns it to the DefaultMemberPermissions field.
func (o *ApplicationCommandUpdateRequest) SetDefaultMemberPermissions(v int32) {
	o.DefaultMemberPermissions.Set(&v)
}
// SetDefaultMemberPermissionsNil sets the value for DefaultMemberPermissions to be an explicit nil
func (o *ApplicationCommandUpdateRequest) SetDefaultMemberPermissionsNil() {
	o.DefaultMemberPermissions.Set(nil)
}

// UnsetDefaultMemberPermissions ensures that no value is present for DefaultMemberPermissions, not even an explicit nil
func (o *ApplicationCommandUpdateRequest) UnsetDefaultMemberPermissions() {
	o.DefaultMemberPermissions.Unset()
}

// GetDmPermission returns the DmPermission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetDmPermission() bool {
	if o == nil || IsNil(o.DmPermission.Get()) {
		var ret bool
		return ret
	}
	return *o.DmPermission.Get()
}

// GetDmPermissionOk returns a tuple with the DmPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetDmPermissionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DmPermission.Get(), o.DmPermission.IsSet()
}

// HasDmPermission returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasDmPermission() bool {
	if o != nil && o.DmPermission.IsSet() {
		return true
	}

	return false
}

// SetDmPermission gets a reference to the given NullableBool and assigns it to the DmPermission field.
func (o *ApplicationCommandUpdateRequest) SetDmPermission(v bool) {
	o.DmPermission.Set(&v)
}
// SetDmPermissionNil sets the value for DmPermission to be an explicit nil
func (o *ApplicationCommandUpdateRequest) SetDmPermissionNil() {
	o.DmPermission.Set(nil)
}

// UnsetDmPermission ensures that no value is present for DmPermission, not even an explicit nil
func (o *ApplicationCommandUpdateRequest) UnsetDmPermission() {
	o.DmPermission.Unset()
}

// GetContexts returns the Contexts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetContexts() []InteractionContextType {
	if o == nil {
		var ret []InteractionContextType
		return ret
	}
	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetContextsOk() ([]InteractionContextType, bool) {
	if o == nil || IsNil(o.Contexts) {
		return nil, false
	}
	return o.Contexts, true
}

// HasContexts returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasContexts() bool {
	if o != nil && !IsNil(o.Contexts) {
		return true
	}

	return false
}

// SetContexts gets a reference to the given []InteractionContextType and assigns it to the Contexts field.
func (o *ApplicationCommandUpdateRequest) SetContexts(v []InteractionContextType) {
	o.Contexts = v
}

// GetIntegrationTypes returns the IntegrationTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetIntegrationTypes() []ApplicationIntegrationType {
	if o == nil {
		var ret []ApplicationIntegrationType
		return ret
	}
	return o.IntegrationTypes
}

// GetIntegrationTypesOk returns a tuple with the IntegrationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetIntegrationTypesOk() ([]ApplicationIntegrationType, bool) {
	if o == nil || IsNil(o.IntegrationTypes) {
		return nil, false
	}
	return o.IntegrationTypes, true
}

// HasIntegrationTypes returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasIntegrationTypes() bool {
	if o != nil && !IsNil(o.IntegrationTypes) {
		return true
	}

	return false
}

// SetIntegrationTypes gets a reference to the given []ApplicationIntegrationType and assigns it to the IntegrationTypes field.
func (o *ApplicationCommandUpdateRequest) SetIntegrationTypes(v []ApplicationIntegrationType) {
	o.IntegrationTypes = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandUpdateRequest) GetType() ApplicationCommandType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret ApplicationCommandType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandUpdateRequest) GetTypeOk() (*ApplicationCommandType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableApplicationCommandType and assigns it to the Type field.
func (o *ApplicationCommandUpdateRequest) SetType(v ApplicationCommandType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ApplicationCommandUpdateRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ApplicationCommandUpdateRequest) UnsetType() {
	o.Type.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationCommandUpdateRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationCommandUpdateRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationCommandUpdateRequest) SetId(v string) {
	o.Id = &v
}

func (o ApplicationCommandUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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
	if o.Contexts != nil {
		toSerialize["contexts"] = o.Contexts
	}
	if o.IntegrationTypes != nil {
		toSerialize["integration_types"] = o.IntegrationTypes
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

func (o *ApplicationCommandUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varApplicationCommandUpdateRequest := _ApplicationCommandUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandUpdateRequest)

	if err != nil {
		return err
	}

	*o = ApplicationCommandUpdateRequest(varApplicationCommandUpdateRequest)

	return err
}

type NullableApplicationCommandUpdateRequest struct {
	value *ApplicationCommandUpdateRequest
	isSet bool
}

func (v NullableApplicationCommandUpdateRequest) Get() *ApplicationCommandUpdateRequest {
	return v.value
}

func (v *NullableApplicationCommandUpdateRequest) Set(val *ApplicationCommandUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandUpdateRequest(val *ApplicationCommandUpdateRequest) *NullableApplicationCommandUpdateRequest {
	return &NullableApplicationCommandUpdateRequest{value: val, isSet: true}
}

func (v NullableApplicationCommandUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


