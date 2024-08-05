# Button

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**CustomId** | Pointer to **NullableString** |  | [optional] 
**Style** | [**ButtonStyleTypes**](ButtonStyleTypes.md) |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Emoji** | Pointer to [**NullableEmoji**](Emoji.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 

## Methods

### NewButton

`func NewButton(type_ MessageComponentTypes, style ButtonStyleTypes, ) *Button`

NewButton instantiates a new Button object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewButtonWithDefaults

`func NewButtonWithDefaults() *Button`

NewButtonWithDefaults instantiates a new Button object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Button) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Button) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Button) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetCustomId

`func (o *Button) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *Button) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *Button) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *Button) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### SetCustomIdNil

`func (o *Button) SetCustomIdNil(b bool)`

 SetCustomIdNil sets the value for CustomId to be an explicit nil

### UnsetCustomId
`func (o *Button) UnsetCustomId()`

UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
### GetStyle

`func (o *Button) GetStyle() ButtonStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *Button) GetStyleOk() (*ButtonStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *Button) SetStyle(v ButtonStyleTypes)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *Button) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Button) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Button) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Button) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Button) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Button) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDisabled

`func (o *Button) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Button) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Button) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Button) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *Button) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *Button) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetEmoji

`func (o *Button) GetEmoji() Emoji`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *Button) GetEmojiOk() (*Emoji, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *Button) SetEmoji(v Emoji)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *Button) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *Button) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *Button) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetUrl

`func (o *Button) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Button) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Button) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Button) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Button) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Button) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSkuId

`func (o *Button) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *Button) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *Button) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *Button) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


