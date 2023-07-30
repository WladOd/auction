# \IndexesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndex**](IndexesAPI.md#CreateIndex) | **Post** /ledger/{ledger}/collection/{collection}/indexes | Create a new index in the collection
[**DeleteIndex**](IndexesAPI.md#DeleteIndex) | **Put** /ledger/{ledger}/collection/{collection}/indexes | Delete index in the collection



## CreateIndex

> map[string]interface{} CreateIndex(ctx, ledger, collection).IndexCreateRequest(indexCreateRequest).Execute()

Create a new index in the collection



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
    indexCreateRequest := *openapiclient.NewIndexCreateRequest([]string{"Fields_example"}, false) // IndexCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexesAPI.CreateIndex(context.Background(), ledger, collection).IndexCreateRequest(indexCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexesAPI.CreateIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndex`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IndexesAPI.CreateIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **indexCreateRequest** | [**IndexCreateRequest**](IndexCreateRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndex

> map[string]interface{} DeleteIndex(ctx, ledger, collection).IndexDeleteRequest(indexDeleteRequest).Execute()

Delete index in the collection



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
    indexDeleteRequest := *openapiclient.NewIndexDeleteRequest("Collection_example", []string{"Fields_example"}) // IndexDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexesAPI.DeleteIndex(context.Background(), ledger, collection).IndexDeleteRequest(indexDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexesAPI.DeleteIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndex`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IndexesAPI.DeleteIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **indexDeleteRequest** | [**IndexDeleteRequest**](IndexDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

