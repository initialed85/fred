# Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | Pointer to **string** |  | [optional] 
**ChangeIdObject** | Pointer to [**Change**](Change.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ReferencedByExecutionTriggerIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**ReferencedByTriggerHasExecutionTriggerIdObjects** | Pointer to [**[]TriggerHasExecution**](TriggerHasExecution.md) |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**RuleIdObject** | Pointer to [**Rule**](Rule.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTrigger

`func NewTrigger() *Trigger`

NewTrigger instantiates a new Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWithDefaults

`func NewTriggerWithDefaults() *Trigger`

NewTriggerWithDefaults instantiates a new Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *Trigger) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *Trigger) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *Trigger) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *Trigger) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetChangeIdObject

`func (o *Trigger) GetChangeIdObject() Change`

GetChangeIdObject returns the ChangeIdObject field if non-nil, zero value otherwise.

### GetChangeIdObjectOk

`func (o *Trigger) GetChangeIdObjectOk() (*Change, bool)`

GetChangeIdObjectOk returns a tuple with the ChangeIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeIdObject

`func (o *Trigger) SetChangeIdObject(v Change)`

SetChangeIdObject sets ChangeIdObject field to given value.

### HasChangeIdObject

`func (o *Trigger) HasChangeIdObject() bool`

HasChangeIdObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Trigger) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Trigger) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Trigger) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Trigger) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Trigger) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Trigger) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Trigger) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Trigger) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *Trigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Trigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferencedByExecutionTriggerIdObjects

`func (o *Trigger) GetReferencedByExecutionTriggerIdObjects() []Execution`

GetReferencedByExecutionTriggerIdObjects returns the ReferencedByExecutionTriggerIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionTriggerIdObjectsOk

`func (o *Trigger) GetReferencedByExecutionTriggerIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionTriggerIdObjectsOk returns a tuple with the ReferencedByExecutionTriggerIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionTriggerIdObjects

`func (o *Trigger) SetReferencedByExecutionTriggerIdObjects(v []Execution)`

SetReferencedByExecutionTriggerIdObjects sets ReferencedByExecutionTriggerIdObjects field to given value.

### HasReferencedByExecutionTriggerIdObjects

`func (o *Trigger) HasReferencedByExecutionTriggerIdObjects() bool`

HasReferencedByExecutionTriggerIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionTriggerIdObjectsNil

`func (o *Trigger) SetReferencedByExecutionTriggerIdObjectsNil(b bool)`

 SetReferencedByExecutionTriggerIdObjectsNil sets the value for ReferencedByExecutionTriggerIdObjects to be an explicit nil

### UnsetReferencedByExecutionTriggerIdObjects
`func (o *Trigger) UnsetReferencedByExecutionTriggerIdObjects()`

UnsetReferencedByExecutionTriggerIdObjects ensures that no value is present for ReferencedByExecutionTriggerIdObjects, not even an explicit nil
### GetReferencedByTriggerHasExecutionTriggerIdObjects

`func (o *Trigger) GetReferencedByTriggerHasExecutionTriggerIdObjects() []TriggerHasExecution`

GetReferencedByTriggerHasExecutionTriggerIdObjects returns the ReferencedByTriggerHasExecutionTriggerIdObjects field if non-nil, zero value otherwise.

### GetReferencedByTriggerHasExecutionTriggerIdObjectsOk

`func (o *Trigger) GetReferencedByTriggerHasExecutionTriggerIdObjectsOk() (*[]TriggerHasExecution, bool)`

GetReferencedByTriggerHasExecutionTriggerIdObjectsOk returns a tuple with the ReferencedByTriggerHasExecutionTriggerIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByTriggerHasExecutionTriggerIdObjects

`func (o *Trigger) SetReferencedByTriggerHasExecutionTriggerIdObjects(v []TriggerHasExecution)`

SetReferencedByTriggerHasExecutionTriggerIdObjects sets ReferencedByTriggerHasExecutionTriggerIdObjects field to given value.

### HasReferencedByTriggerHasExecutionTriggerIdObjects

`func (o *Trigger) HasReferencedByTriggerHasExecutionTriggerIdObjects() bool`

HasReferencedByTriggerHasExecutionTriggerIdObjects returns a boolean if a field has been set.

### SetReferencedByTriggerHasExecutionTriggerIdObjectsNil

`func (o *Trigger) SetReferencedByTriggerHasExecutionTriggerIdObjectsNil(b bool)`

 SetReferencedByTriggerHasExecutionTriggerIdObjectsNil sets the value for ReferencedByTriggerHasExecutionTriggerIdObjects to be an explicit nil

### UnsetReferencedByTriggerHasExecutionTriggerIdObjects
`func (o *Trigger) UnsetReferencedByTriggerHasExecutionTriggerIdObjects()`

UnsetReferencedByTriggerHasExecutionTriggerIdObjects ensures that no value is present for ReferencedByTriggerHasExecutionTriggerIdObjects, not even an explicit nil
### GetRuleId

`func (o *Trigger) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *Trigger) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *Trigger) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *Trigger) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleIdObject

`func (o *Trigger) GetRuleIdObject() Rule`

GetRuleIdObject returns the RuleIdObject field if non-nil, zero value otherwise.

### GetRuleIdObjectOk

`func (o *Trigger) GetRuleIdObjectOk() (*Rule, bool)`

GetRuleIdObjectOk returns a tuple with the RuleIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIdObject

`func (o *Trigger) SetRuleIdObject(v Rule)`

SetRuleIdObject sets RuleIdObject field to given value.

### HasRuleIdObject

`func (o *Trigger) HasRuleIdObject() bool`

HasRuleIdObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Trigger) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Trigger) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Trigger) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Trigger) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


