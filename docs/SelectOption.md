# SelectOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Value** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Emoji** | Pointer to [**NullableEmoji**](Emoji.md) |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSelectOption

`func NewSelectOption(label string, value string, ) *SelectOption`

NewSelectOption instantiates a new SelectOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectOptionWithDefaults

`func NewSelectOptionWithDefaults() *SelectOption`

NewSelectOptionWithDefaults instantiates a new SelectOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *SelectOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SelectOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SelectOption) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *SelectOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SelectOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SelectOption) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *SelectOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SelectOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SelectOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SelectOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SelectOption) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SelectOption) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmoji

`func (o *SelectOption) GetEmoji() Emoji`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *SelectOption) GetEmojiOk() (*Emoji, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *SelectOption) SetEmoji(v Emoji)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *SelectOption) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### SetEmojiNil

`func (o *SelectOption) SetEmojiNil(b bool)`

 SetEmojiNil sets the value for Emoji to be an explicit nil

### UnsetEmoji
`func (o *SelectOption) UnsetEmoji()`

UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
### GetDefault

`func (o *SelectOption) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SelectOption) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SelectOption) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SelectOption) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *SelectOption) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *SelectOption) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


