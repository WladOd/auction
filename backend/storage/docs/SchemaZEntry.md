# SchemaZEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **string** | Name of the sorted set | [optional] 
**Key** | Pointer to **string** | Referenced key | [optional] 
**Entry** | Pointer to [**SchemaEntry**](SchemaEntry.md) |  | [optional] 
**Score** | Pointer to **float64** | Sorted set element&#39;s score | [optional] 
**AtTx** | Pointer to **string** | At which transaction the key is bound 0 if reference is not bound and should read the most recent reference | [optional] 

## Methods

### NewSchemaZEntry

`func NewSchemaZEntry() *SchemaZEntry`

NewSchemaZEntry instantiates a new SchemaZEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaZEntryWithDefaults

`func NewSchemaZEntryWithDefaults() *SchemaZEntry`

NewSchemaZEntryWithDefaults instantiates a new SchemaZEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *SchemaZEntry) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *SchemaZEntry) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *SchemaZEntry) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *SchemaZEntry) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetKey

`func (o *SchemaZEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SchemaZEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SchemaZEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SchemaZEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetEntry

`func (o *SchemaZEntry) GetEntry() SchemaEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *SchemaZEntry) GetEntryOk() (*SchemaEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *SchemaZEntry) SetEntry(v SchemaEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *SchemaZEntry) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetScore

`func (o *SchemaZEntry) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SchemaZEntry) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SchemaZEntry) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *SchemaZEntry) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetAtTx

`func (o *SchemaZEntry) GetAtTx() string`

GetAtTx returns the AtTx field if non-nil, zero value otherwise.

### GetAtTxOk

`func (o *SchemaZEntry) GetAtTxOk() (*string, bool)`

GetAtTxOk returns a tuple with the AtTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtTx

`func (o *SchemaZEntry) SetAtTx(v string)`

SetAtTx sets AtTx field to given value.

### HasAtTx

`func (o *SchemaZEntry) HasAtTx() bool`

HasAtTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


