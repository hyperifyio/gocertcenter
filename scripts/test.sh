#!/bin/bash
set -e
set -x

gh issue create --title "TASK Project Setup"                                                                               --body "### Goal
Initialize the project and define its structure.

### Subtasks
* [ ] Initialize the project repository.
* [ ] Define the project structure and conventions." --label task
