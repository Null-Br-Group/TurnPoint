# Encurtador De URL's

> **Simple URL redirector for file downloads — project under construction**

## Overview

This project is a minimalist **URL redirector** written in Go. The concept is straightforward: create **URL + ID** pairs that point to files or executables hosted elsewhere. The person who controls the server determines where each ID points, allowing the destination to be changed at any time without altering the link that has already been distributed.

There is no administrative panel, external database, or heavy dependencies. It is a single service—easy to spin up and easy to tear down.

## Objectives

- Create short URLs associated with IDs that serve as an entry point for **downloading files and executables**.
- Keep control of the destination in the hands of those who operate the service: the file can be replaced, renamed, or moved without breaking the public link.
- Log the basics of the requests (IP, User-Agent, timestamp) to determine whether and when a link was accessed.

## Ephemeral Infrastructure

The greatest asset of this redirector is its **simplicity**, which translates into operational flexibility and OPSEC (Operational Security):

1. **Stateless by design**: all redirection logic resides within the binary itself and in memory. There is no state coupled to the machine.
2. **Total mobility**: when necessary, the service can be moved to **any server and any domain/DNS** in a matter of minutes. Simply point the new DNS to the new host, and the same set of URLs will continue to work.
3. **Low-cost forensics**: with no disk persistence, a powered-off node leaves no easily accessible audit trail.
4. **Temporary redirects**: use of `302`/`307` instead of `301`, preventing proxies and browsers from caching the destination after the operation has concluded.

## Architecture

|Layer | Current implementation |
|---|---|
| HTTPserver | pure `net/http` (Go stdlib) |
| URL Repository | In-memory (`map` protected by a mutex) |
| Persistence | None (by choice) |
| Deploy | Single binary, no external dependencies |

## Project status

> [!WARNING]
> **Under Development.** The project is in an active phase of development. The API, the ID format, and the behavior of the redirector may change without notice. Do not use in production.

*Roadmap and known issues:*
- [ ] Rotation of destinations by ID at runtime (via API or configuration file)
- [ ] Structured logs (JSON) with optional rotation
- [ ] Configuration via command-line flags or a `.env` file (remove hardcoded values)
- [ ] Unit and integration tests
- [ ] Detailed usage documentation and deployment examples

## ⚖️ Notice Regarding Ethical and Responsible Use

This software was developed **exclusively for security research, testing in controlled and/or authorized environments** (such as authorized *Red Team* laboratories or infrastructure studies).

Using this tool to obfuscate malicious traffic, direct users to harmful content, or carry out activities without explicit authorization is illegal and violates cybercrime laws such as the United States Computer Fraud and Abuse Act (18 U.S.C. § 1030) and Brazil's Law 12.737/2012 (Digital Crimes Law). The author assumes no responsibility for the misuse of this code.
