# ActionRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Components** | [**[]ActionRowComponentsInner**](ActionRowComponentsInner.md) |  | 

## Methods

### NewActionRow

`func NewActionRow(type_ MessageComponentTypes, components []ActionRowComponentsInner, ) *ActionRow`

NewActionRow instantiates a new ActionRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRowWithDefaults

`func NewActionRowWithDefaults() *ActionRow`

NewActionRowWithDefaults instantiates a new ActionRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionRow) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionRow) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionRow) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetComponents

`func (o *ActionRow) GetComponents() []ActionRowComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ActionRow) GetComponentsOk() (*[]ActionRowComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ActionRow) SetComponents(v []ActionRowComponentsInner)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


