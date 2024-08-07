# ApplicationCommandInteractionMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**InteractionTypes**](InteractionTypes.md) |  | 
**User** | Pointer to [**NullableUserResponse**](UserResponse.md) |  | [optional] 
**AuthorizingIntegrationOwners** | **map[string]string** |  | 
**OriginalResponseMessageId** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationCommandInteractionMetadataResponse

`func NewApplicationCommandInteractionMetadataResponse(id string, type_ InteractionTypes, authorizingIntegrationOwners map[string]string, ) *ApplicationCommandInteractionMetadataResponse`

NewApplicationCommandInteractionMetadataResponse instantiates a new ApplicationCommandInteractionMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandInteractionMetadataResponseWithDefaults

`func NewApplicationCommandInteractionMetadataResponseWithDefaults() *ApplicationCommandInteractionMetadataResponse`

NewApplicationCommandInteractionMetadataResponseWithDefaults instantiates a new ApplicationCommandInteractionMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCommandInteractionMetadataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCommandInteractionMetadataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCommandInteractionMetadataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ApplicationCommandInteractionMetadataResponse) GetType() InteractionTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandInteractionMetadataResponse) GetTypeOk() (*InteractionTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandInteractionMetadataResponse) SetType(v InteractionTypes)`

SetType sets Type field to given value.


### GetUser

`func (o *ApplicationCommandInteractionMetadataResponse) GetUser() UserResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApplicationCommandInteractionMetadataResponse) GetUserOk() (*UserResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApplicationCommandInteractionMetadataResponse) SetUser(v UserResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *ApplicationCommandInteractionMetadataResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ApplicationCommandInteractionMetadataResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ApplicationCommandInteractionMetadataResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetAuthorizingIntegrationOwners

`func (o *ApplicationCommandInteractionMetadataResponse) GetAuthorizingIntegrationOwners() map[string]string`

GetAuthorizingIntegrationOwners returns the AuthorizingIntegrationOwners field if non-nil, zero value otherwise.

### GetAuthorizingIntegrationOwnersOk

`func (o *ApplicationCommandInteractionMetadataResponse) GetAuthorizingIntegrationOwnersOk() (*map[string]string, bool)`

GetAuthorizingIntegrationOwnersOk returns a tuple with the AuthorizingIntegrationOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizingIntegrationOwners

`func (o *ApplicationCommandInteractionMetadataResponse) SetAuthorizingIntegrationOwners(v map[string]string)`

SetAuthorizingIntegrationOwners sets AuthorizingIntegrationOwners field to given value.


### GetOriginalResponseMessageId

`func (o *ApplicationCommandInteractionMetadataResponse) GetOriginalResponseMessageId() string`

GetOriginalResponseMessageId returns the OriginalResponseMessageId field if non-nil, zero value otherwise.

### GetOriginalResponseMessageIdOk

`func (o *ApplicationCommandInteractionMetadataResponse) GetOriginalResponseMessageIdOk() (*string, bool)`

GetOriginalResponseMessageIdOk returns a tuple with the OriginalResponseMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalResponseMessageId

`func (o *ApplicationCommandInteractionMetadataResponse) SetOriginalResponseMessageId(v string)`

SetOriginalResponseMessageId sets OriginalResponseMessageId field to given value.

### HasOriginalResponseMessageId

`func (o *ApplicationCommandInteractionMetadataResponse) HasOriginalResponseMessageId() bool`

HasOriginalResponseMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


