# \ChangeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteChange**](ChangeAPI.md#DeleteChange) | **Delete** /api/changes/{primaryKey} | 
[**GetChange**](ChangeAPI.md#GetChange) | **Get** /api/changes/{primaryKey} | 
[**GetChanges**](ChangeAPI.md#GetChanges) | **Get** /api/changes | 
[**PatchChange**](ChangeAPI.md#PatchChange) | **Patch** /api/changes/{primaryKey} | 
[**PostChanges**](ChangeAPI.md#PostChanges) | **Post** /api/changes | 



## DeleteChange

> DeleteChange(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.ChangeAPI.DeleteChange(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeAPI.DeleteChange``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteChangeRequest struct via the builder pattern


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


## GetChange

> ResponseWithGenericOfChange GetChange(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.ChangeAPI.GetChange(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeAPI.GetChange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChange`: ResponseWithGenericOfChange
	fmt.Fprintf(os.Stdout, "Response from `ChangeAPI.GetChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfChange**](ResponseWithGenericOfChange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges

> ResponseWithGenericOfChange GetChanges(ctx).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).BranchNameEq(branchNameEq).BranchNameNe(branchNameNe).BranchNameGt(branchNameGt).BranchNameGte(branchNameGte).BranchNameLt(branchNameLt).BranchNameLte(branchNameLte).BranchNameIn(branchNameIn).BranchNameNotin(branchNameNotin).BranchNameLike(branchNameLike).BranchNameNotlike(branchNameNotlike).BranchNameIlike(branchNameIlike).BranchNameNotilike(branchNameNotilike).BranchNameDesc(branchNameDesc).BranchNameAsc(branchNameAsc).CommitHashEq(commitHashEq).CommitHashNe(commitHashNe).CommitHashGt(commitHashGt).CommitHashGte(commitHashGte).CommitHashLt(commitHashLt).CommitHashLte(commitHashLte).CommitHashIn(commitHashIn).CommitHashNotin(commitHashNotin).CommitHashLike(commitHashLike).CommitHashNotlike(commitHashNotlike).CommitHashIlike(commitHashIlike).CommitHashNotilike(commitHashNotilike).CommitHashDesc(commitHashDesc).CommitHashAsc(commitHashAsc).MessageEq(messageEq).MessageNe(messageNe).MessageGt(messageGt).MessageGte(messageGte).MessageLt(messageLt).MessageLte(messageLte).MessageIn(messageIn).MessageNotin(messageNotin).MessageLike(messageLike).MessageNotlike(messageNotlike).MessageIlike(messageIlike).MessageNotilike(messageNotilike).MessageDesc(messageDesc).MessageAsc(messageAsc).AuthoredByEq(authoredByEq).AuthoredByNe(authoredByNe).AuthoredByGt(authoredByGt).AuthoredByGte(authoredByGte).AuthoredByLt(authoredByLt).AuthoredByLte(authoredByLte).AuthoredByIn(authoredByIn).AuthoredByNotin(authoredByNotin).AuthoredByLike(authoredByLike).AuthoredByNotlike(authoredByNotlike).AuthoredByIlike(authoredByIlike).AuthoredByNotilike(authoredByNotilike).AuthoredByDesc(authoredByDesc).AuthoredByAsc(authoredByAsc).AuthoredAtEq(authoredAtEq).AuthoredAtNe(authoredAtNe).AuthoredAtGt(authoredAtGt).AuthoredAtGte(authoredAtGte).AuthoredAtLt(authoredAtLt).AuthoredAtLte(authoredAtLte).AuthoredAtIn(authoredAtIn).AuthoredAtNotin(authoredAtNotin).AuthoredAtLike(authoredAtLike).AuthoredAtNotlike(authoredAtNotlike).AuthoredAtIlike(authoredAtIlike).AuthoredAtNotilike(authoredAtNotilike).AuthoredAtDesc(authoredAtDesc).AuthoredAtAsc(authoredAtAsc).CommittedByEq(committedByEq).CommittedByNe(committedByNe).CommittedByGt(committedByGt).CommittedByGte(committedByGte).CommittedByLt(committedByLt).CommittedByLte(committedByLte).CommittedByIn(committedByIn).CommittedByNotin(committedByNotin).CommittedByLike(committedByLike).CommittedByNotlike(committedByNotlike).CommittedByIlike(committedByIlike).CommittedByNotilike(committedByNotilike).CommittedByDesc(committedByDesc).CommittedByAsc(committedByAsc).CommittedAtEq(committedAtEq).CommittedAtNe(committedAtNe).CommittedAtGt(committedAtGt).CommittedAtGte(committedAtGte).CommittedAtLt(committedAtLt).CommittedAtLte(committedAtLte).CommittedAtIn(committedAtIn).CommittedAtNotin(committedAtNotin).CommittedAtLike(committedAtLike).CommittedAtNotlike(committedAtNotlike).CommittedAtIlike(committedAtIlike).CommittedAtNotilike(committedAtNotilike).CommittedAtDesc(committedAtDesc).CommittedAtAsc(committedAtAsc).RepositoryIdEq(repositoryIdEq).RepositoryIdNe(repositoryIdNe).RepositoryIdGt(repositoryIdGt).RepositoryIdGte(repositoryIdGte).RepositoryIdLt(repositoryIdLt).RepositoryIdLte(repositoryIdLte).RepositoryIdIn(repositoryIdIn).RepositoryIdNotin(repositoryIdNotin).RepositoryIdLike(repositoryIdLike).RepositoryIdNotlike(repositoryIdNotlike).RepositoryIdIlike(repositoryIdIlike).RepositoryIdNotilike(repositoryIdNotilike).RepositoryIdDesc(repositoryIdDesc).RepositoryIdAsc(repositoryIdAsc).RepositoryIdObjectDesc(repositoryIdObjectDesc).RepositoryIdObjectAsc(repositoryIdObjectAsc).ReferencedByTriggerChangeIdObjectsDesc(referencedByTriggerChangeIdObjectsDesc).ReferencedByTriggerChangeIdObjectsAsc(referencedByTriggerChangeIdObjectsAsc).Execute()



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
	branchNameEq := "branchNameEq_example" // string | SQL = comparison (optional)
	branchNameNe := "branchNameNe_example" // string | SQL != comparison (optional)
	branchNameGt := "branchNameGt_example" // string | SQL > comparison, may not work with all column types (optional)
	branchNameGte := "branchNameGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	branchNameLt := "branchNameLt_example" // string | SQL < comparison, may not work with all column types (optional)
	branchNameLte := "branchNameLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	branchNameIn := "branchNameIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	branchNameNotin := "branchNameNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	branchNameLike := "branchNameLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchNameNotlike := "branchNameNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchNameIlike := "branchNameIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchNameNotilike := "branchNameNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchNameDesc := "branchNameDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	branchNameAsc := "branchNameAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	commitHashEq := "commitHashEq_example" // string | SQL = comparison (optional)
	commitHashNe := "commitHashNe_example" // string | SQL != comparison (optional)
	commitHashGt := "commitHashGt_example" // string | SQL > comparison, may not work with all column types (optional)
	commitHashGte := "commitHashGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	commitHashLt := "commitHashLt_example" // string | SQL < comparison, may not work with all column types (optional)
	commitHashLte := "commitHashLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	commitHashIn := "commitHashIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	commitHashNotin := "commitHashNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	commitHashLike := "commitHashLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	commitHashNotlike := "commitHashNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	commitHashIlike := "commitHashIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	commitHashNotilike := "commitHashNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	commitHashDesc := "commitHashDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	commitHashAsc := "commitHashAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	messageEq := "messageEq_example" // string | SQL = comparison (optional)
	messageNe := "messageNe_example" // string | SQL != comparison (optional)
	messageGt := "messageGt_example" // string | SQL > comparison, may not work with all column types (optional)
	messageGte := "messageGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	messageLt := "messageLt_example" // string | SQL < comparison, may not work with all column types (optional)
	messageLte := "messageLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	messageIn := "messageIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	messageNotin := "messageNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	messageLike := "messageLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	messageNotlike := "messageNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	messageIlike := "messageIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	messageNotilike := "messageNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	messageDesc := "messageDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	messageAsc := "messageAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	authoredByEq := "authoredByEq_example" // string | SQL = comparison (optional)
	authoredByNe := "authoredByNe_example" // string | SQL != comparison (optional)
	authoredByGt := "authoredByGt_example" // string | SQL > comparison, may not work with all column types (optional)
	authoredByGte := "authoredByGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	authoredByLt := "authoredByLt_example" // string | SQL < comparison, may not work with all column types (optional)
	authoredByLte := "authoredByLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	authoredByIn := "authoredByIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	authoredByNotin := "authoredByNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	authoredByLike := "authoredByLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredByNotlike := "authoredByNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredByIlike := "authoredByIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredByNotilike := "authoredByNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredByDesc := "authoredByDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	authoredByAsc := "authoredByAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	authoredAtEq := time.Now() // time.Time | SQL = comparison (optional)
	authoredAtNe := time.Now() // time.Time | SQL != comparison (optional)
	authoredAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	authoredAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	authoredAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	authoredAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	authoredAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	authoredAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	authoredAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	authoredAtDesc := "authoredAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	authoredAtAsc := "authoredAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	committedByEq := "committedByEq_example" // string | SQL = comparison (optional)
	committedByNe := "committedByNe_example" // string | SQL != comparison (optional)
	committedByGt := "committedByGt_example" // string | SQL > comparison, may not work with all column types (optional)
	committedByGte := "committedByGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	committedByLt := "committedByLt_example" // string | SQL < comparison, may not work with all column types (optional)
	committedByLte := "committedByLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	committedByIn := "committedByIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	committedByNotin := "committedByNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	committedByLike := "committedByLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedByNotlike := "committedByNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedByIlike := "committedByIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedByNotilike := "committedByNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedByDesc := "committedByDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	committedByAsc := "committedByAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	committedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	committedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	committedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	committedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	committedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	committedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	committedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	committedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	committedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	committedAtDesc := "committedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	committedAtAsc := "committedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	repositoryIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	repositoryIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	repositoryIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	repositoryIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	repositoryIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	repositoryIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	repositoryIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	repositoryIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdDesc := "repositoryIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdAsc := "repositoryIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdObjectDesc := "repositoryIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdObjectAsc := "repositoryIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTriggerChangeIdObjectsDesc := "referencedByTriggerChangeIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTriggerChangeIdObjectsAsc := "referencedByTriggerChangeIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeAPI.GetChanges(context.Background()).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).BranchNameEq(branchNameEq).BranchNameNe(branchNameNe).BranchNameGt(branchNameGt).BranchNameGte(branchNameGte).BranchNameLt(branchNameLt).BranchNameLte(branchNameLte).BranchNameIn(branchNameIn).BranchNameNotin(branchNameNotin).BranchNameLike(branchNameLike).BranchNameNotlike(branchNameNotlike).BranchNameIlike(branchNameIlike).BranchNameNotilike(branchNameNotilike).BranchNameDesc(branchNameDesc).BranchNameAsc(branchNameAsc).CommitHashEq(commitHashEq).CommitHashNe(commitHashNe).CommitHashGt(commitHashGt).CommitHashGte(commitHashGte).CommitHashLt(commitHashLt).CommitHashLte(commitHashLte).CommitHashIn(commitHashIn).CommitHashNotin(commitHashNotin).CommitHashLike(commitHashLike).CommitHashNotlike(commitHashNotlike).CommitHashIlike(commitHashIlike).CommitHashNotilike(commitHashNotilike).CommitHashDesc(commitHashDesc).CommitHashAsc(commitHashAsc).MessageEq(messageEq).MessageNe(messageNe).MessageGt(messageGt).MessageGte(messageGte).MessageLt(messageLt).MessageLte(messageLte).MessageIn(messageIn).MessageNotin(messageNotin).MessageLike(messageLike).MessageNotlike(messageNotlike).MessageIlike(messageIlike).MessageNotilike(messageNotilike).MessageDesc(messageDesc).MessageAsc(messageAsc).AuthoredByEq(authoredByEq).AuthoredByNe(authoredByNe).AuthoredByGt(authoredByGt).AuthoredByGte(authoredByGte).AuthoredByLt(authoredByLt).AuthoredByLte(authoredByLte).AuthoredByIn(authoredByIn).AuthoredByNotin(authoredByNotin).AuthoredByLike(authoredByLike).AuthoredByNotlike(authoredByNotlike).AuthoredByIlike(authoredByIlike).AuthoredByNotilike(authoredByNotilike).AuthoredByDesc(authoredByDesc).AuthoredByAsc(authoredByAsc).AuthoredAtEq(authoredAtEq).AuthoredAtNe(authoredAtNe).AuthoredAtGt(authoredAtGt).AuthoredAtGte(authoredAtGte).AuthoredAtLt(authoredAtLt).AuthoredAtLte(authoredAtLte).AuthoredAtIn(authoredAtIn).AuthoredAtNotin(authoredAtNotin).AuthoredAtLike(authoredAtLike).AuthoredAtNotlike(authoredAtNotlike).AuthoredAtIlike(authoredAtIlike).AuthoredAtNotilike(authoredAtNotilike).AuthoredAtDesc(authoredAtDesc).AuthoredAtAsc(authoredAtAsc).CommittedByEq(committedByEq).CommittedByNe(committedByNe).CommittedByGt(committedByGt).CommittedByGte(committedByGte).CommittedByLt(committedByLt).CommittedByLte(committedByLte).CommittedByIn(committedByIn).CommittedByNotin(committedByNotin).CommittedByLike(committedByLike).CommittedByNotlike(committedByNotlike).CommittedByIlike(committedByIlike).CommittedByNotilike(committedByNotilike).CommittedByDesc(committedByDesc).CommittedByAsc(committedByAsc).CommittedAtEq(committedAtEq).CommittedAtNe(committedAtNe).CommittedAtGt(committedAtGt).CommittedAtGte(committedAtGte).CommittedAtLt(committedAtLt).CommittedAtLte(committedAtLte).CommittedAtIn(committedAtIn).CommittedAtNotin(committedAtNotin).CommittedAtLike(committedAtLike).CommittedAtNotlike(committedAtNotlike).CommittedAtIlike(committedAtIlike).CommittedAtNotilike(committedAtNotilike).CommittedAtDesc(committedAtDesc).CommittedAtAsc(committedAtAsc).RepositoryIdEq(repositoryIdEq).RepositoryIdNe(repositoryIdNe).RepositoryIdGt(repositoryIdGt).RepositoryIdGte(repositoryIdGte).RepositoryIdLt(repositoryIdLt).RepositoryIdLte(repositoryIdLte).RepositoryIdIn(repositoryIdIn).RepositoryIdNotin(repositoryIdNotin).RepositoryIdLike(repositoryIdLike).RepositoryIdNotlike(repositoryIdNotlike).RepositoryIdIlike(repositoryIdIlike).RepositoryIdNotilike(repositoryIdNotilike).RepositoryIdDesc(repositoryIdDesc).RepositoryIdAsc(repositoryIdAsc).RepositoryIdObjectDesc(repositoryIdObjectDesc).RepositoryIdObjectAsc(repositoryIdObjectAsc).ReferencedByTriggerChangeIdObjectsDesc(referencedByTriggerChangeIdObjectsDesc).ReferencedByTriggerChangeIdObjectsAsc(referencedByTriggerChangeIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeAPI.GetChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChanges`: ResponseWithGenericOfChange
	fmt.Fprintf(os.Stdout, "Response from `ChangeAPI.GetChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangesRequest struct via the builder pattern


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
 **branchNameEq** | **string** | SQL &#x3D; comparison | 
 **branchNameNe** | **string** | SQL !&#x3D; comparison | 
 **branchNameGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **branchNameGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **branchNameLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **branchNameLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **branchNameIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **branchNameNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **branchNameLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchNameNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchNameIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchNameNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchNameDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **branchNameAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **commitHashEq** | **string** | SQL &#x3D; comparison | 
 **commitHashNe** | **string** | SQL !&#x3D; comparison | 
 **commitHashGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **commitHashGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **commitHashLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **commitHashLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **commitHashIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **commitHashNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **commitHashLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **commitHashNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **commitHashIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **commitHashNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **commitHashDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **commitHashAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **messageEq** | **string** | SQL &#x3D; comparison | 
 **messageNe** | **string** | SQL !&#x3D; comparison | 
 **messageGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **messageGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **messageLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **messageLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **messageIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **messageNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **messageLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **messageNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **messageIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **messageNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **messageDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **messageAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **authoredByEq** | **string** | SQL &#x3D; comparison | 
 **authoredByNe** | **string** | SQL !&#x3D; comparison | 
 **authoredByGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **authoredByGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **authoredByLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **authoredByLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **authoredByIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **authoredByNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **authoredByLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredByNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredByIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredByNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredByDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **authoredByAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **authoredAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **authoredAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **authoredAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **authoredAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **authoredAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **authoredAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **authoredAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **authoredAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **authoredAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **authoredAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **authoredAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **committedByEq** | **string** | SQL &#x3D; comparison | 
 **committedByNe** | **string** | SQL !&#x3D; comparison | 
 **committedByGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **committedByGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **committedByLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **committedByLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **committedByIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **committedByNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **committedByLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedByNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedByIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedByNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedByDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **committedByAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **committedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **committedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **committedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **committedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **committedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **committedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **committedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **committedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **committedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **committedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **committedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdEq** | **string** | SQL &#x3D; comparison | 
 **repositoryIdNe** | **string** | SQL !&#x3D; comparison | 
 **repositoryIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **repositoryIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **repositoryIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **repositoryIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **repositoryIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **repositoryIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **repositoryIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTriggerChangeIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTriggerChangeIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfChange**](ResponseWithGenericOfChange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchChange

> ResponseWithGenericOfChange PatchChange(ctx, primaryKey).Change(change).Depth(depth).Execute()



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
	change := *openapiclient.NewChange() // Change | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeAPI.PatchChange(context.Background(), primaryKey).Change(change).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeAPI.PatchChange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchChange`: ResponseWithGenericOfChange
	fmt.Fprintf(os.Stdout, "Response from `ChangeAPI.PatchChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **change** | [**Change**](Change.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfChange**](ResponseWithGenericOfChange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostChanges

> ResponseWithGenericOfChange PostChanges(ctx).Change(change).Depth(depth).Execute()



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
	change := []openapiclient.Change{*openapiclient.NewChange()} // []Change | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeAPI.PostChanges(context.Background()).Change(change).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeAPI.PostChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostChanges`: ResponseWithGenericOfChange
	fmt.Fprintf(os.Stdout, "Response from `ChangeAPI.PostChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **change** | [**[]Change**](Change.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfChange**](ResponseWithGenericOfChange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

