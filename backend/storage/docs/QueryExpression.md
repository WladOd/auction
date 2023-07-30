# QueryExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldComparisons** | Pointer to [**[]FieldComparison**](FieldComparison.md) |  | [optional] 

## Methods

### NewQueryExpression

`func NewQueryExpression() *QueryExpression`

NewQueryExpression instantiates a new QueryExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryExpressionWithDefaults

`func NewQueryExpressionWithDefaults() *QueryExpression`

NewQueryExpressionWithDefaults instantiates a new QueryExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldComparisons

`func (o *QueryExpression) GetFieldComparisons() []FieldComparison`

GetFieldComparisons returns the FieldComparisons field if non-nil, zero value otherwise.

### GetFieldComparisonsOk

`func (o *QueryExpression) GetFieldComparisonsOk() (*[]FieldComparison, bool)`

GetFieldComparisonsOk returns a tuple with the FieldComparisons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComparisons

`func (o *QueryExpression) SetFieldComparisons(v []FieldComparison)`

SetFieldComparisons sets FieldComparisons field to given value.

### HasFieldComparisons

`func (o *QueryExpression) HasFieldComparisons() bool`

HasFieldComparisons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


