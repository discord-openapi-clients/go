# MessageComponentEmojiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Animated** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMessageComponentEmojiResponse

`func NewMessageComponentEmojiResponse(name string, ) *MessageComponentEmojiResponse`

NewMessageComponentEmojiResponse instantiates a new MessageComponentEmojiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageComponentEmojiResponseWithDefaults

`func NewMessageComponentEmojiResponseWithDefaults() *MessageComponentEmojiResponse`

NewMessageComponentEmojiResponseWithDefaults instantiates a new MessageComponentEmojiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageComponentEmojiResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageComponentEmojiResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageComponentEmojiResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageComponentEmojiResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MessageComponentEmojiResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageComponentEmojiResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageComponentEmojiResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAnimated

`func (o *MessageComponentEmojiResponse) GetAnimated() bool`

GetAnimated returns the Animated field if non-nil, zero value otherwise.

### GetAnimatedOk

`func (o *MessageComponentEmojiResponse) GetAnimatedOk() (*bool, bool)`

GetAnimatedOk returns a tuple with the Animated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimated

`func (o *MessageComponentEmojiResponse) SetAnimated(v bool)`

SetAnimated sets Animated field to given value.

### HasAnimated

`func (o *MessageComponentEmojiResponse) HasAnimated() bool`

HasAnimated returns a boolean if a field has been set.

### SetAnimatedNil

`func (o *MessageComponentEmojiResponse) SetAnimatedNil(b bool)`

 SetAnimatedNil sets the value for Animated to be an explicit nil

### UnsetAnimated
`func (o *MessageComponentEmojiResponse) UnsetAnimated()`

UnsetAnimated ensures that no value is present for Animated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


