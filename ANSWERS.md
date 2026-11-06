# Lab 6 – Software Process and Workflow Automation
Tools: GitHub Actions (CI/CD) + bpmn.io (BPMN)

## Task 1 – Manual vs Automated Activities
Workflow: a small Go application.

| Activity | Classification | Explanation |
|---|---|---|
| Code development | Manual | A developer writes and fixes the code. |
| Build | Automated | `go build` runs by itself on every push. |
| Testing | Automated | `go test` runs by itself after a successful build. |
| Approval / review of failures | Partially automated | The tool reports pass/fail, but a human decides what to fix. |
| Deployment | Automated | Runs automatically only if tests pass. |

## Task 2 – Pipeline
File: `.github/workflows/ci.yml`, name `CI-Pipeline-<UNIVERSITY ID>`.
Stages: Build → Test → Deploy. In the GitHub **Actions** tab each stage shows its status and logs.

## Task 3 – Pipeline Behavior
1. **Success run:** push the code. Build, Test and Deploy are all green.
2. **Failure demo:** in `main.go` change `return a + b` to `return a - b`, then commit and push.
   - Failed stage: **Test** (`TestAdd`).
   - Deploy was **skipped** (not run), because it `needs` Test.
   - Logs show: `--- FAIL: TestAdd`, `Add(2,3) = -1; want 5`, and `exit code 1`.
3. **Fix:** change back to `a + b`, commit, push. Pipeline is green again.

## Task 4 – BPMN
File: `workflow.bpmn` (open it at https://demo.bpmn.io → "Open").
Start → Write/Fix Code (User task) → Build (Service task) → Run Tests (Service task) → Gateway "Build and tests passed?"
- Yes → Deploy (Service task) → End
- No → Review Logs and Correct (User task) → back to Write/Fix Code

## Task 5 – Pipeline ↔ BPMN
| BPMN activity | Pipeline stage |
|---|---|
| Build (service task) | `build` job |
| Run Tests (service task) | `test` job |
| Gateway (passed?) | Success/failure of the `test` job |
| Deploy (service task) | `deploy` job |
| Write/Fix Code, Review Logs (user tasks) | Developer work outside the pipeline |

- **Automated:** build, test, deployment, and the pass/fail decision.
- **Human involvement:** writing code, reading logs, fixing bugs.
- **When an automated activity fails:** the pipeline stops, later stages are skipped, logs show the cause, and the flow goes to "Review Logs and Correct", then back to code and a new run.

## Video script (3–5 min)
1. (0:00) Introduce yourself and your university ID. Show the repo and the pipeline name in the Actions tab.
2. (0:30) Push a good commit; show Build → Test → Deploy going green; open a stage to show its logs.
3. (1:30) Push the `a - b` change; show Test failing, Deploy skipped, and the error in the logs.
4. (2:30) Fix it, push, show green again.
5. (3:15) Open the BPMN in bpmn.io; walk through it and explain the mapping table in Task 5.
