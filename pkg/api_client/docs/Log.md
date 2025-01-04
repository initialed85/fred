# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buffer** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OutputId** | Pointer to **string** |  | [optional] 
**OutputIdObject** | Pointer to [**Output**](Output.md) |  | [optional] 
**ReferencedByOutputLogIdObjects** | Pointer to [**[]Output**](Output.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLog

`func NewLog() *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuffer

`func (o *Log) GetBuffer() string`

GetBuffer returns the Buffer field if non-nil, zero value otherwise.

### GetBufferOk

`func (o *Log) GetBufferOk() (*string, bool)`

GetBufferOk returns a tuple with the Buffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuffer

`func (o *Log) SetBuffer(v string)`

SetBuffer sets Buffer field to given value.

### HasBuffer

`func (o *Log) HasBuffer() bool`

HasBuffer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Log) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Log) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Log) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Log) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Log) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Log) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Log) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Log) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *Log) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Log) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Log) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Log) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOutputId

`func (o *Log) GetOutputId() string`

GetOutputId returns the OutputId field if non-nil, zero value otherwise.

### GetOutputIdOk

`func (o *Log) GetOutputIdOk() (*string, bool)`

GetOutputIdOk returns a tuple with the OutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputId

`func (o *Log) SetOutputId(v string)`

SetOutputId sets OutputId field to given value.

### HasOutputId

`func (o *Log) HasOutputId() bool`

HasOutputId returns a boolean if a field has been set.

### GetOutputIdObject

`func (o *Log) GetOutputIdObject() Output`

GetOutputIdObject returns the OutputIdObject field if non-nil, zero value otherwise.

### GetOutputIdObjectOk

`func (o *Log) GetOutputIdObjectOk() (*Output, bool)`

GetOutputIdObjectOk returns a tuple with the OutputIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIdObject

`func (o *Log) SetOutputIdObject(v Output)`

SetOutputIdObject sets OutputIdObject field to given value.

### HasOutputIdObject

`func (o *Log) HasOutputIdObject() bool`

HasOutputIdObject returns a boolean if a field has been set.

### GetReferencedByOutputLogIdObjects

`func (o *Log) GetReferencedByOutputLogIdObjects() []Output`

GetReferencedByOutputLogIdObjects returns the ReferencedByOutputLogIdObjects field if non-nil, zero value otherwise.

### GetReferencedByOutputLogIdObjectsOk

`func (o *Log) GetReferencedByOutputLogIdObjectsOk() (*[]Output, bool)`

GetReferencedByOutputLogIdObjectsOk returns a tuple with the ReferencedByOutputLogIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByOutputLogIdObjects

`func (o *Log) SetReferencedByOutputLogIdObjects(v []Output)`

SetReferencedByOutputLogIdObjects sets ReferencedByOutputLogIdObjects field to given value.

### HasReferencedByOutputLogIdObjects

`func (o *Log) HasReferencedByOutputLogIdObjects() bool`

HasReferencedByOutputLogIdObjects returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Log) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Log) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Log) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Log) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


