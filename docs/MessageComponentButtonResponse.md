# MessageComponentButtonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | **int32** |  | 
**CustomId** | Pointer to **NullableString** |  | [optional] 
**Style** | [**ButtonStyleTypes**](ButtonStyleTypes.md) |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**Emoji** | Pointer to [**NullableMessageComponentEmojiResponse**](MessageComponentEmojiResponse.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 

## Methods

### NewMessageComponentButtonResponse

`func NewMessageComponentButtonResponse(type_ MessageComponentTypes, id int32, style ButtonStyleTypes, ) *MessageComponentButtonResponse`

NewMessageComponentButtonResponse instantiates a new MessageComponentButtonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentButtonResponseWithDefaults

`func NewMessageComponentButtonResponseWithDefaults() *MessageComponentButtonResponse`

NewMessageComponentButtonResponseWithDefaults instantiates a new MessageComponentButtonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageComponentButtonResponse) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageComponentButtonResponse) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageComponentButtonResponse) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *MessageComponentButtonResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageComponentButtonResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageComponentButtonResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCustomId

`func (o *MessageComponentButtonResponse) GetCustomId() string`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *MessageComponentButtonResponse) GetCustomIdOk() (*string, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *MessageComponentButtonResponse) SetCustomId(v string)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *MessageComponentButtonResponse) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### SetCustomIdNil

`func (o *MessageComponentButtonResponse) SetCustomIdNil(b bool)`

 SetCustomIdNil sets the value for CustomId to be an explicit nil

### UnsetCustomId
`func (o *MessageComponentButtonResponse) UnsetCustomId()`

UnsetCustomId ensures that no value is present for CustomId, not even an explicit nil
### GetStyle

`func (o *MessageComponentButtonResponse) GetStyle() ButtonStyleTypes`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *MessageComponentButtonResponse) GetStyleOk() (*ButtonStyleTypes, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *MessageComponentButtonResponse) SetStyle(v ButtonStyleTypes)`

SetStyle sets Style field to given value.


### GetLabel

`func (o *MessageComponentButtonResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MessageComponentButtonResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MessageComponentButtonResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MessageComponentButtonResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *MessageComponentButtonResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *MessageComponentButtonResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDisabled

`func (o *MessageComponentButtonResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MessageComponentButtonResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MessageComponentButtonResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MessageComponentButtonResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *MessageComponentButtonResponse) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *MessageComponentButtonResponse) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetEmoji

`func (o *MessageComponentButtonResponse) GetEmoji() MessageComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *MessageComponentButtonResponse) GetEmojiOk() (*MessageComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *MessageComponentButtonResponse) SetEmoji(v MessageComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *MessageComponentButtonResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *MessageComponentButtonResponse) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *MessageComponentButtonResponse) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetUrl

`func (o *MessageComponentButtonResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessageComponentButtonResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessageComponentButtonResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessageComponentButtonResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MessageComponentButtonResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MessageComponentButtonResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSkuId

`func (o *MessageComponentButtonResponse) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MessageComponentButtonResponse) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *MessageComponentButtonResponse) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *MessageComponentButtonResponse) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


