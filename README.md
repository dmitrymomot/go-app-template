# Go App Template

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/dmitrymomot/go-app-template)](https://github.com/dmitrymomot/go-app-template)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/go-app-template.svg)](https://pkg.go.dev/github.com/dmitrymomot/go-app-template)
[![License](https://img.shields.io/github/license/dmitrymomot/go-app-template)](https://github.com/dmitrymomot/go-app-template/blob/main/LICENSE)

[![Tests](https://github.com/dmitrymomot/go-app-template/actions/workflows/tests.yml/badge.svg)](https://github.com/dmitrymomot/go-app-template/actions/workflows/tests.yml)
[![CodeQL Analysis](https://github.com/dmitrymomot/go-app-template/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dmitrymomot/go-app-template/actions/workflows/codeql-analysis.yml)
[![GolangCI Lint](https://github.com/dmitrymomot/go-app-template/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/dmitrymomot/go-app-template/actions/workflows/golangci-lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/go-app-template)](https://goreportcard.com/report/github.com/dmitrymomot/go-app-template)

A **full-stack** app template based on Golang and featuring TailwindCSS for styling, HTMX for dynamic UIs, Postgres/SQLite/TursoDB (libSQL) for efficient data management, and Templ for component-style templating  (similar to ReactJs), streamlining web development workflows.

## Motivaton

The main goal of this project is to provide a simple and efficient way to build web applications using Golang. The project is designed to be a starting point for building web applications, providing a basic structure and a set of tools to help you get started.

### Why HTMX?

[HTMX](https://htmx.org) stands out by simplifying web development compared to JavaScript frameworks like React, Angular, Vue, or Svelte. Its main advantage lies in its ease of use and integration, allowing developers to add interactive features using familiar HTML, without extensive JavaScript expertise. This results in quicker development times and potentially better performance due to less code being sent to the browser. HTMX's server-centric approach also simplifies state management, making it an efficient choice for enhancing traditional multi-page applications with dynamic elements, without the complexity of full-scale SPA frameworks.

## Features / Roadmap

- [x] [Go-chi](https://go-chi.io/#/) router with pre-configured middleware and graceful shutdown
- [x] [TailwindCSS](https://tailwindcss.com) for styling
- [x] [HTMX](https://htmx.org) for dynamic UIs
- [x] Postgres/SQLite/TursoDB (libSQL) for efficient data management
- [x] [Templ](https://templ.guide) for component-style templating (similar to ReactJs)
- [x] Session management (using [alexedwards/scs](https://github.com/alexedwards/scs))
- [x] Hot-reloading for development (using [air](https://github.com/cosmtrek/air))
- [x] Static file serving
- [x] Jobs and workers for background tasks (using [asyncer](https://github.com/dmitrymomot/asyncer))
- [x] Cron jobs for scheduled tasks (using [asyncer](https://github.com/dmitrymomot/asyncer) too)
- [x] Migration tool for database schema changes (using [rubenv/sql-migrate](https://github.com/rubenv/sql-migrate))
- [x] [SQLC](https://sqlc.dev) for type-safe SQL queries
- [ ] Websockets for real-time communication (using [melody](https://github.com/olahol/melody))
- [ ] Sub-domain routing
- [-] Multi-language support (using [i18n](https://github.com/nicksnyder/go-i18n))
- [ ] File uploads and storage in the s3-compatible object storage
- [ ] Testing and benchmarking
- [ ] Continuous integration and deployment to DigitalOcean App Platform
- [ ] Monitoring and logging via [betterstack](https://betterstack.com/)
- [ ] Authentication and authorization via Email/Password, OAuth2 (Google, Facebook, Twitter, GitHub, etc.).
- [ ] Forgot password and email verification
- [ ] Two-factor authentication


## Getting Started

### Prerequisites

// TODO: Add prerequisites

### Installation

// TODO: Add installation

### Usage

// TODO: Add usage

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
