load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary",)

go_library(
    name = "render",
    srcs = ["render.go"],
    importpath = "github.com/cripplet/vangogh/core/render",
    deps = [
        "//lib/api:api",
        "//lib/api/proto:base_go_proto",
        "//lib/core/proto:core_extension_go_proto",
    ],
)

go_binary(
    name = "main",
    srcs = ["main.go"],
    importpath = "github.com/cripplet/vangogh/core/main",
    data = [
        "//lib/core/template:templates",
    ],
    deps = [
    ],
)
