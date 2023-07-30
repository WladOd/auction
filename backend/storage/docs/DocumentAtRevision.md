# DocumentAtRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** |  | 
**Revision** | **string** |  | 
**Document** | **map[string]interface{}** |  | 

## Methods

### NewDocumentAtRevision

`func NewDocumentAtRevision(transactionId string, revision string, document map[string]interface{}, ) *DocumentAtRevision`

NewDocumentAtRevision instantiates a new DocumentAtRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentAtRevisionWithDefaults

`func NewDocumentAtRevisionWithDefaults() *DocumentAtRevision`

NewDocumentAtRevisionWithDefaults instantiates a new DocumentAtRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *DocumentAtRevision) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DocumentAtRevision) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DocumentAtRevision) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetRevision

`func (o *DocumentAtRevision) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DocumentAtRevision) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DocumentAtRevision) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetDocument

`func (o *DocumentAtRevision) GetDocument() map[string]interface{}`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *DocumentAtRevision) GetDocumentOk() (*map[string]interface{}, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *DocumentAtRevision) SetDocument(v map[string]interface{})`

SetDocument sets Document field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


