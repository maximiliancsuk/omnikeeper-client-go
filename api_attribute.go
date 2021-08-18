/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"reflect"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// AttributeApiService AttributeApi service
type AttributeApiService service

type ApiBulkReplaceAttributesInLayerRequest struct {
	ctx _context.Context
	ApiService *AttributeApiService
	version string
	bulkCIAttributeLayerScopeDTO *BulkCIAttributeLayerScopeDTO
}

func (r ApiBulkReplaceAttributesInLayerRequest) BulkCIAttributeLayerScopeDTO(bulkCIAttributeLayerScopeDTO BulkCIAttributeLayerScopeDTO) ApiBulkReplaceAttributesInLayerRequest {
	r.bulkCIAttributeLayerScopeDTO = &bulkCIAttributeLayerScopeDTO
	return r
}

func (r ApiBulkReplaceAttributesInLayerRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.BulkReplaceAttributesInLayerExecute(r)
}

/*
BulkReplaceAttributesInLayer bulk replace all attributes in specified layer

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiBulkReplaceAttributesInLayerRequest
*/
func (a *AttributeApiService) BulkReplaceAttributesInLayer(ctx _context.Context, version string) ApiBulkReplaceAttributesInLayerRequest {
	return ApiBulkReplaceAttributesInLayerRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
func (a *AttributeApiService) BulkReplaceAttributesInLayerExecute(r ApiBulkReplaceAttributesInLayerRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttributeApiService.BulkReplaceAttributesInLayer")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Attribute/bulkReplaceAttributesInLayer"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.bulkCIAttributeLayerScopeDTO == nil {
		return nil, reportError("bulkCIAttributeLayerScopeDTO is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "application/json-patch+json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkCIAttributeLayerScopeDTO
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFindMergedAttributesByNameRequest struct {
	ctx _context.Context
	ApiService *AttributeApiService
	regex *string
	layerIDs *[]int64
	version string
	ciids *[]string
	atTime *time.Time
}

func (r ApiFindMergedAttributesByNameRequest) Regex(regex string) ApiFindMergedAttributesByNameRequest {
	r.regex = &regex
	return r
}
func (r ApiFindMergedAttributesByNameRequest) LayerIDs(layerIDs []int64) ApiFindMergedAttributesByNameRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiFindMergedAttributesByNameRequest) Ciids(ciids []string) ApiFindMergedAttributesByNameRequest {
	r.ciids = &ciids
	return r
}
func (r ApiFindMergedAttributesByNameRequest) AtTime(atTime time.Time) ApiFindMergedAttributesByNameRequest {
	r.atTime = &atTime
	return r
}

func (r ApiFindMergedAttributesByNameRequest) Execute() ([]CIAttributeDTO, *_nethttp.Response, error) {
	return r.ApiService.FindMergedAttributesByNameExecute(r)
}

/*
FindMergedAttributesByName Method for FindMergedAttributesByName

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiFindMergedAttributesByNameRequest
*/
func (a *AttributeApiService) FindMergedAttributesByName(ctx _context.Context, version string) ApiFindMergedAttributesByNameRequest {
	return ApiFindMergedAttributesByNameRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []CIAttributeDTO
func (a *AttributeApiService) FindMergedAttributesByNameExecute(r ApiFindMergedAttributesByNameRequest) ([]CIAttributeDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []CIAttributeDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttributeApiService.FindMergedAttributesByName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Attribute/findMergedAttributesByName"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.regex == nil {
		return localVarReturnValue, nil, reportError("regex is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	localVarQueryParams.Add("regex", parameterToString(*r.regex, ""))
	if r.ciids != nil {
		t := *r.ciids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ciids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ciids", parameterToString(t, "multi"))
		}
	}
	{
		t := *r.layerIDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("layerIDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("layerIDs", parameterToString(t, "multi"))
		}
	}
	if r.atTime != nil {
		localVarQueryParams.Add("atTime", parameterToString(*r.atTime, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "text/plain", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMergedAttributeRequest struct {
	ctx _context.Context
	ApiService *AttributeApiService
	ciid *string
	name *string
	layerIDs *[]int64
	version string
	atTime *time.Time
}

func (r ApiGetMergedAttributeRequest) Ciid(ciid string) ApiGetMergedAttributeRequest {
	r.ciid = &ciid
	return r
}
func (r ApiGetMergedAttributeRequest) Name(name string) ApiGetMergedAttributeRequest {
	r.name = &name
	return r
}
func (r ApiGetMergedAttributeRequest) LayerIDs(layerIDs []int64) ApiGetMergedAttributeRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetMergedAttributeRequest) AtTime(atTime time.Time) ApiGetMergedAttributeRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetMergedAttributeRequest) Execute() (CIAttributeDTO, *_nethttp.Response, error) {
	return r.ApiService.GetMergedAttributeExecute(r)
}

/*
GetMergedAttribute Method for GetMergedAttribute

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetMergedAttributeRequest
*/
func (a *AttributeApiService) GetMergedAttribute(ctx _context.Context, version string) ApiGetMergedAttributeRequest {
	return ApiGetMergedAttributeRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return CIAttributeDTO
func (a *AttributeApiService) GetMergedAttributeExecute(r ApiGetMergedAttributeRequest) (CIAttributeDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CIAttributeDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttributeApiService.GetMergedAttribute")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Attribute/getMergedAttribute"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ciid == nil {
		return localVarReturnValue, nil, reportError("ciid is required and must be specified")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	localVarQueryParams.Add("ciid", parameterToString(*r.ciid, ""))
	localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	{
		t := *r.layerIDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("layerIDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("layerIDs", parameterToString(t, "multi"))
		}
	}
	if r.atTime != nil {
		localVarQueryParams.Add("atTime", parameterToString(*r.atTime, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "text/plain", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMergedAttributesRequest struct {
	ctx _context.Context
	ApiService *AttributeApiService
	ciids *[]string
	layerIDs *[]int64
	version string
	atTime *time.Time
}

func (r ApiGetMergedAttributesRequest) Ciids(ciids []string) ApiGetMergedAttributesRequest {
	r.ciids = &ciids
	return r
}
func (r ApiGetMergedAttributesRequest) LayerIDs(layerIDs []int64) ApiGetMergedAttributesRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetMergedAttributesRequest) AtTime(atTime time.Time) ApiGetMergedAttributesRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetMergedAttributesRequest) Execute() ([]CIAttributeDTO, *_nethttp.Response, error) {
	return r.ApiService.GetMergedAttributesExecute(r)
}

/*
GetMergedAttributes Method for GetMergedAttributes

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetMergedAttributesRequest
*/
func (a *AttributeApiService) GetMergedAttributes(ctx _context.Context, version string) ApiGetMergedAttributesRequest {
	return ApiGetMergedAttributesRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []CIAttributeDTO
func (a *AttributeApiService) GetMergedAttributesExecute(r ApiGetMergedAttributesRequest) ([]CIAttributeDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []CIAttributeDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttributeApiService.GetMergedAttributes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Attribute/getMergedAttributes"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ciids == nil {
		return localVarReturnValue, nil, reportError("ciids is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	{
		t := *r.ciids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ciids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ciids", parameterToString(t, "multi"))
		}
	}
	{
		t := *r.layerIDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("layerIDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("layerIDs", parameterToString(t, "multi"))
		}
	}
	if r.atTime != nil {
		localVarQueryParams.Add("atTime", parameterToString(*r.atTime, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "text/plain", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMergedAttributesWithNameRequest struct {
	ctx _context.Context
	ApiService *AttributeApiService
	name *string
	layerIDs *[]int64
	version string
	atTime *time.Time
}

func (r ApiGetMergedAttributesWithNameRequest) Name(name string) ApiGetMergedAttributesWithNameRequest {
	r.name = &name
	return r
}
func (r ApiGetMergedAttributesWithNameRequest) LayerIDs(layerIDs []int64) ApiGetMergedAttributesWithNameRequest {
	r.layerIDs = &layerIDs
	return r
}
func (r ApiGetMergedAttributesWithNameRequest) AtTime(atTime time.Time) ApiGetMergedAttributesWithNameRequest {
	r.atTime = &atTime
	return r
}

func (r ApiGetMergedAttributesWithNameRequest) Execute() ([]CIAttributeDTO, *_nethttp.Response, error) {
	return r.ApiService.GetMergedAttributesWithNameExecute(r)
}

/*
GetMergedAttributesWithName Method for GetMergedAttributesWithName

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param version
 @return ApiGetMergedAttributesWithNameRequest
*/
func (a *AttributeApiService) GetMergedAttributesWithName(ctx _context.Context, version string) ApiGetMergedAttributesWithNameRequest {
	return ApiGetMergedAttributesWithNameRequest{
		ApiService: a,
		ctx: ctx,
		version: version,
	}
}

// Execute executes the request
//  @return []CIAttributeDTO
func (a *AttributeApiService) GetMergedAttributesWithNameExecute(r ApiGetMergedAttributesWithNameRequest) ([]CIAttributeDTO, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []CIAttributeDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AttributeApiService.GetMergedAttributesWithName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v{version}/Attribute/getMergedAttributesWithName"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", _neturl.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if r.layerIDs == nil {
		return localVarReturnValue, nil, reportError("layerIDs is required and must be specified")
	}

	localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	{
		t := *r.layerIDs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("layerIDs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("layerIDs", parameterToString(t, "multi"))
		}
	}
	if r.atTime != nil {
		localVarQueryParams.Add("atTime", parameterToString(*r.atTime, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;odata.metadata=minimal;odata.streaming=true", "application/json;odata.metadata=minimal;odata.streaming=false", "application/json;odata.metadata=minimal", "application/json;odata.metadata=full;odata.streaming=true", "application/json;odata.metadata=full;odata.streaming=false", "application/json;odata.metadata=full", "application/json;odata.metadata=none;odata.streaming=true", "application/json;odata.metadata=none;odata.streaming=false", "application/json;odata.metadata=none", "application/json;odata.streaming=true", "application/json;odata.streaming=false", "application/json", "application/xml", "application/odata", "text/plain", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
