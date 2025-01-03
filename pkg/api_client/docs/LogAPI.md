# \LogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLog**](LogAPI.md#DeleteLog) | **Delete** /api/logs/{primaryKey} | 
[**GetLog**](LogAPI.md#GetLog) | **Get** /api/logs/{primaryKey} | 
[**GetLogs**](LogAPI.md#GetLogs) | **Get** /api/logs | 
[**PatchLog**](LogAPI.md#PatchLog) | **Patch** /api/logs/{primaryKey} | 
[**PostLogs**](LogAPI.md#PostLogs) | **Post** /api/logs | 



## DeleteLog

> DeleteLog(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.LogAPI.DeleteLog(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.DeleteLog``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLogRequest struct via the builder pattern


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


## GetLog

> ResponseWithGenericOfLog GetLog(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.LogAPI.GetLog(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLog`: ResponseWithGenericOfLog
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfLog**](ResponseWithGenericOfLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> ResponseWithGenericOfLog GetLogs(ctx).Limit(limit).Offset(offset).Depth(depth).OutputLoad(outputLoad).ReferencedByOutputLoad(referencedByOutputLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).BufferEq(bufferEq).BufferNe(bufferNe).BufferGt(bufferGt).BufferGte(bufferGte).BufferLt(bufferLt).BufferLte(bufferLte).BufferIn(bufferIn).BufferNotin(bufferNotin).BufferLike(bufferLike).BufferNotlike(bufferNotlike).BufferIlike(bufferIlike).BufferNotilike(bufferNotilike).BufferDesc(bufferDesc).BufferAsc(bufferAsc).OutputIdEq(outputIdEq).OutputIdNe(outputIdNe).OutputIdGt(outputIdGt).OutputIdGte(outputIdGte).OutputIdLt(outputIdLt).OutputIdLte(outputIdLte).OutputIdIn(outputIdIn).OutputIdNotin(outputIdNotin).OutputIdLike(outputIdLike).OutputIdNotlike(outputIdNotlike).OutputIdIlike(outputIdIlike).OutputIdNotilike(outputIdNotilike).OutputIdDesc(outputIdDesc).OutputIdAsc(outputIdAsc).OutputIdObjectDesc(outputIdObjectDesc).OutputIdObjectAsc(outputIdObjectAsc).ReferencedByOutputLogidObjectsDesc(referencedByOutputLogidObjectsDesc).ReferencedByOutputLogidObjectsAsc(referencedByOutputLogidObjectsAsc).Execute()



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
	outputLoad := "outputLoad_example" // string | load the given directly related Djangolang object, value is ignored (presence of key is sufficient) (optional)
	referencedByOutputLoad := "referencedByOutputLoad_example" // string | load the given indirectly related Djangolang objects, value is ignored (presence of key is sufficient) (optional)
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
	bufferEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	bufferNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	bufferGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	bufferGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	bufferLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	bufferLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	bufferIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	bufferNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	bufferLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	bufferNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	bufferIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	bufferNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	bufferDesc := "bufferDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	bufferAsc := "bufferAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	outputIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	outputIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	outputIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	outputIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	outputIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	outputIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	outputIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	outputIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	outputIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	outputIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	outputIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	outputIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	outputIdDesc := "outputIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	outputIdAsc := "outputIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	outputIdObjectDesc := "outputIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	outputIdObjectAsc := "outputIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByOutputLogidObjectsDesc := "referencedByOutputLogidObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByOutputLogidObjectsAsc := "referencedByOutputLogidObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.GetLogs(context.Background()).Limit(limit).Offset(offset).Depth(depth).OutputLoad(outputLoad).ReferencedByOutputLoad(referencedByOutputLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).BufferEq(bufferEq).BufferNe(bufferNe).BufferGt(bufferGt).BufferGte(bufferGte).BufferLt(bufferLt).BufferLte(bufferLte).BufferIn(bufferIn).BufferNotin(bufferNotin).BufferLike(bufferLike).BufferNotlike(bufferNotlike).BufferIlike(bufferIlike).BufferNotilike(bufferNotilike).BufferDesc(bufferDesc).BufferAsc(bufferAsc).OutputIdEq(outputIdEq).OutputIdNe(outputIdNe).OutputIdGt(outputIdGt).OutputIdGte(outputIdGte).OutputIdLt(outputIdLt).OutputIdLte(outputIdLte).OutputIdIn(outputIdIn).OutputIdNotin(outputIdNotin).OutputIdLike(outputIdLike).OutputIdNotlike(outputIdNotlike).OutputIdIlike(outputIdIlike).OutputIdNotilike(outputIdNotilike).OutputIdDesc(outputIdDesc).OutputIdAsc(outputIdAsc).OutputIdObjectDesc(outputIdObjectDesc).OutputIdObjectAsc(outputIdObjectAsc).ReferencedByOutputLogidObjectsDesc(referencedByOutputLogidObjectsDesc).ReferencedByOutputLogidObjectsAsc(referencedByOutputLogidObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.GetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogs`: ResponseWithGenericOfLog
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | SQL LIMIT operator | 
 **offset** | **int32** | SQL OFFSET operator | 
 **depth** | **int32** | Max recursion depth for loading foreign objects; default &#x3D; 1  (0 &#x3D; recurse until graph cycle detected, 1 &#x3D; this object only, 2 &#x3D; this object + neighbours, 3 &#x3D; this object + neighbours + their neighbours... etc) | 
 **outputLoad** | **string** | load the given directly related Djangolang object, value is ignored (presence of key is sufficient) | 
 **referencedByOutputLoad** | **string** | load the given indirectly related Djangolang objects, value is ignored (presence of key is sufficient) | 
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
 **bufferEq** | **string** | SQL &#x3D; comparison | 
 **bufferNe** | **string** | SQL !&#x3D; comparison | 
 **bufferGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **bufferGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **bufferLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **bufferLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **bufferIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **bufferNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **bufferLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **bufferNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **bufferIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **bufferNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **bufferDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **bufferAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **outputIdEq** | **string** | SQL &#x3D; comparison | 
 **outputIdNe** | **string** | SQL !&#x3D; comparison | 
 **outputIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **outputIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **outputIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **outputIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **outputIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **outputIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **outputIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **outputIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **outputIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **outputIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **outputIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **outputIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **outputIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **outputIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByOutputLogidObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByOutputLogidObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfLog**](ResponseWithGenericOfLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLog

> ResponseWithGenericOfLog PatchLog(ctx, primaryKey).Log(log).Depth(depth).Execute()



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
	log := *openapiclient.NewLog() // Log | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.PatchLog(context.Background(), primaryKey).Log(log).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.PatchLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchLog`: ResponseWithGenericOfLog
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.PatchLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **log** | [**Log**](Log.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfLog**](ResponseWithGenericOfLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLogs

> ResponseWithGenericOfLog PostLogs(ctx).Log(log).Depth(depth).Execute()



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
	log := []openapiclient.Log{*openapiclient.NewLog()} // []Log | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.PostLogs(context.Background()).Log(log).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.PostLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLogs`: ResponseWithGenericOfLog
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.PostLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **log** | [**[]Log**](Log.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfLog**](ResponseWithGenericOfLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

