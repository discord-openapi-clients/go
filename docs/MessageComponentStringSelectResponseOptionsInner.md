# MessageComponentStringSelectResponseOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Value** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Emoji** | Pointer to [**MessageComponentEmojiResponse**](MessageComponentEmojiResponse.md) |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMessageComponentStringSelectResponseOptionsInner

`func NewMessageComponentStringSelectResponseOptionsInner(label string, value string, ) *MessageComponentStringSelectResponseOptionsInner`

NewMessageComponentStringSelectResponseOptionsInner instantiates a new MessageComponentStringSelectResponseOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentStringSelectResponseOptionsInnerWithDefaults

`func NewMessageComponentStringSelectResponseOptionsInnerWithDefaults() *MessageComponentStringSelectResponseOptionsInner`

NewMessageComponentStringSelectResponseOptionsInnerWithDefaults instantiates a new MessageComponentStringSelectResponseOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *MessageComponentStringSelectResponseOptionsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MessageComponentStringSelectResponseOptionsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MessageComponentStringSelectResponseOptionsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *MessageComponentStringSelectResponseOptionsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MessageComponentStringSelectResponseOptionsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MessageComponentStringSelectResponseOptionsInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *MessageComponentStringSelectResponseOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageComponentStringSelectResponseOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageComponentStringSelectResponseOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageComponentStringSelectResponseOptionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MessageComponentStringSelectResponseOptionsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MessageComponentStringSelectResponseOptionsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmoji

`func (o *MessageComponentStringSelectResponseOptionsInner) GetEmoji() MessageComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *MessageComponentStringSelectResponseOptionsInner) GetEmojiOk() (*MessageComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *MessageComponentStringSelectResponseOptionsInner) SetEmoji(v MessageComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *MessageComponentStringSelectResponseOptionsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetDefault

`func (o *MessageComponentStringSelectResponseOptionsInner) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *MessageComponentStringSelectResponseOptionsInner) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *MessageComponentStringSelectResponseOptionsInner) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *MessageComponentStringSelectResponseOptionsInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *MessageComponentStringSelectResponseOptionsInner) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *MessageComponentStringSelectResponseOptionsInner) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


