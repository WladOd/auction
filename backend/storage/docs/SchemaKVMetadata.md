# SchemaKVMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** | True if this entry denotes a logical deletion | [optional] 
**Expiration** | Pointer to [**SchemaExpiration**](SchemaExpiration.md) |  | [optional] 
**NonIndexable** | Pointer to **bool** | If set to true, this entry will not be indexed and will only be accessed through GetAt calls | [optional] 

## Methods

### NewSchemaKVMetadata

`func NewSchemaKVMetadata() *SchemaKVMetadata`

NewSchemaKVMetadata instantiates a new SchemaKVMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaKVMetadataWithDefaults

`func NewSchemaKVMetadataWithDefaults() *SchemaKVMetadata`

NewSchemaKVMetadataWithDefaults instantiates a new SchemaKVMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *SchemaKVMetadata) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SchemaKVMetadata) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SchemaKVMetadata) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *SchemaKVMetadata) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetExpiration

`func (o *SchemaKVMetadata) GetExpiration() SchemaExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *SchemaKVMetadata) GetExpirationOk() (*SchemaExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *SchemaKVMetadata) SetExpiration(v SchemaExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *SchemaKVMetadata) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetNonIndexable

`func (o *SchemaKVMetadata) GetNonIndexable() bool`

GetNonIndexable returns the NonIndexable field if non-nil, zero value otherwise.

### GetNonIndexableOk

`func (o *SchemaKVMetadata) GetNonIndexableOk() (*bool, bool)`

GetNonIndexableOk returns a tuple with the NonIndexable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonIndexable

`func (o *SchemaKVMetadata) SetNonIndexable(v bool)`

SetNonIndexable sets NonIndexable field to given value.

### HasNonIndexable

`func (o *SchemaKVMetadata) HasNonIndexable() bool`

HasNonIndexable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


