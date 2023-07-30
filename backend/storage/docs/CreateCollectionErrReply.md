# CreateCollectionErrReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | application-specific error code | 
**Error** | **string** | application-level error message, for debugging | 
**Status** | **string** | user-level status message | 
**MaxNumberOfCollections** | **int32** | the maximum number of collections a tenant can create with the current license | 
**MaxLedgerDbSize** | **float64** | ledger DB max size exceed | 
**LicenseType** | **string** | the type of the license of the tenant | 

## Methods

### NewCreateCollectionErrReply

`func NewCreateCollectionErrReply(code int32, error_ string, status string, maxNumberOfCollections int32, maxLedgerDbSize float64, licenseType string, ) *CreateCollectionErrReply`

NewCreateCollectionErrReply instantiates a new CreateCollectionErrReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionErrReplyWithDefaults

`func NewCreateCollectionErrReplyWithDefaults() *CreateCollectionErrReply`

NewCreateCollectionErrReplyWithDefaults instantiates a new CreateCollectionErrReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateCollectionErrReply) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateCollectionErrReply) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateCollectionErrReply) SetCode(v int32)`

SetCode sets Code field to given value.


### GetError

`func (o *CreateCollectionErrReply) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateCollectionErrReply) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateCollectionErrReply) SetError(v string)`

SetError sets Error field to given value.


### GetStatus

`func (o *CreateCollectionErrReply) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCollectionErrReply) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCollectionErrReply) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMaxNumberOfCollections

`func (o *CreateCollectionErrReply) GetMaxNumberOfCollections() int32`

GetMaxNumberOfCollections returns the MaxNumberOfCollections field if non-nil, zero value otherwise.

### GetMaxNumberOfCollectionsOk

`func (o *CreateCollectionErrReply) GetMaxNumberOfCollectionsOk() (*int32, bool)`

GetMaxNumberOfCollectionsOk returns a tuple with the MaxNumberOfCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfCollections

`func (o *CreateCollectionErrReply) SetMaxNumberOfCollections(v int32)`

SetMaxNumberOfCollections sets MaxNumberOfCollections field to given value.


### GetMaxLedgerDbSize

`func (o *CreateCollectionErrReply) GetMaxLedgerDbSize() float64`

GetMaxLedgerDbSize returns the MaxLedgerDbSize field if non-nil, zero value otherwise.

### GetMaxLedgerDbSizeOk

`func (o *CreateCollectionErrReply) GetMaxLedgerDbSizeOk() (*float64, bool)`

GetMaxLedgerDbSizeOk returns a tuple with the MaxLedgerDbSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLedgerDbSize

`func (o *CreateCollectionErrReply) SetMaxLedgerDbSize(v float64)`

SetMaxLedgerDbSize sets MaxLedgerDbSize field to given value.


### GetLicenseType

`func (o *CreateCollectionErrReply) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *CreateCollectionErrReply) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *CreateCollectionErrReply) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


