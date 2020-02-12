load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "generate",
    srcs = ["generate.go"],
    importpath = "github.com/cripplet/vangogh",
    deps = [
        "//lib/proto:blog_go_proto",
    ],
)
