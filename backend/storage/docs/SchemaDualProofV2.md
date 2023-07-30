# SchemaDualProofV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceTxHeader** | Pointer to [**SchemaTxHeader**](SchemaTxHeader.md) |  | [optional] 
**TargetTxHeader** | Pointer to [**SchemaTxHeader**](SchemaTxHeader.md) |  | [optional] 
**InclusionProof** | Pointer to **[]string** | Inclusion proof of the source transaction hash in the main Merkle Tree | [optional] 
**ConsistencyProof** | Pointer to **[]string** | Consistency proof between Merkle Trees in the source and target transactions | [optional] 

## Methods

### NewSchemaDualProofV2

`func NewSchemaDualProofV2() *SchemaDualProofV2`

NewSchemaDualProofV2 instantiates a new SchemaDualProofV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDualProofV2WithDefaults

`func NewSchemaDualProofV2WithDefaults() *SchemaDualProofV2`

NewSchemaDualProofV2WithDefaults instantiates a new SchemaDualProofV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceTxHeader

`func (o *SchemaDualProofV2) GetSourceTxHeader() SchemaTxHeader`

GetSourceTxHeader returns the SourceTxHeader field if non-nil, zero value otherwise.

### GetSourceTxHeaderOk

`func (o *SchemaDualProofV2) GetSourceTxHeaderOk() (*SchemaTxHeader, bool)`

GetSourceTxHeaderOk returns a tuple with the SourceTxHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTxHeader

`func (o *SchemaDualProofV2) SetSourceTxHeader(v SchemaTxHeader)`

SetSourceTxHeader sets SourceTxHeader field to given value.

### HasSourceTxHeader

`func (o *SchemaDualProofV2) HasSourceTxHeader() bool`

HasSourceTxHeader returns a boolean if a field has been set.

### GetTargetTxHeader

`func (o *SchemaDualProofV2) GetTargetTxHeader() SchemaTxHeader`

GetTargetTxHeader returns the TargetTxHeader field if non-nil, zero value otherwise.

### GetTargetTxHeaderOk

`func (o *SchemaDualProofV2) GetTargetTxHeaderOk() (*SchemaTxHeader, bool)`

GetTargetTxHeaderOk returns a tuple with the TargetTxHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTxHeader

`func (o *SchemaDualProofV2) SetTargetTxHeader(v SchemaTxHeader)`

SetTargetTxHeader sets TargetTxHeader field to given value.

### HasTargetTxHeader

`func (o *SchemaDualProofV2) HasTargetTxHeader() bool`

HasTargetTxHeader returns a boolean if a field has been set.

### GetInclusionProof

`func (o *SchemaDualProofV2) GetInclusionProof() []string`

GetInclusionProof returns the InclusionProof field if non-nil, zero value otherwise.

### GetInclusionProofOk

`func (o *SchemaDualProofV2) GetInclusionProofOk() (*[]string, bool)`

GetInclusionProofOk returns a tuple with the InclusionProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusionProof

`func (o *SchemaDualProofV2) SetInclusionProof(v []string)`

SetInclusionProof sets InclusionProof field to given value.

### HasInclusionProof

`func (o *SchemaDualProofV2) HasInclusionProof() bool`

HasInclusionProof returns a boolean if a field has been set.

### GetConsistencyProof

`func (o *SchemaDualProofV2) GetConsistencyProof() []string`

GetConsistencyProof returns the ConsistencyProof field if non-nil, zero value otherwise.

### GetConsistencyProofOk

`func (o *SchemaDualProofV2) GetConsistencyProofOk() (*[]string, bool)`

GetConsistencyProofOk returns a tuple with the ConsistencyProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyProof

`func (o *SchemaDualProofV2) SetConsistencyProof(v []string)`

SetConsistencyProof sets ConsistencyProof field to given value.

### HasConsistencyProof

`func (o *SchemaDualProofV2) HasConsistencyProof() bool`

HasConsistencyProof returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


