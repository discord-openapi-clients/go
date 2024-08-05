# UpdateGuildRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableInt32** |  | [optional] 
**Color** | Pointer to **NullableInt32** |  | [optional] 
**Hoist** | Pointer to **NullableBool** |  | [optional] 
**Mentionable** | Pointer to **NullableBool** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**UnicodeEmoji** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateGuildRoleRequest

`func NewUpdateGuildRoleRequest() *UpdateGuildRoleRequest`

NewUpdateGuildRoleRequest instantiates a new UpdateGuildRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGuildRoleRequestWithDefaults

`func NewUpdateGuildRoleRequestWithDefaults() *UpdateGuildRoleRequest`

NewUpdateGuildRoleRequestWithDefaults instantiates a new UpdateGuildRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGuildRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGuildRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGuildRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGuildRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateGuildRoleRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateGuildRoleRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *UpdateGuildRoleRequest) GetPermissions() int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateGuildRoleRequest) GetPermissionsOk() (*int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateGuildRoleRequest) SetPermissions(v int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateGuildRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *UpdateGuildRoleRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *UpdateGuildRoleRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetColor

`func (o *UpdateGuildRoleRequest) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateGuildRoleRequest) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateGuildRoleRequest) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateGuildRoleRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *UpdateGuildRoleRequest) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *UpdateGuildRoleRequest) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetHoist

`func (o *UpdateGuildRoleRequest) GetHoist() bool`

GetHoist returns the Hoist field if non-nil, zero value otherwise.

### GetHoistOk

`func (o *UpdateGuildRoleRequest) GetHoistOk() (*bool, bool)`

GetHoistOk returns a tuple with the Hoist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoist

`func (o *UpdateGuildRoleRequest) SetHoist(v bool)`

SetHoist sets Hoist field to given value.

### HasHoist

`func (o *UpdateGuildRoleRequest) HasHoist() bool`

HasHoist returns a boolean if a field has been set.

### SetHoistNil

`func (o *UpdateGuildRoleRequest) SetHoistNil(b bool)`

 SetHoistNil sets the value for Hoist to be an explicit nil

### UnsetHoist
`func (o *UpdateGuildRoleRequest) UnsetHoist()`

UnsetHoist ensures that no value is present for Hoist, not even an explicit nil
### GetMentionable

`func (o *UpdateGuildRoleRequest) GetMentionable() bool`

GetMentionable returns the Mentionable field if non-nil, zero value otherwise.

### GetMentionableOk

`func (o *UpdateGuildRoleRequest) GetMentionableOk() (*bool, bool)`

GetMentionableOk returns a tuple with the Mentionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionable

`func (o *UpdateGuildRoleRequest) SetMentionable(v bool)`

SetMentionable sets Mentionable field to given value.

### HasMentionable

`func (o *UpdateGuildRoleRequest) HasMentionable() bool`

HasMentionable returns a boolean if a field has been set.

### SetMentionableNil

`func (o *UpdateGuildRoleRequest) SetMentionableNil(b bool)`

 SetMentionableNil sets the value for Mentionable to be an explicit nil

### UnsetMentionable
`func (o *UpdateGuildRoleRequest) UnsetMentionable()`

UnsetMentionable ensures that no value is present for Mentionable, not even an explicit nil
### GetIcon

`func (o *UpdateGuildRoleRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UpdateGuildRoleRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UpdateGuildRoleRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UpdateGuildRoleRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *UpdateGuildRoleRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *UpdateGuildRoleRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetUnicodeEmoji

`func (o *UpdateGuildRoleRequest) GetUnicodeEmoji() string`

GetUnicodeEmoji returns the UnicodeEmoji field if non-nil, zero value otherwise.

### GetUnicodeEmojiOk

`func (o *UpdateGuildRoleRequest) GetUnicodeEmojiOk() (*string, bool)`

GetUnicodeEmojiOk returns a tuple with the UnicodeEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeEmoji

`func (o *UpdateGuildRoleRequest) SetUnicodeEmoji(v string)`

SetUnicodeEmoji sets UnicodeEmoji field to given value.

### HasUnicodeEmoji

`func (o *UpdateGuildRoleRequest) HasUnicodeEmoji() bool`

HasUnicodeEmoji returns a boolean if a field has been set.

### SetUnicodeEmojiNil

`func (o *UpdateGuildRoleRequest) SetUnicodeEmojiNil(b bool)`

 SetUnicodeEmojiNil sets the value for UnicodeEmoji to be an explicit nil

### UnsetUnicodeEmoji
`func (o *UpdateGuildRoleRequest) UnsetUnicodeEmoji()`

UnsetUnicodeEmoji ensures that no value is present for UnicodeEmoji, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


