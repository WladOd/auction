# DocumentInsertManyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** |  | [optional] 
**DocumentIds** | **[]string** |  | 

## Methods

### NewDocumentInsertManyResponse

`func NewDocumentInsertManyResponse(documentIds []string, ) *DocumentInsertManyResponse`

NewDocumentInsertManyResponse instantiates a new DocumentInsertManyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentInsertManyResponseWithDefaults

`func NewDocumentInsertManyResponseWithDefaults() *DocumentInsertManyResponse`

NewDocumentInsertManyResponseWithDefaults instantiates a new DocumentInsertManyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DocumentInsertManyResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DocumentInsertManyResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DocumentInsertManyResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DocumentInsertManyResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetDocumentIds

`func (o *DocumentInsertManyResponse) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *DocumentInsertManyResponse) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *DocumentInsertManyResponse) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


