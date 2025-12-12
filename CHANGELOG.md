# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.0.0] - 2025-12-12

### Added

*   Core tracing functionality with a `Trace` type for storing trace data.
*   Context-aware tracing with `GetFromContext` and `SetInContext` functions.
*   HTTP client and server tracing with `SetHTTPHeaders` and `GetFromHTTPRequest` functions.
*   JSON serialization for traces with a custom `MarshalJSON` method.
