# SelectOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Value** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Emoji** | Pointer to [**NullableMessageComponentEmojiResponse**](MessageComponentEmojiResponse.md) |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSelectOptionResponse

`func NewSelectOptionResponse(label string, value string, ) *SelectOptionResponse`

NewSelectOptionResponse instantiates a new SelectOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectOptionResponseWithDefaults

`func NewSelectOptionResponseWithDefaults() *SelectOptionResponse`

NewSelectOptionResponseWithDefaults instantiates a new SelectOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *SelectOptionResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SelectOptionResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SelectOptionResponse) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *SelectOptionResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SelectOptionResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SelectOptionResponse) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *SelectOptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SelectOptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SelectOptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SelectOptionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SelectOptionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SelectOptionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmoji

`func (o *SelectOptionResponse) GetEmoji() MessageComponentEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *SelectOptionResponse) GetEmojiOk() (*MessageComponentEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *SelectOptionResponse) SetEmoji(v MessageComponentEmojiResponse)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *SelectOptionResponse) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *SelectOptionResponse) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *SelectOptionResponse) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetDefault

`func (o *SelectOptionResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SelectOptionResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SelectOptionResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SelectOptionResponse) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *SelectOptionResponse) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *SelectOptionResponse) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


