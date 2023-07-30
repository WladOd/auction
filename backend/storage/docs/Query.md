# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expressions** | Pointer to [**[]QueryExpression**](QueryExpression.md) |  | [optional] 
**OrderBy** | Pointer to [**[]OrderBy**](OrderBy.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressions

`func (o *Query) GetExpressions() []QueryExpression`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *Query) GetExpressionsOk() (*[]QueryExpression, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *Query) SetExpressions(v []QueryExpression)`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *Query) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetOrderBy

`func (o *Query) GetOrderBy() []OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *Query) GetOrderByOk() (*[]OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *Query) SetOrderBy(v []OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *Query) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetLimit

`func (o *Query) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Query) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Query) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Query) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


