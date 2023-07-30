# IndexCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | **[]string** |  | 
**IsUnique** | **bool** |  | 

## Methods

### NewIndexCreateRequest

`func NewIndexCreateRequest(fields []string, isUnique bool, ) *IndexCreateRequest`

NewIndexCreateRequest instantiates a new IndexCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexCreateRequestWithDefaults

`func NewIndexCreateRequestWithDefaults() *IndexCreateRequest`

NewIndexCreateRequestWithDefaults instantiates a new IndexCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *IndexCreateRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IndexCreateRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IndexCreateRequest) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetIsUnique

`func (o *IndexCreateRequest) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *IndexCreateRequest) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *IndexCreateRequest) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


