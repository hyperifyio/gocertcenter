# Example API

::: app-desc
No description provided (generated by Swagger Codegen
https://github.com/swagger-api/swagger-codegen)
:::

::: app-desc
More information: <https://helloreverb.com>
:::

::: app-desc
Contact Info: [hello@helloreverb.com](hello@helloreverb.com)
:::

::: app-desc
Version: 0.0.0
:::

::: license-info
All rights reserved
:::

::: license-url
http://apache.org/licenses/LICENSE-2.0.html
:::

## Access

## [Methods]{#__Methods}

\[ Jump to [Models](#__Models) \]

### Table of Contents

::: method-summary
:::

#### [Default](#Default)

-   [[`get`]{.http-method}` /organizations`](#organizationsGet)
-   [[`get`]{.http-method}` /organizations/{organization}/certificates`](#organizationsOrganizationCertificatesGet)
-   [[`post`]{.http-method}` /organizations/{organization}/certificates`](#organizationsOrganizationCertificatesPost)
-   [[`get`]{.http-method}` /organizations/{organization}/certificates/{rootSerialNumber}/certificates`](#organizationsOrganizationCertificatesRootSerialNumberCertificatesGet)
-   [[`post`]{.http-method}` /organizations/{organization}/certificates/{rootSerialNumber}/certificates`](#organizationsOrganizationCertificatesRootSerialNumberCertificatesPost)
-   [[`get`]{.http-method}` /organizations/{organization}/certificates/{rootSerialNumber}/certificates/{serialNumber}`](#organizationsOrganizationCertificatesRootSerialNumberCertificatesSerialNumberGet)
-   [[`get`]{.http-method}` /organizations/{organization}/certificates/{rootSerialNumber}`](#organizationsOrganizationCertificatesRootSerialNumberGet)
-   [[`get`]{.http-method}` /organizations/{organization}`](#organizationsOrganizationGet)
-   [[`post`]{.http-method}` /organizations`](#organizationsPost)
-   [[`get`]{.http-method}` /`](#rootGet)

# [Default]{#Default}

:::::::: method
[]{#organizationsGet}

::: method-path
[Up](#__Methods){.up}

``` get
get /organizations
```
:::

::: method-summary
Returns a specific root certificate ([organizationsGet]{.nickname})
:::

::: method-notes
:::

### Return type {#return-type .field-label}

::: return-type
[inline_response_200_1](#inline_response_200_1)
:::

### Example data {#example-data .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "payload" : [ {
    "allNames" : [ "allNames", "allNames" ],
    "name" : "name",
    "id" : "id"
  }, {
    "allNames" : [ "allNames", "allNames" ],
    "name" : "name",
    "id" : "id"
  } ]
}
```

### Produces {#produces .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses .field-label}

#### 200 {#section .field-label}

[inline_response_200_1](#inline_response_200_1)
::::::::

------------------------------------------------------------------------

::::::::::: method
[]{#organizationsOrganizationCertificatesGet}

::: method-path
[Up](#__Methods){.up}

``` get
get /organizations/{organization}/certificates
```
:::

::: method-summary
Returns a collection of root certificate entities
([organizationsOrganizationCertificatesGet]{.nickname})
:::

::: method-notes
:::

### Path parameters {#path-parameters .field-label}

::::: field-items
::: param
organization (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::
:::::

### Return type {#return-type-1 .field-label}

::: return-type
[inline_response_200_2](#inline_response_200_2)
:::

### Example data {#example-data-1 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "payload" : [ {
    "commonName" : "commonName",
    "serialNumber" : "serialNumber",
    "signedBy" : "signedBy",
    "isRootCertificate" : true,
    "organization" : "organization",
    "certificate" : "certificate",
    "isCA" : true,
    "isServerCertificate" : true,
    "isIntermediateCertificate" : true,
    "isClientCertificate" : true,
    "parents" : [ "parents", "parents" ]
  }, {
    "commonName" : "commonName",
    "serialNumber" : "serialNumber",
    "signedBy" : "signedBy",
    "isRootCertificate" : true,
    "organization" : "organization",
    "certificate" : "certificate",
    "isCA" : true,
    "isServerCertificate" : true,
    "isIntermediateCertificate" : true,
    "isClientCertificate" : true,
    "parents" : [ "parents", "parents" ]
  } ]
}
```

### Produces {#produces-1 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-1 .field-label}

#### 200 {#section-1 .field-label}

[inline_response_200_2](#inline_response_200_2)
:::::::::::

------------------------------------------------------------------------

:::::::::::::: method
[]{#organizationsOrganizationCertificatesPost}

::: method-path
[Up](#__Methods){.up}

``` post
post /organizations/{organization}/certificates
```
:::

::: method-summary
Returns a collection of organization entities
([organizationsOrganizationCertificatesPost]{.nickname})
:::

::: method-notes
:::

### Path parameters {#path-parameters-1 .field-label}

::::: field-items
::: param
organization (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::
:::::

### Consumes {#consumes .field-label}

This API call consumes the following media types via the
[Content-Type]{.header} request header:

-   `application/json`

### Request body {#request-body .field-label}

::::: field-items
::: param
body [organization_certificates_body](#organization_certificates_body)
(optional)
:::

::: param-desc
[Body Parameter]{.param-type} --- Certificate request data
:::
:::::

### Return type {#return-type-2 .field-label}

::: return-type
[inline_response_200_3](#inline_response_200_3)
:::

### Example data {#example-data-2 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "commonName" : "commonName",
  "serialNumber" : "serialNumber",
  "signedBy" : "signedBy",
  "isRootCertificate" : true,
  "organization" : "organization",
  "certificate" : "certificate",
  "isCA" : true,
  "isServerCertificate" : true,
  "isIntermediateCertificate" : true,
  "isClientCertificate" : true,
  "parents" : [ "parents", "parents" ]
}
```

### Produces {#produces-2 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-2 .field-label}

#### 200 {#section-2 .field-label}

[inline_response_200_3](#inline_response_200_3)
::::::::::::::

------------------------------------------------------------------------

::::::::::::: method
[]{#organizationsOrganizationCertificatesRootSerialNumberCertificatesGet}

::: method-path
[Up](#__Methods){.up}

``` get
get /organizations/{organization}/certificates/{rootSerialNumber}/certificates
```
:::

::: method-summary
Returns a collection of root certificate entities
([organizationsOrganizationCertificatesRootSerialNumberCertificatesGet]{.nickname})
:::

::: method-notes
:::

### Path parameters {#path-parameters-2 .field-label}

::::::: field-items
::: param
organization (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::

::: param
rootSerialNumber (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::
:::::::

### Return type {#return-type-3 .field-label}

::: return-type
[inline_response_200_2](#inline_response_200_2)
:::

### Example data {#example-data-3 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "payload" : [ {
    "commonName" : "commonName",
    "serialNumber" : "serialNumber",
    "signedBy" : "signedBy",
    "isRootCertificate" : true,
    "organization" : "organization",
    "certificate" : "certificate",
    "isCA" : true,
    "isServerCertificate" : true,
    "isIntermediateCertificate" : true,
    "isClientCertificate" : true,
    "parents" : [ "parents", "parents" ]
  }, {
    "commonName" : "commonName",
    "serialNumber" : "serialNumber",
    "signedBy" : "signedBy",
    "isRootCertificate" : true,
    "organization" : "organization",
    "certificate" : "certificate",
    "isCA" : true,
    "isServerCertificate" : true,
    "isIntermediateCertificate" : true,
    "isClientCertificate" : true,
    "parents" : [ "parents", "parents" ]
  } ]
}
```

### Produces {#produces-3 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-3 .field-label}

#### 200 {#section-3 .field-label}

[inline_response_200_2](#inline_response_200_2)
:::::::::::::

------------------------------------------------------------------------

:::::::::::::::: method
[]{#organizationsOrganizationCertificatesRootSerialNumberCertificatesPost}

::: method-path
[Up](#__Methods){.up}

``` post
post /organizations/{organization}/certificates/{rootSerialNumber}/certificates
```
:::

::: method-summary
Creates another certificate under a root certificate
([organizationsOrganizationCertificatesRootSerialNumberCertificatesPost]{.nickname})
:::

::: method-notes
:::

### Path parameters {#path-parameters-3 .field-label}

::::::: field-items
::: param
organization (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::

::: param
rootSerialNumber (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::
:::::::

### Consumes {#consumes-1 .field-label}

This API call consumes the following media types via the
[Content-Type]{.header} request header:

-   `application/json`

### Request body {#request-body-1 .field-label}

::::: field-items
::: param
body
[rootSerialNumber_certificates_body](#rootSerialNumber_certificates_body)
(optional)
:::

::: param-desc
[Body Parameter]{.param-type} --- Certificate request data
:::
:::::

### Return type {#return-type-4 .field-label}

::: return-type
[inline_response_200_3](#inline_response_200_3)
:::

### Example data {#example-data-4 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "commonName" : "commonName",
  "serialNumber" : "serialNumber",
  "signedBy" : "signedBy",
  "isRootCertificate" : true,
  "organization" : "organization",
  "certificate" : "certificate",
  "isCA" : true,
  "isServerCertificate" : true,
  "isIntermediateCertificate" : true,
  "isClientCertificate" : true,
  "parents" : [ "parents", "parents" ]
}
```

### Produces {#produces-4 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-4 .field-label}

#### 200 {#section-4 .field-label}

[inline_response_200_3](#inline_response_200_3)
::::::::::::::::

------------------------------------------------------------------------

::::::::::::::: method
[]{#organizationsOrganizationCertificatesRootSerialNumberCertificatesSerialNumberGet}

::: method-path
[Up](#__Methods){.up}

``` get
get /organizations/{organization}/certificates/{rootSerialNumber}/certificates/{serialNumber}
```
:::

::: method-summary
Returns a certificate entity owned by a root certificate
([organizationsOrganizationCertificatesRootSerialNumberCertificatesSerialNumberGet]{.nickname})
:::

::: method-notes
:::

### Path parameters {#path-parameters-4 .field-label}

::::::::: field-items
::: param
organization (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::

::: param
rootSerialNumber (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::

::: param
serialNumber (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::
:::::::::

### Return type {#return-type-5 .field-label}

::: return-type
[organizations_body](#organizations_body)
:::

### Example data {#example-data-5 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "allNames" : [ "allNames", "allNames" ],
  "name" : "name",
  "id" : "id"
}
```

### Produces {#produces-5 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-5 .field-label}

#### 200 {#section-5 .field-label}

[organizations_body](#organizations_body)
:::::::::::::::

------------------------------------------------------------------------

::::::::::::: method
[]{#organizationsOrganizationCertificatesRootSerialNumberGet}

::: method-path
[Up](#__Methods){.up}

``` get
get /organizations/{organization}/certificates/{rootSerialNumber}
```
:::

::: method-summary
Returns an root certificate
([organizationsOrganizationCertificatesRootSerialNumberGet]{.nickname})
:::

::: method-notes
:::

### Path parameters {#path-parameters-5 .field-label}

::::::: field-items
::: param
organization (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::

::: param
rootSerialNumber (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::
:::::::

### Return type {#return-type-6 .field-label}

::: return-type
[inline_response_200_3](#inline_response_200_3)
:::

### Example data {#example-data-6 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "commonName" : "commonName",
  "serialNumber" : "serialNumber",
  "signedBy" : "signedBy",
  "isRootCertificate" : true,
  "organization" : "organization",
  "certificate" : "certificate",
  "isCA" : true,
  "isServerCertificate" : true,
  "isIntermediateCertificate" : true,
  "isClientCertificate" : true,
  "parents" : [ "parents", "parents" ]
}
```

### Produces {#produces-6 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-6 .field-label}

#### 200 {#section-6 .field-label}

[inline_response_200_3](#inline_response_200_3)
:::::::::::::

------------------------------------------------------------------------

::::::::::: method
[]{#organizationsOrganizationGet}

::: method-path
[Up](#__Methods){.up}

``` get
get /organizations/{organization}
```
:::

::: method-summary
Returns an organization entity
([organizationsOrganizationGet]{.nickname})
:::

::: method-notes
:::

### Path parameters {#path-parameters-6 .field-label}

::::: field-items
::: param
organization (required)
:::

::: param-desc
[Path Parameter]{.param-type} ---
:::
:::::

### Return type {#return-type-7 .field-label}

::: return-type
[organizations_body](#organizations_body)
:::

### Example data {#example-data-7 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "allNames" : [ "allNames", "allNames" ],
  "name" : "name",
  "id" : "id"
}
```

### Produces {#produces-7 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-7 .field-label}

#### 200 {#section-7 .field-label}

[organizations_body](#organizations_body)
:::::::::::

------------------------------------------------------------------------

::::::::::: method
[]{#organizationsPost}

::: method-path
[Up](#__Methods){.up}

``` post
post /organizations
```
:::

::: method-summary
Creates an organization ([organizationsPost]{.nickname})
:::

::: method-notes
:::

### Consumes {#consumes-2 .field-label}

This API call consumes the following media types via the
[Content-Type]{.header} request header:

-   `application/json`

### Request body {#request-body-2 .field-label}

::::: field-items
::: param
body [organizations_body](#organizations_body) (optional)
:::

::: param-desc
[Body Parameter]{.param-type} --- Organization data
:::
:::::

### Return type {#return-type-8 .field-label}

::: return-type
[organizations_body](#organizations_body)
:::

### Example data {#example-data-8 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "allNames" : [ "allNames", "allNames" ],
  "name" : "name",
  "id" : "id"
}
```

### Produces {#produces-8 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-8 .field-label}

#### 200 {#section-8 .field-label}

[organizations_body](#organizations_body)
:::::::::::

------------------------------------------------------------------------

:::::::: method
[]{#rootGet}

::: method-path
[Up](#__Methods){.up}

``` get
get /
```
:::

::: method-summary
Returns information about the running server ([rootGet]{.nickname})
:::

::: method-notes
This includes the software name and a version
:::

### Return type {#return-type-9 .field-label}

::: return-type
[inline_response_200](#inline_response_200)
:::

### Example data {#example-data-9 .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "name" : "gocertcenter",
  "version" : "0.0.1"
}
```

### Produces {#produces-9 .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses-9 .field-label}

#### 200 {#section-9 .field-label}

[inline_response_200](#inline_response_200)
::::::::

------------------------------------------------------------------------

## [Models]{#__Models}

\[ Jump to [Methods](#__Methods) \]

### Table of Contents

1.  [`inline_response_200`](#inline_response_200)
2.  [`inline_response_200_1`](#inline_response_200_1)
3.  [`inline_response_200_2`](#inline_response_200_2)
4.  [`inline_response_200_3`](#inline_response_200_3)
5.  [`organization_certificates_body`](#organization_certificates_body)
6.  [`organizations_body`](#organizations_body)
7.  [`rootSerialNumber_certificates_body`](#rootSerialNumber_certificates_body)

:::::::::: model
### [`inline_response_200`]{#inline_response_200} [Up](#__Models){.up}

::::::::: field-items
::: param
name
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param-desc
[example: gocertcenter]{.param-type}
:::

::: param
version
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param-desc
[example: 0.0.1]{.param-type}
:::
:::::::::
::::::::::

:::::: model
### [`inline_response_200_1`]{#inline_response_200_1} [Up](#__Models){.up}

::::: field-items
::: param
payload
:::

::: param-desc
[[array\[Object\]](#object)]{.param-type}
:::
:::::
::::::

:::::: model
### [`inline_response_200_2`]{#inline_response_200_2} [Up](#__Models){.up}

::::: field-items
::: param
payload
:::

::: param-desc
[[array\[Object\]](#object)]{.param-type}
:::
:::::
::::::

:::::::::::::::::::::::::: model
### [`inline_response_200_3`]{#inline_response_200_3} [Up](#__Models){.up}

::::::::::::::::::::::::: field-items
::: param
certificate
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param
commonName
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param
isCA
:::

::: param-desc
[[Boolean](#boolean)]{.param-type}
:::

::: param
isClientCertificate
:::

::: param-desc
[[Boolean](#boolean)]{.param-type}
:::

::: param
isIntermediateCertificate
:::

::: param-desc
[[Boolean](#boolean)]{.param-type}
:::

::: param
isRootCertificate
:::

::: param-desc
[[Boolean](#boolean)]{.param-type}
:::

::: param
isServerCertificate
:::

::: param-desc
[[Boolean](#boolean)]{.param-type}
:::

::: param
organization
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param
parents
:::

::: param-desc
[[array\[String\]](#string)]{.param-type}
:::

::: param
serialNumber
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param
signedBy
:::

::: param-desc
[[String](#string)]{.param-type}
:::
:::::::::::::::::::::::::
::::::::::::::::::::::::::

:::::::::::: model
### [`organization_certificates_body`]{#organization_certificates_body} [Up](#__Models){.up}

::::::::::: field-items
::: param
commonName
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param
dnsNames
:::

::: param-desc
[[array\[String\]](#string)]{.param-type}
:::

::: param
expiration
:::

::: param-desc
[[Integer](#integer)]{.param-type}
:::

::: param
type
:::

::: param-desc
[[String](#string)]{.param-type}
:::
:::::::::::
::::::::::::

:::::::::: model
### [`organizations_body`]{#organizations_body} [Up](#__Models){.up}

::::::::: field-items
::: param
allNames
:::

::: param-desc
[[array\[String\]](#string)]{.param-type}
:::

::: param
id
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param
name
:::

::: param-desc
[[String](#string)]{.param-type}
:::
:::::::::
::::::::::

:::::::::::: model
### [`rootSerialNumber_certificates_body`]{#rootSerialNumber_certificates_body} [Up](#__Models){.up}

::::::::::: field-items
::: param
commonName
:::

::: param-desc
[[String](#string)]{.param-type}
:::

::: param
dnsNames
:::

::: param-desc
[[array\[String\]](#string)]{.param-type}
:::

::: param
expiration
:::

::: param-desc
[[Integer](#integer)]{.param-type}
:::

::: param
type
:::

::: param-desc
[[String](#string)]{.param-type}
:::
:::::::::::
::::::::::::
