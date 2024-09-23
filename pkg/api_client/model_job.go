/*
Djangolang

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_client

import (
	"encoding/json"
	"time"
)

// checks if the Job type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Job{}

// Job struct for Job
type Job struct {
	BuildTaskId                             *string           `json:"build_task_id,omitempty"`
	BuildTaskIdObject                       *Task             `json:"build_task_id_object,omitempty"`
	CreatedAt                               *time.Time        `json:"created_at,omitempty"`
	DeletedAt                               *time.Time        `json:"deleted_at,omitempty"`
	DeployTaskId                            *string           `json:"deploy_task_id,omitempty"`
	DeployTaskIdObject                      *Task             `json:"deploy_task_id_object,omitempty"`
	Id                                      *string           `json:"id,omitempty"`
	PublishTaskId                           *string           `json:"publish_task_id,omitempty"`
	PublishTaskIdObject                     *Task             `json:"publish_task_id_object,omitempty"`
	ReferencedByExecutionJobIdObjects       []Execution       `json:"referenced_by_execution_job_id_objects,omitempty"`
	ReferencedByRuleRequiresJobJobIdObjects []RuleRequiresJob `json:"referenced_by_rule_requires_job_job_id_objects,omitempty"`
	RuleId                                  *string           `json:"rule_id,omitempty"`
	RuleIdObject                            *Rule             `json:"rule_id_object,omitempty"`
	TestTaskId                              *string           `json:"test_task_id,omitempty"`
	TestTaskIdObject                        *Task             `json:"test_task_id_object,omitempty"`
	UpdatedAt                               *time.Time        `json:"updated_at,omitempty"`
	ValidateTaskId                          *string           `json:"validate_task_id,omitempty"`
	ValidateTaskIdObject                    *Task             `json:"validate_task_id_object,omitempty"`
}

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob() *Job {
	this := Job{}
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetBuildTaskId returns the BuildTaskId field value if set, zero value otherwise.
func (o *Job) GetBuildTaskId() string {
	if o == nil || IsNil(o.BuildTaskId) {
		var ret string
		return ret
	}
	return *o.BuildTaskId
}

// GetBuildTaskIdOk returns a tuple with the BuildTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetBuildTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuildTaskId) {
		return nil, false
	}
	return o.BuildTaskId, true
}

// HasBuildTaskId returns a boolean if a field has been set.
func (o *Job) HasBuildTaskId() bool {
	if o != nil && !IsNil(o.BuildTaskId) {
		return true
	}

	return false
}

// SetBuildTaskId gets a reference to the given string and assigns it to the BuildTaskId field.
func (o *Job) SetBuildTaskId(v string) {
	o.BuildTaskId = &v
}

// GetBuildTaskIdObject returns the BuildTaskIdObject field value if set, zero value otherwise.
func (o *Job) GetBuildTaskIdObject() Task {
	if o == nil || IsNil(o.BuildTaskIdObject) {
		var ret Task
		return ret
	}
	return *o.BuildTaskIdObject
}

// GetBuildTaskIdObjectOk returns a tuple with the BuildTaskIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetBuildTaskIdObjectOk() (*Task, bool) {
	if o == nil || IsNil(o.BuildTaskIdObject) {
		return nil, false
	}
	return o.BuildTaskIdObject, true
}

// HasBuildTaskIdObject returns a boolean if a field has been set.
func (o *Job) HasBuildTaskIdObject() bool {
	if o != nil && !IsNil(o.BuildTaskIdObject) {
		return true
	}

	return false
}

// SetBuildTaskIdObject gets a reference to the given Task and assigns it to the BuildTaskIdObject field.
func (o *Job) SetBuildTaskIdObject(v Task) {
	o.BuildTaskIdObject = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Job) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Job) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Job) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Job) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Job) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Job) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDeployTaskId returns the DeployTaskId field value if set, zero value otherwise.
func (o *Job) GetDeployTaskId() string {
	if o == nil || IsNil(o.DeployTaskId) {
		var ret string
		return ret
	}
	return *o.DeployTaskId
}

// GetDeployTaskIdOk returns a tuple with the DeployTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetDeployTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeployTaskId) {
		return nil, false
	}
	return o.DeployTaskId, true
}

// HasDeployTaskId returns a boolean if a field has been set.
func (o *Job) HasDeployTaskId() bool {
	if o != nil && !IsNil(o.DeployTaskId) {
		return true
	}

	return false
}

// SetDeployTaskId gets a reference to the given string and assigns it to the DeployTaskId field.
func (o *Job) SetDeployTaskId(v string) {
	o.DeployTaskId = &v
}

// GetDeployTaskIdObject returns the DeployTaskIdObject field value if set, zero value otherwise.
func (o *Job) GetDeployTaskIdObject() Task {
	if o == nil || IsNil(o.DeployTaskIdObject) {
		var ret Task
		return ret
	}
	return *o.DeployTaskIdObject
}

// GetDeployTaskIdObjectOk returns a tuple with the DeployTaskIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetDeployTaskIdObjectOk() (*Task, bool) {
	if o == nil || IsNil(o.DeployTaskIdObject) {
		return nil, false
	}
	return o.DeployTaskIdObject, true
}

// HasDeployTaskIdObject returns a boolean if a field has been set.
func (o *Job) HasDeployTaskIdObject() bool {
	if o != nil && !IsNil(o.DeployTaskIdObject) {
		return true
	}

	return false
}

// SetDeployTaskIdObject gets a reference to the given Task and assigns it to the DeployTaskIdObject field.
func (o *Job) SetDeployTaskIdObject(v Task) {
	o.DeployTaskIdObject = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Job) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Job) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Job) SetId(v string) {
	o.Id = &v
}

// GetPublishTaskId returns the PublishTaskId field value if set, zero value otherwise.
func (o *Job) GetPublishTaskId() string {
	if o == nil || IsNil(o.PublishTaskId) {
		var ret string
		return ret
	}
	return *o.PublishTaskId
}

// GetPublishTaskIdOk returns a tuple with the PublishTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetPublishTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublishTaskId) {
		return nil, false
	}
	return o.PublishTaskId, true
}

// HasPublishTaskId returns a boolean if a field has been set.
func (o *Job) HasPublishTaskId() bool {
	if o != nil && !IsNil(o.PublishTaskId) {
		return true
	}

	return false
}

// SetPublishTaskId gets a reference to the given string and assigns it to the PublishTaskId field.
func (o *Job) SetPublishTaskId(v string) {
	o.PublishTaskId = &v
}

// GetPublishTaskIdObject returns the PublishTaskIdObject field value if set, zero value otherwise.
func (o *Job) GetPublishTaskIdObject() Task {
	if o == nil || IsNil(o.PublishTaskIdObject) {
		var ret Task
		return ret
	}
	return *o.PublishTaskIdObject
}

// GetPublishTaskIdObjectOk returns a tuple with the PublishTaskIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetPublishTaskIdObjectOk() (*Task, bool) {
	if o == nil || IsNil(o.PublishTaskIdObject) {
		return nil, false
	}
	return o.PublishTaskIdObject, true
}

// HasPublishTaskIdObject returns a boolean if a field has been set.
func (o *Job) HasPublishTaskIdObject() bool {
	if o != nil && !IsNil(o.PublishTaskIdObject) {
		return true
	}

	return false
}

// SetPublishTaskIdObject gets a reference to the given Task and assigns it to the PublishTaskIdObject field.
func (o *Job) SetPublishTaskIdObject(v Task) {
	o.PublishTaskIdObject = &v
}

// GetReferencedByExecutionJobIdObjects returns the ReferencedByExecutionJobIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetReferencedByExecutionJobIdObjects() []Execution {
	if o == nil {
		var ret []Execution
		return ret
	}
	return o.ReferencedByExecutionJobIdObjects
}

// GetReferencedByExecutionJobIdObjectsOk returns a tuple with the ReferencedByExecutionJobIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetReferencedByExecutionJobIdObjectsOk() ([]Execution, bool) {
	if o == nil || IsNil(o.ReferencedByExecutionJobIdObjects) {
		return nil, false
	}
	return o.ReferencedByExecutionJobIdObjects, true
}

// HasReferencedByExecutionJobIdObjects returns a boolean if a field has been set.
func (o *Job) HasReferencedByExecutionJobIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByExecutionJobIdObjects) {
		return true
	}

	return false
}

// SetReferencedByExecutionJobIdObjects gets a reference to the given []Execution and assigns it to the ReferencedByExecutionJobIdObjects field.
func (o *Job) SetReferencedByExecutionJobIdObjects(v []Execution) {
	o.ReferencedByExecutionJobIdObjects = v
}

// GetReferencedByRuleRequiresJobJobIdObjects returns the ReferencedByRuleRequiresJobJobIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetReferencedByRuleRequiresJobJobIdObjects() []RuleRequiresJob {
	if o == nil {
		var ret []RuleRequiresJob
		return ret
	}
	return o.ReferencedByRuleRequiresJobJobIdObjects
}

// GetReferencedByRuleRequiresJobJobIdObjectsOk returns a tuple with the ReferencedByRuleRequiresJobJobIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetReferencedByRuleRequiresJobJobIdObjectsOk() ([]RuleRequiresJob, bool) {
	if o == nil || IsNil(o.ReferencedByRuleRequiresJobJobIdObjects) {
		return nil, false
	}
	return o.ReferencedByRuleRequiresJobJobIdObjects, true
}

// HasReferencedByRuleRequiresJobJobIdObjects returns a boolean if a field has been set.
func (o *Job) HasReferencedByRuleRequiresJobJobIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByRuleRequiresJobJobIdObjects) {
		return true
	}

	return false
}

// SetReferencedByRuleRequiresJobJobIdObjects gets a reference to the given []RuleRequiresJob and assigns it to the ReferencedByRuleRequiresJobJobIdObjects field.
func (o *Job) SetReferencedByRuleRequiresJobJobIdObjects(v []RuleRequiresJob) {
	o.ReferencedByRuleRequiresJobJobIdObjects = v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *Job) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *Job) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *Job) SetRuleId(v string) {
	o.RuleId = &v
}

// GetRuleIdObject returns the RuleIdObject field value if set, zero value otherwise.
func (o *Job) GetRuleIdObject() Rule {
	if o == nil || IsNil(o.RuleIdObject) {
		var ret Rule
		return ret
	}
	return *o.RuleIdObject
}

// GetRuleIdObjectOk returns a tuple with the RuleIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetRuleIdObjectOk() (*Rule, bool) {
	if o == nil || IsNil(o.RuleIdObject) {
		return nil, false
	}
	return o.RuleIdObject, true
}

// HasRuleIdObject returns a boolean if a field has been set.
func (o *Job) HasRuleIdObject() bool {
	if o != nil && !IsNil(o.RuleIdObject) {
		return true
	}

	return false
}

// SetRuleIdObject gets a reference to the given Rule and assigns it to the RuleIdObject field.
func (o *Job) SetRuleIdObject(v Rule) {
	o.RuleIdObject = &v
}

// GetTestTaskId returns the TestTaskId field value if set, zero value otherwise.
func (o *Job) GetTestTaskId() string {
	if o == nil || IsNil(o.TestTaskId) {
		var ret string
		return ret
	}
	return *o.TestTaskId
}

// GetTestTaskIdOk returns a tuple with the TestTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetTestTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestTaskId) {
		return nil, false
	}
	return o.TestTaskId, true
}

// HasTestTaskId returns a boolean if a field has been set.
func (o *Job) HasTestTaskId() bool {
	if o != nil && !IsNil(o.TestTaskId) {
		return true
	}

	return false
}

// SetTestTaskId gets a reference to the given string and assigns it to the TestTaskId field.
func (o *Job) SetTestTaskId(v string) {
	o.TestTaskId = &v
}

// GetTestTaskIdObject returns the TestTaskIdObject field value if set, zero value otherwise.
func (o *Job) GetTestTaskIdObject() Task {
	if o == nil || IsNil(o.TestTaskIdObject) {
		var ret Task
		return ret
	}
	return *o.TestTaskIdObject
}

// GetTestTaskIdObjectOk returns a tuple with the TestTaskIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetTestTaskIdObjectOk() (*Task, bool) {
	if o == nil || IsNil(o.TestTaskIdObject) {
		return nil, false
	}
	return o.TestTaskIdObject, true
}

// HasTestTaskIdObject returns a boolean if a field has been set.
func (o *Job) HasTestTaskIdObject() bool {
	if o != nil && !IsNil(o.TestTaskIdObject) {
		return true
	}

	return false
}

// SetTestTaskIdObject gets a reference to the given Task and assigns it to the TestTaskIdObject field.
func (o *Job) SetTestTaskIdObject(v Task) {
	o.TestTaskIdObject = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Job) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Job) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Job) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValidateTaskId returns the ValidateTaskId field value if set, zero value otherwise.
func (o *Job) GetValidateTaskId() string {
	if o == nil || IsNil(o.ValidateTaskId) {
		var ret string
		return ret
	}
	return *o.ValidateTaskId
}

// GetValidateTaskIdOk returns a tuple with the ValidateTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetValidateTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.ValidateTaskId) {
		return nil, false
	}
	return o.ValidateTaskId, true
}

// HasValidateTaskId returns a boolean if a field has been set.
func (o *Job) HasValidateTaskId() bool {
	if o != nil && !IsNil(o.ValidateTaskId) {
		return true
	}

	return false
}

// SetValidateTaskId gets a reference to the given string and assigns it to the ValidateTaskId field.
func (o *Job) SetValidateTaskId(v string) {
	o.ValidateTaskId = &v
}

// GetValidateTaskIdObject returns the ValidateTaskIdObject field value if set, zero value otherwise.
func (o *Job) GetValidateTaskIdObject() Task {
	if o == nil || IsNil(o.ValidateTaskIdObject) {
		var ret Task
		return ret
	}
	return *o.ValidateTaskIdObject
}

// GetValidateTaskIdObjectOk returns a tuple with the ValidateTaskIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetValidateTaskIdObjectOk() (*Task, bool) {
	if o == nil || IsNil(o.ValidateTaskIdObject) {
		return nil, false
	}
	return o.ValidateTaskIdObject, true
}

// HasValidateTaskIdObject returns a boolean if a field has been set.
func (o *Job) HasValidateTaskIdObject() bool {
	if o != nil && !IsNil(o.ValidateTaskIdObject) {
		return true
	}

	return false
}

// SetValidateTaskIdObject gets a reference to the given Task and assigns it to the ValidateTaskIdObject field.
func (o *Job) SetValidateTaskIdObject(v Task) {
	o.ValidateTaskIdObject = &v
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Job) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildTaskId) {
		toSerialize["build_task_id"] = o.BuildTaskId
	}
	if !IsNil(o.BuildTaskIdObject) {
		toSerialize["build_task_id_object"] = o.BuildTaskIdObject
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.DeployTaskId) {
		toSerialize["deploy_task_id"] = o.DeployTaskId
	}
	if !IsNil(o.DeployTaskIdObject) {
		toSerialize["deploy_task_id_object"] = o.DeployTaskIdObject
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PublishTaskId) {
		toSerialize["publish_task_id"] = o.PublishTaskId
	}
	if !IsNil(o.PublishTaskIdObject) {
		toSerialize["publish_task_id_object"] = o.PublishTaskIdObject
	}
	if o.ReferencedByExecutionJobIdObjects != nil {
		toSerialize["referenced_by_execution_job_id_objects"] = o.ReferencedByExecutionJobIdObjects
	}
	if o.ReferencedByRuleRequiresJobJobIdObjects != nil {
		toSerialize["referenced_by_rule_requires_job_job_id_objects"] = o.ReferencedByRuleRequiresJobJobIdObjects
	}
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	if !IsNil(o.RuleIdObject) {
		toSerialize["rule_id_object"] = o.RuleIdObject
	}
	if !IsNil(o.TestTaskId) {
		toSerialize["test_task_id"] = o.TestTaskId
	}
	if !IsNil(o.TestTaskIdObject) {
		toSerialize["test_task_id_object"] = o.TestTaskIdObject
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ValidateTaskId) {
		toSerialize["validate_task_id"] = o.ValidateTaskId
	}
	if !IsNil(o.ValidateTaskIdObject) {
		toSerialize["validate_task_id_object"] = o.ValidateTaskIdObject
	}
	return toSerialize, nil
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
