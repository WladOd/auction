# \ExportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExportInformation**](ExportAPI.md#GetExportInformation) | **Get** /ledger/{ledger}/export/status | Return information about export
[**S3Export**](ExportAPI.md#S3Export) | **Post** /ledger/{ledger}/export/s3 | Start export request



## GetExportInformation

> ExportInformation GetExportInformation(ctx, ledger).Id(id).Execute()

Return information about export



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
    id := "id_example" // string | ID of export
    ledger := "ledger_example" // string | Explicit ledger name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPI.GetExportInformation(context.Background(), ledger).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.GetExportInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportInformation`: ExportInformation
    fmt.Fprintf(os.Stdout, "Response from `ExportAPI.GetExportInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | ID of export | 


### Return type

[**ExportInformation**](ExportInformation.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3Export

> ExportID S3Export(ctx, ledger).ExportS3(exportS3).Execute()

Start export request



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
    exportS3 := *openapiclient.NewExportS3("Region_example", "AccessKey_example", "SecretKey_example", "Token_example", "Bucket_example", "UploadKey_example") // ExportS3 | S3 request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPI.S3Export(context.Background(), ledger).ExportS3(exportS3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.S3Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3Export`: ExportID
    fmt.Fprintf(os.Stdout, "Response from `ExportAPI.S3Export`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledger** | **string** | Explicit ledger name | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3ExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportS3** | [**ExportS3**](ExportS3.md) | S3 request | 

### Return type

[**ExportID**](ExportID.md)

### Authorization

[PassetoAuth](../README.md#PassetoAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

