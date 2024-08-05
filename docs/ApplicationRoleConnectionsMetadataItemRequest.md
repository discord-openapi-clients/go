# ApplicationRoleConnectionsMetadataItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MetadataItemTypes**](MetadataItemTypes.md) |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApplicationRoleConnectionsMetadataItemRequest

`func NewApplicationRoleConnectionsMetadataItemRequest(type_ MetadataItemTypes, key string, name string, description string, ) *ApplicationRoleConnectionsMetadataItemRequest`

NewApplicationRoleConnectionsMetadataItemRequest instantiates a new ApplicationRoleConnectionsMetadataItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRoleConnectionsMetadataItemRequestWithDefaults

`func NewApplicationRoleConnectionsMetadataItemRequestWithDefaults() *ApplicationRoleConnectionsMetadataItemRequest`

NewApplicationRoleConnectionsMetadataItemRequestWithDefaults instantiates a new ApplicationRoleConnectionsMetadataItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetType() MetadataItemTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetTypeOk() (*MetadataItemTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetType(v MetadataItemTypes)`

SetType sets Type field to given value.


### GetKey

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemRequest) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### SetNameLocalizationsNil

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetNameLocalizationsNil(b bool)`

 SetNameLocalizationsNil sets the value for NameLocalizations to be an explicit nil

### UnsetNameLocalizations
`func (o *ApplicationRoleConnectionsMetadataItemRequest) UnsetNameLocalizations()`

UnsetNameLocalizations ensures that no value is present for NameLocalizations, not even an explicit nil
### GetDescription

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemRequest) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### SetDescriptionLocalizationsNil

`func (o *ApplicationRoleConnectionsMetadataItemRequest) SetDescriptionLocalizationsNil(b bool)`

 SetDescriptionLocalizationsNil sets the value for DescriptionLocalizations to be an explicit nil

### UnsetDescriptionLocalizations
`func (o *ApplicationRoleConnectionsMetadataItemRequest) UnsetDescriptionLocalizations()`

UnsetDescriptionLocalizations ensures that no value is present for DescriptionLocalizations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


