# SchemaTxHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Transaction ID | [optional] 
**PrevAlh** | Pointer to **string** | State value (Accumulative Hash - Alh) of the previous transaction | [optional] 
**Ts** | Pointer to **string** | Unix timestamp of the transaction (in seconds) | [optional] 
**Nentries** | Pointer to **int32** | Number of entries in a transaction | [optional] 
**EH** | Pointer to **string** | Entries Hash - cumulative hash of all entries in the transaction | [optional] 
**BlTxId** | Pointer to **string** | Binary linking tree transaction ID (ID of last transaction already in the main Merkle Tree) | [optional] 
**BlRoot** | Pointer to **string** | Binary linking tree root (Root hash of the Merkle Tree) | [optional] 
**Version** | Pointer to **int32** | Header version | [optional] 
**Metadata** | Pointer to [**SchemaTxMetadata**](SchemaTxMetadata.md) |  | [optional] 

## Methods

### NewSchemaTxHeader

`func NewSchemaTxHeader() *SchemaTxHeader`

NewSchemaTxHeader instantiates a new SchemaTxHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTxHeaderWithDefaults

`func NewSchemaTxHeaderWithDefaults() *SchemaTxHeader`

NewSchemaTxHeaderWithDefaults instantiates a new SchemaTxHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemaTxHeader) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaTxHeader) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaTxHeader) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaTxHeader) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrevAlh

`func (o *SchemaTxHeader) GetPrevAlh() string`

GetPrevAlh returns the PrevAlh field if non-nil, zero value otherwise.

### GetPrevAlhOk

`func (o *SchemaTxHeader) GetPrevAlhOk() (*string, bool)`

GetPrevAlhOk returns a tuple with the PrevAlh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevAlh

`func (o *SchemaTxHeader) SetPrevAlh(v string)`

SetPrevAlh sets PrevAlh field to given value.

### HasPrevAlh

`func (o *SchemaTxHeader) HasPrevAlh() bool`

HasPrevAlh returns a boolean if a field has been set.

### GetTs

`func (o *SchemaTxHeader) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *SchemaTxHeader) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *SchemaTxHeader) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *SchemaTxHeader) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetNentries

`func (o *SchemaTxHeader) GetNentries() int32`

GetNentries returns the Nentries field if non-nil, zero value otherwise.

### GetNentriesOk

`func (o *SchemaTxHeader) GetNentriesOk() (*int32, bool)`

GetNentriesOk returns a tuple with the Nentries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNentries

`func (o *SchemaTxHeader) SetNentries(v int32)`

SetNentries sets Nentries field to given value.

### HasNentries

`func (o *SchemaTxHeader) HasNentries() bool`

HasNentries returns a boolean if a field has been set.

### GetEH

`func (o *SchemaTxHeader) GetEH() string`

GetEH returns the EH field if non-nil, zero value otherwise.

### GetEHOk

`func (o *SchemaTxHeader) GetEHOk() (*string, bool)`

GetEHOk returns a tuple with the EH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEH

`func (o *SchemaTxHeader) SetEH(v string)`

SetEH sets EH field to given value.

### HasEH

`func (o *SchemaTxHeader) HasEH() bool`

HasEH returns a boolean if a field has been set.

### GetBlTxId

`func (o *SchemaTxHeader) GetBlTxId() string`

GetBlTxId returns the BlTxId field if non-nil, zero value otherwise.

### GetBlTxIdOk

`func (o *SchemaTxHeader) GetBlTxIdOk() (*string, bool)`

GetBlTxIdOk returns a tuple with the BlTxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlTxId

`func (o *SchemaTxHeader) SetBlTxId(v string)`

SetBlTxId sets BlTxId field to given value.

### HasBlTxId

`func (o *SchemaTxHeader) HasBlTxId() bool`

HasBlTxId returns a boolean if a field has been set.

### GetBlRoot

`func (o *SchemaTxHeader) GetBlRoot() string`

GetBlRoot returns the BlRoot field if non-nil, zero value otherwise.

### GetBlRootOk

`func (o *SchemaTxHeader) GetBlRootOk() (*string, bool)`

GetBlRootOk returns a tuple with the BlRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlRoot

`func (o *SchemaTxHeader) SetBlRoot(v string)`

SetBlRoot sets BlRoot field to given value.

### HasBlRoot

`func (o *SchemaTxHeader) HasBlRoot() bool`

HasBlRoot returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaTxHeader) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaTxHeader) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaTxHeader) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaTxHeader) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *SchemaTxHeader) GetMetadata() SchemaTxMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchemaTxHeader) GetMetadataOk() (*SchemaTxMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchemaTxHeader) SetMetadata(v SchemaTxMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchemaTxHeader) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


