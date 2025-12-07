# Policy Catalog Requirements and Issues

## Requirements

This repository serves as a centralized catalog for various policies, including rate-limiter and jwt-validator. Each policy has multiple versions, and consumers should be able to download and use specific versions of individual policies.

### Key Requirements:
- **Single Repository**: All policies are hosted in one Git repository for simplicity.
- **Versioned Policies**: Each policy version is a separate Go module.
- **Exact Version Download**: Consumers can download a specific policy version exactly, e.g., jwt-validator v2.1.0.
- **Go Module Compatibility**: Policies are distributed as Go modules, allowing `go get` to fetch specific versions.

## Expected Behavior

- A consumer project can import and use a specific policy version by running `go get <module-path>@<version>`.
- For example:
  ```
  go get github.com/DakshithaS/test-policy-hub/jwt-validator/v2@v2.1.0
  ```
- This should download only the jwt-validator v2.1.0 module without fetching the entire repository or other policies.
- The imported module provides the policy functionality, and the consumer can instantiate and use it.

## Current Implementation

- Policies are organized in subfolders: `rate-limiter/` and `jwt-validator/`.
- Each version is in its own subfolder, e.g., `jwt-validator/v2/` for v2.1.0.
- Each version folder contains a `go.mod` file defining the module, e.g., `module github.com/DakshithaS/test-policy-hub/jwt-validator/v2`.
- Versions are tagged with prefixed names, e.g., `jwt-validator/v2.1.0`.

## Directory Structure

```
.
├── README.md
├── REQUIREMENTS.md
├── rate-limiter
│   └── v1
│       ├── go.mod
│       ├── rate_limiter.go
│       ├── policy-definition.yaml
│       ├── metadata.json
│       └── README.md
└── jwt-validator
    ├── v1
    │   ├── go.mod
    │   ├── jwt_validator.go
    │   ├── policy-definition.yaml
    │   ├── metadata.json
    │   └── README.md
    └── v2
        ├── go.mod
        ├── jwt_validator.go
        ├── policy-definition.yaml
        ├── metadata.json
        └── README.md
```

## Issues

### Primary Issue: Go Module Versioning with Prefixed Tags
- Go's module system expects version tags to be semantic version strings (e.g., `v2.1.0`), not prefixed (e.g., `jwt-validator/v2.1.0`).
- When attempting `go get github.com/DakshithaS/test-policy-hub/jwt-validator/v2@v2.1.0`, Go fails because it cannot find a valid semver tag.
- Error: `invalid version: missing github.com/DakshithaS/test-policy-hub/jwt-validator/go.mod and .../v2/go.mod at revision jwt-validator/v2.1.0`

### Workaround: Using Commit Hashes
- To achieve the expected behavior, consumers must use commit hashes instead of version tags.
- Example: `go get github.com/DakshithaS/test-policy-hub/jwt-validator/v2@4a7dc7c`
- This works because Go fetches the repository at the specific commit and locates the sub-module.
- However, this is not user-friendly, as commit hashes are not semantic and hard to remember.

### How to Use Specific Versions
To download a specific policy version:
1. Find the commit hash for the desired tag:
   ```
   git rev-parse <policy>/<version>
   ```
   For example:
   - jwt-validator v2.1.0: `git rev-parse jwt-validator/v2.1.0` → `4a7dc7c6dd55d62b8112bedb2c871db9981eb728`
   - jwt-validator v1.1.0: `git rev-parse jwt-validator/v1.1.0` → `767fa3f2f813e55ba79b93bd3263d047dbb7d6fd`
   - rate-limiter v1.0.0: `git rev-parse rate-limiter/v1.0.0` → `4a7dc7c6dd55d62b8112bedb2c871db9981eb728`

2. Use the commit hash in `go get`:
   ```
   go get github.com/DakshithaS/test-policy-hub/<policy>/<major-version>@<commit-hash>
   ```
   For example:
   ```
   go get github.com/DakshithaS/test-policy-hub/jwt-validator/v2@4a7dc7c
   ```

This fetches only the specific policy module at the exact version, achieving the expected behavior within the constraints of a single repository.

### Limitations
- Cannot use semantic versions directly due to tag prefixing required to avoid conflicts between policies in the same repo.
- Separate repositories per policy would allow clean semver tagging (e.g., `v2.1.0` for each), but the requirement is to use a single repo.
- Go does not support prefixed semver tags for sub-modules in a single repository.

## Conclusion

The current setup achieves the expected behavior using commit hashes for exact version downloads. While not as user-friendly as semantic versions, it allows consumers to fetch specific policy versions without downloading unnecessary parts of the repository. For better usability, consider separate repositories per policy.