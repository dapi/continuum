# Brownfield Quick Start

Continuum is designed to attach to an existing codebase without pretending the repository is simpler than it is.

The brownfield path is intentionally evidence-first: the agent studies the real project before Memory Bank becomes an authority.

## Goal

From an existing repository to the first governed agent task with two operator prompts.

## Prerequisites

Install the CLIs you want to use:

```sh
go install github.com/dapi/memory-bank-cli/cmd/memory-bank-cli@latest
go install github.com/dapi/start-issue/v2/cmd/start-issue@latest
go install github.com/dapi/code-converge/cmd/code-converge@latest
```

Use pinned versions in reproducible environments.

The repository should already be a Git repository. `gh` should be authenticated for issue and PR workflows.

## Prompt 1 — attach Continuum

Run your coding agent in the repository root and give it this prompt:

```text
Attach Continuum to this brownfield repository.

Before installing or consulting Memory Bank, read the existing repository instructions and inspect the actual project sources: README and docs, code, manifests, CI/CD, configuration, runbooks, historical ADRs, and other evidence available in the repository.

Create an evidence-backed ./brownfield-intake-prd.md with facts, source references, confidence, conflicts, assumptions, open questions, and owner/freshness where known. Do not invent architecture, requirements, metrics, or delivery plans.

Then install Memory Bank with memory-bank-cli, adapt product/, domain/, engineering/, and ops/ from the same evidence, convert the intake into the governed PRD, preserve existing repository instructions, and run memory-bank-cli lint and memory-bank-cli doctor.

Do not implement product changes during this adoption task. Finish with a concise report of sources inspected, canonical documents changed, unresolved gaps, and verification results.
```

This is the compressed operator form of the Memory Bank brownfield adaptation protocol. The important invariant is that generic Memory Bank content is not treated as evidence about the existing codebase.

## Prompt 2 — execute a real issue

Create or choose one real GitHub issue. Then start it:

```sh
start-issue 2841
```

The exact issue number is not important. The point is that the second operator instruction is a real unit of product or engineering work, not another setup exercise.

The agent now starts inside an isolated worktree with repository instructions and the adapted durable context available to it.

A useful issue body should describe intent and acceptance evidence, not prescribe every implementation detail.

Example:

```text
Fix duplicate settlement creation when a provider retries the same callback.

Expected behavior:
- the callback remains safe under retries;
- existing successful settlements are not duplicated;
- relevant tests cover the retry path;
- preserve current external API behavior.
```

## Converge the change

After implementation, run:

```sh
code-converge
```

The intended loop is:

```text
issue
  ↓
isolated agent execution
  ↓
implementation
  ↓
review
  ↓
fix findings
  ↓
review again
  ↓
publish
  ↓
CI
  ↓
CI recovery when applicable
  ↓
verified change
```

Generation is not the completion condition. Convergence is.

## What “attached” means

A brownfield repository is meaningfully attached to Continuum when:

- project context reflects repository evidence instead of template assumptions;
- product, domain, engineering, and operational knowledge have explicit canonical owners;
- conflicts and unknowns remain visible rather than silently guessed;
- `memory-bank-cli lint` and `doctor` pass, or remaining verification gaps are explicit;
- at least one real issue has been executed through the adapted context;
- the resulting change can be reviewed and verified through the normal repository delivery path.

## The two-prompt claim

“Two prompts” does **not** mean that a mature codebase can be understood without work.

It means the operator interface can remain small:

```text
1. Attach Continuum to this repository.
2. Execute this real issue.
```

The system — not the operator — carries the discovery, context construction, isolation, task execution, review, and verification mechanics.
