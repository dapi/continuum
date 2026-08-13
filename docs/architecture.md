# Continuum Architecture

Continuum is a composition of independent tools with one shared invariant: **project intent must survive individual agent sessions**.

## System model

```text
┌───────────────────────────────────────────────────────────────┐
│                    BUSINESS / PRODUCT                         │
│  Requirements · incidents · bugs · research · features       │
└──────────────────────────────┬────────────────────────────────┘
                               │
                               ▼
                         GitHub Issue
                               │
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                        MEMORY BANK                            │
│                                                               │
│  product/      domain/        engineering/        ops/        │
│  prd/          epics/         use-cases/          adr/        │
│  flows/        governance     verification        SSoT        │
└──────────────────────────────┬────────────────────────────────┘
                               │ context + rules
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                        START-ISSUE                            │
│                                                               │
│  issue → branch → worktree → initialization → agent session  │
└──────────────────────────────┬────────────────────────────────┘
                               │
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                       CODING AGENT                            │
│                                                               │
│     implementation · tests · docs · local verification       │
└──────────────────────────────┬────────────────────────────────┘
                               │
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                       CODE-CONVERGE                           │
│                                                               │
│ review → findings → fix → review → publish → CI → recovery   │
└──────────────────────────────┬────────────────────────────────┘
                               │
                               ▼
                         Verified change
                               │
                               ▼
┌───────────────────────────────────────────────────────────────┐
│                       DURABLE LEARNING                         │
│                                                               │
│ Requirements · decisions · contracts · runbooks · rules      │
│                     return to Memory Bank                     │
└───────────────────────────────────────────────────────────────┘
```

## Layer responsibilities

### 1. Context and governance — Memory Bank

Memory Bank owns durable development knowledge rather than implementation details.

Typical ownership:

- product intent;
- domain language and invariants;
- requirements and use cases;
- architectural rationale;
- engineering constraints;
- operational rules;
- lifecycle gates;
- verification contracts.

Code remains the source of truth for implementation.

This separation is critical: agents can inspect code to understand *what exists*, while Memory Bank preserves *why it exists and what must remain true*.

### 2. Distribution and adoption — Memory Bank CLI

The CLI makes governance operational rather than aspirational.

Responsibilities include:

- installing the upstream Memory Bank template;
- preserving local project ownership;
- updating managed files safely;
- detecting drift and broken adoption;
- validating navigation and governance;
- supporting reproducible versioned rollout.

This enables Continuum to evolve across many downstream repositories without treating each project as a forked documentation island.

### 3. Task isolation and execution bootstrap — start-issue

`start-issue` turns an issue into a reproducible execution unit.

The key abstraction is not “launch an agent.” It is:

```text
Task identity
    +
Repository state
    +
Dedicated branch
    +
Isolated worktree
    +
Project initialization
    +
Agent prompt/context
```

This gives each task a clean workspace while allowing many agent tasks to run concurrently against the same repository.

### 4. Convergence and quality — code-converge

`code-converge` defines completion as convergence, not generation.

A change is not complete merely because an agent stopped editing files.

The convergence loop is:

```text
Review
  ↓
Findings?
 ├─ yes → Fix → Review
 └─ no  → Publish → CI
                    ↓
                 CI failed?
                  ├─ yes → Fix CI → Review
                  └─ no  → Done
```

Important properties:

- bounded fix cycles;
- explicit CI recovery budget;
- structured review output;
- fail-closed semantics;
- exact published SHA verification;
- host-controlled publication rather than allowing the agent to define success.

## Cross-cutting invariants

Continuum should preserve these invariants as the platform evolves.

### Durable context beats session memory

A fresh agent session must be able to resume work without reconstructing the project from old chats.

### Intent and implementation have different owners

Requirements, rationale, contracts, and constraints belong in durable project context. Implementation belongs in code.

### Every task has an explicit lifecycle

Bug, incident, research task, small change, refactoring, and feature work should not all use the same delivery ceremony.

### Execution is isolated

Parallel tasks should not share mutable workspaces by default.

### Completion is verified externally

The agent does not get to declare success solely by saying “done.” Review, tests, CI, and lifecycle contracts decide completion.

### New knowledge returns to the system

Discoveries made during implementation should not remain trapped in an issue thread or agent transcript when they affect future work.

## Deployment model

Continuum is intentionally composable.

A downstream repository can adopt only the layers it needs, but the full model is:

```text
repository/
├── memory-bank/       # durable context and governance
├── .start-issue/      # execution defaults
├── .code-converge/    # convergence policy/config
├── init.sh            # optional worktree initialization
└── application code
```

The tools remain separate executables and repositories so each layer can evolve independently while sharing one development model.

## Platform boundary

Continuum sits between coding agents and the software delivery substrate.

```text
                MODELS
       Claude / GPT / Gemini / ...
                   │
             CODING AGENTS
                   │
                   ▼
       ┌─────────────────────┐
       │      CONTINUUM      │
       │                     │
       │ Context             │
       │ Governance          │
       │ Task orchestration  │
       │ Isolation           │
       │ Verification        │
       │ Convergence         │
       └──────────┬──────────┘
                  │
                  ▼
        Git / GitHub / CI / Prod
```

Continuum therefore stays useful as models and agent CLIs change: the durable value is the development system around them.
