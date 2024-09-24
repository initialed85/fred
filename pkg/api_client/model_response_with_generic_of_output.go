/*
Djangolang

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_client

import (
	"encoding/json"
)

// checks if the ResponseWithGenericOfOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseWithGenericOfOutput{}

// ResponseWithGenericOfOutput struct for ResponseWithGenericOfOutput
type ResponseWithGenericOfOutput struct {
	Count      *int64   `json:"count,omitempty"`
	Error      []string `json:"error,omitempty"`
	Limit      *int64   `json:"limit,omitempty"`
	Objects    []Output `json:"objects,omitempty"`
	Offset     *int64   `json:"offset,omitempty"`
	Status     *int64   `json:"status,omitempty"`
	Success    *bool    `json:"success,omitempty"`
	TotalCount *int64   `json:"total_count,omitempty"`
}

// NewResponseWithGenericOfOutput instantiates a new ResponseWithGenericOfOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseWithGenericOfOutput() *ResponseWithGenericOfOutput {
	this := ResponseWithGenericOfOutput{}
	return &this
}

// NewResponseWithGenericOfOutputWithDefaults instantiates a new ResponseWithGenericOfOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithGenericOfOutputWithDefaults() *ResponseWithGenericOfOutput {
	this := ResponseWithGenericOfOutput{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ResponseWithGenericOfOutput) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfOutput) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ResponseWithGenericOfOutput) SetCount(v int64) {
	o.Count = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseWithGenericOfOutput) GetError() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseWithGenericOfOutput) GetErrorOk() ([]string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []string and assigns it to the Error field.
func (o *ResponseWithGenericOfOutput) SetError(v []string) {
	o.Error = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseWithGenericOfOutput) GetLimit() int64 {
	if o == nil || IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfOutput) GetLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ResponseWithGenericOfOutput) SetLimit(v int64) {
	o.Limit = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseWithGenericOfOutput) GetObjects() []Output {
	if o == nil {
		var ret []Output
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseWithGenericOfOutput) GetObjectsOk() ([]Output, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []Output and assigns it to the Objects field.
func (o *ResponseWithGenericOfOutput) SetObjects(v []Output) {
	o.Objects = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ResponseWithGenericOfOutput) GetOffset() int64 {
	if o == nil || IsNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfOutput) GetOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ResponseWithGenericOfOutput) SetOffset(v int64) {
	o.Offset = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseWithGenericOfOutput) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfOutput) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ResponseWithGenericOfOutput) SetStatus(v int64) {
	o.Status = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ResponseWithGenericOfOutput) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfOutput) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ResponseWithGenericOfOutput) SetSuccess(v bool) {
	o.Success = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ResponseWithGenericOfOutput) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfOutput) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ResponseWithGenericOfOutput) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ResponseWithGenericOfOutput) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o ResponseWithGenericOfOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseWithGenericOfOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableResponseWithGenericOfOutput struct {
	value *ResponseWithGenericOfOutput
	isSet bool
}

func (v NullableResponseWithGenericOfOutput) Get() *ResponseWithGenericOfOutput {
	return v.value
}

func (v *NullableResponseWithGenericOfOutput) Set(val *ResponseWithGenericOfOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseWithGenericOfOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseWithGenericOfOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseWithGenericOfOutput(val *ResponseWithGenericOfOutput) *NullableResponseWithGenericOfOutput {
	return &NullableResponseWithGenericOfOutput{value: val, isSet: true}
}

func (v NullableResponseWithGenericOfOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseWithGenericOfOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
