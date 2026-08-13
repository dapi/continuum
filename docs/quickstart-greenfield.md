# Greenfield Quick Start

Continuum can be attached at the beginning of a project so the product, domain, engineering rules, operational constraints, and delivery evidence become durable before the codebase grows around disposable agent sessions.

## Goal

From a new Git repository to the first governed agent task with two operator prompts.

## Prerequisites

Install the CLIs you want to use:

```sh
go install github.com/dapi/memory-bank-cli/cmd/memory-bank-cli@latest
go install github.com/dapi/start-issue/v2/cmd/start-issue@latest
go install github.com/dapi/code-converge/cmd/code-converge@latest
```

Use pinned versions in reproducible environments.

Initialize the Git repository and put the initial product brief, README, or other project evidence into it. `gh` should be authenticated for issue and PR workflows.

## Prompt 1 — attach Continuum

Run your coding agent in the repository root and give it this prompt:

```text
Attach Continuum to this greenfield repository.

Read the existing repository instructions, README, product brief, docs, configuration, manifests, CI, and code that already exist. Treat repository sources as authoritative for project-specific facts and do not invent missing requirements, users, metrics, domain rules, architecture, or operational procedures.

Install Memory Bank with memory-bank-cli if it is not already present. Adapt the canonical product/, domain/, engineering/, and ops/ owners to the confirmed project facts. Add stable use cases and ADRs only when they are supported by existing evidence.

Create the initial PRD in memory-bank/prd/ using the project identifier or PRD-001. Keep unknowns and conflicts as explicit open questions rather than guessing. Update navigation and derived_from links, then run memory-bank-cli lint and memory-bank-cli doctor.

Do not implement product features during this adoption task. Finish with a concise report of sources used, canonical documents changed, the created PRD, unresolved questions, and verification results.
```

## Prompt 2 — execute the first real issue

Create the first real GitHub issue from the product work that should happen next, then run:

```sh
start-issue 1
```

Example issue:

```text
Implement customer sign-in with email and one-time code.

Expected behavior:
- a customer can request a one-time code by email;
- a valid code creates an authenticated session;
- expired or reused codes are rejected;
- tests cover the successful and rejected flows.
```

The issue is the execution unit. Memory Bank remains the durable source for project context, product intent, domain language, engineering constraints, and verification contracts.

## Converge the change

After implementation, run:

```sh
code-converge
```

The development loop becomes:

```text
intent
  ↓
durable project context
  ↓
GitHub issue
  ↓
isolated worktree + agent
  ↓
implementation
  ↓
review / fix loop
  ↓
CI
  ↓
verified delivery
  ↓
new durable knowledge
```

## What “attached” means

A greenfield repository is meaningfully attached to Continuum when:

- initial project intent has canonical owners instead of being trapped in chat history;
- an initial PRD records goals, non-goals, product scope, rules, risks, metrics where known, and open questions;
- architecture and operational facts are documented only to the extent they are actually decided;
- navigation and governance validate;
- the first real issue can be executed in an isolated agent workspace and delivered through review and CI.

## The two-prompt claim

The operator surface is intentionally small:

```text
1. Attach Continuum to this project.
2. Execute this issue.
```

The durable context, task isolation, agent execution, review, fixes, and verification happen inside the development system rather than being manually reconstructed for each session.
