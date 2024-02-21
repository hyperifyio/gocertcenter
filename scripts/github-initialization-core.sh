#!/bin/bash
set -e
set -x

#gh issue create --title "TASK Project Setup"                                                                               --body "### Goal
#Initialize the project and define its structure.
#
#### Subtasks
#* [ ] Initialize the project repository.
#* [ ] Define the project structure and conventions." --label task

gh issue create --title "TASK Design Storage Schema"                                                                       --body "### Goal
Design a flexible schema for efficiently storing certificates and keys.

### Subtasks
* [ ] Define schema for certificates.
* [ ] Define schema for keys." --label task
gh issue create --title "TASK Implement In-Memory Storage Module"                                                          --body "### Goal
Develop an in-memory storage solution for rapid access and manipulation of certificate and key data during runtime.

### Subtasks
* [ ] Design in-memory data structures.
* [ ] Implement CRUD operations in memory." --label task
gh issue create --title "TASK Develop Certificate Management Service Layer"                                                --body "### Goal
Create a service layer to encapsulate the business logic for managing the lifecycle of certificates.

### Subtasks
* [ ] Design the service layer architecture.
* [ ] Implement logic for certificate generation, renewal, and revocation." --label task
gh issue create --title "TASK Process Certificate Signing Requests (CSRs)"                                                 --body "### Goal
Implement functionality to handle and process CSRs, integrating with the service layer for certificate issuance.

### Subtasks
* [ ] Design CSR processing mechanism.
* [ ] Implement CSR parsing and certificate issuance based on valid requests." --label task
gh issue create --title "TASK Setup Logging System"                                                                        --body "### Goal
Implement a comprehensive logging system for the microservice, focusing on security, auditability, and troubleshooting.

### Subtasks
* [ ] Define logging standards and formats.
* [ ] Integrate logging framework into the microservice." --label task

