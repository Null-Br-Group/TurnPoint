# TurnPoint

>> A lightweight URL redirection service written in Go.

**TurnPoint** is an independent security engineering project by **Null-Br Group**.

Its purpose is to provide an identifier-based redirection mechanism: a short URL points to a destination managed by the service operator. This makes it possible to change the destination associated with an identifier without having to redistribute the original link.

The project is being developed with simplicity, portability, and operational flexibility in mind.

## Overview

TurnPoint is designed as a lightweight HTTP service with a small footprint and minimal external dependencies.

The core concept is simple:

1. Register a destination URL.
2. Associate it with a unique identifier.
3. Distribute the generated redirection link.
4. Manage the destination associated with that identifier.

The intended use cases include controlled file delivery, security research, and authorized Red Team infrastructure.

## Design Goals

- **Lightweight**: keep the service simple and minimize external dependencies.
- Portable: make deployment straightforward across different environments.
- **Flexible redirection**: allow destinations to be changed without replacing previously distributed links.
- **Operational visibility**: provide useful request information, such as source IP, User-Agent, and timestamps.
- **Controlled deployment**: support temporary and disposable infrastructure when appropriate.

These are design goals; not all of them are implemented in the current prototype.

## Architecture

The project is written in Go and uses the standard "net/http" package for HTTP handling.

The intended architecture is deliberately small:

Component| Purpose
HTTP server| Handle registration and redirection requests
URL repository| Associate identifiers with destination URLs
Redirection handler| Resolve identifiers and redirect requests
Request logging| Record relevant access information

The current prototype is still evolving. Its internal interfaces and implementation details may change as the project develops.

## Project Status

«[!WARNING]
Early development — not production-ready.

The current repository contains an incomplete prototype. Core behavior, API contracts, configuration, and internal components are still being developed and validated.»

### Roadmap

- [ ] Establish a consistent URL registration and redirection API
- [ ] Complete destination lookup and identifier management
- [ ] Implement safe, concurrency-aware in-memory storage
- [ ] Support destination updates at runtime
- [ ] Add configurable server address and public base URL
- [ ] Implement structured request logging
- [ ] Add input validation and request handling safeguards
- [ ] Add unit and integration tests
- [ ] Document usage, configuration, and deployment

## Responsible Use

TurnPoint is intended for security research, controlled testing, and explicitly authorized security operations.

The project must not be used to distribute harmful content, deceive users, or facilitate unauthorized activity.

Operators are responsible for ensuring that their use of the software complies with applicable laws, regulations, and authorization requirements.

## About Null-Br Group

TurnPoint is developed under Null-Br Group, an independent security research and development collective focused on offensive security, vulnerability research, reverse engineering, and systems engineering.

Our work is driven by research, not visibility.

**Research. Engineering. Exploration.**