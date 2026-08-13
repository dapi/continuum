# Live Demo — Two Prompts to a Verified PR

This walkthrough demonstrates the core Continuum claim on a small but real GitHub workflow.

## Preconditions

Install the Continuum components you want to demonstrate:

```sh
go install github.com/dapi/memory-bank-cli/cmd/memory-bank-cli@latest
go install github.com/dapi/start-issue/v2/cmd/start-issue@latest
go install github.com/dapi/code-converge/cmd/code-converge@latest
```

Authenticate GitHub CLI and make sure your coding agent CLI is available.

Clone the repository and start from `main`:

```sh
git clone https://github.com/dapi/continuum.git
cd continuum
git checkout main
```

The demo application lives in `examples/two-prompt-demo` and already has a passing test and CI workflow.

## Prompt 1 — Attach Continuum

Give your coding agent this high-level instruction:

```text
Attach Continuum to this repository as a brownfield project. Follow the Continuum brownfield quick start and the Memory Bank brownfield adaptation protocol. Discover the repository from evidence before treating Memory Bank as project truth. Adapt only supported durable context, validate the result, and stop before implementing product changes.
```

Expected internal work:

1. inspect repository instructions, docs, code and CI;
2. create evidence-backed brownfield intake;
3. install Memory Bank;
4. adapt canonical project owners without inventing facts;
5. validate with `memory-bank-cli lint` and `memory-bank-cli doctor`;
6. leave the repository ready for issue-driven work.

The operator does not manually reconstruct project context for the agent.

## Prompt 2 — Execute the issue

The canonical task is GitHub Issue #1:

```text
Execute GitHub issue #1 through Continuum. Use an isolated issue worktree, follow the applicable Memory Bank delivery flow, implement and test the requirement, then run the convergence loop until review is clean and CI for the published change is green. Produce a PR; stop only on a real human gate or a failed closed condition.
```

A concrete terminal-oriented execution may use:

```sh
start-issue 1 --repo dapi/continuum --agent codex
```

and, from the resulting issue worktree after implementation:

```sh
code-converge
```

Expected end state:

```text
GitHub Issue #1
      ↓
issue branch + isolated worktree
      ↓
Memory Bank context + delivery flow
      ↓
agent implementation + tests
      ↓
code-converge
      ↓
review → fix → review → CI
      ↓
verified PR
```

## What to show during the demo

Do not spend the demo explaining basic agent capabilities. Show the lifecycle boundaries:

- the issue is the task boundary;
- Memory Bank is durable context, not chat history;
- the worktree isolates execution;
- review findings cause another fix cycle rather than becoming manual cleanup;
- CI is checked against the published SHA;
- failure does not silently become success;
- project knowledge survives the session.

## The point

The demo is intentionally small. The claim is not that adding one endpoint is difficult.

The claim is that the same operator interface stays small while the repository and its accumulated requirements become large:

> **Attach the system. Execute the issue.**

Continuum handles the lifecycle between those two commands and a verified change.
