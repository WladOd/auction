# ActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**StatusCode**](StatusCode.md) |  | 
**Message** | **string** |  | 
**Code** | Pointer to **int32** | HTTP status code | [optional] 

## Methods

### NewActionResult

`func NewActionResult(status StatusCode, message string, ) *ActionResult`

NewActionResult instantiates a new ActionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResultWithDefaults

`func NewActionResultWithDefaults() *ActionResult`

NewActionResultWithDefaults instantiates a new ActionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionResult) GetStatus() StatusCode`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionResult) GetStatusOk() (*StatusCode, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionResult) SetStatus(v StatusCode)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ActionResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ActionResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ActionResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ActionResult) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ActionResult) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ActionResult) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ActionResult) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


