# CreateDm200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ChannelTypes**](ChannelTypes.md) |  | 
**LastMessageId** | Pointer to **string** |  | [optional] 
**Flags** | **int32** |  | 
**LastPinTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Recipients** | [**[]UserResponse**](UserResponse.md) |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **NullableBool** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDm200Response

`func NewCreateDm200Response(id string, type_ ChannelTypes, flags int32, recipients []UserResponse, ) *CreateDm200Response`

NewCreateDm200Response instantiates a new CreateDm200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDm200ResponseWithDefaults

`func NewCreateDm200ResponseWithDefaults() *CreateDm200Response`

NewCreateDm200ResponseWithDefaults instantiates a new CreateDm200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateDm200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateDm200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateDm200Response) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CreateDm200Response) GetType() ChannelTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDm200Response) GetTypeOk() (*ChannelTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDm200Response) SetType(v ChannelTypes)`

SetType sets Type field to given value.


### GetLastMessageId

`func (o *CreateDm200Response) GetLastMessageId() string`

GetLastMessageId returns the LastMessageId field if non-nil, zero value otherwise.

### GetLastMessageIdOk

`func (o *CreateDm200Response) GetLastMessageIdOk() (*string, bool)`

GetLastMessageIdOk returns a tuple with the LastMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageId

`func (o *CreateDm200Response) SetLastMessageId(v string)`

SetLastMessageId sets LastMessageId field to given value.

### HasLastMessageId

`func (o *CreateDm200Response) HasLastMessageId() bool`

HasLastMessageId returns a boolean if a field has been set.

### GetFlags

`func (o *CreateDm200Response) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *CreateDm200Response) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *CreateDm200Response) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetLastPinTimestamp

`func (o *CreateDm200Response) GetLastPinTimestamp() time.Time`

GetLastPinTimestamp returns the LastPinTimestamp field if non-nil, zero value otherwise.

### GetLastPinTimestampOk

`func (o *CreateDm200Response) GetLastPinTimestampOk() (*time.Time, bool)`

GetLastPinTimestampOk returns a tuple with the LastPinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPinTimestamp

`func (o *CreateDm200Response) SetLastPinTimestamp(v time.Time)`

SetLastPinTimestamp sets LastPinTimestamp field to given value.

### HasLastPinTimestamp

`func (o *CreateDm200Response) HasLastPinTimestamp() bool`

HasLastPinTimestamp returns a boolean if a field has been set.

### SetLastPinTimestampNil

`func (o *CreateDm200Response) SetLastPinTimestampNil(b bool)`

 SetLastPinTimestampNil sets the value for LastPinTimestamp to be an explicit nil

### UnsetLastPinTimestamp
`func (o *CreateDm200Response) UnsetLastPinTimestamp()`

UnsetLastPinTimestamp ensures that no value is present for LastPinTimestamp, not even an explicit nil
### GetRecipients

`func (o *CreateDm200Response) GetRecipients() []UserResponse`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CreateDm200Response) GetRecipientsOk() (*[]UserResponse, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CreateDm200Response) SetRecipients(v []UserResponse)`

SetRecipients sets Recipients field to given value.


### GetName

`func (o *CreateDm200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDm200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDm200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDm200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateDm200Response) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateDm200Response) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIcon

`func (o *CreateDm200Response) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateDm200Response) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateDm200Response) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateDm200Response) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CreateDm200Response) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateDm200Response) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetOwnerId

`func (o *CreateDm200Response) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateDm200Response) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateDm200Response) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CreateDm200Response) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetManaged

`func (o *CreateDm200Response) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *CreateDm200Response) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *CreateDm200Response) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *CreateDm200Response) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *CreateDm200Response) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *CreateDm200Response) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetApplicationId

`func (o *CreateDm200Response) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CreateDm200Response) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CreateDm200Response) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CreateDm200Response) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


