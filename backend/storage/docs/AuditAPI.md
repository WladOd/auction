# \AuditAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditDocument**](AuditAPI.md#AuditDocument) | **Post** /ledger/{ledger}/collection/{collection}/document/{document-id}/audit | Search for audit items for provided document
[**DiffDocument**](AuditAPI.md#DiffDocument) | **Put** /ledger/{ledger}/collection/{collection}/document/{document-id}/audit | Return diff for document revisions
[**GetCurrentState**](AuditAPI.md#GetCurrentState) | **Get** /ledger/{ledger}/state | Return current state of immudb ledger
[**GetDocumentProof**](AuditAPI.md#GetDocumentProof) | **Post** /ledger/{ledger}/collection/{collection}/document/{document-id}/proof | Return a proof for a document
[**GetLedgerDbSize**](AuditAPI.md#GetLedgerDbSize) | **Get** /ledger/{ledger}/size | Return ledger DB size



## AuditDocument

> DocumentAuditResponse AuditDocument(ctx, ledger, collection, documentId).DocumentAuditRequest(documentAuditRequest).Execute()

Search for audit items for provided document



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
    documentId := "documentId_example" // string | Explicit document ID
    documentAuditRequest := *openapiclient.NewDocumentAuditRequest(int32(123), int32(123), false) // DocumentAuditRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditAPI.AuditDocument(context.Background(), ledger, collection, documentId).DocumentAuditRequest(documentAuditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.AuditDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditDocument`: DocumentAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `AuditAPI.AuditDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 
**documentId** | **string** | Explicit document ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **documentAuditRequest** | [**DocumentAuditRequest**](DocumentAuditRequest.md) |  | 

### Return type

[**DocumentAuditResponse**](DocumentAuditResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiffDocument

> DocumentDiffResponse DiffDocument(ctx, ledger, collection, documentId).DocumentDiffRequest(documentDiffRequest).Execute()

Return diff for document revisions



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
    documentId := "documentId_example" // string | Explicit document ID
    documentDiffRequest := *openapiclient.NewDocumentDiffRequest(int32(123), int32(123)) // DocumentDiffRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditAPI.DiffDocument(context.Background(), ledger, collection, documentId).DocumentDiffRequest(documentDiffRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.DiffDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiffDocument`: DocumentDiffResponse
    fmt.Fprintf(os.Stdout, "Response from `AuditAPI.DiffDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 
**documentId** | **string** | Explicit document ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiffDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **documentDiffRequest** | [**DocumentDiffRequest**](DocumentDiffRequest.md) |  | 

### Return type

[**DocumentDiffResponse**](DocumentDiffResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentState

> SchemaImmutableState GetCurrentState(ctx, ledger).Execute()

Return current state of immudb ledger



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
    resp, r, err := apiClient.AuditAPI.GetCurrentState(context.Background(), ledger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.GetCurrentState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentState`: SchemaImmutableState
    fmt.Fprintf(os.Stdout, "Response from `AuditAPI.GetCurrentState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SchemaImmutableState**](SchemaImmutableState.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentProof

> DocumentProofResponse GetDocumentProof(ctx, ledger, collection, documentId).DocumentProofRequest(documentProofRequest).Execute()

Return a proof for a document



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
    documentId := "documentId_example" // string | Explicit document ID
    documentProofRequest := *openapiclient.NewDocumentProofRequest(int32(123)) // DocumentProofRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditAPI.GetDocumentProof(context.Background(), ledger, collection, documentId).DocumentProofRequest(documentProofRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.GetDocumentProof``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentProof`: DocumentProofResponse
    fmt.Fprintf(os.Stdout, "Response from `AuditAPI.GetDocumentProof`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 
**documentId** | **string** | Explicit document ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **documentProofRequest** | [**DocumentProofRequest**](DocumentProofRequest.md) |  | 

### Return type

[**DocumentProofResponse**](DocumentProofResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLedgerDbSize

> LedgerDBSize GetLedgerDbSize(ctx, ledger).Execute()

Return ledger DB size



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
    resp, r, err := apiClient.AuditAPI.GetLedgerDbSize(context.Background(), ledger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.GetLedgerDbSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLedgerDbSize`: LedgerDBSize
    fmt.Fprintf(os.Stdout, "Response from `AuditAPI.GetLedgerDbSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerDbSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LedgerDBSize**](LedgerDBSize.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

