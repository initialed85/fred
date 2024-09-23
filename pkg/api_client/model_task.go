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

// checks if the Task type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Task{}

// Task struct for Task
type Task struct {
	CreatedAt                            *time.Time  `json:"created_at,omitempty"`
	DeletedAt                            *time.Time  `json:"deleted_at,omitempty"`
	Id                                   *string     `json:"id,omitempty"`
	ReferencedByJobBuildTaskIdObjects    []Job       `json:"referenced_by_job_build_task_id_objects,omitempty"`
	ReferencedByJobDeployTaskIdObjects   []Job       `json:"referenced_by_job_deploy_task_id_objects,omitempty"`
	ReferencedByJobPublishTaskIdObjects  []Job       `json:"referenced_by_job_publish_task_id_objects,omitempty"`
	ReferencedByJobTestTaskIdObjects     []Job       `json:"referenced_by_job_test_task_id_objects,omitempty"`
	ReferencedByJobValidateTaskIdObjects []Job       `json:"referenced_by_job_validate_task_id_objects,omitempty"`
	ReferencedByOutputTaskIdObjects      []Execution `json:"referenced_by_output_task_id_objects,omitempty"`
	UpdatedAt                            *time.Time  `json:"updated_at,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Task) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Task) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Task) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Task) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Task) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Task) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Task) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Task) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Task) SetId(v string) {
	o.Id = &v
}

// GetReferencedByJobBuildTaskIdObjects returns the ReferencedByJobBuildTaskIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetReferencedByJobBuildTaskIdObjects() []Job {
	if o == nil {
		var ret []Job
		return ret
	}
	return o.ReferencedByJobBuildTaskIdObjects
}

// GetReferencedByJobBuildTaskIdObjectsOk returns a tuple with the ReferencedByJobBuildTaskIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetReferencedByJobBuildTaskIdObjectsOk() ([]Job, bool) {
	if o == nil || IsNil(o.ReferencedByJobBuildTaskIdObjects) {
		return nil, false
	}
	return o.ReferencedByJobBuildTaskIdObjects, true
}

// HasReferencedByJobBuildTaskIdObjects returns a boolean if a field has been set.
func (o *Task) HasReferencedByJobBuildTaskIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByJobBuildTaskIdObjects) {
		return true
	}

	return false
}

// SetReferencedByJobBuildTaskIdObjects gets a reference to the given []Job and assigns it to the ReferencedByJobBuildTaskIdObjects field.
func (o *Task) SetReferencedByJobBuildTaskIdObjects(v []Job) {
	o.ReferencedByJobBuildTaskIdObjects = v
}

// GetReferencedByJobDeployTaskIdObjects returns the ReferencedByJobDeployTaskIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetReferencedByJobDeployTaskIdObjects() []Job {
	if o == nil {
		var ret []Job
		return ret
	}
	return o.ReferencedByJobDeployTaskIdObjects
}

// GetReferencedByJobDeployTaskIdObjectsOk returns a tuple with the ReferencedByJobDeployTaskIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetReferencedByJobDeployTaskIdObjectsOk() ([]Job, bool) {
	if o == nil || IsNil(o.ReferencedByJobDeployTaskIdObjects) {
		return nil, false
	}
	return o.ReferencedByJobDeployTaskIdObjects, true
}

// HasReferencedByJobDeployTaskIdObjects returns a boolean if a field has been set.
func (o *Task) HasReferencedByJobDeployTaskIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByJobDeployTaskIdObjects) {
		return true
	}

	return false
}

// SetReferencedByJobDeployTaskIdObjects gets a reference to the given []Job and assigns it to the ReferencedByJobDeployTaskIdObjects field.
func (o *Task) SetReferencedByJobDeployTaskIdObjects(v []Job) {
	o.ReferencedByJobDeployTaskIdObjects = v
}

// GetReferencedByJobPublishTaskIdObjects returns the ReferencedByJobPublishTaskIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetReferencedByJobPublishTaskIdObjects() []Job {
	if o == nil {
		var ret []Job
		return ret
	}
	return o.ReferencedByJobPublishTaskIdObjects
}

// GetReferencedByJobPublishTaskIdObjectsOk returns a tuple with the ReferencedByJobPublishTaskIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetReferencedByJobPublishTaskIdObjectsOk() ([]Job, bool) {
	if o == nil || IsNil(o.ReferencedByJobPublishTaskIdObjects) {
		return nil, false
	}
	return o.ReferencedByJobPublishTaskIdObjects, true
}

// HasReferencedByJobPublishTaskIdObjects returns a boolean if a field has been set.
func (o *Task) HasReferencedByJobPublishTaskIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByJobPublishTaskIdObjects) {
		return true
	}

	return false
}

// SetReferencedByJobPublishTaskIdObjects gets a reference to the given []Job and assigns it to the ReferencedByJobPublishTaskIdObjects field.
func (o *Task) SetReferencedByJobPublishTaskIdObjects(v []Job) {
	o.ReferencedByJobPublishTaskIdObjects = v
}

// GetReferencedByJobTestTaskIdObjects returns the ReferencedByJobTestTaskIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetReferencedByJobTestTaskIdObjects() []Job {
	if o == nil {
		var ret []Job
		return ret
	}
	return o.ReferencedByJobTestTaskIdObjects
}

// GetReferencedByJobTestTaskIdObjectsOk returns a tuple with the ReferencedByJobTestTaskIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetReferencedByJobTestTaskIdObjectsOk() ([]Job, bool) {
	if o == nil || IsNil(o.ReferencedByJobTestTaskIdObjects) {
		return nil, false
	}
	return o.ReferencedByJobTestTaskIdObjects, true
}

// HasReferencedByJobTestTaskIdObjects returns a boolean if a field has been set.
func (o *Task) HasReferencedByJobTestTaskIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByJobTestTaskIdObjects) {
		return true
	}

	return false
}

// SetReferencedByJobTestTaskIdObjects gets a reference to the given []Job and assigns it to the ReferencedByJobTestTaskIdObjects field.
func (o *Task) SetReferencedByJobTestTaskIdObjects(v []Job) {
	o.ReferencedByJobTestTaskIdObjects = v
}

// GetReferencedByJobValidateTaskIdObjects returns the ReferencedByJobValidateTaskIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetReferencedByJobValidateTaskIdObjects() []Job {
	if o == nil {
		var ret []Job
		return ret
	}
	return o.ReferencedByJobValidateTaskIdObjects
}

// GetReferencedByJobValidateTaskIdObjectsOk returns a tuple with the ReferencedByJobValidateTaskIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetReferencedByJobValidateTaskIdObjectsOk() ([]Job, bool) {
	if o == nil || IsNil(o.ReferencedByJobValidateTaskIdObjects) {
		return nil, false
	}
	return o.ReferencedByJobValidateTaskIdObjects, true
}

// HasReferencedByJobValidateTaskIdObjects returns a boolean if a field has been set.
func (o *Task) HasReferencedByJobValidateTaskIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByJobValidateTaskIdObjects) {
		return true
	}

	return false
}

// SetReferencedByJobValidateTaskIdObjects gets a reference to the given []Job and assigns it to the ReferencedByJobValidateTaskIdObjects field.
func (o *Task) SetReferencedByJobValidateTaskIdObjects(v []Job) {
	o.ReferencedByJobValidateTaskIdObjects = v
}

// GetReferencedByOutputTaskIdObjects returns the ReferencedByOutputTaskIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Task) GetReferencedByOutputTaskIdObjects() []Execution {
	if o == nil {
		var ret []Execution
		return ret
	}
	return o.ReferencedByOutputTaskIdObjects
}

// GetReferencedByOutputTaskIdObjectsOk returns a tuple with the ReferencedByOutputTaskIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Task) GetReferencedByOutputTaskIdObjectsOk() ([]Execution, bool) {
	if o == nil || IsNil(o.ReferencedByOutputTaskIdObjects) {
		return nil, false
	}
	return o.ReferencedByOutputTaskIdObjects, true
}

// HasReferencedByOutputTaskIdObjects returns a boolean if a field has been set.
func (o *Task) HasReferencedByOutputTaskIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByOutputTaskIdObjects) {
		return true
	}

	return false
}

// SetReferencedByOutputTaskIdObjects gets a reference to the given []Execution and assigns it to the ReferencedByOutputTaskIdObjects field.
func (o *Task) SetReferencedByOutputTaskIdObjects(v []Execution) {
	o.ReferencedByOutputTaskIdObjects = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Task) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Task) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Task) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Task) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ReferencedByJobBuildTaskIdObjects != nil {
		toSerialize["referenced_by_job_build_task_id_objects"] = o.ReferencedByJobBuildTaskIdObjects
	}
	if o.ReferencedByJobDeployTaskIdObjects != nil {
		toSerialize["referenced_by_job_deploy_task_id_objects"] = o.ReferencedByJobDeployTaskIdObjects
	}
	if o.ReferencedByJobPublishTaskIdObjects != nil {
		toSerialize["referenced_by_job_publish_task_id_objects"] = o.ReferencedByJobPublishTaskIdObjects
	}
	if o.ReferencedByJobTestTaskIdObjects != nil {
		toSerialize["referenced_by_job_test_task_id_objects"] = o.ReferencedByJobTestTaskIdObjects
	}
	if o.ReferencedByJobValidateTaskIdObjects != nil {
		toSerialize["referenced_by_job_validate_task_id_objects"] = o.ReferencedByJobValidateTaskIdObjects
	}
	if o.ReferencedByOutputTaskIdObjects != nil {
		toSerialize["referenced_by_output_task_id_objects"] = o.ReferencedByOutputTaskIdObjects
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
