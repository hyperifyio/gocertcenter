# gocertcenter

::: app-desc
A microservice for managing digital certificates, designed for private
services and embedded application connections.
:::

::: app-desc
More information: <https://github.com/hyperifyio/gocertcenter>
:::

::: app-desc
Contact Info: [info@hg.fi](info@hg.fi)
:::

::: app-desc
Version: 0.0.1
:::

::: license-info
FSL-1.1-MIT
:::

::: license-url
https://fsl.software/
:::

## Access

## [Methods]{#__Methods}

\[ Jump to [Models](#__Models) \]

### Table of Contents

::: method-summary
:::

#### [Default](#Default)

-   [[`get`]{.http-method}` /`](#rootGet)

# [Default]{#Default}

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

### Return type {#return-type .field-label}

::: return-type
[inline_response_200](#inline_response_200)
:::

### Example data {#example-data .field-label}

::: example-data-content-type
Content-Type: application/json
:::

``` example
{
  "name" : "gocertcenter",
  "version" : "0.0.1"
}
```

### Produces {#produces .field-label}

This API call produces the following media types according to the
[Accept]{.header} request header; the media type will be conveyed by the
[Content-Type]{.header} response header.

-   `application/json`

### Responses {#responses .field-label}

#### 200 {#section .field-label}

[inline_response_200](#inline_response_200)
::::::::

------------------------------------------------------------------------

## [Models]{#__Models}

\[ Jump to [Methods](#__Methods) \]

### Table of Contents

1.  [`inline_response_200`](#inline_response_200)

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
