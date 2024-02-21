#!/bin/bash
set -e
set -x






gh issue create --title "TASK CLI: Setup CLI Framework"                            --body "### Goal
Select and set up a Go-based framework suitable for rendering simple CLIs.

### Subtasks
* [ ] Evaluate Go CLI frameworks for simplicity and extensibility.
* [ ] Integrate chosen framework into the microservice." --label task
gh issue create --title "TASK CLI: Design Layout and Common Components"            --body "### Goal
Design a basic layout and common components for the CLI, ensuring consistency across different commands.

### Subtasks
* [ ] Define a simple and intuitive layout.
* [ ] Implement common UI components shared across commands." --label task
gh issue create --title "TASK CLI: Implement mTLS Authentication for CLI"          --body "### Goal
Ensure the CLI is accessible only with valid certificate and keys, leveraging client certificates for authentication.

### Subtasks
* [ ] Implement client certificate validation." --label task
gh issue create --title "TASK CLI: Implement \"List Certificates\" Command"        --body "### Goal
Create a view for listing all certificates with detailed information and options for renewal or revocation.

### Subtasks
* [ ] Design the \"List Certificates\" page layout.
* [ ] Fetch and display certificate data from the backend." --label task
gh issue create --title "TASK CLI: Implement \"Certificate Details\" Command"      --body "### Goal
Develop a detailed view for individual certificates, showing comprehensive information and actions such as renew or revoke.

### Subtasks
* [ ] Design the \"Certificate Details\" page layout.
* [ ] Implement fetching and displaying detailed information for a selected certificate." --label task
gh issue create --title "TASK CLI: Implement \"List Keys\" Command"                --body "### Goal
Create a view for listing all keys with options to view details and delete keys if no longer needed.

### Subtasks
* [ ] Design the \"List Keys\" page layout.
* [ ] Implement functionality to list and manage keys." --label task
gh issue create --title "TASK CLI: Implement \"Generate Certificate\" Command"     --body "### Goal
Develop a form enabling users to generate new certificates, specifying the type and other relevant details.

### Subtasks
* [ ] Design the \"Generate Certificate\" form layout.
* [ ] Implement form handling and certificate generation logic." --label task
gh issue create --title "TASK CLI: Implement \"Generate Key\" Command"             --body "### Goal
Provide a simple interface for generating new keys, including options for key type and usage.

### Subtasks
* [ ] Design the \"Generate Key\" page layout.
* [ ] Implement key generation functionality on the backend." --label task
gh issue create --title "TASK CLI: Implement \"Renew Certificate\" Functionality"  --body "### Goal
Create a mechanism within the CLI allowing users to initiate certificate renewal processes.

### Subtasks
* [ ] Design UI elements for initiating renewal.
* [ ] Integrate with backend renewal services." --label task
gh issue create --title "TASK CLI: Implement \"Revoke Certificate\" Functionality" --body "### Goal
Enable certificate revocation directly from the CLI, providing a simple workflow for users.

### Subtasks
* [ ] Design UI elements for certificate revocation.
* [ ] Integrate with backend revocation services." --label task
gh issue create --title "TASK CLI: Testing and Validation of CLI"                  --body "### Goal
Conduct thorough testing of the CLI to ensure functionality, usability, and security.

### Subtasks
* [ ] Perform functional testing of each view and action.
* [ ] Validate mTLS setup and access controls." --label task


