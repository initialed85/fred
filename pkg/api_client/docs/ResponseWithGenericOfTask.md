# ResponseWithGenericOfTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Objects** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewResponseWithGenericOfTask

`func NewResponseWithGenericOfTask() *ResponseWithGenericOfTask`

NewResponseWithGenericOfTask instantiates a new ResponseWithGenericOfTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithGenericOfTaskWithDefaults

`func NewResponseWithGenericOfTaskWithDefaults() *ResponseWithGenericOfTask`

NewResponseWithGenericOfTaskWithDefaults instantiates a new ResponseWithGenericOfTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ResponseWithGenericOfTask) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResponseWithGenericOfTask) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResponseWithGenericOfTask) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResponseWithGenericOfTask) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetError

`func (o *ResponseWithGenericOfTask) GetError() []string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseWithGenericOfTask) GetErrorOk() (*[]string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseWithGenericOfTask) SetError(v []string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseWithGenericOfTask) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ResponseWithGenericOfTask) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ResponseWithGenericOfTask) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetLimit

`func (o *ResponseWithGenericOfTask) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseWithGenericOfTask) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseWithGenericOfTask) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseWithGenericOfTask) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetObjects

`func (o *ResponseWithGenericOfTask) GetObjects() []Task`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ResponseWithGenericOfTask) GetObjectsOk() (*[]Task, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ResponseWithGenericOfTask) SetObjects(v []Task)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ResponseWithGenericOfTask) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ResponseWithGenericOfTask) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ResponseWithGenericOfTask) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetOffset

`func (o *ResponseWithGenericOfTask) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ResponseWithGenericOfTask) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ResponseWithGenericOfTask) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ResponseWithGenericOfTask) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseWithGenericOfTask) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseWithGenericOfTask) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseWithGenericOfTask) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseWithGenericOfTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccess

`func (o *ResponseWithGenericOfTask) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseWithGenericOfTask) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseWithGenericOfTask) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseWithGenericOfTask) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTotalCount

`func (o *ResponseWithGenericOfTask) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseWithGenericOfTask) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseWithGenericOfTask) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResponseWithGenericOfTask) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


