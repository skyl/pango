load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.12.0/rules_go-0.12.0.tar.gz"],
    sha256 = "c1f52b8789218bb1542ed362c4f7de7052abcf254d865d96fb7ba6d44bc15ee3",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
# load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
# gazelle_dependencies()


# load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
#
# # You *must* import the Go rules before setting up the go_image rules.
# # http_archive(
# #     name = "io_bazel_rules_go",
# #     # Replace with a real SHA256 checksum
# #     sha256 = "{SHA256}"
# #     # Replace with a real commit SHA
# #     strip_prefix = "rules_go-{HEAD}",
# #     urls = ["https://github.com/bazelbuild/rules_go/archive/{HEAD}.tar.gz"],
# # )
#
# load("@io_bazel_rules_go//go:def.bzl", "go_repositories")
#
# go_repositories()
#
# load(
#     "@io_bazel_rules_docker//go:image.bzl",
#     _go_image_repos = "repositories",
# )
#
# _go_image_repos()
