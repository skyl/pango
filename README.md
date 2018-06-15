# Prerequisites

Assuming you have the following installed:

* Bazel


# Install Dependencies

```
# TODO: why can't we put the .toml and lock file in the root?
cd src/pango
dep ensure
cd ../..
```


# Test

```
bazel test //...
```


# Build

```
bazel build //...
```
