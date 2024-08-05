# MessageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Filename** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsRemix** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMessageAttachmentRequest

`func NewMessageAttachmentRequest(id string, ) *MessageAttachmentRequest`

NewMessageAttachmentRequest instantiates a new MessageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageAttachmentRequestWithDefaults

`func NewMessageAttachmentRequestWithDefaults() *MessageAttachmentRequest`

NewMessageAttachmentRequestWithDefaults instantiates a new MessageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageAttachmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageAttachmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageAttachmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetFilename

`func (o *MessageAttachmentRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MessageAttachmentRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MessageAttachmentRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MessageAttachmentRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *MessageAttachmentRequest) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *MessageAttachmentRequest) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetDescription

`func (o *MessageAttachmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageAttachmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageAttachmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageAttachmentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MessageAttachmentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MessageAttachmentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsRemix

`func (o *MessageAttachmentRequest) GetIsRemix() bool`

GetIsRemix returns the IsRemix field if non-nil, zero value otherwise.

### GetIsRemixOk

`func (o *MessageAttachmentRequest) GetIsRemixOk() (*bool, bool)`

GetIsRemixOk returns a tuple with the IsRemix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemix

`func (o *MessageAttachmentRequest) SetIsRemix(v bool)`

SetIsRemix sets IsRemix field to given value.

### HasIsRemix

`func (o *MessageAttachmentRequest) HasIsRemix() bool`

HasIsRemix returns a boolean if a field has been set.

### SetIsRemixNil

`func (o *MessageAttachmentRequest) SetIsRemixNil(b bool)`

 SetIsRemixNil sets the value for IsRemix to be an explicit nil

### UnsetIsRemix
`func (o *MessageAttachmentRequest) UnsetIsRemix()`

UnsetIsRemix ensures that no value is present for IsRemix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


