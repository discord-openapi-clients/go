# MessageComponentRoleSelectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | **int32** |  | 
**CustomId** | **string** |  | 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**MinValues** | Pointer to **NullableInt32** |  | [optional] 
**MaxValues** | Pointer to **NullableInt32** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**DefaultValues** | Pointer to [**[]RoleSelectDefaultValueResponse**](RoleSelectDefaultValueResponse.md) |  | [optional] 

## Methods

### NewMessageComponentRoleSelectResponse

`func NewMessageComponentRoleSelectResponse(type_ MessageComponentTypes, id int32, customId string, ) *MessageComponentRoleSelectResponse`

NewMessageComponentRoleSelectResponse instantiates a new MessageComponentRoleSelectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentRoleSelectResponseWithDefaults

`func NewMessageComponentRoleSelectResponseWithDefaults() *MessageComponentRoleSelectResponse`

NewMessageComponentRoleSelectResponseWithDefaults instantiates a new MessageComponentRoleSelectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageComponentRoleSelectResponse) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageComponentRoleSelectResponse) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageComponentRoleSelectResponse) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *MessageComponentRoleSelectResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageComponentRoleSelectResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageComponentRoleSelectResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *MessageComponentRoleSelectResponse) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *MessageComponentRoleSelectResponse) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *MessageComponentRoleSelectResponse) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetPlaceholder

`func (o *MessageComponentRoleSelectResponse) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *MessageComponentRoleSelectResponse) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *MessageComponentRoleSelectResponse) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *MessageComponentRoleSelectResponse) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *MessageComponentRoleSelectResponse) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *MessageComponentRoleSelectResponse) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetMinValues

`func (o *MessageComponentRoleSelectResponse) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *MessageComponentRoleSelectResponse) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *MessageComponentRoleSelectResponse) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *MessageComponentRoleSelectResponse) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### SetMinValuesNil

`func (o *MessageComponentRoleSelectResponse) SetMinValuesNil(b bool)`

 SetMinValuesNil sets the value for MinValues to be an explicit nil

### UnsetMinValues
`func (o *MessageComponentRoleSelectResponse) UnsetMinValues()`

UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
### GetMaxValues

`func (o *MessageComponentRoleSelectResponse) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *MessageComponentRoleSelectResponse) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *MessageComponentRoleSelectResponse) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *MessageComponentRoleSelectResponse) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### SetMaxValuesNil

`func (o *MessageComponentRoleSelectResponse) SetMaxValuesNil(b bool)`

 SetMaxValuesNil sets the value for MaxValues to be an explicit nil

### UnsetMaxValues
`func (o *MessageComponentRoleSelectResponse) UnsetMaxValues()`

UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
### GetDisabled

`func (o *MessageComponentRoleSelectResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MessageComponentRoleSelectResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MessageComponentRoleSelectResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MessageComponentRoleSelectResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *MessageComponentRoleSelectResponse) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *MessageComponentRoleSelectResponse) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetDefaultValues

`func (o *MessageComponentRoleSelectResponse) GetDefaultValues() []RoleSelectDefaultValueResponse`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *MessageComponentRoleSelectResponse) GetDefaultValuesOk() (*[]RoleSelectDefaultValueResponse, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *MessageComponentRoleSelectResponse) SetDefaultValues(v []RoleSelectDefaultValueResponse)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *MessageComponentRoleSelectResponse) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### SetDefaultValuesNil

`func (o *MessageComponentRoleSelectResponse) SetDefaultValuesNil(b bool)`

 SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil

### UnsetDefaultValues
`func (o *MessageComponentRoleSelectResponse) UnsetDefaultValues()`

UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


