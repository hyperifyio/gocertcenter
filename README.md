# `gocertcenter` by [HyperifyIO](https://github.com/hyperifyio): Elevating PKI with Certificate Management

`gocertcenter` embarks on HyperifyIO's ambitious journey to redefine secure, 
custom Public Key Infrastructure (PKI) solutions. As a crucial component of our 
ecosystem, it addresses the nuanced demands of digital certificate management 
for private networks and applications.

## Elevating Security through Simplified Certificate Management

At the heart of `gocertcenter` is a commitment to streamline the complexities 
of managing digital certificates for microservices, embedded systems, and 
private applications. Designed with a focus on security enhancements, 
`gocertcenter` excels in generating, renewing, and revoking certificates across 
the spectrum—from CA and intermediate to server and client certificates. Its 
adaptability and lightweight architecture make it an optimal choice for direct 
integration into Go applications, ensuring developers a seamless experience.

`gocertcenter` is purpose-built for organizations seeking to fortify their 
internal communications and application security using mTLS and PKI, 
sidestepping the challenges posed by traditional public Certificate Authorities 
(CAs). Ideal for scenarios demanding heightened security, trust, and privacy, 
it presents a perfect fit for both on-premise solutions and as a hosted SaaS 
offering at [https://cert.center](https://cert.center), providing an automated, 
effortless certificate management platform tailored for the intricacies of 
internal and bespoke application requirements.

**Project Status**: Advancing towards the initial MVP v1 release, our 
dedication lies in crafting a service that stands synonymous with security 
excellence and operational superiority. Stay updated with our progress and 
contribute to shaping the future of PKI management by following the [MVP v1 
development journey](https://github.com/hyperifyio/gocertcenter/issues/1).

## Key Features

- **Certificate Generation**: Streamline the creation of CA, intermediate, 
  server, and client certificates, specifically for internal and embedded 
  application use.

- **Renewal and Revocation**: Facilitate the automated renewal process and 
  straightforward revocation of certificates to maintain a secure and trusted 
  environment.

- **REST API Interface**: Utilize a fully documented REST API for seamless 
  certificate management, tailored for private service architectures.

- **Security First**: Ensure top-tier security with TLS encryption across all 
  communications, bolstered by the integrity of mTLS and PKI for private 
  connections.

- **Monitoring Integration**: Employ Prometheus for detailed, real-time metrics 
  and performance insights, enabling proactive management and operational 
  oversight.

## Designed for Developers

With developers in mind, CMM offers straightforward integration into Go 
applications, empowering microservices to self-manage their certificates 
efficiently. This approach significantly enhances security postures without 
compromising performance, catering exclusively to the needs of securing private 
services and customer-specific application connections.

## Installation

**Note**: Installation instructions will be detailed upon the MVP v1 release, 
including guidance for both on-premise and SaaS deployments.

## Usage

Upon project completion, we will provide extensive usage documentation, 
covering REST API interactions and guidance on integrating CMM into your 
projects, whether for on-premise use or as a part of our SaaS offering at 
[https://cert.center](https://cert.center).

### OpenAPI

Available from http://localhost:8080/documentation/json

## Development

### Internal modules

#### `./internal/app/` - Internal modules for application business

| Module            | Description                                                                 | Depends on                               |
|-------------------|-----------------------------------------------------------------------------|------------------------------------------|
| `appendpoints`    | Application end-point implementations (the main package)                    | `appmodels`, `appdtos`, `appcontrollers` |
| `appcontrollers`  | Controllers for application logic                                           | `appmodels`, `appdtos`, `apputils`       |
| `apprepositories` | Repository implementations for storing application state (the main package) | `appmodels`, `appdtos`, `apputils`       |
| `apputils`        | Lower level utilities for application logic                                 | `appmodels`, `appdtos`                   |
| `appmodels`       | Models for application state                                                | -                                        |
| `appdtos`         | DTOs for transferring application state                                     | -                                        |
| `appmocks`        | Mocks for testing application components                                    |                                          |

#### `./internal/app/apprepositories/` - Internal modules for storing application state

| Module               | Description                                                                 |
|----------------------|-----------------------------------------------------------------------------|
| `memoryrepository`   | Memory based volatile repository for storing application data               |
| `filerepository`     | File based persistent repository for storing application data               |

#### `./internal/app/appendpoints/` - Internal modules for REST API end-points

| Module            | Description                                              |
|-------------------|----------------------------------------------------------|
| `appendpoints`    | Application end-point implementations (the main package) |
| `indexendpoint`   | Index end-point implementation                           |

#### `./internal/common/` - Internal modules for common use cases

| Module        | Description                                                        |
|---------------|--------------------------------------------------------------------|
| `fsutils`     | Utilities handling higher level file operations                    |
| `hashutils`   | Hashing utils                                                      |
| `mainutils`   | Main utils, eg. environment handling                               |
| `commonmocks` | Mocks for testing                                                  |
| `managers`    | Managers to decouple 3rd party dependencies from application logic |

#### `./internal/common/api/` -- Internal modules for the REST API framework

| Module         | Description                           |
|----------------|---------------------------------------|
| `apidtos`      | Common REST API DTOs (like ErrorDTO)  |
| `apierrors`    | Common REST API error handlers        |
| `apimocks`     | Common REST API Mocks                 |
| `apirequests`  | Common REST API request models        |
| `apiresponses` | Common REST API response models       |
| `apiserver`    | Common REST API server implementation |
| `apitypes`     | Common REST API interfaces and types  |

### Application Design

#### Managers

In `gocertcenter` *managers* are intended to handle all outside dependencies 
which would be otherwise hard to test, like file system operations and x509 
certificate operations.

#### Controllers

| Module                             | Description                    | Child controllers                                 | Parent relations                                    |
|------------------------------------|--------------------------------|---------------------------------------------------|-----------------------------------------------------|
| `ApplicationController`            | Controls application model     | `[]OrganizationController`                        |                                                     |
| `OrganizationController`           | Controls an organization model | `[]CertificateController`                         | `ApplicationController`                             |
| `CertificateController`            | Controls a certificate model   | `[]CertificateController`, `PrivateKeyController` | `OrganizationController` or `CertificateController` |
| `PrivateKeyController`             | Controls a private key model   |                                                   | `CertificateController`                             |

#### Models

#### DTOs


## License

CMM is available under the Functional Source License, Version 1.1, MIT Future 
License (FSL-1.1-MIT), allowing use, copying, modification, and redistribution 
for permitted purposes, excluding competing uses. Two years post-release, the 
software will also be available under MIT license terms. Refer to the 
[LICENSE.md](LICENSE) file for complete license details.

## Support

For support or further inquiries, particularly regarding the SaaS offering and 
its suitability for securing private services and embedded application 
connections, please open an issue on our 
[GitHub repository](https://github.com/hyperifyio/gocertcenter/issues) or contact us 
directly at info@hg.fi.
