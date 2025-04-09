# Unleash Automation Kit

**Unleash Automation Kit** is a collection of tools designed to automate and streamline feature flag operations
using [Unleash](https://www.getunleash.io/).

Unleash Automation Kit is a collection of tools designed to automate operational tasks around feature flags and enables infrastructure-level adoption of feature flags.

---

## 📦 Tools

| Tool                           | Status         | Description                                                                              |
|--------------------------------|----------------|------------------------------------------------------------------------------------------|
| `stale-flag-cleaner`           | ✅ Ready        | Automatically creates tasks to clean up stale feature flags in Unleash                   |
| `feature-flag-killer`          | 🚧 In progress | Disables feature flags in response to alerts from Grafana                                |
| `rollout-controller`           | 🚧 In progress | Automates canary release rollout                                                         |
| `unleash-project-configurator` | 🚧 In progress | Bootstraps Unleash projects, permissions, environments, and API tokens                   |
| `external-auth-service`        | 🚧 In progress | Implements an external authentication service for Envoy (or similar) using Unleash rules |

---

## stale-flag-cleaner

`stale-flag-cleaner` scans all stale flags in Unleash and automatically:

- Creates a ticket in your issue tracker (e.g., Jira)
- Tags the feature to prevent duplicate ticket creation

This tool is designed to be **idempotent** and safe for scheduled execution (e.g., via cron).

### Environment Setup

Copy the example environment file and configure the required values:

```bash
cp .stale-flag-cleaner-env.example .env
```

### Usage

```bash
go run cmd/stale-flag-cleaner/stale-flag-cleaner.go