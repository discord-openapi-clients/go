# MessageComponentActionRowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MessageComponentTypes**](MessageComponentTypes.md) |  | 
**Id** | **int32** |  | 
**Components** | Pointer to [**[]MessageComponentActionRowResponseComponentsInner**](MessageComponentActionRowResponseComponentsInner.md) |  | [optional] 

## Methods

### NewMessageComponentActionRowResponse

`func NewMessageComponentActionRowResponse(type_ MessageComponentTypes, id int32, ) *MessageComponentActionRowResponse`

NewMessageComponentActionRowResponse instantiates a new MessageComponentActionRowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentActionRowResponseWithDefaults

`func NewMessageComponentActionRowResponseWithDefaults() *MessageComponentActionRowResponse`

NewMessageComponentActionRowResponseWithDefaults instantiates a new MessageComponentActionRowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageComponentActionRowResponse) GetType() MessageComponentTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageComponentActionRowResponse) GetTypeOk() (*MessageComponentTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageComponentActionRowResponse) SetType(v MessageComponentTypes)`

SetType sets Type field to given value.


### GetId

`func (o *MessageComponentActionRowResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageComponentActionRowResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageComponentActionRowResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetComponents

`func (o *MessageComponentActionRowResponse) GetComponents() []MessageComponentActionRowResponseComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *MessageComponentActionRowResponse) GetComponentsOk() (*[]MessageComponentActionRowResponseComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *MessageComponentActionRowResponse) SetComponents(v []MessageComponentActionRowResponseComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *MessageComponentActionRowResponse) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### SetComponentsNil

`func (o *MessageComponentActionRowResponse) SetComponentsNil(b bool)`

 SetComponentsNil sets the value for Components to be an explicit nil

### UnsetComponents
`func (o *MessageComponentActionRowResponse) UnsetComponents()`

UnsetComponents ensures that no value is present for Components, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


