# ResponseWithGenericOfRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Objects** | Pointer to [**[]Repository**](Repository.md) |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewResponseWithGenericOfRepository

`func NewResponseWithGenericOfRepository() *ResponseWithGenericOfRepository`

NewResponseWithGenericOfRepository instantiates a new ResponseWithGenericOfRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithGenericOfRepositoryWithDefaults

`func NewResponseWithGenericOfRepositoryWithDefaults() *ResponseWithGenericOfRepository`

NewResponseWithGenericOfRepositoryWithDefaults instantiates a new ResponseWithGenericOfRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ResponseWithGenericOfRepository) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResponseWithGenericOfRepository) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResponseWithGenericOfRepository) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResponseWithGenericOfRepository) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetError

`func (o *ResponseWithGenericOfRepository) GetError() []string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseWithGenericOfRepository) GetErrorOk() (*[]string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseWithGenericOfRepository) SetError(v []string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseWithGenericOfRepository) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ResponseWithGenericOfRepository) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ResponseWithGenericOfRepository) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetLimit

`func (o *ResponseWithGenericOfRepository) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseWithGenericOfRepository) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseWithGenericOfRepository) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseWithGenericOfRepository) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetObjects

`func (o *ResponseWithGenericOfRepository) GetObjects() []Repository`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ResponseWithGenericOfRepository) GetObjectsOk() (*[]Repository, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ResponseWithGenericOfRepository) SetObjects(v []Repository)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ResponseWithGenericOfRepository) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ResponseWithGenericOfRepository) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ResponseWithGenericOfRepository) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetOffset

`func (o *ResponseWithGenericOfRepository) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ResponseWithGenericOfRepository) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ResponseWithGenericOfRepository) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ResponseWithGenericOfRepository) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseWithGenericOfRepository) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseWithGenericOfRepository) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseWithGenericOfRepository) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseWithGenericOfRepository) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccess

`func (o *ResponseWithGenericOfRepository) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseWithGenericOfRepository) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseWithGenericOfRepository) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseWithGenericOfRepository) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTotalCount

`func (o *ResponseWithGenericOfRepository) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseWithGenericOfRepository) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseWithGenericOfRepository) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResponseWithGenericOfRepository) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


