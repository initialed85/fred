# \TriggerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrigger**](TriggerAPI.md#DeleteTrigger) | **Delete** /api/triggers/{primaryKey} | 
[**GetTrigger**](TriggerAPI.md#GetTrigger) | **Get** /api/triggers/{primaryKey} | 
[**GetTriggers**](TriggerAPI.md#GetTriggers) | **Get** /api/triggers | 
[**PatchTrigger**](TriggerAPI.md#PatchTrigger) | **Patch** /api/triggers/{primaryKey} | 
[**PostTriggers**](TriggerAPI.md#PostTriggers) | **Post** /api/triggers | 



## DeleteTrigger

> DeleteTrigger(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.TriggerAPI.DeleteTrigger(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.DeleteTrigger``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTriggerRequest struct via the builder pattern


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


## GetTrigger

> ResponseWithGenericOfTrigger GetTrigger(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.TriggerAPI.GetTrigger(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrigger`: ResponseWithGenericOfTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfTrigger**](ResponseWithGenericOfTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriggers

> ResponseWithGenericOfTrigger GetTriggers(ctx).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).JobExecutorClaimedUntilEq(jobExecutorClaimedUntilEq).JobExecutorClaimedUntilNe(jobExecutorClaimedUntilNe).JobExecutorClaimedUntilGt(jobExecutorClaimedUntilGt).JobExecutorClaimedUntilGte(jobExecutorClaimedUntilGte).JobExecutorClaimedUntilLt(jobExecutorClaimedUntilLt).JobExecutorClaimedUntilLte(jobExecutorClaimedUntilLte).JobExecutorClaimedUntilIn(jobExecutorClaimedUntilIn).JobExecutorClaimedUntilNotin(jobExecutorClaimedUntilNotin).JobExecutorClaimedUntilLike(jobExecutorClaimedUntilLike).JobExecutorClaimedUntilNotlike(jobExecutorClaimedUntilNotlike).JobExecutorClaimedUntilIlike(jobExecutorClaimedUntilIlike).JobExecutorClaimedUntilNotilike(jobExecutorClaimedUntilNotilike).JobExecutorClaimedUntilDesc(jobExecutorClaimedUntilDesc).JobExecutorClaimedUntilAsc(jobExecutorClaimedUntilAsc).JobExecutionStartedAtEq(jobExecutionStartedAtEq).JobExecutionStartedAtNe(jobExecutionStartedAtNe).JobExecutionStartedAtGt(jobExecutionStartedAtGt).JobExecutionStartedAtGte(jobExecutionStartedAtGte).JobExecutionStartedAtLt(jobExecutionStartedAtLt).JobExecutionStartedAtLte(jobExecutionStartedAtLte).JobExecutionStartedAtIn(jobExecutionStartedAtIn).JobExecutionStartedAtNotin(jobExecutionStartedAtNotin).JobExecutionStartedAtLike(jobExecutionStartedAtLike).JobExecutionStartedAtNotlike(jobExecutionStartedAtNotlike).JobExecutionStartedAtIlike(jobExecutionStartedAtIlike).JobExecutionStartedAtNotilike(jobExecutionStartedAtNotilike).JobExecutionStartedAtDesc(jobExecutionStartedAtDesc).JobExecutionStartedAtAsc(jobExecutionStartedAtAsc).RuleIdEq(ruleIdEq).RuleIdNe(ruleIdNe).RuleIdGt(ruleIdGt).RuleIdGte(ruleIdGte).RuleIdLt(ruleIdLt).RuleIdLte(ruleIdLte).RuleIdIn(ruleIdIn).RuleIdNotin(ruleIdNotin).RuleIdLike(ruleIdLike).RuleIdNotlike(ruleIdNotlike).RuleIdIlike(ruleIdIlike).RuleIdNotilike(ruleIdNotilike).RuleIdDesc(ruleIdDesc).RuleIdAsc(ruleIdAsc).RuleIdObjectDesc(ruleIdObjectDesc).RuleIdObjectAsc(ruleIdObjectAsc).ChangeIdEq(changeIdEq).ChangeIdNe(changeIdNe).ChangeIdGt(changeIdGt).ChangeIdGte(changeIdGte).ChangeIdLt(changeIdLt).ChangeIdLte(changeIdLte).ChangeIdIn(changeIdIn).ChangeIdNotin(changeIdNotin).ChangeIdLike(changeIdLike).ChangeIdNotlike(changeIdNotlike).ChangeIdIlike(changeIdIlike).ChangeIdNotilike(changeIdNotilike).ChangeIdDesc(changeIdDesc).ChangeIdAsc(changeIdAsc).ChangeIdObjectDesc(changeIdObjectDesc).ChangeIdObjectAsc(changeIdObjectAsc).ReferencedByTriggerHasExecutionTriggerIdObjectsDesc(referencedByTriggerHasExecutionTriggerIdObjectsDesc).ReferencedByTriggerHasExecutionTriggerIdObjectsAsc(referencedByTriggerHasExecutionTriggerIdObjectsAsc).ReferencedByExecutionTriggerIdObjectsDesc(referencedByExecutionTriggerIdObjectsDesc).ReferencedByExecutionTriggerIdObjectsAsc(referencedByExecutionTriggerIdObjectsAsc).Execute()



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
	jobExecutorClaimedUntilEq := time.Now() // time.Time | SQL = comparison (optional)
	jobExecutorClaimedUntilNe := time.Now() // time.Time | SQL != comparison (optional)
	jobExecutorClaimedUntilGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	jobExecutorClaimedUntilGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	jobExecutorClaimedUntilLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	jobExecutorClaimedUntilLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	jobExecutorClaimedUntilIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	jobExecutorClaimedUntilNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	jobExecutorClaimedUntilLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilDesc := "jobExecutorClaimedUntilDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	jobExecutorClaimedUntilAsc := "jobExecutorClaimedUntilAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	jobExecutionStartedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	jobExecutionStartedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	jobExecutionStartedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	jobExecutionStartedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	jobExecutionStartedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	jobExecutionStartedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	jobExecutionStartedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	jobExecutionStartedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	jobExecutionStartedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutionStartedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutionStartedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutionStartedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutionStartedAtDesc := "jobExecutionStartedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	jobExecutionStartedAtAsc := "jobExecutionStartedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	ruleIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	ruleIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	ruleIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	ruleIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	ruleIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	ruleIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	ruleIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	ruleIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdDesc := "ruleIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdAsc := "ruleIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdObjectDesc := "ruleIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdObjectAsc := "ruleIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	changeIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	changeIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	changeIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	changeIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	changeIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	changeIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	changeIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	changeIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	changeIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	changeIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	changeIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	changeIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	changeIdDesc := "changeIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	changeIdAsc := "changeIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	changeIdObjectDesc := "changeIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	changeIdObjectAsc := "changeIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTriggerHasExecutionTriggerIdObjectsDesc := "referencedByTriggerHasExecutionTriggerIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTriggerHasExecutionTriggerIdObjectsAsc := "referencedByTriggerHasExecutionTriggerIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionTriggerIdObjectsDesc := "referencedByExecutionTriggerIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionTriggerIdObjectsAsc := "referencedByExecutionTriggerIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.GetTriggers(context.Background()).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).JobExecutorClaimedUntilEq(jobExecutorClaimedUntilEq).JobExecutorClaimedUntilNe(jobExecutorClaimedUntilNe).JobExecutorClaimedUntilGt(jobExecutorClaimedUntilGt).JobExecutorClaimedUntilGte(jobExecutorClaimedUntilGte).JobExecutorClaimedUntilLt(jobExecutorClaimedUntilLt).JobExecutorClaimedUntilLte(jobExecutorClaimedUntilLte).JobExecutorClaimedUntilIn(jobExecutorClaimedUntilIn).JobExecutorClaimedUntilNotin(jobExecutorClaimedUntilNotin).JobExecutorClaimedUntilLike(jobExecutorClaimedUntilLike).JobExecutorClaimedUntilNotlike(jobExecutorClaimedUntilNotlike).JobExecutorClaimedUntilIlike(jobExecutorClaimedUntilIlike).JobExecutorClaimedUntilNotilike(jobExecutorClaimedUntilNotilike).JobExecutorClaimedUntilDesc(jobExecutorClaimedUntilDesc).JobExecutorClaimedUntilAsc(jobExecutorClaimedUntilAsc).JobExecutionStartedAtEq(jobExecutionStartedAtEq).JobExecutionStartedAtNe(jobExecutionStartedAtNe).JobExecutionStartedAtGt(jobExecutionStartedAtGt).JobExecutionStartedAtGte(jobExecutionStartedAtGte).JobExecutionStartedAtLt(jobExecutionStartedAtLt).JobExecutionStartedAtLte(jobExecutionStartedAtLte).JobExecutionStartedAtIn(jobExecutionStartedAtIn).JobExecutionStartedAtNotin(jobExecutionStartedAtNotin).JobExecutionStartedAtLike(jobExecutionStartedAtLike).JobExecutionStartedAtNotlike(jobExecutionStartedAtNotlike).JobExecutionStartedAtIlike(jobExecutionStartedAtIlike).JobExecutionStartedAtNotilike(jobExecutionStartedAtNotilike).JobExecutionStartedAtDesc(jobExecutionStartedAtDesc).JobExecutionStartedAtAsc(jobExecutionStartedAtAsc).RuleIdEq(ruleIdEq).RuleIdNe(ruleIdNe).RuleIdGt(ruleIdGt).RuleIdGte(ruleIdGte).RuleIdLt(ruleIdLt).RuleIdLte(ruleIdLte).RuleIdIn(ruleIdIn).RuleIdNotin(ruleIdNotin).RuleIdLike(ruleIdLike).RuleIdNotlike(ruleIdNotlike).RuleIdIlike(ruleIdIlike).RuleIdNotilike(ruleIdNotilike).RuleIdDesc(ruleIdDesc).RuleIdAsc(ruleIdAsc).RuleIdObjectDesc(ruleIdObjectDesc).RuleIdObjectAsc(ruleIdObjectAsc).ChangeIdEq(changeIdEq).ChangeIdNe(changeIdNe).ChangeIdGt(changeIdGt).ChangeIdGte(changeIdGte).ChangeIdLt(changeIdLt).ChangeIdLte(changeIdLte).ChangeIdIn(changeIdIn).ChangeIdNotin(changeIdNotin).ChangeIdLike(changeIdLike).ChangeIdNotlike(changeIdNotlike).ChangeIdIlike(changeIdIlike).ChangeIdNotilike(changeIdNotilike).ChangeIdDesc(changeIdDesc).ChangeIdAsc(changeIdAsc).ChangeIdObjectDesc(changeIdObjectDesc).ChangeIdObjectAsc(changeIdObjectAsc).ReferencedByTriggerHasExecutionTriggerIdObjectsDesc(referencedByTriggerHasExecutionTriggerIdObjectsDesc).ReferencedByTriggerHasExecutionTriggerIdObjectsAsc(referencedByTriggerHasExecutionTriggerIdObjectsAsc).ReferencedByExecutionTriggerIdObjectsDesc(referencedByExecutionTriggerIdObjectsDesc).ReferencedByExecutionTriggerIdObjectsAsc(referencedByExecutionTriggerIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTriggers`: ResponseWithGenericOfTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | SQL LIMIT operator | 
 **offset** | **int32** | SQL OFFSET operator | 
 **depth** | **int32** | Max recursion depth for loading foreign objects; default &#x3D; 1  (0 &#x3D; recurse until graph cycle detected, 1 &#x3D; this object only, 2 &#x3D; this object + neighbours, 3 &#x3D; this object + neighbours + their neighbours... etc) | 
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
 **jobExecutorClaimedUntilEq** | **time.Time** | SQL &#x3D; comparison | 
 **jobExecutorClaimedUntilNe** | **time.Time** | SQL !&#x3D; comparison | 
 **jobExecutorClaimedUntilGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **jobExecutorClaimedUntilGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **jobExecutorClaimedUntilLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **jobExecutorClaimedUntilLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **jobExecutorClaimedUntilIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **jobExecutorClaimedUntilNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **jobExecutorClaimedUntilLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **jobExecutorClaimedUntilAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **jobExecutionStartedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **jobExecutionStartedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **jobExecutionStartedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **jobExecutionStartedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **jobExecutionStartedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **jobExecutionStartedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **jobExecutionStartedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **jobExecutionStartedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **jobExecutionStartedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutionStartedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutionStartedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutionStartedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutionStartedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **jobExecutionStartedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdEq** | **string** | SQL &#x3D; comparison | 
 **ruleIdNe** | **string** | SQL !&#x3D; comparison | 
 **ruleIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **ruleIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **ruleIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **ruleIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **ruleIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **ruleIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **ruleIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **changeIdEq** | **string** | SQL &#x3D; comparison | 
 **changeIdNe** | **string** | SQL !&#x3D; comparison | 
 **changeIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **changeIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **changeIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **changeIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **changeIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **changeIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **changeIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **changeIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **changeIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **changeIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **changeIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **changeIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **changeIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **changeIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTriggerHasExecutionTriggerIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTriggerHasExecutionTriggerIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionTriggerIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionTriggerIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfTrigger**](ResponseWithGenericOfTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTrigger

> ResponseWithGenericOfTrigger PatchTrigger(ctx, primaryKey).Trigger(trigger).Depth(depth).Execute()



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
	trigger := *openapiclient.NewTrigger() // Trigger | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.PatchTrigger(context.Background(), primaryKey).Trigger(trigger).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.PatchTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTrigger`: ResponseWithGenericOfTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.PatchTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trigger** | [**Trigger**](Trigger.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfTrigger**](ResponseWithGenericOfTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTriggers

> ResponseWithGenericOfTrigger PostTriggers(ctx).Trigger(trigger).Depth(depth).Execute()



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
	trigger := []openapiclient.Trigger{*openapiclient.NewTrigger()} // []Trigger | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.PostTriggers(context.Background()).Trigger(trigger).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.PostTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTriggers`: ResponseWithGenericOfTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.PostTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trigger** | [**[]Trigger**](Trigger.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfTrigger**](ResponseWithGenericOfTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

