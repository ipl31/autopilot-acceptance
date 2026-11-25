# Contributing

Read PRODUCT.md and REVIEW.md. Include a regression test for behavior changes.
Run these checks on the final code and test revision. Formatting passes when
the command produces no output.

## Checks

```bash
gofmt -l . # no output
go vet ./...
go run honnef.co/go/tools/cmd/staticcheck@2026.2.2 ./...
mkdir -p /tmp/amtest && env -u TMUX TMUX_TMPDIR=/tmp/amtest go test -race ./...
go build ./...
```
