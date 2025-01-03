# M2MRuleTriggerJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ReferencedByExecutionM2mRuleTriggerJobIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewM2MRuleTriggerJob

`func NewM2MRuleTriggerJob() *M2MRuleTriggerJob`

NewM2MRuleTriggerJob instantiates a new M2MRuleTriggerJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM2MRuleTriggerJobWithDefaults

`func NewM2MRuleTriggerJobWithDefaults() *M2MRuleTriggerJob`

NewM2MRuleTriggerJobWithDefaults instantiates a new M2MRuleTriggerJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *M2MRuleTriggerJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *M2MRuleTriggerJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *M2MRuleTriggerJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *M2MRuleTriggerJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *M2MRuleTriggerJob) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *M2MRuleTriggerJob) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *M2MRuleTriggerJob) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *M2MRuleTriggerJob) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *M2MRuleTriggerJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *M2MRuleTriggerJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *M2MRuleTriggerJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *M2MRuleTriggerJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferencedByExecutionM2mRuleTriggerJobIdObjects

`func (o *M2MRuleTriggerJob) GetReferencedByExecutionM2mRuleTriggerJobIdObjects() []Execution`

GetReferencedByExecutionM2mRuleTriggerJobIdObjects returns the ReferencedByExecutionM2mRuleTriggerJobIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionM2mRuleTriggerJobIdObjectsOk

`func (o *M2MRuleTriggerJob) GetReferencedByExecutionM2mRuleTriggerJobIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionM2mRuleTriggerJobIdObjectsOk returns a tuple with the ReferencedByExecutionM2mRuleTriggerJobIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionM2mRuleTriggerJobIdObjects

`func (o *M2MRuleTriggerJob) SetReferencedByExecutionM2mRuleTriggerJobIdObjects(v []Execution)`

SetReferencedByExecutionM2mRuleTriggerJobIdObjects sets ReferencedByExecutionM2mRuleTriggerJobIdObjects field to given value.

### HasReferencedByExecutionM2mRuleTriggerJobIdObjects

`func (o *M2MRuleTriggerJob) HasReferencedByExecutionM2mRuleTriggerJobIdObjects() bool`

HasReferencedByExecutionM2mRuleTriggerJobIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionM2mRuleTriggerJobIdObjectsNil

`func (o *M2MRuleTriggerJob) SetReferencedByExecutionM2mRuleTriggerJobIdObjectsNil(b bool)`

 SetReferencedByExecutionM2mRuleTriggerJobIdObjectsNil sets the value for ReferencedByExecutionM2mRuleTriggerJobIdObjects to be an explicit nil

### UnsetReferencedByExecutionM2mRuleTriggerJobIdObjects
`func (o *M2MRuleTriggerJob) UnsetReferencedByExecutionM2mRuleTriggerJobIdObjects()`

UnsetReferencedByExecutionM2mRuleTriggerJobIdObjects ensures that no value is present for ReferencedByExecutionM2mRuleTriggerJobIdObjects, not even an explicit nil
### GetUpdatedAt

`func (o *M2MRuleTriggerJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *M2MRuleTriggerJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *M2MRuleTriggerJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *M2MRuleTriggerJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


