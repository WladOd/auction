# SchemaImmutableState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | Pointer to **string** | The db name | [optional] 
**TxId** | Pointer to **string** | Id of the most recent transaction | [optional] 
**TxHash** | Pointer to **string** | State of the most recent transaction | [optional] 
**Signature** | Pointer to [**SchemaSignature**](SchemaSignature.md) |  | [optional] 
**PrecommittedTxId** | Pointer to **string** | Id of the most recent precommitted transaction | [optional] 
**PrecommittedTxHash** | Pointer to **string** | State of the most recent precommitted transaction | [optional] 

## Methods

### NewSchemaImmutableState

`func NewSchemaImmutableState() *SchemaImmutableState`

NewSchemaImmutableState instantiates a new SchemaImmutableState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaImmutableStateWithDefaults

`func NewSchemaImmutableStateWithDefaults() *SchemaImmutableState`

NewSchemaImmutableStateWithDefaults instantiates a new SchemaImmutableState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *SchemaImmutableState) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *SchemaImmutableState) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *SchemaImmutableState) SetDb(v string)`

SetDb sets Db field to given value.

### HasDb

`func (o *SchemaImmutableState) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetTxId

`func (o *SchemaImmutableState) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *SchemaImmutableState) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *SchemaImmutableState) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *SchemaImmutableState) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetTxHash

`func (o *SchemaImmutableState) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *SchemaImmutableState) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *SchemaImmutableState) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *SchemaImmutableState) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetSignature

`func (o *SchemaImmutableState) GetSignature() SchemaSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SchemaImmutableState) GetSignatureOk() (*SchemaSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SchemaImmutableState) SetSignature(v SchemaSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SchemaImmutableState) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetPrecommittedTxId

`func (o *SchemaImmutableState) GetPrecommittedTxId() string`

GetPrecommittedTxId returns the PrecommittedTxId field if non-nil, zero value otherwise.

### GetPrecommittedTxIdOk

`func (o *SchemaImmutableState) GetPrecommittedTxIdOk() (*string, bool)`

GetPrecommittedTxIdOk returns a tuple with the PrecommittedTxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecommittedTxId

`func (o *SchemaImmutableState) SetPrecommittedTxId(v string)`

SetPrecommittedTxId sets PrecommittedTxId field to given value.

### HasPrecommittedTxId

`func (o *SchemaImmutableState) HasPrecommittedTxId() bool`

HasPrecommittedTxId returns a boolean if a field has been set.

### GetPrecommittedTxHash

`func (o *SchemaImmutableState) GetPrecommittedTxHash() string`

GetPrecommittedTxHash returns the PrecommittedTxHash field if non-nil, zero value otherwise.

### GetPrecommittedTxHashOk

`func (o *SchemaImmutableState) GetPrecommittedTxHashOk() (*string, bool)`

GetPrecommittedTxHashOk returns a tuple with the PrecommittedTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecommittedTxHash

`func (o *SchemaImmutableState) SetPrecommittedTxHash(v string)`

SetPrecommittedTxHash sets PrecommittedTxHash field to given value.

### HasPrecommittedTxHash

`func (o *SchemaImmutableState) HasPrecommittedTxHash() bool`

HasPrecommittedTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


