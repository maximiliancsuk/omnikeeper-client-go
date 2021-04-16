# GridViewConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowCIIDColumn** | Pointer to **bool** |  | [optional] 
**WriteLayer** | Pointer to **int64** |  | [optional] 
**ReadLayerset** | Pointer to **[]int64** |  | [optional] 
**Columns** | Pointer to [**[]GridViewColumn**](GridViewColumn.md) |  | [optional] 
**Trait** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGridViewConfiguration

`func NewGridViewConfiguration() *GridViewConfiguration`

NewGridViewConfiguration instantiates a new GridViewConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridViewConfigurationWithDefaults

`func NewGridViewConfigurationWithDefaults() *GridViewConfiguration`

NewGridViewConfigurationWithDefaults instantiates a new GridViewConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowCIIDColumn

`func (o *GridViewConfiguration) GetShowCIIDColumn() bool`

GetShowCIIDColumn returns the ShowCIIDColumn field if non-nil, zero value otherwise.

### GetShowCIIDColumnOk

`func (o *GridViewConfiguration) GetShowCIIDColumnOk() (*bool, bool)`

GetShowCIIDColumnOk returns a tuple with the ShowCIIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCIIDColumn

`func (o *GridViewConfiguration) SetShowCIIDColumn(v bool)`

SetShowCIIDColumn sets ShowCIIDColumn field to given value.

### HasShowCIIDColumn

`func (o *GridViewConfiguration) HasShowCIIDColumn() bool`

HasShowCIIDColumn returns a boolean if a field has been set.

### GetWriteLayer

`func (o *GridViewConfiguration) GetWriteLayer() int64`

GetWriteLayer returns the WriteLayer field if non-nil, zero value otherwise.

### GetWriteLayerOk

`func (o *GridViewConfiguration) GetWriteLayerOk() (*int64, bool)`

GetWriteLayerOk returns a tuple with the WriteLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLayer

`func (o *GridViewConfiguration) SetWriteLayer(v int64)`

SetWriteLayer sets WriteLayer field to given value.

### HasWriteLayer

`func (o *GridViewConfiguration) HasWriteLayer() bool`

HasWriteLayer returns a boolean if a field has been set.

### GetReadLayerset

`func (o *GridViewConfiguration) GetReadLayerset() []int64`

GetReadLayerset returns the ReadLayerset field if non-nil, zero value otherwise.

### GetReadLayersetOk

`func (o *GridViewConfiguration) GetReadLayersetOk() (*[]int64, bool)`

GetReadLayersetOk returns a tuple with the ReadLayerset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLayerset

`func (o *GridViewConfiguration) SetReadLayerset(v []int64)`

SetReadLayerset sets ReadLayerset field to given value.

### HasReadLayerset

`func (o *GridViewConfiguration) HasReadLayerset() bool`

HasReadLayerset returns a boolean if a field has been set.

### SetReadLayersetNil

`func (o *GridViewConfiguration) SetReadLayersetNil(b bool)`

 SetReadLayersetNil sets the value for ReadLayerset to be an explicit nil

### UnsetReadLayerset
`func (o *GridViewConfiguration) UnsetReadLayerset()`

UnsetReadLayerset ensures that no value is present for ReadLayerset, not even an explicit nil
### GetColumns

`func (o *GridViewConfiguration) GetColumns() []GridViewColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *GridViewConfiguration) GetColumnsOk() (*[]GridViewColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *GridViewConfiguration) SetColumns(v []GridViewColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *GridViewConfiguration) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *GridViewConfiguration) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *GridViewConfiguration) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetTrait

`func (o *GridViewConfiguration) GetTrait() string`

GetTrait returns the Trait field if non-nil, zero value otherwise.

### GetTraitOk

`func (o *GridViewConfiguration) GetTraitOk() (*string, bool)`

GetTraitOk returns a tuple with the Trait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrait

`func (o *GridViewConfiguration) SetTrait(v string)`

SetTrait sets Trait field to given value.

### HasTrait

`func (o *GridViewConfiguration) HasTrait() bool`

HasTrait returns a boolean if a field has been set.

### SetTraitNil

`func (o *GridViewConfiguration) SetTraitNil(b bool)`

 SetTraitNil sets the value for Trait to be an explicit nil

### UnsetTrait
`func (o *GridViewConfiguration) UnsetTrait()`

UnsetTrait ensures that no value is present for Trait, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

