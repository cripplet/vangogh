load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "api",
    srcs = ["api.go"],
    importpath = "github.com/cripplet/vangogh/api",
    deps = [
        "//lib/api/proto:base_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
    visibility = ["//visibility:public"],
)
