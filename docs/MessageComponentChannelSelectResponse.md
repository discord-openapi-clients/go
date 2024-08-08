# MessageComponentChannelSelectResponse

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
**ChannelTypes** | Pointer to [**[]ChannelTypes**](ChannelTypes.md) |  | [optional] 
**DefaultValues** | Pointer to [**[]ChannelSelectDefaultValueResponse**](ChannelSelectDefaultValueResponse.md) |  | [optional] 

## Methods

### NewMessageComponentChannelSelectResponse

`func NewMessageComponentChannelSelectResponse(type_ MessageComponentTypes, id int32, customId string, ) *MessageComponentChannelSelectResponse`

NewMessageComponentChannelSelectResponse instantiates a new MessageComponentChannelSelectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentChannelSelectResponseWithDefaults

`func NewMessageComponentChannelSelectResponseWithDefaults() *MessageComponentChannelSelectResponse`

NewMessageComponentChannelSelectResponseWithDefaults instantiates a new MessageComponentChannelSelectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageComponentChannelSelectResponse) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageComponentChannelSelectResponse) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageComponentChannelSelectResponse) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *MessageComponentChannelSelectResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageComponentChannelSelectResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageComponentChannelSelectResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *MessageComponentChannelSelectResponse) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *MessageComponentChannelSelectResponse) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *MessageComponentChannelSelectResponse) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetPlaceholder

`func (o *MessageComponentChannelSelectResponse) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *MessageComponentChannelSelectResponse) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *MessageComponentChannelSelectResponse) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *MessageComponentChannelSelectResponse) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *MessageComponentChannelSelectResponse) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *MessageComponentChannelSelectResponse) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetMinValues

`func (o *MessageComponentChannelSelectResponse) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *MessageComponentChannelSelectResponse) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *MessageComponentChannelSelectResponse) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *MessageComponentChannelSelectResponse) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### SetMinValuesNil

`func (o *MessageComponentChannelSelectResponse) SetMinValuesNil(b bool)`

 SetMinValuesNil sets the value for MinValues to be an explicit nil

### UnsetMinValues
`func (o *MessageComponentChannelSelectResponse) UnsetMinValues()`

UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
### GetMaxValues

`func (o *MessageComponentChannelSelectResponse) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *MessageComponentChannelSelectResponse) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *MessageComponentChannelSelectResponse) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *MessageComponentChannelSelectResponse) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### SetMaxValuesNil

`func (o *MessageComponentChannelSelectResponse) SetMaxValuesNil(b bool)`

 SetMaxValuesNil sets the value for MaxValues to be an explicit nil

### UnsetMaxValues
`func (o *MessageComponentChannelSelectResponse) UnsetMaxValues()`

UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
### GetDisabled

`func (o *MessageComponentChannelSelectResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MessageComponentChannelSelectResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MessageComponentChannelSelectResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MessageComponentChannelSelectResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *MessageComponentChannelSelectResponse) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *MessageComponentChannelSelectResponse) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetChannelTypes

`func (o *MessageComponentChannelSelectResponse) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *MessageComponentChannelSelectResponse) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *MessageComponentChannelSelectResponse) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *MessageComponentChannelSelectResponse) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### SetChannelTypesNil

`func (o *MessageComponentChannelSelectResponse) SetChannelTypesNil(b bool)`

 SetChannelTypesNil sets the value for ChannelTypes to be an explicit nil

### UnsetChannelTypes
`func (o *MessageComponentChannelSelectResponse) UnsetChannelTypes()`

UnsetChannelTypes ensures that no value is present for ChannelTypes, not even an explicit nil
### GetDefaultValues

`func (o *MessageComponentChannelSelectResponse) GetDefaultValues() []ChannelSelectDefaultValueResponse`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *MessageComponentChannelSelectResponse) GetDefaultValuesOk() (*[]ChannelSelectDefaultValueResponse, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *MessageComponentChannelSelectResponse) SetDefaultValues(v []ChannelSelectDefaultValueResponse)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *MessageComponentChannelSelectResponse) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### SetDefaultValuesNil

`func (o *MessageComponentChannelSelectResponse) SetDefaultValuesNil(b bool)`

 SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil

### UnsetDefaultValues
`func (o *MessageComponentChannelSelectResponse) UnsetDefaultValues()`

UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


