# DocumentUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** |  | 
**DocumentId** | **string** |  | 
**Revision** | **string** |  | 

## Methods

### NewDocumentUpdateResponse

`func NewDocumentUpdateResponse(transactionId string, documentId string, revision string, ) *DocumentUpdateResponse`

NewDocumentUpdateResponse instantiates a new DocumentUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentUpdateResponseWithDefaults

`func NewDocumentUpdateResponseWithDefaults() *DocumentUpdateResponse`

NewDocumentUpdateResponseWithDefaults instantiates a new DocumentUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DocumentUpdateResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DocumentUpdateResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DocumentUpdateResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetDocumentId

`func (o *DocumentUpdateResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DocumentUpdateResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DocumentUpdateResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetRevision

`func (o *DocumentUpdateResponse) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DocumentUpdateResponse) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DocumentUpdateResponse) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


