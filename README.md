# pcheck

Provide information on imported golang packages for a given package

## Usage

```go
pcheck /Users/kwalters/go/src/github.com/lostghost/pcheck
```

### Nested Dependency

- Main Package
 - Imported Package 1
  - Sub Package A
 - Imported Package 2
 - Imported Package 3

### List

- Imported Package 1
- Imported Package 2
- Imported Package 3
- Sub Package A

## Todo

- ~~Parse a single package path from command line arg~~
- Parse a single package name from command line arg (path from GOPATH)
- ~~Show a list of imported packages in a single package~~
- Flag to exclude std packages
- Flag to include test packages
- Show a list of imported packages for a package and all of it's imported packages
- Show a nested list of imported packages for a package and all of it's imported packages
- Show missing packages
- Show packages that have newer remote commits than your local version
- Traverse all packages, not including std package
- Traverse a tree of packages such as ./...
