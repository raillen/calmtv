GO_TAGS ?= sqlite_fts5

.PHONY: test build lint vuln image atlas package-smoke ui-smoke target-preflight target-session-check package

test:
	go test -tags "$(GO_TAGS)" ./...

build:
	go build -tags "$(GO_TAGS)" -trimpath -o build/tv-shell ./cmd/tv-shell

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then golangci-lint run; else echo "golangci-lint unavailable"; exit 127; fi

vuln:
	@if command -v govulncheck >/dev/null 2>&1; then govulncheck ./...; else echo "govulncheck unavailable"; exit 127; fi

image:
	@if command -v lb >/dev/null 2>&1; then cd image/live-build && ./auto/config && ./auto/build; else echo "live-build unavailable"; exit 127; fi

atlas:
	atlas validate atlas
	atlas docs check

package-smoke:
	./scripts/package-smoke

ui-smoke:
	./scripts/test-ui-navigation

target-preflight:
	./scripts/target-preflight

target-session-check:
	./scripts/target-session-check

package: build
	@if command -v dpkg-buildpackage >/dev/null 2>&1; then ./scripts/build-deb; else ./scripts/build-deb-portable; fi
