# ChangeJobCoordinatorClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutSeconds** | Pointer to **float64** |  | [optional] 
**Until** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChangeJobCoordinatorClaimRequest

`func NewChangeJobCoordinatorClaimRequest() *ChangeJobCoordinatorClaimRequest`

NewChangeJobCoordinatorClaimRequest instantiates a new ChangeJobCoordinatorClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeJobCoordinatorClaimRequestWithDefaults

`func NewChangeJobCoordinatorClaimRequestWithDefaults() *ChangeJobCoordinatorClaimRequest`

NewChangeJobCoordinatorClaimRequestWithDefaults instantiates a new ChangeJobCoordinatorClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutSeconds

`func (o *ChangeJobCoordinatorClaimRequest) GetTimeoutSeconds() float64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ChangeJobCoordinatorClaimRequest) GetTimeoutSecondsOk() (*float64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ChangeJobCoordinatorClaimRequest) SetTimeoutSeconds(v float64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *ChangeJobCoordinatorClaimRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetUntil

`func (o *ChangeJobCoordinatorClaimRequest) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *ChangeJobCoordinatorClaimRequest) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *ChangeJobCoordinatorClaimRequest) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *ChangeJobCoordinatorClaimRequest) HasUntil() bool`

HasUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


