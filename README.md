# Prerequisites

You must first install the following:

* Bazel
* dep (TODO: wrap with Bazel)

## Install Bazel

On a Mac:

```
brew install bazel
```

On a fresh install of Ubuntu 18.04 this might look like:

```
sudo apt update
sudo apt upgrade
sudo apt-get install pkg-config zip g++ zlib1g-dev unzip python
curl -L https://github.com/bazelbuild/bazel/releases/download/0.14.1/bazel-0.14.1-installer-linux-x86_64.sh > bazel-installer.sh
chmod +x bazel-installer.sh
./bazel-installer.sh --user
echo "export PATH=$HOME/bin:\$PATH" >> ~/.bashrc
```

## Install dep

On a Mac:

```
brew install dep
```

On Ubuntu:

```
sudo apt install go-dep
```

# Clone, Install dependencies and Bazel BUILD files

Clone this repository:

```
git clone https://github.com/skyl/pango
cd pango
```

Set your GOPATH to the root of this repository:

```
export GOPATH=`pwd`
```

Install 3rd party dependencies with dep:

```
# TODO wrap this up completely within Bazel?
cd src/pango
dep ensure
cd ../..
```

Use Gazelle to add Bazel build files in the vendor directory:

```
bazel run //:gazelle
```

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

# Run the JWT token authorization service

```
bazel run //src/pango/services/auth
```

# Run a resource that requires the auth bearer token

```
bazel run //src/pango/services/aresource
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
# or, if already imported, just:
dep ensure
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
