# DocumentInsertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** |  | [optional] 
**DocumentId** | **string** |  | 

## Methods

### NewDocumentInsertResponse

`func NewDocumentInsertResponse(documentId string, ) *DocumentInsertResponse`

NewDocumentInsertResponse instantiates a new DocumentInsertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentInsertResponseWithDefaults

`func NewDocumentInsertResponseWithDefaults() *DocumentInsertResponse`

NewDocumentInsertResponseWithDefaults instantiates a new DocumentInsertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DocumentInsertResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DocumentInsertResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DocumentInsertResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DocumentInsertResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetDocumentId

`func (o *DocumentInsertResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DocumentInsertResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DocumentInsertResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


