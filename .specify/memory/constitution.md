# GoRestController Constitution

## Core Principles

### I. Single-task changes (non-negotiable)

Work is delivered in **one discrete task at a time**. A change set MUST address exactly **one** task from the active task list (for example, a single item in `tasks.md` for the current feature).

- Do not combine unrelated behavior, refactors, dependency upgrades, or “while I’m here” edits with the task in progress.
- If a task reveals necessary cleanup, **finish the current task** with minimal, task-scoped changes, then open a **new** task for the cleanup.
- Scope creep is resolved by **splitting** work into additional tasks, not by expanding the current change set.

### II. Commit when the task is done (non-negotiable)

When a task is **complete** (behavior matches acceptance, tests and checks you rely on are green), you MUST **commit** that work before starting the next task.

- Prefer **one commit per completed task** unless the task was explicitly broken into sub-commits agreed for traceability (still one task, ordered commits).
- Commit messages MUST identify the task (e.g. task ID or title from `tasks.md`) so history maps to planned work.

### III. Pull request per completed task (non-negotiable)

After each task is committed, you MUST open a **pull request** (or, if your workflow uses stacked PRs, a **new PR** that contains only that task’s commits) before implementation of the **next** task begins on the main line of work.

- A PR MUST be scoped to **one completed task** unless an explicit exception is documented in the PR description and approved by reviewers.
- Reviewers treat bundled tasks as a **constitution violation** unless the exception is documented and approved.

### IV. Simplicity and traceability

Prefer the smallest change that satisfies the task. Task boundaries exist so history, review, and rollback stay understandable.

## Task definition

- **Task**: A single planned unit of work with clear acceptance criteria, typically one checkbox or numbered item in the feature `tasks.md` (or equivalent agreed tracker).
- **Task complete**: Acceptance criteria met; automated checks required for this repo pass; work is committed per Principle II.

## Development workflow

1. Pick **one** task; do not start another until this one is merged or explicitly parked per team agreement.
2. Implement **only** that task; keep diffs focused.
3. When the task is complete, **commit** (Principle II).
4. **Open a PR** for that task (Principle III); request review as your process requires.
5. After merge (or after the PR is approved and merged per policy), take the **next** task.

Emergency production fixes MAY follow a documented hotfix path, but must be **retroactively** reconciled with tasks (e.g. follow-up task to align spec and `tasks.md`).

## Governance

- This constitution **supersedes** ad hoc habits when they conflict with these rules.
- **Amendments**: Propose a PR that updates this file, states rationale, and bumps the version line below.
- **Compliance**: Authors and reviewers verify single-task scope, commit boundaries, and PR boundaries before merge.
- **Runtime guidance**: Feature work should still flow through Spec Kit (`spec.md` → plan → `tasks.md` → implement) so tasks remain the source of truth for “one task at a time.”

**Version**: 1.0.0 | **Ratified**: 2026-04-04 | **Last Amended**: 2026-04-04
