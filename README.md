# Certificate Management Microservice (CMM)

## Overview

The Certificate Management Microservice (CMM) is a powerful, in-development 
solution designed to meet the specialized needs of digital certificate 
management for private services and embedded application connections. Focused 
on enhancing security for microservices and applications, CMM specializes in 
the generation, renewal, and revocation of CA, intermediate, server, and client 
certificates. Its modular and lightweight design is particularly suited for 
embedding directly into Go applications, facilitating seamless integration for 
developers.

CMM is engineered for organizations looking to secure their internal 
communications and applications with mTLS and PKI, without the 
complexities and overhead associated with public Certificate Authorities (CAs). 
This makes it an ideal choice for private services and custom software use 
cases where security, trust, and privacy are paramount.

In addition to on-premise deployment, CMM will be available as a commercial 
hosted SaaS solution at [https://cert.center](https://cert.center), offering an 
automated, managed platform for hassle-free certificate management, 
specifically tailored for internal and custom application needs.

**Project Status**: Progressing towards our first release, MVP v1, we are 
committed to delivering a robust, feature-rich service. Our development is 
focused on ensuring unparalleled security and operational efficiency. For the 
latest updates and milestones, visit 
[MVP v1 Issue](https://github.com/hyperifyio/gocertcenter/issues/1).

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
