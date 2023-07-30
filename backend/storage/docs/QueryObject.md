# QueryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Operator** | [**Operator**](Operator.md) |  | [default to EQ]
**Value** | [**QueryObjectValue**](QueryObjectValue.md) |  | 

## Methods

### NewQueryObject

`func NewQueryObject(field string, operator Operator, value QueryObjectValue, ) *QueryObject`

NewQueryObject instantiates a new QueryObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryObjectWithDefaults

`func NewQueryObjectWithDefaults() *QueryObject`

NewQueryObjectWithDefaults instantiates a new QueryObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *QueryObject) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *QueryObject) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *QueryObject) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *QueryObject) GetOperator() Operator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *QueryObject) GetOperatorOk() (*Operator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *QueryObject) SetOperator(v Operator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *QueryObject) GetValue() QueryObjectValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryObject) GetValueOk() (*QueryObjectValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryObject) SetValue(v QueryObjectValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


