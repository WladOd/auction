# DocumentProofRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **int32** |  | 
**ProofSinceTransactionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewDocumentProofRequest

`func NewDocumentProofRequest(transactionId int32, ) *DocumentProofRequest`

NewDocumentProofRequest instantiates a new DocumentProofRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentProofRequestWithDefaults

`func NewDocumentProofRequestWithDefaults() *DocumentProofRequest`

NewDocumentProofRequestWithDefaults instantiates a new DocumentProofRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DocumentProofRequest) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DocumentProofRequest) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DocumentProofRequest) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.


### GetProofSinceTransactionId

`func (o *DocumentProofRequest) GetProofSinceTransactionId() int32`

GetProofSinceTransactionId returns the ProofSinceTransactionId field if non-nil, zero value otherwise.

### GetProofSinceTransactionIdOk

`func (o *DocumentProofRequest) GetProofSinceTransactionIdOk() (*int32, bool)`

GetProofSinceTransactionIdOk returns a tuple with the ProofSinceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofSinceTransactionId

`func (o *DocumentProofRequest) SetProofSinceTransactionId(v int32)`

SetProofSinceTransactionId sets ProofSinceTransactionId field to given value.

### HasProofSinceTransactionId

`func (o *DocumentProofRequest) HasProofSinceTransactionId() bool`

HasProofSinceTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


