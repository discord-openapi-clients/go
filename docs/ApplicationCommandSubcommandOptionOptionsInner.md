# ApplicationCommandSubcommandOptionOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ApplicationCommandOptionType**](ApplicationCommandOptionType.md) |  | 
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**ChannelTypes** | Pointer to [**[]ChannelTypes**](ChannelTypes.md) |  | [optional] 
**Autocomplete** | Pointer to **NullableBool** |  | [optional] 
**Choices** | Pointer to [**[]ApplicationCommandOptionStringChoice**](ApplicationCommandOptionStringChoice.md) |  | [optional] 
**MinValue** | Pointer to **NullableFloat64** |  | [optional] 
**MaxValue** | Pointer to **NullableFloat64** |  | [optional] 
**MinLength** | Pointer to **NullableInt32** |  | [optional] 
**MaxLength** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewApplicationCommandSubcommandOptionOptionsInner

`func NewApplicationCommandSubcommandOptionOptionsInner(type_ ApplicationCommandOptionType, name string, description string, ) *ApplicationCommandSubcommandOptionOptionsInner`

NewApplicationCommandSubcommandOptionOptionsInner instantiates a new ApplicationCommandSubcommandOptionOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandSubcommandOptionOptionsInnerWithDefaults

`func NewApplicationCommandSubcommandOptionOptionsInnerWithDefaults() *ApplicationCommandSubcommandOptionOptionsInner`

NewApplicationCommandSubcommandOptionOptionsInnerWithDefaults instantiates a new ApplicationCommandSubcommandOptionOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetType() ApplicationCommandOptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetTypeOk() (*ApplicationCommandOptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetType(v ApplicationCommandOptionType)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### SetNameLocalizationsNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetNameLocalizationsNil(b bool)`

 SetNameLocalizationsNil sets the value for NameLocalizations to be an explicit nil

### UnsetNameLocalizations
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetNameLocalizations()`

UnsetNameLocalizations ensures that no value is present for NameLocalizations, not even an explicit nil
### GetDescription

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### SetDescriptionLocalizationsNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetDescriptionLocalizationsNil(b bool)`

 SetDescriptionLocalizationsNil sets the value for DescriptionLocalizations to be an explicit nil

### UnsetDescriptionLocalizations
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetDescriptionLocalizations()`

UnsetDescriptionLocalizations ensures that no value is present for DescriptionLocalizations, not even an explicit nil
### GetRequired

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetChannelTypes

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### SetChannelTypesNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetChannelTypesNil(b bool)`

 SetChannelTypesNil sets the value for ChannelTypes to be an explicit nil

### UnsetChannelTypes
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetChannelTypes()`

UnsetChannelTypes ensures that no value is present for ChannelTypes, not even an explicit nil
### GetAutocomplete

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetAutocomplete() bool`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetAutocompleteOk() (*bool, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetAutocomplete(v bool)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### SetAutocompleteNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetAutocompleteNil(b bool)`

 SetAutocompleteNil sets the value for Autocomplete to be an explicit nil

### UnsetAutocomplete
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetAutocomplete()`

UnsetAutocomplete ensures that no value is present for Autocomplete, not even an explicit nil
### GetChoices

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetChoices() []ApplicationCommandOptionStringChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetChoicesOk() (*[]ApplicationCommandOptionStringChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetChoices(v []ApplicationCommandOptionStringChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil
### GetMinValue

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetMaxValue

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
### GetMinLength

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ApplicationCommandSubcommandOptionOptionsInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ApplicationCommandSubcommandOptionOptionsInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *ApplicationCommandSubcommandOptionOptionsInner) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *ApplicationCommandSubcommandOptionOptionsInner) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


