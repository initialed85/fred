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

// checks if the Execution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Execution{}

// Execution struct for Execution
type Execution struct {
	BuildOutputId *string `json:"build_output_id,omitempty"`
	BuildOutputIdObject *Output `json:"build_output_id_object,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	DeployOutputId *string `json:"deploy_output_id,omitempty"`
	DeployOutputIdObject *Output `json:"deploy_output_id_object,omitempty"`
	Id *string `json:"id,omitempty"`
	JobId *string `json:"job_id,omitempty"`
	JobIdObject *Job `json:"job_id_object,omitempty"`
	PublishOutputId *string `json:"publish_output_id,omitempty"`
	PublishOutputIdObject *Output `json:"publish_output_id_object,omitempty"`
	ReferencedByTriggerHasExecutionExecutionIdObjects []TriggerHasExecution `json:"referenced_by_trigger_has_execution_execution_id_objects,omitempty"`
	Status *string `json:"status,omitempty"`
	TestOutputId *string `json:"test_output_id,omitempty"`
	TestOutputIdObject *Output `json:"test_output_id_object,omitempty"`
	TriggerId *string `json:"trigger_id,omitempty"`
	TriggerIdObject *Trigger `json:"trigger_id_object,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	ValidateOutputId *string `json:"validate_output_id,omitempty"`
	ValidateOutputIdObject *Output `json:"validate_output_id_object,omitempty"`
}

// NewExecution instantiates a new Execution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecution() *Execution {
	this := Execution{}
	return &this
}

// NewExecutionWithDefaults instantiates a new Execution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionWithDefaults() *Execution {
	this := Execution{}
	return &this
}

// GetBuildOutputId returns the BuildOutputId field value if set, zero value otherwise.
func (o *Execution) GetBuildOutputId() string {
	if o == nil || IsNil(o.BuildOutputId) {
		var ret string
		return ret
	}
	return *o.BuildOutputId
}

// GetBuildOutputIdOk returns a tuple with the BuildOutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetBuildOutputIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuildOutputId) {
		return nil, false
	}
	return o.BuildOutputId, true
}

// HasBuildOutputId returns a boolean if a field has been set.
func (o *Execution) HasBuildOutputId() bool {
	if o != nil && !IsNil(o.BuildOutputId) {
		return true
	}

	return false
}

// SetBuildOutputId gets a reference to the given string and assigns it to the BuildOutputId field.
func (o *Execution) SetBuildOutputId(v string) {
	o.BuildOutputId = &v
}

// GetBuildOutputIdObject returns the BuildOutputIdObject field value if set, zero value otherwise.
func (o *Execution) GetBuildOutputIdObject() Output {
	if o == nil || IsNil(o.BuildOutputIdObject) {
		var ret Output
		return ret
	}
	return *o.BuildOutputIdObject
}

// GetBuildOutputIdObjectOk returns a tuple with the BuildOutputIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetBuildOutputIdObjectOk() (*Output, bool) {
	if o == nil || IsNil(o.BuildOutputIdObject) {
		return nil, false
	}
	return o.BuildOutputIdObject, true
}

// HasBuildOutputIdObject returns a boolean if a field has been set.
func (o *Execution) HasBuildOutputIdObject() bool {
	if o != nil && !IsNil(o.BuildOutputIdObject) {
		return true
	}

	return false
}

// SetBuildOutputIdObject gets a reference to the given Output and assigns it to the BuildOutputIdObject field.
func (o *Execution) SetBuildOutputIdObject(v Output) {
	o.BuildOutputIdObject = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Execution) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Execution) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Execution) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Execution) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Execution) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Execution) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDeployOutputId returns the DeployOutputId field value if set, zero value otherwise.
func (o *Execution) GetDeployOutputId() string {
	if o == nil || IsNil(o.DeployOutputId) {
		var ret string
		return ret
	}
	return *o.DeployOutputId
}

// GetDeployOutputIdOk returns a tuple with the DeployOutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetDeployOutputIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeployOutputId) {
		return nil, false
	}
	return o.DeployOutputId, true
}

// HasDeployOutputId returns a boolean if a field has been set.
func (o *Execution) HasDeployOutputId() bool {
	if o != nil && !IsNil(o.DeployOutputId) {
		return true
	}

	return false
}

// SetDeployOutputId gets a reference to the given string and assigns it to the DeployOutputId field.
func (o *Execution) SetDeployOutputId(v string) {
	o.DeployOutputId = &v
}

