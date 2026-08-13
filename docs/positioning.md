# Continuum Positioning

## Category

**Continuous Agentic Software Development**

Continuum is not positioned as another coding agent. It is the operating layer that makes coding agents usable for long-running software development.

## Core claim

> **Autonomy without losing control.**

Coding agents are already capable of implementing meaningful changes. The hard problem is preserving project intent, enforcing lifecycle rules, isolating execution, verifying output, and keeping the codebase coherent over months and years.

Continuum addresses that layer.

## Developer positioning

### Short version

> **Issue in. Verified PR out.**

### Expanded version

Give an experienced coding agent a real task, the right durable context, an isolated workspace, and an explicit definition of done. Then let the system drive the change through review and CI until it converges or fails explicitly.

### Pain it replaces

- manually re-explaining the project to every new session;
- watching agents work in one shared checkout;
- repeatedly telling them to review their own changes;
- manually restarting them after CI failures;
- accepting “done” before the change is actually verified;
- letting requirements and rationale disappear into chats.

### Developer slogans

- **Stop babysitting coding agents.**
- **Issue in. Verified PR out.**
- **A coding session is ephemeral. Your development system should not be.**
- **Your codebase needs memory.**
- **Agents write code. Continuum operates the loop.**

## Staff / Principal / CTO positioning

### Short version

> **A control plane for autonomous software development.**

### What matters at this level

The question is not whether an agent can implement one feature. The question is whether hundreds or thousands of agent-driven changes can accumulate without destroying architectural coherence, product intent, operability, and trust.

Continuum introduces durable structures around agent execution:

- explicit source ownership;
- persistent requirements and domain knowledge;
- task-specific delivery flows;
- reproducible isolated workspaces;
- verification contracts;
- bounded review/fix loops;
- CI-backed convergence;
- feedback of newly discovered knowledge into the project.

### Enterprise claim

> The more autonomous development becomes, the more explicit governance must become.

Continuum is designed to provide that governance without turning agentic development back into a manual process.

## Brownfield positioning

### Short version

> **Agentic development for the codebase you already have.**

Most agent demos optimize for greenfield creation. Real engineering organizations own mature systems with undocumented constraints, historical decisions, implicit domain rules, legacy integrations, operational requirements, and large regression surfaces.

Continuum is deliberately oriented toward brownfield adoption:

- attach durable context to the existing repository;
- distinguish verified project facts from generic template assumptions;
- make architectural and domain constraints explicit;
- execute new work against the existing Git history and CI;
- improve project knowledge as tasks are completed.

The intended result is not only faster implementation. It is a codebase that becomes easier for the next human or agent to understand.

## Competitive framing

### Coding agent alone

```text
Prompt → Agent → Code
```

Fast, but heavily dependent on session context and human supervision.

### Continuum

```text
Intent
  ↓
Persistent project context
  ↓
Task lifecycle
  ↓
Isolated agent execution
  ↓
Implementation
  ↓
Review / Fix
  ↓
CI / Recovery
  ↓
Verified change
  ↓
Durable learning
```

The differentiator is not a smarter model. It is a more reliable development system.

## Model-agnostic story

Models will continue to improve and coding-agent interfaces will continue to change.

Continuum should remain above that competition.

> **Models generate code. Continuum operates software development.**

This makes model replacement a configuration choice rather than a platform rewrite.

## Three concise presentation narratives

### 1. Stop Babysitting Agents

1. You already know agents can code.
2. The bottleneck is supervision, context reconstruction, and verification.
3. Persist project intent in Git.
4. Turn issues into isolated agent sessions.
5. Converge every change through review and CI.
6. **Autonomy without losing control.**

### 2. Issue In → Verified PR Out

1. Start from a normal GitHub issue.
2. `start-issue` creates the branch, worktree, context, and session.
3. The coding agent implements the task.
4. `code-converge` repeatedly reviews and fixes the delta.
5. CI is checked on the published SHA and recovered when possible.
6. **Issue in. Verified PR out.**

### 3. Your Codebase Needs Memory

1. Agent sessions forget; codebases live for years.
2. Code alone does not preserve intent, constraints, or rationale.
3. Memory Bank stores authoritative project knowledge next to the code.
4. Every agent starts from the same durable context.
5. New discoveries return to the project instead of disappearing into chats.
6. **A coding session is ephemeral. Your development system should not be.**

## One-paragraph description

Continuum is a continuous agentic development platform for real, long-lived software projects. It gives coding agents durable project context and governance, turns GitHub issues into isolated execution environments, drives changes through bounded review/fix loops and CI, and preserves newly discovered requirements and decisions for future work. It is designed for teams that already know how to use coding agents and now need to make autonomous development reliable at codebase scale.
