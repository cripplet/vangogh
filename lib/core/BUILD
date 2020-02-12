load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "render",
    srcs = ["render.go"],
    importpath = "github.com/cripplet/vangogh/core/render",
    deps = [
        "//lib/api:api",
        "//lib/core/proto:core_extension_go_proto",
        "//lib/proto:base_go_proto",
    ],
)

