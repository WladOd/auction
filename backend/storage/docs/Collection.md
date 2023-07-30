# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IdFieldName** | **string** |  | 
**Fields** | [**[]Field**](Field.md) |  | 
**Indexes** | [**[]Index**](Index.md) |  | 

## Methods

### NewCollection

`func NewCollection(name string, idFieldName string, fields []Field, indexes []Index, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Collection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collection) SetName(v string)`

SetName sets Name field to given value.


### GetIdFieldName

`func (o *Collection) GetIdFieldName() string`

GetIdFieldName returns the IdFieldName field if non-nil, zero value otherwise.

### GetIdFieldNameOk

`func (o *Collection) GetIdFieldNameOk() (*string, bool)`

GetIdFieldNameOk returns a tuple with the IdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFieldName

`func (o *Collection) SetIdFieldName(v string)`

SetIdFieldName sets IdFieldName field to given value.


### GetFields

`func (o *Collection) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Collection) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Collection) SetFields(v []Field)`

SetFields sets Fields field to given value.


### GetIndexes

`func (o *Collection) GetIndexes() []Index`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *Collection) GetIndexesOk() (*[]Index, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *Collection) SetIndexes(v []Index)`

SetIndexes sets Indexes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


