# MessageComponentActionRowResponseComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | **int32** |  | 
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
**Value** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**MinLength** | Pointer to **NullableInt32** |  | [optional] 
**MaxLength** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to [**[]MessageComponentStringSelectResponseOptionsInner**](MessageComponentStringSelectResponseOptionsInner.md) |  | [optional] 

## Methods

### NewMessageComponentActionRowResponseComponentsInner

`func NewMessageComponentActionRowResponseComponentsInner(type_ MessageComponentTypes, id int32, customId string, style TextStyleTypes, ) *MessageComponentActionRowResponseComponentsInner`

NewMessageComponentActionRowResponseComponentsInner instantiates a new MessageComponentActionRowResponseComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentActionRowResponseComponentsInnerWithDefaults

`func NewMessageComponentActionRowResponseComponentsInnerWithDefaults() *MessageComponentActionRowResponseComponentsInner`

NewMessageComponentActionRowResponseComponentsInnerWithDefaults instantiates a new MessageComponentActionRowResponseComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageComponentActionRowResponseComponentsInner) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageComponentActionRowResponseComponentsInner) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *MessageComponentActionRowResponseComponentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageComponentActionRowResponseComponentsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *MessageComponentActionRowResponseComponentsInner) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *MessageComponentActionRowResponseComponentsInner) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.


### GetStyle

`func (o *MessageComponentActionRowResponseComponentsInner) GetStyle() TextStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetStyleOk() (*TextStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *MessageComponentActionRowResponseComponentsInner) SetStyle(v TextStyleTypes)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *MessageComponentActionRowResponseComponentsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MessageComponentActionRowResponseComponentsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MessageComponentActionRowResponseComponentsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDisabled

`func (o *MessageComponentActionRowResponseComponentsInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MessageComponentActionRowResponseComponentsInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MessageComponentActionRowResponseComponentsInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetEmoji

`func (o *MessageComponentActionRowResponseComponentsInner) GetEmoji() MessageComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetEmojiOk() (*MessageComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *MessageComponentActionRowResponseComponentsInner) SetEmoji(v MessageComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *MessageComponentActionRowResponseComponentsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetUrl

`func (o *MessageComponentActionRowResponseComponentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessageComponentActionRowResponseComponentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessageComponentActionRowResponseComponentsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSkuId

`func (o *MessageComponentActionRowResponseComponentsInner) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *MessageComponentActionRowResponseComponentsInner) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *MessageComponentActionRowResponseComponentsInner) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetPlaceholder

`func (o *MessageComponentActionRowResponseComponentsInner) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *MessageComponentActionRowResponseComponentsInner) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *MessageComponentActionRowResponseComponentsInner) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetMinValues

`func (o *MessageComponentActionRowResponseComponentsInner) GetMinValues() int32`

GetMinValues returns the MinValues field if non-nil, zero value otherwise.

### GetMinValuesOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetMinValuesOk() (*int32, bool)`

GetMinValuesOk returns a tuple with the MinValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValues

`func (o *MessageComponentActionRowResponseComponentsInner) SetMinValues(v int32)`

SetMinValues sets MinValues field to given value.

### HasMinValues

`func (o *MessageComponentActionRowResponseComponentsInner) HasMinValues() bool`

HasMinValues returns a boolean if a field has been set.

### SetMinValuesNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetMinValuesNil(b bool)`

 SetMinValuesNil sets the value for MinValues to be an explicit nil

### UnsetMinValues
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetMinValues()`

UnsetMinValues ensures that no value is present for MinValues, not even an explicit nil
### GetMaxValues

`func (o *MessageComponentActionRowResponseComponentsInner) GetMaxValues() int32`

GetMaxValues returns the MaxValues field if non-nil, zero value otherwise.

### GetMaxValuesOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetMaxValuesOk() (*int32, bool)`

GetMaxValuesOk returns a tuple with the MaxValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValues

`func (o *MessageComponentActionRowResponseComponentsInner) SetMaxValues(v int32)`

SetMaxValues sets MaxValues field to given value.

### HasMaxValues

`func (o *MessageComponentActionRowResponseComponentsInner) HasMaxValues() bool`

HasMaxValues returns a boolean if a field has been set.

### SetMaxValuesNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetMaxValuesNil(b bool)`

 SetMaxValuesNil sets the value for MaxValues to be an explicit nil

### UnsetMaxValues
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetMaxValues()`

UnsetMaxValues ensures that no value is present for MaxValues, not even an explicit nil
### GetChannelTypes

`func (o *MessageComponentActionRowResponseComponentsInner) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *MessageComponentActionRowResponseComponentsInner) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *MessageComponentActionRowResponseComponentsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### SetChannelTypesNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetChannelTypesNil(b bool)`

 SetChannelTypesNil sets the value for ChannelTypes to be an explicit nil

### UnsetChannelTypes
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetChannelTypes()`

UnsetChannelTypes ensures that no value is present for ChannelTypes, not even an explicit nil
### GetValue

`func (o *MessageComponentActionRowResponseComponentsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MessageComponentActionRowResponseComponentsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MessageComponentActionRowResponseComponentsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRequired

`func (o *MessageComponentActionRowResponseComponentsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *MessageComponentActionRowResponseComponentsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *MessageComponentActionRowResponseComponentsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetMinLength

`func (o *MessageComponentActionRowResponseComponentsInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *MessageComponentActionRowResponseComponentsInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *MessageComponentActionRowResponseComponentsInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *MessageComponentActionRowResponseComponentsInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *MessageComponentActionRowResponseComponentsInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *MessageComponentActionRowResponseComponentsInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetOptions

`func (o *MessageComponentActionRowResponseComponentsInner) GetOptions() []MessageComponentStringSelectResponseOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MessageComponentActionRowResponseComponentsInner) GetOptionsOk() (*[]MessageComponentStringSelectResponseOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MessageComponentActionRowResponseComponentsInner) SetOptions(v []MessageComponentStringSelectResponseOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MessageComponentActionRowResponseComponentsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MessageComponentActionRowResponseComponentsInner) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MessageComponentActionRowResponseComponentsInner) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


