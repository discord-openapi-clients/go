# BulkUpdateGuildRolesRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBulkUpdateGuildRolesRequestInner

`func NewBulkUpdateGuildRolesRequestInner() *BulkUpdateGuildRolesRequestInner`

NewBulkUpdateGuildRolesRequestInner instantiates a new BulkUpdateGuildRolesRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateGuildRolesRequestInnerWithDefaults

`func NewBulkUpdateGuildRolesRequestInnerWithDefaults() *BulkUpdateGuildRolesRequestInner`

NewBulkUpdateGuildRolesRequestInnerWithDefaults instantiates a new BulkUpdateGuildRolesRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkUpdateGuildRolesRequestInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkUpdateGuildRolesRequestInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkUpdateGuildRolesRequestInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkUpdateGuildRolesRequestInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *BulkUpdateGuildRolesRequestInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BulkUpdateGuildRolesRequestInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BulkUpdateGuildRolesRequestInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BulkUpdateGuildRolesRequestInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *BulkUpdateGuildRolesRequestInner) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *BulkUpdateGuildRolesRequestInner) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


