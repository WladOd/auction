# SchemaTxEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Raw key value (contains 1-byte prefix for kind of the key) | [optional] 
**HValue** | Pointer to **string** | Value hash | [optional] 
**VLen** | Pointer to **int32** | Value length | [optional] 
**Metadata** | Pointer to [**SchemaKVMetadata**](SchemaKVMetadata.md) |  | [optional] 
**Value** | Pointer to **string** | Value, must be ignored when len(value) &#x3D;&#x3D; 0 and vLen &gt; 0, otherwise sha256(value) must be equal to hValue | [optional] 

## Methods

### NewSchemaTxEntry

`func NewSchemaTxEntry() *SchemaTxEntry`

NewSchemaTxEntry instantiates a new SchemaTxEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTxEntryWithDefaults

`func NewSchemaTxEntryWithDefaults() *SchemaTxEntry`

NewSchemaTxEntryWithDefaults instantiates a new SchemaTxEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SchemaTxEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SchemaTxEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SchemaTxEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SchemaTxEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetHValue

`func (o *SchemaTxEntry) GetHValue() string`

GetHValue returns the HValue field if non-nil, zero value otherwise.

### GetHValueOk

`func (o *SchemaTxEntry) GetHValueOk() (*string, bool)`

GetHValueOk returns a tuple with the HValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHValue

`func (o *SchemaTxEntry) SetHValue(v string)`

SetHValue sets HValue field to given value.

### HasHValue

`func (o *SchemaTxEntry) HasHValue() bool`

HasHValue returns a boolean if a field has been set.

### GetVLen

`func (o *SchemaTxEntry) GetVLen() int32`

GetVLen returns the VLen field if non-nil, zero value otherwise.

### GetVLenOk

`func (o *SchemaTxEntry) GetVLenOk() (*int32, bool)`

GetVLenOk returns a tuple with the VLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVLen

`func (o *SchemaTxEntry) SetVLen(v int32)`

SetVLen sets VLen field to given value.

### HasVLen

`func (o *SchemaTxEntry) HasVLen() bool`

HasVLen returns a boolean if a field has been set.

### GetMetadata

`func (o *SchemaTxEntry) GetMetadata() SchemaKVMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchemaTxEntry) GetMetadataOk() (*SchemaKVMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchemaTxEntry) SetMetadata(v SchemaKVMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchemaTxEntry) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetValue

`func (o *SchemaTxEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SchemaTxEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SchemaTxEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SchemaTxEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