// GetDeployOutputIdObject returns the DeployOutputIdObject field value if set, zero value otherwise.
func (o *Execution) GetDeployOutputIdObject() Output {
	if o == nil || IsNil(o.DeployOutputIdObject) {
		var ret Output
		return ret
	}
	return *o.DeployOutputIdObject
}

// GetDeployOutputIdObjectOk returns a tuple with the DeployOutputIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetDeployOutputIdObjectOk() (*Output, bool) {
	if o == nil || IsNil(o.DeployOutputIdObject) {
		return nil, false
	}
	return o.DeployOutputIdObject, true
}

// HasDeployOutputIdObject returns a boolean if a field has been set.
func (o *Execution) HasDeployOutputIdObject() bool {
	if o != nil && !IsNil(o.DeployOutputIdObject) {
		return true
	}

	return false
}

// SetDeployOutputIdObject gets a reference to the given Output and assigns it to the DeployOutputIdObject field.
func (o *Execution) SetDeployOutputIdObject(v Output) {
	o.DeployOutputIdObject = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Execution) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Execution) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Execution) SetId(v string) {
	o.Id = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *Execution) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *Execution) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *Execution) SetJobId(v string) {
	o.JobId = &v
}

// GetJobIdObject returns the JobIdObject field value if set, zero value otherwise.
func (o *Execution) GetJobIdObject() Job {
	if o == nil || IsNil(o.JobIdObject) {
		var ret Job
		return ret
	}
	return *o.JobIdObject
}

// GetJobIdObjectOk returns a tuple with the JobIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetJobIdObjectOk() (*Job, bool) {
	if o == nil || IsNil(o.JobIdObject) {
		return nil, false
	}
	return o.JobIdObject, true
}

// HasJobIdObject returns a boolean if a field has been set.
func (o *Execution) HasJobIdObject() bool {
	if o != nil && !IsNil(o.JobIdObject) {
		return true
	}

	return false
}

// SetJobIdObject gets a reference to the given Job and assigns it to the JobIdObject field.
func (o *Execution) SetJobIdObject(v Job) {
	o.JobIdObject = &v
}

// GetPublishOutputId returns the PublishOutputId field value if set, zero value otherwise.
func (o *Execution) GetPublishOutputId() string {
	if o == nil || IsNil(o.PublishOutputId) {
		var ret string
		return ret
	}
	return *o.PublishOutputId
}

// GetPublishOutputIdOk returns a tuple with the PublishOutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetPublishOutputIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublishOutputId) {
		return nil, false
	}
	return o.PublishOutputId, true
}

// HasPublishOutputId returns a boolean if a field has been set.
func (o *Execution) HasPublishOutputId() bool {
	if o != nil && !IsNil(o.PublishOutputId) {
		return true
	}

	return false
}

// SetPublishOutputId gets a reference to the given string and assigns it to the PublishOutputId field.
func (o *Execution) SetPublishOutputId(v string) {
	o.PublishOutputId = &v
}

// GetPublishOutputIdObject returns the PublishOutputIdObject field value if set, zero value otherwise.
func (o *Execution) GetPublishOutputIdObject() Output {
	if o == nil || IsNil(o.PublishOutputIdObject) {
		var ret Output
		return ret
	}
	return *o.PublishOutputIdObject
}

// GetPublishOutputIdObjectOk returns a tuple with the PublishOutputIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetPublishOutputIdObjectOk() (*Output, bool) {
	if o == nil || IsNil(o.PublishOutputIdObject) {
		return nil, false
	}
	return o.PublishOutputIdObject, true
}

// HasPublishOutputIdObject returns a boolean if a field has been set.
func (o *Execution) HasPublishOutputIdObject() bool {
	if o != nil && !IsNil(o.PublishOutputIdObject) {
		return true
	}

	return false
}

// SetPublishOutputIdObject gets a reference to the given Output and assigns it to the PublishOutputIdObject field.
func (o *Execution) SetPublishOutputIdObject(v Output) {
	o.PublishOutputIdObject = &v
}

// GetReferencedByTriggerHasExecutionExecutionIdObjects returns the ReferencedByTriggerHasExecutionExecutionIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Execution) GetReferencedByTriggerHasExecutionExecutionIdObjects() []TriggerHasExecution {
	if o == nil {
		var ret []TriggerHasExecution
		return ret
	}
	return o.ReferencedByTriggerHasExecutionExecutionIdObjects
}

