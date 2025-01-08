# \RepositoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRepository**](RepositoryAPI.md#DeleteRepository) | **Delete** /api/repositories/{primaryKey} | 
[**GetRepositories**](RepositoryAPI.md#GetRepositories) | **Get** /api/repositories | 
[**GetRepository**](RepositoryAPI.md#GetRepository) | **Get** /api/repositories/{primaryKey} | 
[**PatchRepository**](RepositoryAPI.md#PatchRepository) | **Patch** /api/repositories/{primaryKey} | 
[**PostRepositories**](RepositoryAPI.md#PostRepositories) | **Post** /api/repositories | 
[**PostRepositoriesRepositorySyncerClaim**](RepositoryAPI.md#PostRepositoriesRepositorySyncerClaim) | **Post** /api/repositories/{primaryKey}/repository-syncer-claim | 



## DeleteRepository

> DeleteRepository(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.RepositoryAPI.DeleteRepository(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteRepository``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


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


## GetRepositories

> ResponseWithGenericOfRepository GetRepositories(ctx).Limit(limit).Offset(offset).Depth(depth).ReferencedByChangeLoad(referencedByChangeLoad).ReferencedByJobLoad(referencedByJobLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).UrlEq(urlEq).UrlNe(urlNe).UrlGt(urlGt).UrlGte(urlGte).UrlLt(urlLt).UrlLte(urlLte).UrlIn(urlIn).UrlNotin(urlNotin).UrlContains(urlContains).UrlNotcontains(urlNotcontains).UrlLike(urlLike).UrlNotlike(urlNotlike).UrlIlike(urlIlike).UrlNotilike(urlNotilike).UrlDesc(urlDesc).UrlAsc(urlAsc).NameEq(nameEq).NameNe(nameNe).NameGt(nameGt).NameGte(nameGte).NameLt(nameLt).NameLte(nameLte).NameIn(nameIn).NameNotin(nameNotin).NameContains(nameContains).NameNotcontains(nameNotcontains).NameLike(nameLike).NameNotlike(nameNotlike).NameIlike(nameIlike).NameNotilike(nameNotilike).NameDesc(nameDesc).NameAsc(nameAsc).HandledAtEq(handledAtEq).HandledAtNe(handledAtNe).HandledAtGt(handledAtGt).HandledAtGte(handledAtGte).HandledAtLt(handledAtLt).HandledAtLte(handledAtLte).HandledAtIn(handledAtIn).HandledAtNotin(handledAtNotin).HandledAtContains(handledAtContains).HandledAtNotcontains(handledAtNotcontains).HandledAtLike(handledAtLike).HandledAtNotlike(handledAtNotlike).HandledAtIlike(handledAtIlike).HandledAtNotilike(handledAtNotilike).HandledAtDesc(handledAtDesc).HandledAtAsc(handledAtAsc).RepositorySyncerClaimedUntilEq(repositorySyncerClaimedUntilEq).RepositorySyncerClaimedUntilNe(repositorySyncerClaimedUntilNe).RepositorySyncerClaimedUntilGt(repositorySyncerClaimedUntilGt).RepositorySyncerClaimedUntilGte(repositorySyncerClaimedUntilGte).RepositorySyncerClaimedUntilLt(repositorySyncerClaimedUntilLt).RepositorySyncerClaimedUntilLte(repositorySyncerClaimedUntilLte).RepositorySyncerClaimedUntilIn(repositorySyncerClaimedUntilIn).RepositorySyncerClaimedUntilNotin(repositorySyncerClaimedUntilNotin).RepositorySyncerClaimedUntilContains(repositorySyncerClaimedUntilContains).RepositorySyncerClaimedUntilNotcontains(repositorySyncerClaimedUntilNotcontains).RepositorySyncerClaimedUntilLike(repositorySyncerClaimedUntilLike).RepositorySyncerClaimedUntilNotlike(repositorySyncerClaimedUntilNotlike).RepositorySyncerClaimedUntilIlike(repositorySyncerClaimedUntilIlike).RepositorySyncerClaimedUntilNotilike(repositorySyncerClaimedUntilNotilike).RepositorySyncerClaimedUntilDesc(repositorySyncerClaimedUntilDesc).RepositorySyncerClaimedUntilAsc(repositorySyncerClaimedUntilAsc).ReferencedByChangeRepositoryIdObjectsContains(referencedByChangeRepositoryIdObjectsContains).ReferencedByChangeRepositoryIdObjectsNotcontains(referencedByChangeRepositoryIdObjectsNotcontains).ReferencedByChangeRepositoryIdObjectsDesc(referencedByChangeRepositoryIdObjectsDesc).ReferencedByChangeRepositoryIdObjectsAsc(referencedByChangeRepositoryIdObjectsAsc).ReferencedByJobRepositoryIdObjectsContains(referencedByJobRepositoryIdObjectsContains).ReferencedByJobRepositoryIdObjectsNotcontains(referencedByJobRepositoryIdObjectsNotcontains).ReferencedByJobRepositoryIdObjectsDesc(referencedByJobRepositoryIdObjectsDesc).ReferencedByJobRepositoryIdObjectsAsc(referencedByJobRepositoryIdObjectsAsc).Execute()



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
	referencedByChangeLoad := "referencedByChangeLoad_example" // string | load the given indirectly related objects, value is ignored (presence of key is sufficient) (optional)
	referencedByJobLoad := "referencedByJobLoad_example" // string | load the given indirectly related objects, value is ignored (presence of key is sufficient) (optional)
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
	urlEq := "urlEq_example" // string | SQL = comparison (optional)
	urlNe := "urlNe_example" // string | SQL != comparison (optional)
	urlGt := "urlGt_example" // string | SQL > comparison, may not work with all column types (optional)
	urlGte := "urlGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	urlLt := "urlLt_example" // string | SQL < comparison, may not work with all column types (optional)
	urlLte := "urlLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	urlIn := "urlIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	urlNotin := "urlNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	urlContains := "urlContains_example" // string | SQL @> comparison (optional)
	urlNotcontains := "urlNotcontains_example" // string | SQL NOT @> comparison (optional)
	urlLike := "urlLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlNotlike := "urlNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlIlike := "urlIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlNotilike := "urlNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlDesc := "urlDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	urlAsc := "urlAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	nameEq := "nameEq_example" // string | SQL = comparison (optional)
	nameNe := "nameNe_example" // string | SQL != comparison (optional)
	nameGt := "nameGt_example" // string | SQL > comparison, may not work with all column types (optional)
	nameGte := "nameGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	nameLt := "nameLt_example" // string | SQL < comparison, may not work with all column types (optional)
	nameLte := "nameLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	nameIn := "nameIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	nameNotin := "nameNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	nameContains := "nameContains_example" // string | SQL @> comparison (optional)
	nameNotcontains := "nameNotcontains_example" // string | SQL NOT @> comparison (optional)
	nameLike := "nameLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameNotlike := "nameNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameIlike := "nameIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameNotilike := "nameNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameDesc := "nameDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	nameAsc := "nameAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	handledAtEq := time.Now() // time.Time | SQL = comparison (optional)
	handledAtNe := time.Now() // time.Time | SQL != comparison (optional)
	handledAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	handledAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	handledAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	handledAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	handledAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	handledAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	handledAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	handledAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
	handledAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	handledAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	handledAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	handledAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	handledAtDesc := "handledAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	handledAtAsc := "handledAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	repositorySyncerClaimedUntilEq := time.Now() // time.Time | SQL = comparison (optional)
	repositorySyncerClaimedUntilNe := time.Now() // time.Time | SQL != comparison (optional)
	repositorySyncerClaimedUntilGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	repositorySyncerClaimedUntilGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	repositorySyncerClaimedUntilLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	repositorySyncerClaimedUntilLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	repositorySyncerClaimedUntilIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	repositorySyncerClaimedUntilNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	repositorySyncerClaimedUntilContains := time.Now() // time.Time | SQL @> comparison (optional)
	repositorySyncerClaimedUntilNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
	repositorySyncerClaimedUntilLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositorySyncerClaimedUntilNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositorySyncerClaimedUntilIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositorySyncerClaimedUntilNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositorySyncerClaimedUntilDesc := "repositorySyncerClaimedUntilDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	repositorySyncerClaimedUntilAsc := "repositorySyncerClaimedUntilAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByChangeRepositoryIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByChangeRepositoryIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByChangeRepositoryIdObjectsDesc := "referencedByChangeRepositoryIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByChangeRepositoryIdObjectsAsc := "referencedByChangeRepositoryIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobRepositoryIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByJobRepositoryIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByJobRepositoryIdObjectsDesc := "referencedByJobRepositoryIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobRepositoryIdObjectsAsc := "referencedByJobRepositoryIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRepositories(context.Background()).Limit(limit).Offset(offset).Depth(depth).ReferencedByChangeLoad(referencedByChangeLoad).ReferencedByJobLoad(referencedByJobLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).UrlEq(urlEq).UrlNe(urlNe).UrlGt(urlGt).UrlGte(urlGte).UrlLt(urlLt).UrlLte(urlLte).UrlIn(urlIn).UrlNotin(urlNotin).UrlContains(urlContains).UrlNotcontains(urlNotcontains).UrlLike(urlLike).UrlNotlike(urlNotlike).UrlIlike(urlIlike).UrlNotilike(urlNotilike).UrlDesc(urlDesc).UrlAsc(urlAsc).NameEq(nameEq).NameNe(nameNe).NameGt(nameGt).NameGte(nameGte).NameLt(nameLt).NameLte(nameLte).NameIn(nameIn).NameNotin(nameNotin).NameContains(nameContains).NameNotcontains(nameNotcontains).NameLike(nameLike).NameNotlike(nameNotlike).NameIlike(nameIlike).NameNotilike(nameNotilike).NameDesc(nameDesc).NameAsc(nameAsc).HandledAtEq(handledAtEq).HandledAtNe(handledAtNe).HandledAtGt(handledAtGt).HandledAtGte(handledAtGte).HandledAtLt(handledAtLt).HandledAtLte(handledAtLte).HandledAtIn(handledAtIn).HandledAtNotin(handledAtNotin).HandledAtContains(handledAtContains).HandledAtNotcontains(handledAtNotcontains).HandledAtLike(handledAtLike).HandledAtNotlike(handledAtNotlike).HandledAtIlike(handledAtIlike).HandledAtNotilike(handledAtNotilike).HandledAtDesc(handledAtDesc).HandledAtAsc(handledAtAsc).RepositorySyncerClaimedUntilEq(repositorySyncerClaimedUntilEq).RepositorySyncerClaimedUntilNe(repositorySyncerClaimedUntilNe).RepositorySyncerClaimedUntilGt(repositorySyncerClaimedUntilGt).RepositorySyncerClaimedUntilGte(repositorySyncerClaimedUntilGte).RepositorySyncerClaimedUntilLt(repositorySyncerClaimedUntilLt).RepositorySyncerClaimedUntilLte(repositorySyncerClaimedUntilLte).RepositorySyncerClaimedUntilIn(repositorySyncerClaimedUntilIn).RepositorySyncerClaimedUntilNotin(repositorySyncerClaimedUntilNotin).RepositorySyncerClaimedUntilContains(repositorySyncerClaimedUntilContains).RepositorySyncerClaimedUntilNotcontains(repositorySyncerClaimedUntilNotcontains).RepositorySyncerClaimedUntilLike(repositorySyncerClaimedUntilLike).RepositorySyncerClaimedUntilNotlike(repositorySyncerClaimedUntilNotlike).RepositorySyncerClaimedUntilIlike(repositorySyncerClaimedUntilIlike).RepositorySyncerClaimedUntilNotilike(repositorySyncerClaimedUntilNotilike).RepositorySyncerClaimedUntilDesc(repositorySyncerClaimedUntilDesc).RepositorySyncerClaimedUntilAsc(repositorySyncerClaimedUntilAsc).ReferencedByChangeRepositoryIdObjectsContains(referencedByChangeRepositoryIdObjectsContains).ReferencedByChangeRepositoryIdObjectsNotcontains(referencedByChangeRepositoryIdObjectsNotcontains).ReferencedByChangeRepositoryIdObjectsDesc(referencedByChangeRepositoryIdObjectsDesc).ReferencedByChangeRepositoryIdObjectsAsc(referencedByChangeRepositoryIdObjectsAsc).ReferencedByJobRepositoryIdObjectsContains(referencedByJobRepositoryIdObjectsContains).ReferencedByJobRepositoryIdObjectsNotcontains(referencedByJobRepositoryIdObjectsNotcontains).ReferencedByJobRepositoryIdObjectsDesc(referencedByJobRepositoryIdObjectsDesc).ReferencedByJobRepositoryIdObjectsAsc(referencedByJobRepositoryIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | SQL LIMIT operator | 
 **offset** | **int32** | SQL OFFSET operator | 
 **depth** | **int32** | Max recursion depth for loading foreign objects; default &#x3D; 1  (0 &#x3D; recurse until graph cycle detected, 1 &#x3D; this object only, 2 &#x3D; this object + neighbours, 3 &#x3D; this object + neighbours + their neighbours... etc) | 
 **referencedByChangeLoad** | **string** | load the given indirectly related objects, value is ignored (presence of key is sufficient) | 
 **referencedByJobLoad** | **string** | load the given indirectly related objects, value is ignored (presence of key is sufficient) | 
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
 **urlEq** | **string** | SQL &#x3D; comparison | 
 **urlNe** | **string** | SQL !&#x3D; comparison | 
 **urlGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **urlGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **urlLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **urlLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **urlIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **urlNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **urlContains** | **string** | SQL @&gt; comparison | 
 **urlNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **urlLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **urlAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **nameEq** | **string** | SQL &#x3D; comparison | 
 **nameNe** | **string** | SQL !&#x3D; comparison | 
 **nameGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **nameGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **nameLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **nameLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **nameIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **nameNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **nameContains** | **string** | SQL @&gt; comparison | 
 **nameNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **nameLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **nameAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **handledAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **handledAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **handledAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **handledAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **handledAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **handledAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **handledAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **handledAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **handledAtContains** | **time.Time** | SQL @&gt; comparison | 
 **handledAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
 **handledAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **handledAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **handledAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **handledAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **handledAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **handledAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **repositorySyncerClaimedUntilEq** | **time.Time** | SQL &#x3D; comparison | 
 **repositorySyncerClaimedUntilNe** | **time.Time** | SQL !&#x3D; comparison | 
 **repositorySyncerClaimedUntilGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **repositorySyncerClaimedUntilGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **repositorySyncerClaimedUntilLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **repositorySyncerClaimedUntilLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **repositorySyncerClaimedUntilIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **repositorySyncerClaimedUntilNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **repositorySyncerClaimedUntilContains** | **time.Time** | SQL @&gt; comparison | 
 **repositorySyncerClaimedUntilNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
 **repositorySyncerClaimedUntilLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositorySyncerClaimedUntilNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositorySyncerClaimedUntilIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositorySyncerClaimedUntilNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositorySyncerClaimedUntilDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **repositorySyncerClaimedUntilAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByChangeRepositoryIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByChangeRepositoryIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **referencedByChangeRepositoryIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByChangeRepositoryIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobRepositoryIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByJobRepositoryIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **referencedByJobRepositoryIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobRepositoryIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> ResponseWithGenericOfRepository GetRepository(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.RepositoryAPI.GetRepository(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRepository

> ResponseWithGenericOfRepository PatchRepository(ctx, primaryKey).Repository(repository).Depth(depth).Execute()



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
	repository := *openapiclient.NewRepository() // Repository | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.PatchRepository(context.Background(), primaryKey).Repository(repository).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.PatchRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRepository`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.PatchRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | [**Repository**](Repository.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRepositories

> ResponseWithGenericOfRepository PostRepositories(ctx).Repository(repository).Depth(depth).Execute()



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
	repository := []openapiclient.Repository{*openapiclient.NewRepository()} // []Repository | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.PostRepositories(context.Background()).Repository(repository).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.PostRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRepositories`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.PostRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | [**[]Repository**](Repository.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRepositoriesRepositorySyncerClaim

> ResponseWithGenericOfRepository PostRepositoriesRepositorySyncerClaim(ctx, primaryKey).RepositoryRepositorySyncerClaimRequest(repositoryRepositorySyncerClaimRequest).Depth(depth).Execute()



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
	repositoryRepositorySyncerClaimRequest := *openapiclient.NewRepositoryRepositorySyncerClaimRequest() // RepositoryRepositorySyncerClaimRequest | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.PostRepositoriesRepositorySyncerClaim(context.Background(), primaryKey).RepositoryRepositorySyncerClaimRequest(repositoryRepositorySyncerClaimRequest).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.PostRepositoriesRepositorySyncerClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRepositoriesRepositorySyncerClaim`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.PostRepositoriesRepositorySyncerClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRepositoriesRepositorySyncerClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryRepositorySyncerClaimRequest** | [**RepositoryRepositorySyncerClaimRequest**](RepositoryRepositorySyncerClaimRequest.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

