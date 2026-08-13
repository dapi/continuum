# Two-Prompt Demo

A tiny brownfield-style service used to demonstrate the Continuum workflow end to end.

## Baseline

The service exposes:

```text
GET /health
```

Response:

```json
{"service":"continuum-demo","status":"ok"}
```

Run it:

```sh
go run .
```

Test it:

```sh
go test ./...
```

## Demo goal

Use a real GitHub issue in `dapi/continuum` as the task source. The operator should need only two high-level prompts:

1. Attach Continuum to this repository as a brownfield project.
2. Execute the demo issue and converge it to a verified PR.

The agent is expected to perform the internal discovery, Memory Bank adaptation, isolated issue execution, implementation, review/fix cycles, and CI verification required by the platform.

See [`DEMO.md`](DEMO.md) for the live walkthrough.
