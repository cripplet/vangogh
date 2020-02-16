load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility=[
    "//visibility:public",
])

go_library(
    name = "api",
    srcs = ["api.go"],
    importpath = "github.com/cripplet/vangogh/api",
    deps = [
        "//lib/api/proto:base_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)

go_library(
    name = "api_util",
    srcs = ["api_util.go"],
    importpath = "github.com/cripplet/vangogh/api/api_util",
    deps = [
        ":api",
        "//lib/api/proto:base_go_proto",
    ],
)