// GetReferencedByTriggerHasExecutionExecutionIdObjectsOk returns a tuple with the ReferencedByTriggerHasExecutionExecutionIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Execution) GetReferencedByTriggerHasExecutionExecutionIdObjectsOk() ([]TriggerHasExecution, bool) {
	if o == nil || IsNil(o.ReferencedByTriggerHasExecutionExecutionIdObjects) {
		return nil, false
	}
	return o.ReferencedByTriggerHasExecutionExecutionIdObjects, true
}

// HasReferencedByTriggerHasExecutionExecutionIdObjects returns a boolean if a field has been set.
func (o *Execution) HasReferencedByTriggerHasExecutionExecutionIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByTriggerHasExecutionExecutionIdObjects) {
		return true
	}

	return false
}

// SetReferencedByTriggerHasExecutionExecutionIdObjects gets a reference to the given []TriggerHasExecution and assigns it to the ReferencedByTriggerHasExecutionExecutionIdObjects field.
func (o *Execution) SetReferencedByTriggerHasExecutionExecutionIdObjects(v []TriggerHasExecution) {
	o.ReferencedByTriggerHasExecutionExecutionIdObjects = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Execution) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Execution) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Execution) SetStatus(v string) {
	o.Status = &v
}

// GetTestOutputId returns the TestOutputId field value if set, zero value otherwise.
func (o *Execution) GetTestOutputId() string {
	if o == nil || IsNil(o.TestOutputId) {
		var ret string
		return ret
	}
	return *o.TestOutputId
}

// GetTestOutputIdOk returns a tuple with the TestOutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetTestOutputIdOk() (*string, bool) {
	if o == nil || IsNil(o.TestOutputId) {
		return nil, false
	}
	return o.TestOutputId, true
}

// HasTestOutputId returns a boolean if a field has been set.
func (o *Execution) HasTestOutputId() bool {
	if o != nil && !IsNil(o.TestOutputId) {
		return true
	}

	return false
}

// SetTestOutputId gets a reference to the given string and assigns it to the TestOutputId field.
func (o *Execution) SetTestOutputId(v string) {
	o.TestOutputId = &v
}

// GetTestOutputIdObject returns the TestOutputIdObject field value if set, zero value otherwise.
func (o *Execution) GetTestOutputIdObject() Output {
	if o == nil || IsNil(o.TestOutputIdObject) {
		var ret Output
		return ret
	}
	return *o.TestOutputIdObject
}

// GetTestOutputIdObjectOk returns a tuple with the TestOutputIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetTestOutputIdObjectOk() (*Output, bool) {
	if o == nil || IsNil(o.TestOutputIdObject) {
		return nil, false
	}
	return o.TestOutputIdObject, true
}

// HasTestOutputIdObject returns a boolean if a field has been set.
func (o *Execution) HasTestOutputIdObject() bool {
	if o != nil && !IsNil(o.TestOutputIdObject) {
		return true
	}

	return false
}

// SetTestOutputIdObject gets a reference to the given Output and assigns it to the TestOutputIdObject field.
func (o *Execution) SetTestOutputIdObject(v Output) {
	o.TestOutputIdObject = &v
}

// GetTriggerId returns the TriggerId field value if set, zero value otherwise.
func (o *Execution) GetTriggerId() string {
	if o == nil || IsNil(o.TriggerId) {
		var ret string
		return ret
	}
	return *o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetTriggerIdOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerId) {
		return nil, false
	}
	return o.TriggerId, true
}

// HasTriggerId returns a boolean if a field has been set.
func (o *Execution) HasTriggerId() bool {
	if o != nil && !IsNil(o.TriggerId) {
		return true
	}

	return false
}

// SetTriggerId gets a reference to the given string and assigns it to the TriggerId field.
func (o *Execution) SetTriggerId(v string) {
	o.TriggerId = &v
}

// GetTriggerIdObject returns the TriggerIdObject field value if set, zero value otherwise.
func (o *Execution) GetTriggerIdObject() Trigger {
	if o == nil || IsNil(o.TriggerIdObject) {
		var ret Trigger
		return ret
	}
	return *o.TriggerIdObject
}

// GetTriggerIdObjectOk returns a tuple with the TriggerIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetTriggerIdObjectOk() (*Trigger, bool) {
	if o == nil || IsNil(o.TriggerIdObject) {
		return nil, false
	}
	return o.TriggerIdObject, true
}

// HasTriggerIdObject returns a boolean if a field has been set.
func (o *Execution) HasTriggerIdObject() bool {
	if o != nil && !IsNil(o.TriggerIdObject) {
		return true
	}

	return false
}

