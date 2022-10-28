# Go API client for okclient

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 24.5.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import okclient "github.com/max-bytes/omnikeeper-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), okclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), okclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), okclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), okclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnsibleInventoryScanIngestApi* | [**AnsibleInventoryScanIngestIngestAnsibleInventoryScan**](docs/AnsibleInventoryScanIngestApi.md#ansibleinventoryscaningestingestansibleinventoryscan) | **Post** /api/v{version}/Ingest/AnsibleInventoryScan | 
*AuthRedirectApi* | [**AuthRedirectIndex**](docs/AuthRedirectApi.md#authredirectindex) | **Get** /.well-known/openid-configuration | 
*CytoscapeApi* | [**CytoscapeTraitCentric**](docs/CytoscapeApi.md#cytoscapetraitcentric) | **Get** /api/v{version}/Cytoscape/traitCentric | 
*GraphQLApi* | [**GraphQLDebug**](docs/GraphQLApi.md#graphqldebug) | **Post** /graphql-debug | 
*GraphQLApi* | [**GraphQLGet**](docs/GraphQLApi.md#graphqlget) | **Get** /graphql | 
*GraphQLApi* | [**GraphQLIndex**](docs/GraphQLApi.md#graphqlindex) | **Post** /graphql | 
*GraphvizDotApi* | [**GraphvizDotLayerCentric**](docs/GraphvizDotApi.md#graphvizdotlayercentric) | **Get** /api/v{version}/GraphvizDot/layerCentric | 
*GraphvizDotApi* | [**GraphvizDotTraitCentric**](docs/GraphvizDotApi.md#graphvizdottraitcentric) | **Get** /api/v{version}/GraphvizDot/traitCentric | 
*GridViewApi* | [**GridViewAddContext**](docs/GridViewApi.md#gridviewaddcontext) | **Post** /api/v{version}/GridView/context | Adds new context
*GridViewApi* | [**GridViewChangeData**](docs/GridViewApi.md#gridviewchangedata) | **Post** /api/v{version}/GridView/contexts/{context}/change | Saves grid view row changes and returns change results
*GridViewApi* | [**GridViewDeleteContext**](docs/GridViewApi.md#gridviewdeletecontext) | **Delete** /api/v{version}/GridView/context/{name} | Deletes specific context
*GridViewApi* | [**GridViewEditContext**](docs/GridViewApi.md#gridvieweditcontext) | **Put** /api/v{version}/GridView/context/{name} | Edits specific context
*GridViewApi* | [**GridViewGetData**](docs/GridViewApi.md#gridviewgetdata) | **Get** /api/v{version}/GridView/contexts/{context}/data | Returns grid view data for specific context
*GridViewApi* | [**GridViewGetGridViewContext**](docs/GridViewApi.md#gridviewgetgridviewcontext) | **Get** /api/v{version}/GridView/context/{name} | Returns a single context in full
*GridViewApi* | [**GridViewGetGridViewContexts**](docs/GridViewApi.md#gridviewgetgridviewcontexts) | **Get** /api/v{version}/GridView/contexts | Returns a list of contexts for grid view.
*GridViewApi* | [**GridViewGetSchema**](docs/GridViewApi.md#gridviewgetschema) | **Get** /api/v{version}/GridView/contexts/{context}/schema | Returns grid view schema for specific context
*ImportExportLayerApi* | [**ImportExportLayerExportLayer**](docs/ImportExportLayerApi.md#importexportlayerexportlayer) | **Get** /api/v{version}/ImportExportLayer/exportLayer | 
*MetadataApi* | [**MetadataGetMetadata**](docs/MetadataApi.md#metadatagetmetadata) | **Get** /api/odata/{context}/$metadata | 
*MetadataApi* | [**MetadataGetServiceDocument**](docs/MetadataApi.md#metadatagetservicedocument) | **Get** /api/odata/{context} | 
*OKPluginGenericJSONIngestApi* | [**ManageContextGetAllContexts**](docs/OKPluginGenericJSONIngestApi.md#managecontextgetallcontexts) | **Get** /api/v{version}/ingest/genericJSON/manage/context | 
*OKPluginGenericJSONIngestApi* | [**ManageContextGetContext**](docs/OKPluginGenericJSONIngestApi.md#managecontextgetcontext) | **Get** /api/v{version}/ingest/genericJSON/manage/context/{id} | 
*OKPluginGenericJSONIngestApi* | [**ManageContextRemoveContext**](docs/OKPluginGenericJSONIngestApi.md#managecontextremovecontext) | **Delete** /api/v{version}/ingest/genericJSON/manage/context/{id} | 
*OKPluginGenericJSONIngestApi* | [**ManageContextUpsertContext**](docs/OKPluginGenericJSONIngestApi.md#managecontextupsertcontext) | **Post** /api/v{version}/ingest/genericJSON/manage/context | 
*OKPluginGenericJSONIngestApi* | [**PassiveDataIngest**](docs/OKPluginGenericJSONIngestApi.md#passivedataingest) | **Post** /api/v{version}/ingest/genericJSON/data | 
*RestartApplicationApi* | [**RestartApplicationRestart**](docs/RestartApplicationApi.md#restartapplicationrestart) | **Get** /api/v{version}/RestartApplication/restart | 
*UsageStatsApi* | [**UsageStatsFetch**](docs/UsageStatsApi.md#usagestatsfetch) | **Get** /api/v{version}/UsageStats/fetch | 


## Documentation For Models

 - [AbstractInboundIDMethod](docs/AbstractInboundIDMethod.md)
 - [AddContextRequest](docs/AddContextRequest.md)
 - [AnsibleInventoryScanDTO](docs/AnsibleInventoryScanDTO.md)
 - [AttributeValueDTO](docs/AttributeValueDTO.md)
 - [AttributeValueType](docs/AttributeValueType.md)
 - [ChangeDataCell](docs/ChangeDataCell.md)
 - [ChangeDataRequest](docs/ChangeDataRequest.md)
 - [EditContextRequest](docs/EditContextRequest.md)
 - [EdmContainerElementKind](docs/EdmContainerElementKind.md)
 - [EdmExpressionKind](docs/EdmExpressionKind.md)
 - [EdmSchemaElementKind](docs/EdmSchemaElementKind.md)
 - [EdmTypeKind](docs/EdmTypeKind.md)
 - [GenericInboundAttribute](docs/GenericInboundAttribute.md)
 - [GenericInboundCI](docs/GenericInboundCI.md)
 - [GenericInboundCIIdMethod](docs/GenericInboundCIIdMethod.md)
 - [GenericInboundData](docs/GenericInboundData.md)
 - [GenericInboundRelation](docs/GenericInboundRelation.md)
 - [GraphQLQuery](docs/GraphQLQuery.md)
 - [GridViewColumn](docs/GridViewColumn.md)
 - [GridViewConfiguration](docs/GridViewConfiguration.md)
 - [IEdmEntityContainer](docs/IEdmEntityContainer.md)
 - [IEdmEntityContainerElement](docs/IEdmEntityContainerElement.md)
 - [IEdmExpression](docs/IEdmExpression.md)
 - [IEdmModel](docs/IEdmModel.md)
 - [IEdmSchemaElement](docs/IEdmSchemaElement.md)
 - [IEdmTerm](docs/IEdmTerm.md)
 - [IEdmType](docs/IEdmType.md)
 - [IEdmTypeReference](docs/IEdmTypeReference.md)
 - [IEdmVocabularyAnnotation](docs/IEdmVocabularyAnnotation.md)
 - [InboundIDMethodByAttribute](docs/InboundIDMethodByAttribute.md)
 - [InboundIDMethodByAttributeModifiers](docs/InboundIDMethodByAttributeModifiers.md)
 - [InboundIDMethodByByUnion](docs/InboundIDMethodByByUnion.md)
 - [InboundIDMethodByData](docs/InboundIDMethodByData.md)
 - [InboundIDMethodByIntersect](docs/InboundIDMethodByIntersect.md)
 - [InboundIDMethodByRelatedTempID](docs/InboundIDMethodByRelatedTempID.md)
 - [InboundIDMethodByTemporaryCIID](docs/InboundIDMethodByTemporaryCIID.md)
 - [NoFoundTargetCIHandling](docs/NoFoundTargetCIHandling.md)
 - [ODataEntitySetInfo](docs/ODataEntitySetInfo.md)
 - [ODataFunctionImportInfo](docs/ODataFunctionImportInfo.md)
 - [ODataServiceDocument](docs/ODataServiceDocument.md)
 - [ODataSingletonInfo](docs/ODataSingletonInfo.md)
 - [ODataTypeAnnotation](docs/ODataTypeAnnotation.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [SameTargetCIHandling](docs/SameTargetCIHandling.md)
 - [SameTempIDHandling](docs/SameTempIDHandling.md)
 - [SparseRow](docs/SparseRow.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://[keycloak-url]/auth/realms/acme/protocol/openid-connect/auth
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



