#!/bin/bash
set -e
set -x


#gh issue create --title "TASK Implement Authentication System"                                                             --body "### Goal
#Develop an authentication mechanism for the REST API, considering security and usability.
#
#### Subtasks
#* [ ] Design the authentication flow.
#* [ ] Implement mTLS based authentication." --label task
#gh issue create --title "TASK Implement TLS for REST API"                                                                  --body "### Goal
#Configure TLS for the API server.
#
#### Subtasks
#* [ ] Configure TLS for the API server." --label task
#gh issue create --title "TASK Design REST API Interface"                                                                   --body "### Goal
#Design the REST API interface for the microservice, including endpoint definitions and JSON schema designs.
#
#### Subtasks
#* [ ] Define RESTful endpoints.
#* [ ] Design the JSON schema for requests and responses.
#* [ ] Document the API endpoints using OpenAPI/Swagger." --label task
#gh issue create --title "TASK POST /ca - Create CA Certificate"                                                            --body "### Goal
#Implement the endpoint for CA certificate generation.
#
#### Subtasks
#* [ ] Implement endpoint for CA certificate generation." --label task
#gh issue create --title "TASK GET /ca/{SerialNumber} - Read CA Certificate Details"                                        --body "### Goal
#Implement the endpoint to retrieve CA certificate details as JSON.
#
#### Subtasks
#* [ ] Implement endpoint to retrieve CA certificate details as JSON." --label task
#gh issue create --title "TASK GET /ca - List CA Certificates"                                                              --body "### Goal
#Implement the endpoint to list all CA certificates.
#
#### Subtasks
#* [ ] Implement endpoint to list all CA certificates." --label task
#gh issue create --title "TASK POST /ca/{SerialNumber}/renew - Renew CA Certificate"                                        --body "### Goal
#Implement the endpoint for CA certificate renewal.
#
#### Subtasks
#* [ ] Implement endpoint for CA certificate renewal." --label task
#gh issue create --title "TASK DELETE /ca/{SerialNumber} - Revoke CA Certificate"                                           --body "### Goal
#Implement the endpoint for CA certificate revocation.
#
#### Subtasks
#* [ ] Implement endpoint for CA certificate revocation." --label task
#gh issue create --title "TASK POST /ca/{SerialNumber}/intermediate - Create Intermediate Certificate"                      --body "### Goal
#Implement the endpoint for intermediate certificate generation and signing.
#
#### Subtasks
#* [ ] Implement endpoint for intermediate certificate generation and signing." --label task
#gh issue create --title "TASK GET /ca/{SerialNumber}/intermediate/{SerialNumber} - Read Intermediate Certificate Details"  --body "### Goal
#Implement the endpoint to retrieve intermediate certificate details as JSON.
#
#### Subtasks
#* [ ] Implement endpoint to retrieve intermediate certificate details as JSON." --label task
#gh issue create --title "TASK GET /ca/{SerialNumber}/intermediate - List Intermediate Certificates"                        --body "### Goal
#Implement the endpoint to list all intermediate certificates.
#
#### Subtasks
#* [ ] Implement endpoint to list all intermediate certificates." --label task
#gh issue create --title "TASK POST /ca/{SerialNumber}/intermediate/{SerialNumber}/renew - Renew Intermediate Certificate"  --body "### Goal
#Implement the endpoint for intermediate certificate renewal.
#
#### Subtasks
#* [ ] Implement endpoint for intermediate certificate renewal." --label task
#gh issue create --title "TASK DELETE /ca/{SerialNumber}/intermediate/{SerialNumber} - Revoke Intermediate Certificate"     --body "### Goal
#Implement the endpoint for intermediate certificate revocation.
#
#### Subtasks
#* [ ] Implement endpoint for intermediate certificate revocation." --label task
#gh issue create --title "TASK POST /ca/{SerialNumber}/server - Create Server Certificate"                                  --body "### Goal
#Implement the endpoint for server certificate generation.
#
#### Subtasks
#* [ ] Implement endpoint for server certificate generation." --label task
#gh issue create --title "TASK GET /ca/{SerialNumber}/server/{SerialNumber} - Read Server Certificate Details"              --body "### Goal
#Implement the endpoint to retrieve server certificate details as JSON.
#
#### Subtasks
#* [ ] Implement endpoint to retrieve server certificate details as JSON." --label task
#gh issue create --title "TASK GET /ca/{SerialNumber}/server - List Server Certificates"                                    --body "### Goal
#Implement the endpoint to list all server certificates.
#
#### Subtasks
#* [ ] Implement endpoint to list all server certificates." --label task
#gh issue create --title "TASK POST /ca/{SerialNumber}/server/{SerialNumber}/renew - Renew Server Certificate"              --body "### Goal
#Implement the endpoint for server certificate renewal.
#
#### Subtasks
#* [ ] Implement endpoint for server certificate renewal." --label task
#gh issue create --title "TASK DELETE /ca/{SerialNumber}/server/{SerialNumber} - Revoke Server Certificate"                 --body "### Goal
#Implement the endpoint for server certificate revocation.
#
#### Subtasks
#* [ ] Implement endpoint for server certificate revocation." --label task
#gh issue create --title "TASK POST /ca/{SerialNumber}/client - Create Client Certificate"                                  --body "### Goal
#Implement the endpoint for client certificate generation.
#
#### Subtasks
#* [ ] Implement endpoint for client certificate generation." --label task
#gh issue create --title "TASK GET /ca/{SerialNumber}/client/{SerialNumber} - Read Client Certificate Details"              --body "### Goal
#Implement the endpoint to retrieve client certificate details as JSON.
#
#### Subtasks
#* [ ] Implement endpoint to retrieve client certificate details as JSON." --label task

gh issue create --title "TASK GET /ca/{SerialNumber}/client - List Client Certificates"                                    --body "### Goal
Implement the endpoint to list all client certificates.

### Subtasks
* [ ] Implement endpoint to list all client certificates." --label task

gh issue create --title "TASK POST /ca/{SerialNumber}/client/{SerialNumber}/renew - Renew Client Certificate"              --body "### Goal
Implement the endpoint for client certificate renewal.

### Subtasks
* [ ] Implement endpoint for client certificate renewal." --label task
gh issue create --title "TASK DELETE /ca/{SerialNumber}/client/{SerialNumber} - Revoke Client Certificate"                 --body "### Goal
Implement the endpoint for client certificate revocation.

### Subtasks
* [ ] Implement endpoint for client certificate revocation." --label task

