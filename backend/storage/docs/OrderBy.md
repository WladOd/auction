# OrderBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Desc** | **bool** |  | 

## Methods

### NewOrderBy

`func NewOrderBy(field string, desc bool, ) *OrderBy`

NewOrderBy instantiates a new OrderBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderByWithDefaults

`func NewOrderByWithDefaults() *OrderBy`

NewOrderByWithDefaults instantiates a new OrderBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *OrderBy) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *OrderBy) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *OrderBy) SetField(v string)`

SetField sets Field field to given value.


### GetDesc

`func (o *OrderBy) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *OrderBy) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *OrderBy) SetDesc(v bool)`

SetDesc sets Desc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


