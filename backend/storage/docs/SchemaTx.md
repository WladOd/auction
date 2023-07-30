# SchemaTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to [**SchemaTxHeader**](SchemaTxHeader.md) |  | [optional] 
**Entries** | Pointer to [**[]SchemaTxEntry**](SchemaTxEntry.md) | Raw entry values | [optional] 
**KvEntries** | Pointer to [**[]SchemaEntry**](SchemaEntry.md) | KV entries in the transaction (parsed) | [optional] 
**ZEntries** | Pointer to [**[]SchemaZEntry**](SchemaZEntry.md) | Sorted Set entries in the transaction (parsed) | [optional] 

## Methods

### NewSchemaTx

`func NewSchemaTx() *SchemaTx`

NewSchemaTx instantiates a new SchemaTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTxWithDefaults

`func NewSchemaTxWithDefaults() *SchemaTx`

NewSchemaTxWithDefaults instantiates a new SchemaTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *SchemaTx) GetHeader() SchemaTxHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SchemaTx) GetHeaderOk() (*SchemaTxHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SchemaTx) SetHeader(v SchemaTxHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SchemaTx) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetEntries

`func (o *SchemaTx) GetEntries() []SchemaTxEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *SchemaTx) GetEntriesOk() (*[]SchemaTxEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *SchemaTx) SetEntries(v []SchemaTxEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *SchemaTx) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetKvEntries

`func (o *SchemaTx) GetKvEntries() []SchemaEntry`

GetKvEntries returns the KvEntries field if non-nil, zero value otherwise.

### GetKvEntriesOk

`func (o *SchemaTx) GetKvEntriesOk() (*[]SchemaEntry, bool)`

GetKvEntriesOk returns a tuple with the KvEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvEntries

`func (o *SchemaTx) SetKvEntries(v []SchemaEntry)`

SetKvEntries sets KvEntries field to given value.

### HasKvEntries

`func (o *SchemaTx) HasKvEntries() bool`

HasKvEntries returns a boolean if a field has been set.

### GetZEntries

`func (o *SchemaTx) GetZEntries() []SchemaZEntry`

GetZEntries returns the ZEntries field if non-nil, zero value otherwise.

### GetZEntriesOk

`func (o *SchemaTx) GetZEntriesOk() (*[]SchemaZEntry, bool)`

GetZEntriesOk returns a tuple with the ZEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZEntries

`func (o *SchemaTx) SetZEntries(v []SchemaZEntry)`

SetZEntries sets ZEntries field to given value.

### HasZEntries

`func (o *SchemaTx) HasZEntries() bool`

HasZEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


