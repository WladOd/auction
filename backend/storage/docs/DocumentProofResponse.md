# DocumentProofResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | **string** |  | 
**CollectionId** | **int64** |  | 
**IdFieldName** | **string** |  | 
**EncodedDocument** | **string** |  | 
**VerifiableTx** | [**SchemaVerifiableTxV2**](SchemaVerifiableTxV2.md) |  | 

## Methods

### NewDocumentProofResponse

`func NewDocumentProofResponse(database string, collectionId int64, idFieldName string, encodedDocument string, verifiableTx SchemaVerifiableTxV2, ) *DocumentProofResponse`

NewDocumentProofResponse instantiates a new DocumentProofResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentProofResponseWithDefaults

`func NewDocumentProofResponseWithDefaults() *DocumentProofResponse`

NewDocumentProofResponseWithDefaults instantiates a new DocumentProofResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *DocumentProofResponse) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DocumentProofResponse) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DocumentProofResponse) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetCollectionId

`func (o *DocumentProofResponse) GetCollectionId() int64`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *DocumentProofResponse) GetCollectionIdOk() (*int64, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *DocumentProofResponse) SetCollectionId(v int64)`

SetCollectionId sets CollectionId field to given value.


### GetIdFieldName

`func (o *DocumentProofResponse) GetIdFieldName() string`

GetIdFieldName returns the IdFieldName field if non-nil, zero value otherwise.

### GetIdFieldNameOk

`func (o *DocumentProofResponse) GetIdFieldNameOk() (*string, bool)`

GetIdFieldNameOk returns a tuple with the IdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFieldName

`func (o *DocumentProofResponse) SetIdFieldName(v string)`

SetIdFieldName sets IdFieldName field to given value.


### GetEncodedDocument

`func (o *DocumentProofResponse) GetEncodedDocument() string`

GetEncodedDocument returns the EncodedDocument field if non-nil, zero value otherwise.

### GetEncodedDocumentOk

`func (o *DocumentProofResponse) GetEncodedDocumentOk() (*string, bool)`

GetEncodedDocumentOk returns a tuple with the EncodedDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedDocument

`func (o *DocumentProofResponse) SetEncodedDocument(v string)`

SetEncodedDocument sets EncodedDocument field to given value.


### GetVerifiableTx

`func (o *DocumentProofResponse) GetVerifiableTx() SchemaVerifiableTxV2`

GetVerifiableTx returns the VerifiableTx field if non-nil, zero value otherwise.

### GetVerifiableTxOk

`func (o *DocumentProofResponse) GetVerifiableTxOk() (*SchemaVerifiableTxV2, bool)`

GetVerifiableTxOk returns a tuple with the VerifiableTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiableTx

`func (o *DocumentProofResponse) SetVerifiableTx(v SchemaVerifiableTxV2)`

SetVerifiableTx sets VerifiableTx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


