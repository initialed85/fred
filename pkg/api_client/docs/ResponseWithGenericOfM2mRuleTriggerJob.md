# ResponseWithGenericOfM2mRuleTriggerJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Objects** | Pointer to [**[]M2MRuleTriggerJob**](M2MRuleTriggerJob.md) |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewResponseWithGenericOfM2mRuleTriggerJob

`func NewResponseWithGenericOfM2mRuleTriggerJob() *ResponseWithGenericOfM2mRuleTriggerJob`

NewResponseWithGenericOfM2mRuleTriggerJob instantiates a new ResponseWithGenericOfM2mRuleTriggerJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithGenericOfM2mRuleTriggerJobWithDefaults

`func NewResponseWithGenericOfM2mRuleTriggerJobWithDefaults() *ResponseWithGenericOfM2mRuleTriggerJob`

NewResponseWithGenericOfM2mRuleTriggerJobWithDefaults instantiates a new ResponseWithGenericOfM2mRuleTriggerJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetError

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetError() []string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetErrorOk() (*[]string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetError(v []string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ResponseWithGenericOfM2mRuleTriggerJob) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetLimit

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetObjects

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetObjects() []M2MRuleTriggerJob`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetObjectsOk() (*[]M2MRuleTriggerJob, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetObjects(v []M2MRuleTriggerJob)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ResponseWithGenericOfM2mRuleTriggerJob) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetOffset

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccess

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTotalCount

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResponseWithGenericOfM2mRuleTriggerJob) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


