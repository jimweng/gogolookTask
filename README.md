# virtualFileSystem

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Overview](#overview)
    - [To Start Application](#to-start-application)
- [Architecture](#architecture)
- [Endpoints](#endpoints)
  - [GET `/tasks`](#get-tasks)
  - [GET `/tasks/{id}`](#get-tasksid)
  - [POST `/tasks`](#post-tasks)
  - [PUT `/tasks/{id}`](#put-tasksid)
  - [DELETE `/tasks/{id}`](#delete-tasksid)
  - [To View Your Endpoints Locally](#to-view-your-endpoints-locally)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Overview

This application provides a set of instructions that allow for some type of operation to implement a task management api on local machine.

The root folder contains the Go source files that comprises the application. The `Dockerfile` in root folder can be used to generate the proper working environment.

This application use `golangci.yml` and `pre-commit-config.yaml` to align the coding style, which is convenient to be reuse when it's necessary to run cicd process.

#### To Start Application
1. In your terminal, start the server: `go run cmd/main.go`

## Architecture

The task management api is using hexagonal architecture. The api is separated to Service domain and Repository domain. The database can be easily replace by implementing file system interface for other database system without huge code refactoring.

## Endpoints

### GET `/tasks`
  - Description: get all tasks in the system
  - Response
    - List of all tasks
    - Error: reading file

### GET `/tasks/{id}`
  - Description: get the task matched the id in the system
  - Response
    - The details of the task
    - Error: failed to read file
    - Error: not found

### POST `/tasks`
  - Description: Add task into the system
  - Body
  ```json
  {
    "name": "taskName",
    "status": 0 // should be 0 or 1
  }
  ```
  - Response
    - uuid of the task
    - Error: failed to read file
    - Error: the taskname is existed
    - Error: failed to save file

### PUT `/tasks/{id}`
  - Description: Update the existing task
  - Body
  ```json
  {
    "name": "newTaskName",
    "status": 0 // should be 0 or 1
  }
  ```
  - Response
    - return statusOK with nothing in payload if success
    - Error: failed to read file
    - Error: id not found


### DELETE `/tasks/{id}`
  - Description: Delete the task in {id}
  - Response:
    - return statusOK with nothing in payload if success
    - Error: failed to read file

### To View Your Endpoints Locally

1. In your terminal, start the server: `go run ./cmd/main.go`
2. Click [Here](http://localhost:8080/docs/index.html)
