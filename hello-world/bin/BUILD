# load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_path")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "app/bin",
    visibility = ["//visibility:private"],
    deps = [
        "//lib/uuid:go_default_library",
        "@com_github_satori_go_uuid//:go_default_library",
    ],
)

go_binary(
    name = "bin",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

# go_path(
#     name = "mypath",
#     # mode = "archive",
#     deps = ["bin"]
# )

