# \DocumentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditDocument**](DocumentsAPI.md#AuditDocument) | **Post** /ledger/{ledger}/collection/{collection}/document/{document-id}/audit | Search for audit items for provided document
[**CountDocuments**](DocumentsAPI.md#CountDocuments) | **Post** /ledger/{ledger}/collection/{collection}/documents/count | Count documents inside collection
[**DiffDocument**](DocumentsAPI.md#DiffDocument) | **Put** /ledger/{ledger}/collection/{collection}/document/{document-id}/audit | Return diff for document revisions
[**DocumentCreate**](DocumentsAPI.md#DocumentCreate) | **Put** /ledger/{ledger}/collection/{collection}/document | Create new document inside collection
[**DocumentCreateMany**](DocumentsAPI.md#DocumentCreateMany) | **Put** /ledger/{ledger}/collection/{collection}/documents | Create multiple documents inside collection
[**GetDocumentProof**](DocumentsAPI.md#GetDocumentProof) | **Post** /ledger/{ledger}/collection/{collection}/document/{document-id}/proof | Return a proof for a document
[**SearchDocument**](DocumentsAPI.md#SearchDocument) | **Post** /ledger/{ledger}/collection/{collection}/documents/search | Search for a document inside collection
[**UpdateDocument**](DocumentsAPI.md#UpdateDocument) | **Post** /ledger/{ledger}/collection/{collection}/document | Replace whole document with new provided



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
    resp, r, err := apiClient.DocumentsAPI.AuditDocument(context.Background(), ledger, collection, documentId).DocumentAuditRequest(documentAuditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.AuditDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditDocument`: DocumentAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.AuditDocument`: %v\n", resp)
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


## CountDocuments

> DocumentsCountResponse CountDocuments(ctx, ledger, collection).DocumentCountRequest(documentCountRequest).Execute()

Count documents inside collection



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
    documentCountRequest := *openapiclient.NewDocumentCountRequest() // DocumentCountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.CountDocuments(context.Background(), ledger, collection).DocumentCountRequest(documentCountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CountDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountDocuments`: DocumentsCountResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CountDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **documentCountRequest** | [**DocumentCountRequest**](DocumentCountRequest.md) |  | 

### Return type

[**DocumentsCountResponse**](DocumentsCountResponse.md)

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
    resp, r, err := apiClient.DocumentsAPI.DiffDocument(context.Background(), ledger, collection, documentId).DocumentDiffRequest(documentDiffRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DiffDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiffDocument`: DocumentDiffResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DiffDocument`: %v\n", resp)
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


## DocumentCreate

> DocumentInsertResponse DocumentCreate(ctx, ledger, collection).Body(body).Execute()

Create new document inside collection



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
    body := interface{}(987) // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.DocumentCreate(context.Background(), ledger, collection).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocumentCreate`: DocumentInsertResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **interface{}** |  | 

### Return type

[**DocumentInsertResponse**](DocumentInsertResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentCreateMany

> DocumentInsertManyResponse DocumentCreateMany(ctx, ledger, collection).DocumentInsertManyRequest(documentInsertManyRequest).Execute()

Create multiple documents inside collection



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
    documentInsertManyRequest := *openapiclient.NewDocumentInsertManyRequest([]map[string]interface{}{map[string]interface{}(123)}) // DocumentInsertManyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.DocumentCreateMany(context.Background(), ledger, collection).DocumentInsertManyRequest(documentInsertManyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentCreateMany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocumentCreateMany`: DocumentInsertManyResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentCreateMany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumentCreateManyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **documentInsertManyRequest** | [**DocumentInsertManyRequest**](DocumentInsertManyRequest.md) |  | 

### Return type

[**DocumentInsertManyResponse**](DocumentInsertManyResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
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
    resp, r, err := apiClient.DocumentsAPI.GetDocumentProof(context.Background(), ledger, collection, documentId).DocumentProofRequest(documentProofRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocumentProof``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentProof`: DocumentProofResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocumentProof`: %v\n", resp)
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


## SearchDocument

> DocumentSearchResponse SearchDocument(ctx, ledger, collection).DocumentSearchRequest(documentSearchRequest).Execute()

Search for a document inside collection



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
    documentSearchRequest := *openapiclient.NewDocumentSearchRequest(int32(123), int32(123)) // DocumentSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.SearchDocument(context.Background(), ledger, collection).DocumentSearchRequest(documentSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.SearchDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDocument`: DocumentSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.SearchDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **documentSearchRequest** | [**DocumentSearchRequest**](DocumentSearchRequest.md) |  | 

### Return type

[**DocumentSearchResponse**](DocumentSearchResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> DocumentUpdateResponse UpdateDocument(ctx, ledger, collection).DocumentUpdateRequest(documentUpdateRequest).Execute()

Replace whole document with new provided



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
    documentUpdateRequest := *openapiclient.NewDocumentUpdateRequest(map[string]interface{}(123), *openapiclient.NewQuery()) // DocumentUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.UpdateDocument(context.Background(), ledger, collection).DocumentUpdateRequest(documentUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.UpdateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocument`: DocumentUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 
**collection** | **string** | Explicit collection name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **documentUpdateRequest** | [**DocumentUpdateRequest**](DocumentUpdateRequest.md) |  | 

### Return type

[**DocumentUpdateResponse**](DocumentUpdateResponse.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

