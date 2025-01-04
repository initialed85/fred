# \TriggerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrigger**](TriggerAPI.md#DeleteTrigger) | **Delete** /api/triggers/{primaryKey} | 
[**GetTrigger**](TriggerAPI.md#GetTrigger) | **Get** /api/triggers/{primaryKey} | 
[**GetTriggers**](TriggerAPI.md#GetTriggers) | **Get** /api/triggers | 
[**PatchTrigger**](TriggerAPI.md#PatchTrigger) | **Patch** /api/triggers/{primaryKey} | 
[**PostTriggers**](TriggerAPI.md#PostTriggers) | **Post** /api/triggers | 
[**PostTriggersJobExecutorClaim**](TriggerAPI.md#PostTriggersJobExecutorClaim) | **Post** /api/triggers/{primaryKey}/job-executor-claim | 



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

> ResponseWithGenericOfTrigger GetTriggers(ctx).Limit(limit).Offset(offset).Depth(depth).RuleLoad(ruleLoad).JobLoad(jobLoad).ReferencedByExecutionLoad(referencedByExecutionLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).JobExecutorClaimedUntilEq(jobExecutorClaimedUntilEq).JobExecutorClaimedUntilNe(jobExecutorClaimedUntilNe).JobExecutorClaimedUntilGt(jobExecutorClaimedUntilGt).JobExecutorClaimedUntilGte(jobExecutorClaimedUntilGte).JobExecutorClaimedUntilLt(jobExecutorClaimedUntilLt).JobExecutorClaimedUntilLte(jobExecutorClaimedUntilLte).JobExecutorClaimedUntilIn(jobExecutorClaimedUntilIn).JobExecutorClaimedUntilNotin(jobExecutorClaimedUntilNotin).JobExecutorClaimedUntilContains(jobExecutorClaimedUntilContains).JobExecutorClaimedUntilNotcontains(jobExecutorClaimedUntilNotcontains).JobExecutorClaimedUntilLike(jobExecutorClaimedUntilLike).JobExecutorClaimedUntilNotlike(jobExecutorClaimedUntilNotlike).JobExecutorClaimedUntilIlike(jobExecutorClaimedUntilIlike).JobExecutorClaimedUntilNotilike(jobExecutorClaimedUntilNotilike).JobExecutorClaimedUntilDesc(jobExecutorClaimedUntilDesc).JobExecutorClaimedUntilAsc(jobExecutorClaimedUntilAsc).RuleIdEq(ruleIdEq).RuleIdNe(ruleIdNe).RuleIdGt(ruleIdGt).RuleIdGte(ruleIdGte).RuleIdLt(ruleIdLt).RuleIdLte(ruleIdLte).RuleIdIn(ruleIdIn).RuleIdNotin(ruleIdNotin).RuleIdContains(ruleIdContains).RuleIdNotcontains(ruleIdNotcontains).RuleIdLike(ruleIdLike).RuleIdNotlike(ruleIdNotlike).RuleIdIlike(ruleIdIlike).RuleIdNotilike(ruleIdNotilike).RuleIdDesc(ruleIdDesc).RuleIdAsc(ruleIdAsc).RuleIdObjectContains(ruleIdObjectContains).RuleIdObjectNotcontains(ruleIdObjectNotcontains).RuleIdObjectDesc(ruleIdObjectDesc).RuleIdObjectAsc(ruleIdObjectAsc).JobIdEq(jobIdEq).JobIdNe(jobIdNe).JobIdGt(jobIdGt).JobIdGte(jobIdGte).JobIdLt(jobIdLt).JobIdLte(jobIdLte).JobIdIn(jobIdIn).JobIdNotin(jobIdNotin).JobIdContains(jobIdContains).JobIdNotcontains(jobIdNotcontains).JobIdLike(jobIdLike).JobIdNotlike(jobIdNotlike).JobIdIlike(jobIdIlike).JobIdNotilike(jobIdNotilike).JobIdDesc(jobIdDesc).JobIdAsc(jobIdAsc).JobIdObjectContains(jobIdObjectContains).JobIdObjectNotcontains(jobIdObjectNotcontains).JobIdObjectDesc(jobIdObjectDesc).JobIdObjectAsc(jobIdObjectAsc).ReferencedByExecutionTriggerIdObjectsContains(referencedByExecutionTriggerIdObjectsContains).ReferencedByExecutionTriggerIdObjectsNotcontains(referencedByExecutionTriggerIdObjectsNotcontains).ReferencedByExecutionTriggerIdObjectsDesc(referencedByExecutionTriggerIdObjectsDesc).ReferencedByExecutionTriggerIdObjectsAsc(referencedByExecutionTriggerIdObjectsAsc).Execute()



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
	ruleLoad := "ruleLoad_example" // string | load the given directly related object, value is ignored (presence of key is sufficient) (optional)
	jobLoad := "jobLoad_example" // string | load the given directly related object, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionLoad := "referencedByExecutionLoad_example" // string | load the given indirectly related objects, value is ignored (presence of key is sufficient) (optional)
	idEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	idNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	idGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	idGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	idLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	idLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	idIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	idNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	idContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	idNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
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
	createdAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	createdAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
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
	updatedAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	updatedAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
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
	deletedAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	deletedAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
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
	jobExecutorClaimedUntilContains := time.Now() // time.Time | SQL @> comparison (optional)
	jobExecutorClaimedUntilNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
	jobExecutorClaimedUntilLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobExecutorClaimedUntilDesc := "jobExecutorClaimedUntilDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	jobExecutorClaimedUntilAsc := "jobExecutorClaimedUntilAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	ruleIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	ruleIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	ruleIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	ruleIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	ruleIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	ruleIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	ruleIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	ruleIdContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	ruleIdNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
	ruleIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdDesc := "ruleIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdAsc := "ruleIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdObjectContains := TODO // interface{} | SQL @> comparison (optional)
	ruleIdObjectNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	ruleIdObjectDesc := "ruleIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdObjectAsc := "ruleIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	jobIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	jobIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	jobIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	jobIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	jobIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	jobIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	jobIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	jobIdContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	jobIdNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
	jobIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdDesc := "jobIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdAsc := "jobIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdObjectContains := TODO // interface{} | SQL @> comparison (optional)
	jobIdObjectNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	jobIdObjectDesc := "jobIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdObjectAsc := "jobIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionTriggerIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByExecutionTriggerIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByExecutionTriggerIdObjectsDesc := "referencedByExecutionTriggerIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionTriggerIdObjectsAsc := "referencedByExecutionTriggerIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.GetTriggers(context.Background()).Limit(limit).Offset(offset).Depth(depth).RuleLoad(ruleLoad).JobLoad(jobLoad).ReferencedByExecutionLoad(referencedByExecutionLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).JobExecutorClaimedUntilEq(jobExecutorClaimedUntilEq).JobExecutorClaimedUntilNe(jobExecutorClaimedUntilNe).JobExecutorClaimedUntilGt(jobExecutorClaimedUntilGt).JobExecutorClaimedUntilGte(jobExecutorClaimedUntilGte).JobExecutorClaimedUntilLt(jobExecutorClaimedUntilLt).JobExecutorClaimedUntilLte(jobExecutorClaimedUntilLte).JobExecutorClaimedUntilIn(jobExecutorClaimedUntilIn).JobExecutorClaimedUntilNotin(jobExecutorClaimedUntilNotin).JobExecutorClaimedUntilContains(jobExecutorClaimedUntilContains).JobExecutorClaimedUntilNotcontains(jobExecutorClaimedUntilNotcontains).JobExecutorClaimedUntilLike(jobExecutorClaimedUntilLike).JobExecutorClaimedUntilNotlike(jobExecutorClaimedUntilNotlike).JobExecutorClaimedUntilIlike(jobExecutorClaimedUntilIlike).JobExecutorClaimedUntilNotilike(jobExecutorClaimedUntilNotilike).JobExecutorClaimedUntilDesc(jobExecutorClaimedUntilDesc).JobExecutorClaimedUntilAsc(jobExecutorClaimedUntilAsc).RuleIdEq(ruleIdEq).RuleIdNe(ruleIdNe).RuleIdGt(ruleIdGt).RuleIdGte(ruleIdGte).RuleIdLt(ruleIdLt).RuleIdLte(ruleIdLte).RuleIdIn(ruleIdIn).RuleIdNotin(ruleIdNotin).RuleIdContains(ruleIdContains).RuleIdNotcontains(ruleIdNotcontains).RuleIdLike(ruleIdLike).RuleIdNotlike(ruleIdNotlike).RuleIdIlike(ruleIdIlike).RuleIdNotilike(ruleIdNotilike).RuleIdDesc(ruleIdDesc).RuleIdAsc(ruleIdAsc).RuleIdObjectContains(ruleIdObjectContains).RuleIdObjectNotcontains(ruleIdObjectNotcontains).RuleIdObjectDesc(ruleIdObjectDesc).RuleIdObjectAsc(ruleIdObjectAsc).JobIdEq(jobIdEq).JobIdNe(jobIdNe).JobIdGt(jobIdGt).JobIdGte(jobIdGte).JobIdLt(jobIdLt).JobIdLte(jobIdLte).JobIdIn(jobIdIn).JobIdNotin(jobIdNotin).JobIdContains(jobIdContains).JobIdNotcontains(jobIdNotcontains).JobIdLike(jobIdLike).JobIdNotlike(jobIdNotlike).JobIdIlike(jobIdIlike).JobIdNotilike(jobIdNotilike).JobIdDesc(jobIdDesc).JobIdAsc(jobIdAsc).JobIdObjectContains(jobIdObjectContains).JobIdObjectNotcontains(jobIdObjectNotcontains).JobIdObjectDesc(jobIdObjectDesc).JobIdObjectAsc(jobIdObjectAsc).ReferencedByExecutionTriggerIdObjectsContains(referencedByExecutionTriggerIdObjectsContains).ReferencedByExecutionTriggerIdObjectsNotcontains(referencedByExecutionTriggerIdObjectsNotcontains).ReferencedByExecutionTriggerIdObjectsDesc(referencedByExecutionTriggerIdObjectsDesc).ReferencedByExecutionTriggerIdObjectsAsc(referencedByExecutionTriggerIdObjectsAsc).Execute()
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
 **ruleLoad** | **string** | load the given directly related object, value is ignored (presence of key is sufficient) | 
 **jobLoad** | **string** | load the given directly related object, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionLoad** | **string** | load the given indirectly related objects, value is ignored (presence of key is sufficient) | 
 **idEq** | **string** | SQL &#x3D; comparison | 
 **idNe** | **string** | SQL !&#x3D; comparison | 
 **idGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **idGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **idLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **idLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **idIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **idNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **idContains** | **string** | SQL @&gt; comparison | 
 **idNotcontains** | **string** | SQL NOT @&gt; comparison | 
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
 **createdAtContains** | **time.Time** | SQL @&gt; comparison | 
 **createdAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
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
 **updatedAtContains** | **time.Time** | SQL @&gt; comparison | 
 **updatedAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
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
 **deletedAtContains** | **time.Time** | SQL @&gt; comparison | 
 **deletedAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
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
 **jobExecutorClaimedUntilContains** | **time.Time** | SQL @&gt; comparison | 
 **jobExecutorClaimedUntilNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
 **jobExecutorClaimedUntilLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobExecutorClaimedUntilDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **jobExecutorClaimedUntilAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdEq** | **string** | SQL &#x3D; comparison | 
 **ruleIdNe** | **string** | SQL !&#x3D; comparison | 
 **ruleIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **ruleIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **ruleIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **ruleIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **ruleIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **ruleIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **ruleIdContains** | **string** | SQL @&gt; comparison | 
 **ruleIdNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **ruleIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdObjectContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **ruleIdObjectNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **ruleIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **jobIdEq** | **string** | SQL &#x3D; comparison | 
 **jobIdNe** | **string** | SQL !&#x3D; comparison | 
 **jobIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **jobIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **jobIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **jobIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **jobIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **jobIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **jobIdContains** | **string** | SQL @&gt; comparison | 
 **jobIdNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **jobIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **jobIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **jobIdObjectContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **jobIdObjectNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **jobIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **jobIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionTriggerIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByExecutionTriggerIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
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


## PostTriggersJobExecutorClaim

> ResponseWithGenericOfTrigger PostTriggersJobExecutorClaim(ctx, primaryKey).TriggerJobExecutorClaimRequest(triggerJobExecutorClaimRequest).Depth(depth).Execute()



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
	triggerJobExecutorClaimRequest := *openapiclient.NewTriggerJobExecutorClaimRequest() // TriggerJobExecutorClaimRequest | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.PostTriggersJobExecutorClaim(context.Background(), primaryKey).TriggerJobExecutorClaimRequest(triggerJobExecutorClaimRequest).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.PostTriggersJobExecutorClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTriggersJobExecutorClaim`: ResponseWithGenericOfTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.PostTriggersJobExecutorClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTriggersJobExecutorClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerJobExecutorClaimRequest** | [**TriggerJobExecutorClaimRequest**](TriggerJobExecutorClaimRequest.md) |  | 
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