// SetTriggerIdObject gets a reference to the given Trigger and assigns it to the TriggerIdObject field.
func (o *Execution) SetTriggerIdObject(v Trigger) {
	o.TriggerIdObject = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Execution) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Execution) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Execution) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValidateOutputId returns the ValidateOutputId field value if set, zero value otherwise.
func (o *Execution) GetValidateOutputId() string {
	if o == nil || IsNil(o.ValidateOutputId) {
		var ret string
		return ret
	}
	return *o.ValidateOutputId
}

// GetValidateOutputIdOk returns a tuple with the ValidateOutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetValidateOutputIdOk() (*string, bool) {
	if o == nil || IsNil(o.ValidateOutputId) {
		return nil, false
	}
	return o.ValidateOutputId, true
}

// HasValidateOutputId returns a boolean if a field has been set.
func (o *Execution) HasValidateOutputId() bool {
	if o != nil && !IsNil(o.ValidateOutputId) {
		return true
	}

	return false
}

// SetValidateOutputId gets a reference to the given string and assigns it to the ValidateOutputId field.
func (o *Execution) SetValidateOutputId(v string) {
	o.ValidateOutputId = &v
}

// GetValidateOutputIdObject returns the ValidateOutputIdObject field value if set, zero value otherwise.
func (o *Execution) GetValidateOutputIdObject() Output {
	if o == nil || IsNil(o.ValidateOutputIdObject) {
		var ret Output
		return ret
	}
	return *o.ValidateOutputIdObject
}

// GetValidateOutputIdObjectOk returns a tuple with the ValidateOutputIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetValidateOutputIdObjectOk() (*Output, bool) {
	if o == nil || IsNil(o.ValidateOutputIdObject) {
		return nil, false
	}
	return o.ValidateOutputIdObject, true
}

// HasValidateOutputIdObject returns a boolean if a field has been set.
func (o *Execution) HasValidateOutputIdObject() bool {
	if o != nil && !IsNil(o.ValidateOutputIdObject) {
		return true
	}

	return false
}

// SetValidateOutputIdObject gets a reference to the given Output and assigns it to the ValidateOutputIdObject field.
func (o *Execution) SetValidateOutputIdObject(v Output) {
	o.ValidateOutputIdObject = &v
}

func (o Execution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Execution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildOutputId) {
		toSerialize["build_output_id"] = o.BuildOutputId
	}
	if !IsNil(o.BuildOutputIdObject) {
		toSerialize["build_output_id_object"] = o.BuildOutputIdObject
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.DeployOutputId) {
		toSerialize["deploy_output_id"] = o.DeployOutputId
	}
	if !IsNil(o.DeployOutputIdObject) {
		toSerialize["deploy_output_id_object"] = o.DeployOutputIdObject
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JobId) {
		toSerialize["job_id"] = o.JobId
	}
	if !IsNil(o.JobIdObject) {
		toSerialize["job_id_object"] = o.JobIdObject
	}
	if !IsNil(o.PublishOutputId) {
		toSerialize["publish_output_id"] = o.PublishOutputId
	}
	if !IsNil(o.PublishOutputIdObject) {
		toSerialize["publish_output_id_object"] = o.PublishOutputIdObject
	}
	if o.ReferencedByTriggerHasExecutionExecutionIdObjects != nil {
		toSerialize["referenced_by_trigger_has_execution_execution_id_objects"] = o.ReferencedByTriggerHasExecutionExecutionIdObjects
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TestOutputId) {
		toSerialize["test_output_id"] = o.TestOutputId
	}
	if !IsNil(o.TestOutputIdObject) {
		toSerialize["test_output_id_object"] = o.TestOutputIdObject
	}
	if !IsNil(o.TriggerId) {
		toSerialize["trigger_id"] = o.TriggerId
	}
	if !IsNil(o.TriggerIdObject) {
		toSerialize["trigger_id_object"] = o.TriggerIdObject
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ValidateOutputId) {
		toSerialize["validate_output_id"] = o.ValidateOutputId
	}
	if !IsNil(o.ValidateOutputIdObject) {
		toSerialize["validate_output_id_object"] = o.ValidateOutputIdObject
	}
	return toSerialize, nil
}

type NullableExecution struct {
	value *Execution
	isSet bool
}

func (v NullableExecution) Get() *Execution {
	return v.value
}

func (v *NullableExecution) Set(val *Execution) {
	v.value = val
	v.isSet = true
}

func (v NullableExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecution(val *Execution) *NullableExecution {
	return &NullableExecution{value: val, isSet: true}
}

func (v NullableExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


