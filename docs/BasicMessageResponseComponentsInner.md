# BasicMessageResponseComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | **int32** |  | 
**Components** | Pointer to [**[]MessageComponentActionRowResponseComponentsInner**](MessageComponentActionRowResponseComponentsInner.md) |  | [optional] 
**CustomId** | **string** |  | 
**Style** | [**TextStyleTypes**](TextStyleTypes.md) |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Emoji** | Pointer to [**MessageComponentEmojiResponse**](MessageComponentEmojiResponse.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **NullableString** |  | [optional] 
**MinValues** | Pointer to **NullableInt32** |  | [optional] 
**MaxValues** | Pointer to **NullableInt32** |  | [optional] 
**ChannelTypes** | Pointer to [**[]ChannelTypes**](ChannelTypes.md) |  | [optional] 
**DefaultValues** | Pointer to [**[]UserSelectDefaultValueResponse**](UserSelectDefaultValueResponse.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**MinLength** | Pointer to **NullableInt32** |  | [optional] 
**MaxLength** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to [**[]MessageComponentStringSelectResponseOptionsInner**](MessageComponentStringSelectResponseOptionsInner.md) |  | [optional] 

## Methods

### NewBasicMessageResponseComponentsInner

`func NewBasicMessageResponseComponentsInner(type_ MessageComponentTypes, id int32, customId string, style TextStyleTypes, ) *BasicMessageResponseComponentsInner`

NewBasicMessageResponseComponentsInner instantiates a new BasicMessageResponseComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicMessageResponseComponentsInnerWithDefaults

`func NewBasicMessageResponseComponentsInnerWithDefaults() *BasicMessageResponseComponentsInner`

NewBasicMessageResponseComponentsInnerWithDefaults instantiates a new BasicMessageResponseComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BasicMessageResponseComponentsInner) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicMessageResponseComponentsInner) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicMessageResponseComponentsInner) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *BasicMessageResponseComponentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicMessageResponseComponentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicMessageResponseComponentsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetComponents

`func (o *BasicMessageResponseComponentsInner) GetComponents() []MessageComponentActionRowResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BasicMessageResponseComponentsInner) GetComponentsOk() (*[]MessageComponentActionRowResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BasicMessageResponseComponentsInner) SetComponents(v []MessageComponentActionRowResponseComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *BasicMessageResponseComponentsInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *BasicMessageResponseComponentsInner) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *BasicMessageResponseComponentsInner) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil
### GetCustomId

`func (o *BasicMessageResponseComponentsInner) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *BasicMessageResponseComponentsInner) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *BasicMessageResponseComponentsInner) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetStyle

`func (o *BasicMessageResponseComponentsInner) GetStyle() TextStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *BasicMessageResponseComponentsInner) GetStyleOk() (*TextStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *BasicMessageResponseComponentsInner) SetStyle(v TextStyleTypes)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *BasicMessageResponseComponentsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BasicMessageResponseComponentsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BasicMessageResponseComponentsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BasicMessageResponseComponentsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *BasicMessageResponseComponentsInner) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *BasicMessageResponseComponentsInner) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDisabled

`func (o *BasicMessageResponseComponentsInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *BasicMessageResponseComponentsInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *BasicMessageResponseComponentsInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *BasicMessageResponseComponentsInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *BasicMessageResponseComponentsInner) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *BasicMessageResponseComponentsInner) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetEmoji

`func (o *BasicMessageResponseComponentsInner) GetEmoji() MessageComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *BasicMessageResponseComponentsInner) GetEmojiOk() (*MessageComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *BasicMessageResponseComponentsInner) SetEmoji(v MessageComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *BasicMessageResponseComponentsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetUrl

`func (o *BasicMessageResponseComponentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BasicMessageResponseComponentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BasicMessageResponseComponentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BasicMessageResponseComponentsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *BasicMessageResponseComponentsInner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *BasicMessageResponseComponentsInner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSkuId

`func (o *BasicMessageResponseComponentsInner) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *BasicMessageResponseComponentsInner) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *BasicMessageResponseComponentsInner) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *BasicMessageResponseComponentsInner) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetPlaceholder

`func (o *BasicMessageResponseComponentsInner) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *BasicMessageResponseComponentsInner) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *BasicMessageResponseComponentsInner) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *BasicMessageResponseComponentsInner) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *BasicMessageResponseComponentsInner) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *BasicMessageResponseComponentsInner) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetMinValues

