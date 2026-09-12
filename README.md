# Go Webhook Inbox

A small Go service for receiving, validating, and storing webhook events while preventing duplicate records.

> Status: Work in progress. The features below are planned and are not yet implemented.

## Problem

External systems may send the same webhook event more than once when retrying delivery after a timeout or connection failure.

This service will store each event once and acknowledge repeated deliveries without creating duplicate records.

## Planned behavior

- Receive events through an HTTP API.
- Validate required fields.
- Store events in PostgreSQL with their reception time.
- Identify duplicates using the combination of source and event ID.
- Allow stored events to be retrieved.
- Preserve stored data across application restarts.

## Scope

The first version receives, stores, and exposes events for retrieval. Outbound notifications, message queues, and background processing are outside its scope.

## Planned technologies

- Go and the standard `net/http` package
- PostgreSQL
- Go unit and integration tests
- Docker Compose for local development

## Development approach

The project follows test-driven development: write a failing test, implement the behavior, then refactor while keeping the tests passing.

## Progress

- [x] Create the repository
- [ ] Implement request validation
- [ ] Add event storage and retrieval
- [ ] Handle duplicate and concurrent deliveries
- [ ] Add integration tests
- [ ] Document local setup and usage
