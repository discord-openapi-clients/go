# UpdateMessageInteractionCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InteractionCallbackTypes**](InteractionCallbackTypes.md) |  | 
**Data** | Pointer to [**NullableIncomingWebhookUpdateForInteractionCallbackRequestPartial**](IncomingWebhookUpdateForInteractionCallbackRequestPartial.md) |  | [optional] 

## Methods

### NewUpdateMessageInteractionCallbackRequest

`func NewUpdateMessageInteractionCallbackRequest(type_ InteractionCallbackTypes, ) *UpdateMessageInteractionCallbackRequest`

NewUpdateMessageInteractionCallbackRequest instantiates a new UpdateMessageInteractionCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessageInteractionCallbackRequestWithDefaults

`func NewUpdateMessageInteractionCallbackRequestWithDefaults() *UpdateMessageInteractionCallbackRequest`

NewUpdateMessageInteractionCallbackRequestWithDefaults instantiates a new UpdateMessageInteractionCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateMessageInteractionCallbackRequest) GetType() InteractionCallbackTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMessageInteractionCallbackRequest) GetTypeOk() (*InteractionCallbackTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMessageInteractionCallbackRequest) SetType(v InteractionCallbackTypes)`

SetType sets Type field to given value.


### GetData

`func (o *UpdateMessageInteractionCallbackRequest) GetData() IncomingWebhookUpdateForInteractionCallbackRequestPartial`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateMessageInteractionCallbackRequest) GetDataOk() (*IncomingWebhookUpdateForInteractionCallbackRequestPartial, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateMessageInteractionCallbackRequest) SetData(v IncomingWebhookUpdateForInteractionCallbackRequestPartial)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateMessageInteractionCallbackRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UpdateMessageInteractionCallbackRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UpdateMessageInteractionCallbackRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


