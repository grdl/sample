# Sample

[![build](https://github.com/grdl/sample/actions/workflows/build.yml/badge.svg)](https://github.com/grdl/sample/actions/workflows/build.yml)
[![release](https://github.com/grdl/sample/actions/workflows/release.yml/badge.svg)](https://github.com/grdl/sample/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/grdl/sample)](https://goreportcard.com/report/github.com/grdl/sample)

Sample :golang: Go application to be used as a base template for other projects.


## How it works

### Release notes

Release notes for GitHub releases are generated with:
```bash
sed -n '/^## \['${GITHUB_REF##*\/}'\]/,/^## \[/{//!p;}' CHANGELOG.md | sed -e :a -e '/^\n*$/{$d;N;};/\n$/ba'
```

- It looks at the `CHANGELOG.md` file.
- The first `sed` takes the content between the current tag and the next one ([reference](https://stackoverflow.com/a/38978201/1085632)).
- The second `sed` removes all trailing empty lines ([reference](https://stackoverflow.com/a/7359879/1085632)).
- The `${GITHUB_REF##*\/}` returns just the name of the tag using bash parameter expansion ([reference](https://github.community/t/how-to-get-just-the-tag-name/16241/2)).
- The command is passed to goreleaser with [--release-notes flag](https://goreleaser.com/customization/release/#custom-release-notes).