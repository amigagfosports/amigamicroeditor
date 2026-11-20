#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
release="$root/release"
version=${GFOS_VERSION:-0.1.0-rc1}
epoch=${SOURCE_DATE_EPOCH:-$(git -C "$root" log -1 --format=%ct)}
commit=$(git -C "$root" rev-parse --short=12 HEAD)
compile_date=$(date -u -d "@$epoch" '+%Y-%m-%dT%H:%M:%SZ')

cd "$root"
upstream_version=$(go run tools/build-version.go)
rm -rf "$release"
mkdir -p "$release"

build_arch() {
    arch=$1
    stage="$release/GFOpenMicro-$version-linux-$arch"
    mkdir -p "$stage"

    ldflags="-s -w -X github.com/micro-editor/micro/v2/internal/util.Version=$upstream_version -X github.com/micro-editor/micro/v2/internal/util.CommitHash=$commit -X github.com/micro-editor/micro/v2/internal/util.CompileDate=$compile_date"
    CGO_ENABLED=0 GOOS=linux GOARCH="$arch" \
        go build -trimpath -ldflags "$ldflags" -o "$stage/micro" ./cmd/micro

    cp package/GFOpenMicro package/README-GFOS.md LICENSE LICENSE-THIRD-PARTY "$stage/"
    chmod 0755 "$stage/GFOpenMicro" "$stage/micro"

    archive="$release/GFOpenMicro-$version-linux-$arch.tar.gz"
    tar --sort=name --owner=0 --group=0 --numeric-owner --mtime="@$epoch" \
        -C "$release" -czf "$archive" "$(basename "$stage")"
}

build_arch amd64
build_arch arm64

(cd "$release" && sha256sum ./*.tar.gz > SHA256SUMS)
printf 'Built GFOpenMicro %s from Micro %s (%s)\n' "$version" "$upstream_version" "$commit"
