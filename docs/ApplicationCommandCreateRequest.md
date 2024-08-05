# ApplicationCommandCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Options** | Pointer to [**[]ApplicationCommandCreateRequestOptionsInner**](ApplicationCommandCreateRequestOptionsInner.md) |  | [optional] 
**DefaultMemberPermissions** | Pointer to **NullableInt32** |  | [optional] 
**DmPermission** | Pointer to **NullableBool** |  | [optional] 
**Type** | Pointer to [**NullableApplicationCommandType**](ApplicationCommandType.md) |  | [optional] 

## Methods

### NewApplicationCommandCreateRequest

`func NewApplicationCommandCreateRequest(name string, ) *ApplicationCommandCreateRequest`

NewApplicationCommandCreateRequest instantiates a new ApplicationCommandCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandCreateRequestWithDefaults

`func NewApplicationCommandCreateRequestWithDefaults() *ApplicationCommandCreateRequest`

NewApplicationCommandCreateRequestWithDefaults instantiates a new ApplicationCommandCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCommandCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandCreateRequest) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandCreateRequest) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandCreateRequest) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandCreateRequest) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### SetNameLocalizationsNil

`func (o *ApplicationCommandCreateRequest) SetNameLocalizationsNil(b bool)`

 SetNameLocalizationsNil sets the value for NameLocalizations to be an explicit nil

### UnsetNameLocalizations
`func (o *ApplicationCommandCreateRequest) UnsetNameLocalizations()`

UnsetNameLocalizations ensures that no value is present for NameLocalizations, not even an explicit nil
### GetDescription

`func (o *ApplicationCommandCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationCommandCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationCommandCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationCommandCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionLocalizations

`func (o *ApplicationCommandCreateRequest) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandCreateRequest) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandCreateRequest) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandCreateRequest) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### SetDescriptionLocalizationsNil

`func (o *ApplicationCommandCreateRequest) SetDescriptionLocalizationsNil(b bool)`

 SetDescriptionLocalizationsNil sets the value for DescriptionLocalizations to be an explicit nil

### UnsetDescriptionLocalizations
`func (o *ApplicationCommandCreateRequest) UnsetDescriptionLocalizations()`

UnsetDescriptionLocalizations ensures that no value is present for DescriptionLocalizations, not even an explicit nil
### GetOptions

`func (o *ApplicationCommandCreateRequest) GetOptions() []ApplicationCommandCreateRequestOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApplicationCommandCreateRequest) GetOptionsOk() (*[]ApplicationCommandCreateRequestOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApplicationCommandCreateRequest) SetOptions(v []ApplicationCommandCreateRequestOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApplicationCommandCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ApplicationCommandCreateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApplicationCommandCreateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultMemberPermissions

`func (o *ApplicationCommandCreateRequest) GetDefaultMemberPermissions() int32`

GetDefaultMemberPermissions returns the DefaultMemberPermissions field if non-nil, zero value otherwise.

### GetDefaultMemberPermissionsOk

`func (o *ApplicationCommandCreateRequest) GetDefaultMemberPermissionsOk() (*int32, bool)`

GetDefaultMemberPermissionsOk returns a tuple with the DefaultMemberPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemberPermissions

`func (o *ApplicationCommandCreateRequest) SetDefaultMemberPermissions(v int32)`

SetDefaultMemberPermissions sets DefaultMemberPermissions field to given value.

### HasDefaultMemberPermissions

`func (o *ApplicationCommandCreateRequest) HasDefaultMemberPermissions() bool`

HasDefaultMemberPermissions returns a boolean if a field has been set.

### SetDefaultMemberPermissionsNil

`func (o *ApplicationCommandCreateRequest) SetDefaultMemberPermissionsNil(b bool)`

 SetDefaultMemberPermissionsNil sets the value for DefaultMemberPermissions to be an explicit nil

### UnsetDefaultMemberPermissions
`func (o *ApplicationCommandCreateRequest) UnsetDefaultMemberPermissions()`

UnsetDefaultMemberPermissions ensures that no value is present for DefaultMemberPermissions, not even an explicit nil
### GetDmPermission

`func (o *ApplicationCommandCreateRequest) GetDmPermission() bool`

GetDmPermission returns the DmPermission field if non-nil, zero value otherwise.

### GetDmPermissionOk

`func (o *ApplicationCommandCreateRequest) GetDmPermissionOk() (*bool, bool)`

GetDmPermissionOk returns a tuple with the DmPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPermission

`func (o *ApplicationCommandCreateRequest) SetDmPermission(v bool)`

SetDmPermission sets DmPermission field to given value.

### HasDmPermission

`func (o *ApplicationCommandCreateRequest) HasDmPermission() bool`

HasDmPermission returns a boolean if a field has been set.

### SetDmPermissionNil

`func (o *ApplicationCommandCreateRequest) SetDmPermissionNil(b bool)`

 SetDmPermissionNil sets the value for DmPermission to be an explicit nil

### UnsetDmPermission
`func (o *ApplicationCommandCreateRequest) UnsetDmPermission()`

UnsetDmPermission ensures that no value is present for DmPermission, not even an explicit nil
### GetType

`func (o *ApplicationCommandCreateRequest) GetType() ApplicationCommandType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandCreateRequest) GetTypeOk() (*ApplicationCommandType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandCreateRequest) SetType(v ApplicationCommandType)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationCommandCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApplicationCommandCreateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApplicationCommandCreateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


