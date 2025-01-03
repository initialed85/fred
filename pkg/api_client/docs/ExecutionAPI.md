# \ExecutionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExecution**](ExecutionAPI.md#DeleteExecution) | **Delete** /api/executions/{primaryKey} | 
[**GetExecution**](ExecutionAPI.md#GetExecution) | **Get** /api/executions/{primaryKey} | 
[**GetExecutions**](ExecutionAPI.md#GetExecutions) | **Get** /api/executions | 
[**PatchExecution**](ExecutionAPI.md#PatchExecution) | **Patch** /api/executions/{primaryKey} | 
[**PostExecutions**](ExecutionAPI.md#PostExecutions) | **Post** /api/executions | 



## DeleteExecution

> DeleteExecution(ctx, primaryKey).Depth(depth).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	primaryKey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Path parameter primaryKey
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.DeleteExecution(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.DeleteExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecution

> ResponseWithGenericOfExecution GetExecution(ctx, primaryKey).Depth(depth).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	primaryKey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Path parameter primaryKey
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetExecution(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecution`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfExecution**](ResponseWithGenericOfExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutions

> ResponseWithGenericOfExecution GetExecutions(ctx).Limit(limit).Offset(offset).Depth(depth).TaskLoad(taskLoad).M2mRuleTriggerJobLoad(m2mRuleTriggerJobLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).StatusEq(statusEq).StatusNe(statusNe).StatusGt(statusGt).StatusGte(statusGte).StatusLt(statusLt).StatusLte(statusLte).StatusIn(statusIn).StatusNotin(statusNotin).StatusLike(statusLike).StatusNotlike(statusNotlike).StatusIlike(statusIlike).StatusNotilike(statusNotilike).StatusDesc(statusDesc).StatusAsc(statusAsc).StartedAtEq(startedAtEq).StartedAtNe(startedAtNe).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtIn(startedAtIn).StartedAtNotin(startedAtNotin).StartedAtLike(startedAtLike).StartedAtNotlike(startedAtNotlike).StartedAtIlike(startedAtIlike).StartedAtNotilike(startedAtNotilike).StartedAtDesc(startedAtDesc).StartedAtAsc(startedAtAsc).EndedAtEq(endedAtEq).EndedAtNe(endedAtNe).EndedAtGt(endedAtGt).EndedAtGte(endedAtGte).EndedAtLt(endedAtLt).EndedAtLte(endedAtLte).EndedAtIn(endedAtIn).EndedAtNotin(endedAtNotin).EndedAtLike(endedAtLike).EndedAtNotlike(endedAtNotlike).EndedAtIlike(endedAtIlike).EndedAtNotilike(endedAtNotilike).EndedAtDesc(endedAtDesc).EndedAtAsc(endedAtAsc).TaskIdEq(taskIdEq).TaskIdNe(taskIdNe).TaskIdGt(taskIdGt).TaskIdGte(taskIdGte).TaskIdLt(taskIdLt).TaskIdLte(taskIdLte).TaskIdIn(taskIdIn).TaskIdNotin(taskIdNotin).TaskIdLike(taskIdLike).TaskIdNotlike(taskIdNotlike).TaskIdIlike(taskIdIlike).TaskIdNotilike(taskIdNotilike).TaskIdDesc(taskIdDesc).TaskIdAsc(taskIdAsc).TaskIdObjectDesc(taskIdObjectDesc).TaskIdObjectAsc(taskIdObjectAsc).M2mRuleTriggerJobIdEq(m2mRuleTriggerJobIdEq).M2mRuleTriggerJobIdNe(m2mRuleTriggerJobIdNe).M2mRuleTriggerJobIdGt(m2mRuleTriggerJobIdGt).M2mRuleTriggerJobIdGte(m2mRuleTriggerJobIdGte).M2mRuleTriggerJobIdLt(m2mRuleTriggerJobIdLt).M2mRuleTriggerJobIdLte(m2mRuleTriggerJobIdLte).M2mRuleTriggerJobIdIn(m2mRuleTriggerJobIdIn).M2mRuleTriggerJobIdNotin(m2mRuleTriggerJobIdNotin).M2mRuleTriggerJobIdLike(m2mRuleTriggerJobIdLike).M2mRuleTriggerJobIdNotlike(m2mRuleTriggerJobIdNotlike).M2mRuleTriggerJobIdIlike(m2mRuleTriggerJobIdIlike).M2mRuleTriggerJobIdNotilike(m2mRuleTriggerJobIdNotilike).M2mRuleTriggerJobIdDesc(m2mRuleTriggerJobIdDesc).M2mRuleTriggerJobIdAsc(m2mRuleTriggerJobIdAsc).M2mRuleTriggerJobIdObjectDesc(m2mRuleTriggerJobIdObjectDesc).M2mRuleTriggerJobIdObjectAsc(m2mRuleTriggerJobIdObjectAsc).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int32(56) // int32 | SQL LIMIT operator (optional)
	offset := int32(56) // int32 | SQL OFFSET operator (optional)
	depth := int32(56) // int32 | Max recursion depth for loading foreign objects; default = 1  (0 = recurse until graph cycle detected, 1 = this object only, 2 = this object + neighbours, 3 = this object + neighbours + their neighbours... etc) (optional)
	taskLoad := "taskLoad_example" // string | load the given directly related Djangolang object, value is ignored (presence of key is sufficient) (optional)
	m2mRuleTriggerJobLoad := "m2mRuleTriggerJobLoad_example" // string | load the given directly related Djangolang object, value is ignored (presence of key is sufficient) (optional)
	idEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	idNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	idGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	idGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	idLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	idLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	idIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	idNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	idLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idDesc := "idDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	idAsc := "idAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	createdAtEq := time.Now() // time.Time | SQL = comparison (optional)
	createdAtNe := time.Now() // time.Time | SQL != comparison (optional)
	createdAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	createdAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	createdAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	createdAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	createdAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	createdAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	createdAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtDesc := "createdAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	createdAtAsc := "createdAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	updatedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	updatedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	updatedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	updatedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	updatedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	updatedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	updatedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	updatedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	updatedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtDesc := "updatedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	updatedAtAsc := "updatedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	deletedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	deletedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	deletedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	deletedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	deletedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	deletedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	deletedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	deletedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	deletedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtDesc := "deletedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	deletedAtAsc := "deletedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	statusEq := "statusEq_example" // string | SQL = comparison (optional)
	statusNe := "statusNe_example" // string | SQL != comparison (optional)
	statusGt := "statusGt_example" // string | SQL > comparison, may not work with all column types (optional)
	statusGte := "statusGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	statusLt := "statusLt_example" // string | SQL < comparison, may not work with all column types (optional)
	statusLte := "statusLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	statusIn := "statusIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	statusNotin := "statusNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	statusLike := "statusLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusNotlike := "statusNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusIlike := "statusIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusNotilike := "statusNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusDesc := "statusDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	statusAsc := "statusAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	startedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	startedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	startedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	startedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	startedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	startedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	startedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	startedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	startedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtDesc := "startedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	startedAtAsc := "startedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	endedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	endedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	endedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	endedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	endedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	endedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	endedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	endedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	endedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtDesc := "endedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	endedAtAsc := "endedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	taskIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	taskIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	taskIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	taskIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	taskIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	taskIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	taskIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	taskIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdDesc := "taskIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdAsc := "taskIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdObjectDesc := "taskIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdObjectAsc := "taskIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	m2mRuleTriggerJobIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	m2mRuleTriggerJobIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	m2mRuleTriggerJobIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	m2mRuleTriggerJobIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	m2mRuleTriggerJobIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	m2mRuleTriggerJobIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	m2mRuleTriggerJobIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	m2mRuleTriggerJobIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	m2mRuleTriggerJobIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	m2mRuleTriggerJobIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	m2mRuleTriggerJobIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	m2mRuleTriggerJobIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	m2mRuleTriggerJobIdDesc := "m2mRuleTriggerJobIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	m2mRuleTriggerJobIdAsc := "m2mRuleTriggerJobIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	m2mRuleTriggerJobIdObjectDesc := "m2mRuleTriggerJobIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	m2mRuleTriggerJobIdObjectAsc := "m2mRuleTriggerJobIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetExecutions(context.Background()).Limit(limit).Offset(offset).Depth(depth).TaskLoad(taskLoad).M2mRuleTriggerJobLoad(m2mRuleTriggerJobLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).StatusEq(statusEq).StatusNe(statusNe).StatusGt(statusGt).StatusGte(statusGte).StatusLt(statusLt).StatusLte(statusLte).StatusIn(statusIn).StatusNotin(statusNotin).StatusLike(statusLike).StatusNotlike(statusNotlike).StatusIlike(statusIlike).StatusNotilike(statusNotilike).StatusDesc(statusDesc).StatusAsc(statusAsc).StartedAtEq(startedAtEq).StartedAtNe(startedAtNe).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtIn(startedAtIn).StartedAtNotin(startedAtNotin).StartedAtLike(startedAtLike).StartedAtNotlike(startedAtNotlike).StartedAtIlike(startedAtIlike).StartedAtNotilike(startedAtNotilike).StartedAtDesc(startedAtDesc).StartedAtAsc(startedAtAsc).EndedAtEq(endedAtEq).EndedAtNe(endedAtNe).EndedAtGt(endedAtGt).EndedAtGte(endedAtGte).EndedAtLt(endedAtLt).EndedAtLte(endedAtLte).EndedAtIn(endedAtIn).EndedAtNotin(endedAtNotin).EndedAtLike(endedAtLike).EndedAtNotlike(endedAtNotlike).EndedAtIlike(endedAtIlike).EndedAtNotilike(endedAtNotilike).EndedAtDesc(endedAtDesc).EndedAtAsc(endedAtAsc).TaskIdEq(taskIdEq).TaskIdNe(taskIdNe).TaskIdGt(taskIdGt).TaskIdGte(taskIdGte).TaskIdLt(taskIdLt).TaskIdLte(taskIdLte).TaskIdIn(taskIdIn).TaskIdNotin(taskIdNotin).TaskIdLike(taskIdLike).TaskIdNotlike(taskIdNotlike).TaskIdIlike(taskIdIlike).TaskIdNotilike(taskIdNotilike).TaskIdDesc(taskIdDesc).TaskIdAsc(taskIdAsc).TaskIdObjectDesc(taskIdObjectDesc).TaskIdObjectAsc(taskIdObjectAsc).M2mRuleTriggerJobIdEq(m2mRuleTriggerJobIdEq).M2mRuleTriggerJobIdNe(m2mRuleTriggerJobIdNe).M2mRuleTriggerJobIdGt(m2mRuleTriggerJobIdGt).M2mRuleTriggerJobIdGte(m2mRuleTriggerJobIdGte).M2mRuleTriggerJobIdLt(m2mRuleTriggerJobIdLt).M2mRuleTriggerJobIdLte(m2mRuleTriggerJobIdLte).M2mRuleTriggerJobIdIn(m2mRuleTriggerJobIdIn).M2mRuleTriggerJobIdNotin(m2mRuleTriggerJobIdNotin).M2mRuleTriggerJobIdLike(m2mRuleTriggerJobIdLike).M2mRuleTriggerJobIdNotlike(m2mRuleTriggerJobIdNotlike).M2mRuleTriggerJobIdIlike(m2mRuleTriggerJobIdIlike).M2mRuleTriggerJobIdNotilike(m2mRuleTriggerJobIdNotilike).M2mRuleTriggerJobIdDesc(m2mRuleTriggerJobIdDesc).M2mRuleTriggerJobIdAsc(m2mRuleTriggerJobIdAsc).M2mRuleTriggerJobIdObjectDesc(m2mRuleTriggerJobIdObjectDesc).M2mRuleTriggerJobIdObjectAsc(m2mRuleTriggerJobIdObjectAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecutions`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | SQL LIMIT operator | 
 **offset** | **int32** | SQL OFFSET operator | 
 **depth** | **int32** | Max recursion depth for loading foreign objects; default &#x3D; 1  (0 &#x3D; recurse until graph cycle detected, 1 &#x3D; this object only, 2 &#x3D; this object + neighbours, 3 &#x3D; this object + neighbours + their neighbours... etc) | 
 **taskLoad** | **string** | load the given directly related Djangolang object, value is ignored (presence of key is sufficient) | 
 **m2mRuleTriggerJobLoad** | **string** | load the given directly related Djangolang object, value is ignored (presence of key is sufficient) | 
 **idEq** | **string** | SQL &#x3D; comparison | 
 **idNe** | **string** | SQL !&#x3D; comparison | 
 **idGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **idGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **idLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **idLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **idIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **idNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **idLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **idAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **createdAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **createdAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **createdAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **createdAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **createdAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **createdAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **createdAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **createdAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **createdAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **createdAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **updatedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **updatedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **updatedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **updatedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **updatedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **updatedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **updatedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **updatedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **updatedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **updatedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **deletedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **deletedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **deletedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **deletedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **deletedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **deletedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **deletedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **deletedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **deletedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **deletedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **statusEq** | **string** | SQL &#x3D; comparison | 
 **statusNe** | **string** | SQL !&#x3D; comparison | 
 **statusGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **statusGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **statusLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **statusLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **statusIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **statusNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **statusLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **statusAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **startedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **startedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **startedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **startedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **startedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **startedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **startedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **startedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **startedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **startedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **endedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **endedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **endedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **endedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **endedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **endedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **endedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **endedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **endedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **endedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **taskIdEq** | **string** | SQL &#x3D; comparison | 
 **taskIdNe** | **string** | SQL !&#x3D; comparison | 
 **taskIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **taskIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **taskIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **taskIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **taskIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **taskIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **taskIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **taskIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **taskIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **taskIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **m2mRuleTriggerJobIdEq** | **string** | SQL &#x3D; comparison | 
 **m2mRuleTriggerJobIdNe** | **string** | SQL !&#x3D; comparison | 
 **m2mRuleTriggerJobIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **m2mRuleTriggerJobIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **m2mRuleTriggerJobIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **m2mRuleTriggerJobIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **m2mRuleTriggerJobIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **m2mRuleTriggerJobIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **m2mRuleTriggerJobIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **m2mRuleTriggerJobIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **m2mRuleTriggerJobIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **m2mRuleTriggerJobIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **m2mRuleTriggerJobIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **m2mRuleTriggerJobIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **m2mRuleTriggerJobIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **m2mRuleTriggerJobIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfExecution**](ResponseWithGenericOfExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExecution

> ResponseWithGenericOfExecution PatchExecution(ctx, primaryKey).Execution(execution).Depth(depth).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	primaryKey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Path parameter primaryKey
	execution := *openapiclient.NewExecution() // Execution | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.PatchExecution(context.Background(), primaryKey).Execution(execution).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.PatchExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExecution`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.PatchExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **execution** | [**Execution**](Execution.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfExecution**](ResponseWithGenericOfExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExecutions

> ResponseWithGenericOfExecution PostExecutions(ctx).Execution(execution).Depth(depth).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	execution := []openapiclient.Execution{*openapiclient.NewExecution()} // []Execution | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.PostExecutions(context.Background()).Execution(execution).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.PostExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExecutions`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.PostExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **execution** | [**[]Execution**](Execution.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfExecution**](ResponseWithGenericOfExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

