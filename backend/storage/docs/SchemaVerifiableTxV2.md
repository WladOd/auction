# SchemaVerifiableTxV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | Pointer to [**SchemaTx**](SchemaTx.md) |  | [optional] 
**DualProof** | Pointer to [**SchemaDualProofV2**](SchemaDualProofV2.md) |  | [optional] 
**Signature** | Pointer to [**SchemaSignature**](SchemaSignature.md) |  | [optional] 

## Methods

### NewSchemaVerifiableTxV2

`func NewSchemaVerifiableTxV2() *SchemaVerifiableTxV2`

NewSchemaVerifiableTxV2 instantiates a new SchemaVerifiableTxV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaVerifiableTxV2WithDefaults

`func NewSchemaVerifiableTxV2WithDefaults() *SchemaVerifiableTxV2`

NewSchemaVerifiableTxV2WithDefaults instantiates a new SchemaVerifiableTxV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *SchemaVerifiableTxV2) GetTx() SchemaTx`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *SchemaVerifiableTxV2) GetTxOk() (*SchemaTx, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *SchemaVerifiableTxV2) SetTx(v SchemaTx)`

SetTx sets Tx field to given value.

### HasTx

`func (o *SchemaVerifiableTxV2) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetDualProof

`func (o *SchemaVerifiableTxV2) GetDualProof() SchemaDualProofV2`

GetDualProof returns the DualProof field if non-nil, zero value otherwise.

### GetDualProofOk

`func (o *SchemaVerifiableTxV2) GetDualProofOk() (*SchemaDualProofV2, bool)`

GetDualProofOk returns a tuple with the DualProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualProof

`func (o *SchemaVerifiableTxV2) SetDualProof(v SchemaDualProofV2)`

SetDualProof sets DualProof field to given value.

### HasDualProof

`func (o *SchemaVerifiableTxV2) HasDualProof() bool`

HasDualProof returns a boolean if a field has been set.

### GetSignature

`func (o *SchemaVerifiableTxV2) GetSignature() SchemaSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SchemaVerifiableTxV2) GetSignatureOk() (*SchemaSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SchemaVerifiableTxV2) SetSignature(v SchemaSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SchemaVerifiableTxV2) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


