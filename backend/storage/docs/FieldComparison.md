# FieldComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Operator** | [**Operator**](Operator.md) |  | [default to EQ]
**Value** | **interface{}** |  | 

## Methods

### NewFieldComparison

`func NewFieldComparison(field string, operator Operator, value interface{}, ) *FieldComparison`

NewFieldComparison instantiates a new FieldComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldComparisonWithDefaults

`func NewFieldComparisonWithDefaults() *FieldComparison`

NewFieldComparisonWithDefaults instantiates a new FieldComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldComparison) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldComparison) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldComparison) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldComparison) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldComparison) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldComparison) SetOperator(v Operator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *FieldComparison) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldComparison) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldComparison) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FieldComparison) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FieldComparison) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


