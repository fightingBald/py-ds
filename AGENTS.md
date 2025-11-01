# Repository Guidelines

## Project Structure & Module Organization
The repository is a single Go module (`go.mod`, Go 1.24). Introductory playground code sits in `counter/`, while `go_basic_tuto/` is split by topic: `a_basic_type/` for language fundamentals, `d_concurrency/` for goroutines and synchronization, and `e_web/` for chi-based HTTP examples. Practice problems live in `go_practice/`; each directory (`lru_cache/`, `deque/`, `ordered_dict/`, etc.) is an independent package with its own tests. Place fixtures and helper modules beside the code they serve, and add new exercises under the matching tutorial or practice section to keep the layout predictable.

# 6) CODING & NAMING（通用编码规范）

| Rule | Details                                                                     |
| ---- | --------------------------------------------------------------------------- |
| 格式化  | 强制格式化与静态检查（fmt/vet/linter/ruff/prettier 等）。                                 |
| 命名   | 语义化、可读，避免 `data/tmp/info` 等空词。错误用 `ErrXxx`。                                 |
| 结构   | 小函数小类型，单一职责；参数过多用配置 struct/Options。                                         |
| 上下文  | I/O 必带上下文/超时；禁止将上下文存入结构体。                                                   |
| 日志   | 结构化字段，分级打印，严禁 `print` 系调试遗留。                                                |
| SQL  | 必须参数化/Builder/ORM；禁止字符串拼接；必写事务与隔离意图。                                        |
| 语言特性 | 善用语言的最新特性（如 Go 泛型、Java 结构化并发、Python 3.12 模式匹配、TS 装饰器等），在保证可读性的前提下提升安全性与表达力。 |


| 项     | 必须               | 禁止                  |
| ----- | ---------------- | ------------------- |
| 语言/版本 | 遵循项目声明版本与标准库优先   | 过时 API、私有魔改         |


## Build, Test, and Development Commands
- `go fmt ./...`: canonical formatting across every package; run it before staging.
- `go test ./...`: executes the entire suite, including tutorials and practice sets; required green before submitting.
- `go build ./...`: quick compile check that surfaces missing imports or type errors early.
- `go run go_basic_tuto/e_web/a_go_chi.go`: launches the chi sample server at `http://localhost:8080` for manual verification.

## Coding Style & Naming Conventions
Follow idiomatic Go defaults: tab indentation, mixedCaps for exported identifiers, and short receiver names that hint at the type (e.g., `func (q *Queue)`). Keep packages cohesive and name them for their behavior rather than implementation details. Rely on `go fmt` or editor integrations; add `goimports` if you want automatic import grouping. Document exported APIs with brief sentences and co-locate helper types with the code they support.

## Testing Guidelines
Tests use Go’s `testing` package augmented by `stretchr/testify` assertions. Mirror source files with `_test.go` companions (`counter/counter.go` → `counter/counter_test.go`) and prefer table-driven subtests to cover success, boundary, and failure cases. New concurrency work should follow the patterns in `go_basic_tuto/d_concurrency`, asserting context cancellation, deadlines, and synchronization behavior. When adding data structure challenges, include both happy-path checks and guardrails for invalid input.

## Commit & Pull Request Guidelines
Craft commits with short, imperative subjects (`add stack implementation`, `refactor context tests`) and expand with details in the body if the change is non-trivial. Squash noisy intermediate commits before opening a PR. Every PR should explain the motivation, outline functional changes, list any new commands or configuration steps, and confirm `go test ./...` output. Attach screenshots or terminal snippets whenever changes alter observable behavior.

## Environment & Tooling
Develop against Go 1.24 or newer. Use `go mod tidy` after introducing dependencies and review the diff to avoid unneeded modules. Configure editor or CI hooks to run `go fmt` and `go test` automatically. If a solution needs environment variables or sample payloads, document defaults in the relevant package README and avoid committing secrets.
