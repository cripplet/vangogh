load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "generate",
    srcs = ["generate.go"],
    importpath = "github.com/cripplet/vangogh",
    deps = [
        "//lib/proto:blog_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)

go_library(
    name = "api",
    srcs = ["api.go"],
    importpath = "github.com/cripplet/vangogh/api",
    deps = [
        "//lib/proto:blog_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
