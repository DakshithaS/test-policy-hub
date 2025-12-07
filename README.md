# Versioned Policy Catalog

This repository implements a versioned policy catalog for experimenting with policies such as rate-limiter and jwt-validator.

## Overview

Each policy (e.g., `rate-limiter`, `jwt-validator`) has multiple versions. Each policy version contains:
- A Go module under `/src` (with `go.mod` and sample code)
- A `policy-definition.yaml` (policy schema/configuration definition)
- A `metadata.json` (policy metadata for UI/catalog)
- An optional `README.md` per version

The `/src` directory is the Go module root. Versioning is done by Git tags per policy, not by Go module paths. Each policy's `/src` folder is its own independent Go module.

## Directory Structure

```
.
├── README.md
└── policies
    ├── rate-limiter
    │   ├── v1.0.0
    │   │   ├── src
    │   │   │   ├── go.mod
    │   │   │   └── rate_limiter.go
    │   │   ├── policy-definition.yaml
    │   │   ├── metadata.json
    │   │   └── README.md
    │   ├── v1.1.0
    │   │   ├── src
    │   │   │   ├── go.mod
    │   │   │   └── rate_limiter.go
    │   │   ├── policy-definition.yaml
    │   │   ├── metadata.json
    │   │   └── README.md
    │   └── v2.0.0
    │       ├── src
    │       │   ├── go.mod
    │       │   └── rate_limiter.go
    │       ├── policy-definition.yaml
    │       ├── metadata.json
    │       └── README.md
    └── jwt-validator
        ├── v1.0.0
        │   ├── src
        │   │   ├── go.mod
        │   │   └── jwt_validator.go
        │   ├── policy-definition.yaml
        │   ├── metadata.json
        │   └── README.md
        └── v2.1.0
            ├── src
            │   ├── go.mod
            │   └── jwt_validator.go
            ├── policy-definition.yaml
            ├── metadata.json
            └── README.md
```

## Usage

To use a policy, navigate to its version directory and build the Go module.

For example:
```bash
cd policies/rate-limiter/v1.0.0/src
go build
```

Versioning is managed via Git tags, e.g., `rate-limiter/v1.0.0`.