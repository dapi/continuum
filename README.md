# Continuum

**Continuous Agentic Development Platform**

Continuum is a practical operating layer for long-running software development with coding agents.

It connects durable project context, issue-driven execution, isolated agent workspaces, iterative review/fix loops, CI recovery, and verified delivery into one continuous development system.

The goal is simple:

> **Autonomy without losing control.**

## Why Continuum

Coding agents are already good at producing code. The harder problem is operating them reliably over the lifetime of a real codebase:

- requirements drift or disappear across sessions;
- architectural intent is reconstructed from scratch;
- different agents make inconsistent local decisions;
- humans spend time re-priming context and supervising execution;
- generated changes are treated as “done” before review and CI converge;
- brownfield systems accumulate agent-induced entropy instead of becoming easier to maintain.

Continuum treats agentic development as a software lifecycle, not a sequence of prompts.

```text
Business / Product intent
          ↓
      GitHub Issue
          ↓
┌──────────────────────────────┐
│         Memory Bank          │
│ context · requirements       │
│ domain · architecture        │
│ governance · verification    │
└──────────────┬───────────────┘
               ↓
          start-issue
     issue → branch → worktree
               ↓
          Coding Agent
               ↓
         code-converge
 review → fix → review → CI → fix
               ↓
          Verified change
               ↓
    durable knowledge returns
          to Memory Bank
```

## Components

### [Memory Bank](https://github.com/dapi/memory-bank)

A durable, version-controlled context and governance layer for coding agents.

It keeps product intent, domain knowledge, engineering rules, operational constraints, requirements, architecture decisions, lifecycle rules, and verification contracts next to the code.

Memory Bank is the **control plane** of Continuum.

### [Memory Bank CLI](https://github.com/dapi/memory-bank-cli)

Installs, updates, validates, diagnoses, and synchronizes Memory Bank safely across projects.

It handles ownership, managed template evolution, local customization, drift detection, and reproducible adoption.

Memory Bank CLI is the **distribution and lifecycle layer**.

### [start-issue](https://github.com/dapi/start-issue)

Turns a GitHub issue into a dedicated branch, git worktree, and coding-agent session.

It converts task context into an isolated execution environment and supports multiple coding agents plus explicit human-gate workflows.

`start-issue` is the **execution bootstrap layer**.

### [code-converge](https://github.com/dapi/code-converge)

Closes the development loop by repeatedly reviewing the current change, fixing findings, publishing the result, checking CI, recovering failed CI, and failing closed when convergence is not reached.

`code-converge` is the **quality and convergence layer**.

## The Continuum loop

```text
Intent
  ↓
Persistent context
  ↓
Task
  ↓
Isolated agent execution
  ↓
Implementation
  ↓
Review
  ↓
Fix
  ↓
CI
  ↓
Verified delivery
  ↓
New durable knowledge
  └──────────────→ next task
```

A coding session may be ephemeral. The development system is not.

## Brownfield first

Continuum is designed to attach to existing software, not only greenfield demos.

A mature repository can adopt the system incrementally:

1. install Memory Bank;
2. adapt durable context to the real codebase;
3. route new work through GitHub issues;
4. launch agent work in isolated worktrees;
5. converge implementation through review and CI;
6. preserve newly discovered requirements and decisions as durable project knowledge.

The same model works for new projects, but the differentiator is reliability on long-lived, high-context codebases.

## What Continuum is not

Continuum is not another coding model, IDE, chat UI, or autonomous-agent demo.

It is the layer around coding agents that keeps development coherent over time.

```text
Models / coding agents
Claude · Codex · others
          │
          ▼
──────────────────────────────
          CONTINUUM
──────────────────────────────
Context       Governance
Task routing  Isolation
Execution     Verification
Convergence   Durable memory
──────────────────────────────
          │
          ▼
Git · GitHub · CI · Production
```

## Positioning

**For developers**

> Issue in. Verified PR out.

**For engineering teams**

> Stop babysitting coding agents.

**For long-lived products**

> Your codebase needs memory.

**Category**

> Continuous Agentic Software Development.

See [Architecture](docs/architecture.md) and [Positioning](docs/positioning.md).
