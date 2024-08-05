# Emoji

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Animated** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewEmoji

`func NewEmoji(name string, ) *Emoji`

NewEmoji instantiates a new Emoji object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmojiWithDefaults

`func NewEmojiWithDefaults() *Emoji`

NewEmojiWithDefaults instantiates a new Emoji object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Emoji) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Emoji) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Emoji) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Emoji) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Emoji) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Emoji) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Emoji) SetName(v string)`

SetName sets Name field to given value.


### GetAnimated

`func (o *Emoji) GetAnimated() bool`

GetAnimated returns the Animated field if non-nil, zero value otherwise.

### GetAnimatedOk

`func (o *Emoji) GetAnimatedOk() (*bool, bool)`

GetAnimatedOk returns a tuple with the Animated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimated

`func (o *Emoji) SetAnimated(v bool)`

SetAnimated sets Animated field to given value.

### HasAnimated

`func (o *Emoji) HasAnimated() bool`

HasAnimated returns a boolean if a field has been set.

### SetAnimatedNil

`func (o *Emoji) SetAnimatedNil(b bool)`

 SetAnimatedNil sets the value for Animated to be an explicit nil

### UnsetAnimated
`func (o *Emoji) UnsetAnimated()`

UnsetAnimated ensures that no value is present for Animated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


