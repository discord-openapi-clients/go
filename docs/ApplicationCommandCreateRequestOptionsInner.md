# ApplicationCommandCreateRequestOptionsInner

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
**Options** | Pointer to [**[]ApplicationCommandSubcommandOptionOptionsInner**](ApplicationCommandSubcommandOptionOptionsInner.md) |  | [optional] 

## Methods

### NewApplicationCommandCreateRequestOptionsInner

`func NewApplicationCommandCreateRequestOptionsInner(type_ ApplicationCommandOptionType, name string, description string, ) *ApplicationCommandCreateRequestOptionsInner`

NewApplicationCommandCreateRequestOptionsInner instantiates a new ApplicationCommandCreateRequestOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandCreateRequestOptionsInnerWithDefaults

`func NewApplicationCommandCreateRequestOptionsInnerWithDefaults() *ApplicationCommandCreateRequestOptionsInner`

NewApplicationCommandCreateRequestOptionsInnerWithDefaults instantiates a new ApplicationCommandCreateRequestOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandCreateRequestOptionsInner) GetType() ApplicationCommandOptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetTypeOk() (*ApplicationCommandOptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandCreateRequestOptionsInner) SetType(v ApplicationCommandOptionType)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandCreateRequestOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandCreateRequestOptionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandCreateRequestOptionsInner) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandCreateRequestOptionsInner) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandCreateRequestOptionsInner) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### SetNameLocalizationsNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetNameLocalizationsNil(b bool)`

 SetNameLocalizationsNil sets the value for NameLocalizations to be an explicit nil

### UnsetNameLocalizations
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetNameLocalizations()`

UnsetNameLocalizations ensures that no value is present for NameLocalizations, not even an explicit nil
### GetDescription

`func (o *ApplicationCommandCreateRequestOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandCreateRequestOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalizations

`func (o *ApplicationCommandCreateRequestOptionsInner) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandCreateRequestOptionsInner) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandCreateRequestOptionsInner) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### SetDescriptionLocalizationsNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetDescriptionLocalizationsNil(b bool)`

 SetDescriptionLocalizationsNil sets the value for DescriptionLocalizations to be an explicit nil

### UnsetDescriptionLocalizations
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetDescriptionLocalizations()`

UnsetDescriptionLocalizations ensures that no value is present for DescriptionLocalizations, not even an explicit nil
### GetRequired

`func (o *ApplicationCommandCreateRequestOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandCreateRequestOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandCreateRequestOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetChannelTypes

`func (o *ApplicationCommandCreateRequestOptionsInner) GetChannelTypes() []ChannelTypes`

GetChannelTypes returns the ChannelTypes field if non-nil, zero value otherwise.

### GetChannelTypesOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetChannelTypesOk() (*[]ChannelTypes, bool)`

GetChannelTypesOk returns a tuple with the ChannelTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypes

`func (o *ApplicationCommandCreateRequestOptionsInner) SetChannelTypes(v []ChannelTypes)`

SetChannelTypes sets ChannelTypes field to given value.

### HasChannelTypes

`func (o *ApplicationCommandCreateRequestOptionsInner) HasChannelTypes() bool`

HasChannelTypes returns a boolean if a field has been set.

### SetChannelTypesNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetChannelTypesNil(b bool)`

 SetChannelTypesNil sets the value for ChannelTypes to be an explicit nil

### UnsetChannelTypes
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetChannelTypes()`

UnsetChannelTypes ensures that no value is present for ChannelTypes, not even an explicit nil
### GetAutocomplete

`func (o *ApplicationCommandCreateRequestOptionsInner) GetAutocomplete() bool`

GetAutocomplete returns the Autocomplete field if non-nil, zero value otherwise.

### GetAutocompleteOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetAutocompleteOk() (*bool, bool)`

GetAutocompleteOk returns a tuple with the Autocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocomplete

`func (o *ApplicationCommandCreateRequestOptionsInner) SetAutocomplete(v bool)`

SetAutocomplete sets Autocomplete field to given value.

### HasAutocomplete

`func (o *ApplicationCommandCreateRequestOptionsInner) HasAutocomplete() bool`

HasAutocomplete returns a boolean if a field has been set.

### SetAutocompleteNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetAutocompleteNil(b bool)`

 SetAutocompleteNil sets the value for Autocomplete to be an explicit nil

### UnsetAutocomplete
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetAutocomplete()`

UnsetAutocomplete ensures that no value is present for Autocomplete, not even an explicit nil
### GetChoices

`func (o *ApplicationCommandCreateRequestOptionsInner) GetChoices() []ApplicationCommandOptionStringChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetChoicesOk() (*[]ApplicationCommandOptionStringChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ApplicationCommandCreateRequestOptionsInner) SetChoices(v []ApplicationCommandOptionStringChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ApplicationCommandCreateRequestOptionsInner) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil
### GetMinValue

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ApplicationCommandCreateRequestOptionsInner) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetMaxValue

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ApplicationCommandCreateRequestOptionsInner) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
### GetMinLength

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *ApplicationCommandCreateRequestOptionsInner) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### SetMinLengthNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMinLengthNil(b bool)`

 SetMinLengthNil sets the value for MinLength to be an explicit nil

### UnsetMinLength
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetMinLength()`

UnsetMinLength ensures that no value is present for MinLength, not even an explicit nil
### GetMaxLength

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *ApplicationCommandCreateRequestOptionsInner) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetOptions

`func (o *ApplicationCommandCreateRequestOptionsInner) GetOptions() []ApplicationCommandSubcommandOptionOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApplicationCommandCreateRequestOptionsInner) GetOptionsOk() (*[]ApplicationCommandSubcommandOptionOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApplicationCommandCreateRequestOptionsInner) SetOptions(v []ApplicationCommandSubcommandOptionOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApplicationCommandCreateRequestOptionsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ApplicationCommandCreateRequestOptionsInner) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApplicationCommandCreateRequestOptionsInner) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


