# Prerequisites

Assuming you have the following installed:

* Bazel
* dep (TODO: wrap with Bazel)
<!-- * Go -->

# Test

```
# TODO: the vendor directory messes this up a little bit
bazel test //src/pango/lib/...
bazel test //src/pango/services/...
```

# Build

```
# TODO: the vendor directory messes this up a little bit
bazel build //src/pango/lib/...
bazel build //src/pango/services/...
```

# Run Hello World

```
bazel run //src/pango/services/hello
```

# Run the Goa Service example

```
bazel run //src/pango/services/cellar
```

# Work with `goagen`

We haven't quite figured out how we are going to handle dependencies,
especially dev tools outside of Bazel.
Ideally we are going to use eg dep and/or bazel to have locked dependencies
without checking in the 3rd party libraries into the git history. However,
we need to figure out how this works for binary dev tools such as goagen.
For now, use `go get` as described here: https://goa.design/learn/guide/
. Then, check in the goa artifacts.

# Installing new (Go) dependencies

All 3rd party deps are managed with `dep`.

```
cd src/pango
dep ensure -add foo/bar
```

These are then translated to Bazel with Gazelle:

```
bazel run //:gazelle
```

# Updating the Bazel build files for Go packages

When you add a service or dependency to the monorepo, you need to create
build files so that Bazel will know how to build yourthing and what will
need to be rebuilt when yourthing changes. Luckily this can be automated
with Gazelle:

```
bazel run //:gazelle
```