`func (o *BasicMessageResponseComponentsInner) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *BasicMessageResponseComponentsInner) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *BasicMessageResponseComponentsInner) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *BasicMessageResponseComponentsInner) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### SetMinValuesNil

`func (o *BasicMessageResponseComponentsInner) SetMinValuesNil(b bool)`

 SetMinValuesNil sets the value for MinValues to be an explicit nil

### UnsetMinValues
`func (o *BasicMessageResponseComponentsInner) UnsetMinValues()`

UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
### GetMaxValues

`func (o *BasicMessageResponseComponentsInner) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *BasicMessageResponseComponentsInner) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *BasicMessageResponseComponentsInner) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *BasicMessageResponseComponentsInner) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### SetMaxValuesNil

`func (o *BasicMessageResponseComponentsInner) SetMaxValuesNil(b bool)`

 SetMaxValuesNil sets the value for MaxValues to be an explicit nil

### UnsetMaxValues
`func (o *BasicMessageResponseComponentsInner) UnsetMaxValues()`

UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
### GetChannelTypes

`func (o *BasicMessageResponseComponentsInner) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *BasicMessageResponseComponentsInner) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *BasicMessageResponseComponentsInner) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *BasicMessageResponseComponentsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### SetChannelTypesNil

`func (o *BasicMessageResponseComponentsInner) SetChannelTypesNil(b bool)`

 SetChannelTypesNil sets the value for ChannelTypes to be an explicit nil

### UnsetChannelTypes
`func (o *BasicMessageResponseComponentsInner) UnsetChannelTypes()`

UnsetChannelTypes ensures that no value is present for ChannelTypes, not even an explicit nil
### GetDefaultValues

`func (o *BasicMessageResponseComponentsInner) GetDefaultValues() []UserSelectDefaultValueResponse`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *BasicMessageResponseComponentsInner) GetDefaultValuesOk() (*[]UserSelectDefaultValueResponse, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *BasicMessageResponseComponentsInner) SetDefaultValues(v []UserSelectDefaultValueResponse)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *BasicMessageResponseComponentsInner) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### SetDefaultValuesNil

`func (o *BasicMessageResponseComponentsInner) SetDefaultValuesNil(b bool)`

 SetDefaultValuesNil sets the value for DefaultValues to be an explicit nil

### UnsetDefaultValues
`func (o *BasicMessageResponseComponentsInner) UnsetDefaultValues()`

UnsetDefaultValues ensures that no value is present for DefaultValues, not even an explicit nil
### GetValue

`func (o *BasicMessageResponseComponentsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BasicMessageResponseComponentsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BasicMessageResponseComponentsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BasicMessageResponseComponentsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BasicMessageResponseComponentsInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BasicMessageResponseComponentsInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRequired

`func (o *BasicMessageResponseComponentsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BasicMessageResponseComponentsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BasicMessageResponseComponentsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BasicMessageResponseComponentsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *BasicMessageResponseComponentsInner) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *BasicMessageResponseComponentsInner) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetMinLength

`func (o *BasicMessageResponseComponentsInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *BasicMessageResponseComponentsInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *BasicMessageResponseComponentsInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *BasicMessageResponseComponentsInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *BasicMessageResponseComponentsInner) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *BasicMessageResponseComponentsInner) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *BasicMessageResponseComponentsInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *BasicMessageResponseComponentsInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *BasicMessageResponseComponentsInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *BasicMessageResponseComponentsInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *BasicMessageResponseComponentsInner) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *BasicMessageResponseComponentsInner) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetOptions

`func (o *BasicMessageResponseComponentsInner) GetOptions() []MessageComponentStringSelectResponseOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BasicMessageResponseComponentsInner) GetOptionsOk() (*[]MessageComponentStringSelectResponseOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BasicMessageResponseComponentsInner) SetOptions(v []MessageComponentStringSelectResponseOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BasicMessageResponseComponentsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *BasicMessageResponseComponentsInner) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *BasicMessageResponseComponentsInner) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


