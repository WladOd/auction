# \CollectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionCreate**](CollectionsAPI.md#CollectionCreate) | **Put** /ledger/{ledger}/collection/{collection} | Create collection in the ledger
[**CollectionDelete**](CollectionsAPI.md#CollectionDelete) | **Delete** /ledger/{ledger}/collection/{collection} | Delete collection
[**CollectionGet**](CollectionsAPI.md#CollectionGet) | **Get** /ledger/{ledger}/collection/{collection} | Return information about collection
[**CollectionUpdate**](CollectionsAPI.md#CollectionUpdate) | **Post** /ledger/{ledger}/collection/{collection} | Update collection within the ledger
[**CollectionsList**](CollectionsAPI.md#CollectionsList) | **Get** /ledger/{ledger}/collections | List collection within a ledger



## CollectionCreate

> CollectionCreate(ctx, ledger, collection).CollectionCreateRequest(collectionCreateRequest).Execute()

Create collection in the ledger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/storage"
)

func main() {
    ledger := "ledger_example" // string | Explicit ledger name
    collection := "collection_example" // string | Explicit collection name
    collectionCreateRequest := *openapiclient.NewCollectionCreateRequest() // CollectionCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CollectionsAPI.CollectionCreate(context.Background(), ledger, collection).CollectionCreateRequest(collectionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionCreateRequest** | [**CollectionCreateRequest**](CollectionCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionDelete

> CollectionDelete(ctx, ledger, collection).Execute()

Delete collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/storage"
)

func main() {
    ledger := "ledger_example" // string | Explicit ledger name
    collection := "collection_example" // string | Collection name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CollectionsAPI.CollectionDelete(context.Background(), ledger, collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionGet

> Collection CollectionGet(ctx, ledger, collection).Execute()

Return information about collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/storage"
)

func main() {
    ledger := "ledger_example" // string | Explicit ledger name
    collection := "collection_example" // string | Collection name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsAPI.CollectionGet(context.Background(), ledger, collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionGet`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Collection**](Collection.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionUpdate

> CollectionUpdate(ctx, ledger, collection).CollectionUpdateRequest(collectionUpdateRequest).Execute()

Update collection within the ledger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/storage"
)

func main() {
    ledger := "ledger_example" // string | Explicit ledger name
    collection := "collection_example" // string | Explicit collection name
    collectionUpdateRequest := *openapiclient.NewCollectionUpdateRequest("IdFieldName_example") // CollectionUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CollectionsAPI.CollectionUpdate(context.Background(), ledger, collection).CollectionUpdateRequest(collectionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionUpdateRequest** | [**CollectionUpdateRequest**](CollectionUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsList

> CollectionListResponse CollectionsList(ctx, ledger).Execute()

List collection within a ledger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/storage"
)

func main() {
    ledger := "ledger_example" // string | Explicit ledger name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsAPI.CollectionsList(context.Background(), ledger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CollectionsList`: CollectionListResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionListResponse**](CollectionListResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

