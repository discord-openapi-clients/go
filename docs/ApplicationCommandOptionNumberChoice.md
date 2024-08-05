# ApplicationCommandOptionNumberChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Value** | **float64** |  | 

## Methods

### NewApplicationCommandOptionNumberChoice

`func NewApplicationCommandOptionNumberChoice(name string, value float64, ) *ApplicationCommandOptionNumberChoice`

NewApplicationCommandOptionNumberChoice instantiates a new ApplicationCommandOptionNumberChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandOptionNumberChoiceWithDefaults

`func NewApplicationCommandOptionNumberChoiceWithDefaults() *ApplicationCommandOptionNumberChoice`

NewApplicationCommandOptionNumberChoiceWithDefaults instantiates a new ApplicationCommandOptionNumberChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCommandOptionNumberChoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandOptionNumberChoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandOptionNumberChoice) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationCommandOptionNumberChoice) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandOptionNumberChoice) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandOptionNumberChoice) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandOptionNumberChoice) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### SetNameLocalizationsNil

`func (o *ApplicationCommandOptionNumberChoice) SetNameLocalizationsNil(b bool)`

 SetNameLocalizationsNil sets the value for NameLocalizations to be an explicit nil

### UnsetNameLocalizations
`func (o *ApplicationCommandOptionNumberChoice) UnsetNameLocalizations()`

UnsetNameLocalizations ensures that no value is present for NameLocalizations, not even an explicit nil
### GetValue

`func (o *ApplicationCommandOptionNumberChoice) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationCommandOptionNumberChoice) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationCommandOptionNumberChoice) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


