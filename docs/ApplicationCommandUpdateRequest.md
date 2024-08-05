# ApplicationCommandUpdateRequest

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
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationCommandUpdateRequest

`func NewApplicationCommandUpdateRequest(name string, ) *ApplicationCommandUpdateRequest`

NewApplicationCommandUpdateRequest instantiates a new ApplicationCommandUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandUpdateRequestWithDefaults

`func NewApplicationCommandUpdateRequestWithDefaults() *ApplicationCommandUpdateRequest`

NewApplicationCommandUpdateRequestWithDefaults instantiates a new ApplicationCommandUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCommandUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandUpdateRequest) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandUpdateRequest) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandUpdateRequest) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandUpdateRequest) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### SetNameLocalizationsNil

`func (o *ApplicationCommandUpdateRequest) SetNameLocalizationsNil(b bool)`

 SetNameLocalizationsNil sets the value for NameLocalizations to be an explicit nil

### UnsetNameLocalizations
`func (o *ApplicationCommandUpdateRequest) UnsetNameLocalizations()`

UnsetNameLocalizations ensures that no value is present for NameLocalizations, not even an explicit nil
### GetDescription

`func (o *ApplicationCommandUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationCommandUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationCommandUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationCommandUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionLocalizations

`func (o *ApplicationCommandUpdateRequest) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandUpdateRequest) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandUpdateRequest) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandUpdateRequest) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### SetDescriptionLocalizationsNil

`func (o *ApplicationCommandUpdateRequest) SetDescriptionLocalizationsNil(b bool)`

 SetDescriptionLocalizationsNil sets the value for DescriptionLocalizations to be an explicit nil

### UnsetDescriptionLocalizations
`func (o *ApplicationCommandUpdateRequest) UnsetDescriptionLocalizations()`

UnsetDescriptionLocalizations ensures that no value is present for DescriptionLocalizations, not even an explicit nil
### GetOptions

`func (o *ApplicationCommandUpdateRequest) GetOptions() []ApplicationCommandCreateRequestOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApplicationCommandUpdateRequest) GetOptionsOk() (*[]ApplicationCommandCreateRequestOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApplicationCommandUpdateRequest) SetOptions(v []ApplicationCommandCreateRequestOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApplicationCommandUpdateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ApplicationCommandUpdateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApplicationCommandUpdateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDefaultMemberPermissions

`func (o *ApplicationCommandUpdateRequest) GetDefaultMemberPermissions() int32`

GetDefaultMemberPermissions returns the DefaultMemberPermissions field if non-nil, zero value otherwise.

### GetDefaultMemberPermissionsOk

`func (o *ApplicationCommandUpdateRequest) GetDefaultMemberPermissionsOk() (*int32, bool)`

GetDefaultMemberPermissionsOk returns a tuple with the DefaultMemberPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMemberPermissions

`func (o *ApplicationCommandUpdateRequest) SetDefaultMemberPermissions(v int32)`

SetDefaultMemberPermissions sets DefaultMemberPermissions field to given value.

### HasDefaultMemberPermissions

`func (o *ApplicationCommandUpdateRequest) HasDefaultMemberPermissions() bool`

HasDefaultMemberPermissions returns a boolean if a field has been set.

### SetDefaultMemberPermissionsNil

`func (o *ApplicationCommandUpdateRequest) SetDefaultMemberPermissionsNil(b bool)`

 SetDefaultMemberPermissionsNil sets the value for DefaultMemberPermissions to be an explicit nil

### UnsetDefaultMemberPermissions
`func (o *ApplicationCommandUpdateRequest) UnsetDefaultMemberPermissions()`

UnsetDefaultMemberPermissions ensures that no value is present for DefaultMemberPermissions, not even an explicit nil
### GetDmPermission

`func (o *ApplicationCommandUpdateRequest) GetDmPermission() bool`

GetDmPermission returns the DmPermission field if non-nil, zero value otherwise.

### GetDmPermissionOk

`func (o *ApplicationCommandUpdateRequest) GetDmPermissionOk() (*bool, bool)`

GetDmPermissionOk returns a tuple with the DmPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmPermission

`func (o *ApplicationCommandUpdateRequest) SetDmPermission(v bool)`

SetDmPermission sets DmPermission field to given value.

### HasDmPermission

`func (o *ApplicationCommandUpdateRequest) HasDmPermission() bool`

HasDmPermission returns a boolean if a field has been set.

### SetDmPermissionNil

`func (o *ApplicationCommandUpdateRequest) SetDmPermissionNil(b bool)`

 SetDmPermissionNil sets the value for DmPermission to be an explicit nil

### UnsetDmPermission
`func (o *ApplicationCommandUpdateRequest) UnsetDmPermission()`

UnsetDmPermission ensures that no value is present for DmPermission, not even an explicit nil
### GetType

`func (o *ApplicationCommandUpdateRequest) GetType() ApplicationCommandType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandUpdateRequest) GetTypeOk() (*ApplicationCommandType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandUpdateRequest) SetType(v ApplicationCommandType)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationCommandUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApplicationCommandUpdateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApplicationCommandUpdateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *ApplicationCommandUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCommandUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCommandUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationCommandUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


