# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildTaskId** | Pointer to **string** |  | [optional] 
**BuildTaskIdObject** | Pointer to [**Task**](Task.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**DeployTaskId** | Pointer to **string** |  | [optional] 
**DeployTaskIdObject** | Pointer to [**Task**](Task.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PublishTaskId** | Pointer to **string** |  | [optional] 
**PublishTaskIdObject** | Pointer to [**Task**](Task.md) |  | [optional] 
**ReferencedByExecutionJobIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**ReferencedByRuleRequiresJobJobIdObjects** | Pointer to [**[]RuleRequiresJob**](RuleRequiresJob.md) |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**RuleIdObject** | Pointer to [**Rule**](Rule.md) |  | [optional] 
**TestTaskId** | Pointer to **string** |  | [optional] 
**TestTaskIdObject** | Pointer to [**Task**](Task.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ValidateTaskId** | Pointer to **string** |  | [optional] 
**ValidateTaskIdObject** | Pointer to [**Task**](Task.md) |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildTaskId

`func (o *Job) GetBuildTaskId() string`

GetBuildTaskId returns the BuildTaskId field if non-nil, zero value otherwise.

### GetBuildTaskIdOk

`func (o *Job) GetBuildTaskIdOk() (*string, bool)`

GetBuildTaskIdOk returns a tuple with the BuildTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTaskId

`func (o *Job) SetBuildTaskId(v string)`

SetBuildTaskId sets BuildTaskId field to given value.

### HasBuildTaskId

`func (o *Job) HasBuildTaskId() bool`

HasBuildTaskId returns a boolean if a field has been set.

### GetBuildTaskIdObject

`func (o *Job) GetBuildTaskIdObject() Task`

GetBuildTaskIdObject returns the BuildTaskIdObject field if non-nil, zero value otherwise.

### GetBuildTaskIdObjectOk

`func (o *Job) GetBuildTaskIdObjectOk() (*Task, bool)`

GetBuildTaskIdObjectOk returns a tuple with the BuildTaskIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTaskIdObject

`func (o *Job) SetBuildTaskIdObject(v Task)`

SetBuildTaskIdObject sets BuildTaskIdObject field to given value.

### HasBuildTaskIdObject

`func (o *Job) HasBuildTaskIdObject() bool`

HasBuildTaskIdObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Job) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Job) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Job) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Job) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Job) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeployTaskId

`func (o *Job) GetDeployTaskId() string`

GetDeployTaskId returns the DeployTaskId field if non-nil, zero value otherwise.

### GetDeployTaskIdOk

`func (o *Job) GetDeployTaskIdOk() (*string, bool)`

GetDeployTaskIdOk returns a tuple with the DeployTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployTaskId

`func (o *Job) SetDeployTaskId(v string)`

SetDeployTaskId sets DeployTaskId field to given value.

### HasDeployTaskId

`func (o *Job) HasDeployTaskId() bool`

HasDeployTaskId returns a boolean if a field has been set.

### GetDeployTaskIdObject

`func (o *Job) GetDeployTaskIdObject() Task`

GetDeployTaskIdObject returns the DeployTaskIdObject field if non-nil, zero value otherwise.

### GetDeployTaskIdObjectOk

`func (o *Job) GetDeployTaskIdObjectOk() (*Task, bool)`

GetDeployTaskIdObjectOk returns a tuple with the DeployTaskIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployTaskIdObject

`func (o *Job) SetDeployTaskIdObject(v Task)`

SetDeployTaskIdObject sets DeployTaskIdObject field to given value.

### HasDeployTaskIdObject

`func (o *Job) HasDeployTaskIdObject() bool`

HasDeployTaskIdObject returns a boolean if a field has been set.

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublishTaskId

`func (o *Job) GetPublishTaskId() string`

GetPublishTaskId returns the PublishTaskId field if non-nil, zero value otherwise.

### GetPublishTaskIdOk

`func (o *Job) GetPublishTaskIdOk() (*string, bool)`

GetPublishTaskIdOk returns a tuple with the PublishTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTaskId

`func (o *Job) SetPublishTaskId(v string)`

SetPublishTaskId sets PublishTaskId field to given value.

### HasPublishTaskId

`func (o *Job) HasPublishTaskId() bool`

HasPublishTaskId returns a boolean if a field has been set.

### GetPublishTaskIdObject

`func (o *Job) GetPublishTaskIdObject() Task`

GetPublishTaskIdObject returns the PublishTaskIdObject field if non-nil, zero value otherwise.

### GetPublishTaskIdObjectOk

`func (o *Job) GetPublishTaskIdObjectOk() (*Task, bool)`

GetPublishTaskIdObjectOk returns a tuple with the PublishTaskIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTaskIdObject

`func (o *Job) SetPublishTaskIdObject(v Task)`

SetPublishTaskIdObject sets PublishTaskIdObject field to given value.

### HasPublishTaskIdObject

`func (o *Job) HasPublishTaskIdObject() bool`

HasPublishTaskIdObject returns a boolean if a field has been set.

### GetReferencedByExecutionJobIdObjects

`func (o *Job) GetReferencedByExecutionJobIdObjects() []Execution`

GetReferencedByExecutionJobIdObjects returns the ReferencedByExecutionJobIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionJobIdObjectsOk

`func (o *Job) GetReferencedByExecutionJobIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionJobIdObjectsOk returns a tuple with the ReferencedByExecutionJobIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionJobIdObjects

`func (o *Job) SetReferencedByExecutionJobIdObjects(v []Execution)`

SetReferencedByExecutionJobIdObjects sets ReferencedByExecutionJobIdObjects field to given value.

### HasReferencedByExecutionJobIdObjects

`func (o *Job) HasReferencedByExecutionJobIdObjects() bool`

HasReferencedByExecutionJobIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionJobIdObjectsNil

`func (o *Job) SetReferencedByExecutionJobIdObjectsNil(b bool)`

 SetReferencedByExecutionJobIdObjectsNil sets the value for ReferencedByExecutionJobIdObjects to be an explicit nil

### UnsetReferencedByExecutionJobIdObjects
`func (o *Job) UnsetReferencedByExecutionJobIdObjects()`

UnsetReferencedByExecutionJobIdObjects ensures that no value is present for ReferencedByExecutionJobIdObjects, not even an explicit nil
### GetReferencedByRuleRequiresJobJobIdObjects

`func (o *Job) GetReferencedByRuleRequiresJobJobIdObjects() []RuleRequiresJob`

GetReferencedByRuleRequiresJobJobIdObjects returns the ReferencedByRuleRequiresJobJobIdObjects field if non-nil, zero value otherwise.

### GetReferencedByRuleRequiresJobJobIdObjectsOk

`func (o *Job) GetReferencedByRuleRequiresJobJobIdObjectsOk() (*[]RuleRequiresJob, bool)`

GetReferencedByRuleRequiresJobJobIdObjectsOk returns a tuple with the ReferencedByRuleRequiresJobJobIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByRuleRequiresJobJobIdObjects

`func (o *Job) SetReferencedByRuleRequiresJobJobIdObjects(v []RuleRequiresJob)`

SetReferencedByRuleRequiresJobJobIdObjects sets ReferencedByRuleRequiresJobJobIdObjects field to given value.

### HasReferencedByRuleRequiresJobJobIdObjects

`func (o *Job) HasReferencedByRuleRequiresJobJobIdObjects() bool`

HasReferencedByRuleRequiresJobJobIdObjects returns a boolean if a field has been set.

### SetReferencedByRuleRequiresJobJobIdObjectsNil

`func (o *Job) SetReferencedByRuleRequiresJobJobIdObjectsNil(b bool)`

 SetReferencedByRuleRequiresJobJobIdObjectsNil sets the value for ReferencedByRuleRequiresJobJobIdObjects to be an explicit nil

### UnsetReferencedByRuleRequiresJobJobIdObjects
`func (o *Job) UnsetReferencedByRuleRequiresJobJobIdObjects()`

UnsetReferencedByRuleRequiresJobJobIdObjects ensures that no value is present for ReferencedByRuleRequiresJobJobIdObjects, not even an explicit nil
### GetRuleId

`func (o *Job) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *Job) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *Job) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *Job) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleIdObject

`func (o *Job) GetRuleIdObject() Rule`

GetRuleIdObject returns the RuleIdObject field if non-nil, zero value otherwise.

### GetRuleIdObjectOk

`func (o *Job) GetRuleIdObjectOk() (*Rule, bool)`

GetRuleIdObjectOk returns a tuple with the RuleIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIdObject

`func (o *Job) SetRuleIdObject(v Rule)`

SetRuleIdObject sets RuleIdObject field to given value.

### HasRuleIdObject

`func (o *Job) HasRuleIdObject() bool`

HasRuleIdObject returns a boolean if a field has been set.

### GetTestTaskId

`func (o *Job) GetTestTaskId() string`

GetTestTaskId returns the TestTaskId field if non-nil, zero value otherwise.

### GetTestTaskIdOk

`func (o *Job) GetTestTaskIdOk() (*string, bool)`

GetTestTaskIdOk returns a tuple with the TestTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTaskId

`func (o *Job) SetTestTaskId(v string)`

SetTestTaskId sets TestTaskId field to given value.

### HasTestTaskId

`func (o *Job) HasTestTaskId() bool`

HasTestTaskId returns a boolean if a field has been set.

### GetTestTaskIdObject

`func (o *Job) GetTestTaskIdObject() Task`

GetTestTaskIdObject returns the TestTaskIdObject field if non-nil, zero value otherwise.

### GetTestTaskIdObjectOk

`func (o *Job) GetTestTaskIdObjectOk() (*Task, bool)`

GetTestTaskIdObjectOk returns a tuple with the TestTaskIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTaskIdObject

`func (o *Job) SetTestTaskIdObject(v Task)`

SetTestTaskIdObject sets TestTaskIdObject field to given value.

### HasTestTaskIdObject

`func (o *Job) HasTestTaskIdObject() bool`

HasTestTaskIdObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Job) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Job) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Job) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Job) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValidateTaskId

`func (o *Job) GetValidateTaskId() string`

GetValidateTaskId returns the ValidateTaskId field if non-nil, zero value otherwise.

### GetValidateTaskIdOk

`func (o *Job) GetValidateTaskIdOk() (*string, bool)`

GetValidateTaskIdOk returns a tuple with the ValidateTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTaskId

`func (o *Job) SetValidateTaskId(v string)`

SetValidateTaskId sets ValidateTaskId field to given value.

### HasValidateTaskId

`func (o *Job) HasValidateTaskId() bool`

HasValidateTaskId returns a boolean if a field has been set.

### GetValidateTaskIdObject

`func (o *Job) GetValidateTaskIdObject() Task`

GetValidateTaskIdObject returns the ValidateTaskIdObject field if non-nil, zero value otherwise.

### GetValidateTaskIdObjectOk

`func (o *Job) GetValidateTaskIdObjectOk() (*Task, bool)`

GetValidateTaskIdObjectOk returns a tuple with the ValidateTaskIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTaskIdObject

`func (o *Job) SetValidateTaskIdObject(v Task)`

SetValidateTaskIdObject sets ValidateTaskIdObject field to given value.

### HasValidateTaskIdObject

`func (o *Job) HasValidateTaskIdObject() bool`

HasValidateTaskIdObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


