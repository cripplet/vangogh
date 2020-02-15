load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary",)

go_library(
    name = "render",
    srcs = ["render.go"],
    importpath = "github.com/cripplet/vangogh/core/render",
    deps = [
        # "//lib/api:api",
        "//lib/api/proto:base_go_proto",
        "//lib/core/proto:core_extension_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        # Import the any.Any type from here instead of //ptypes/any.
        "@io_bazel_rules_go//proto/wkt:any_go_proto",
        "@io_bazel_rules_go//proto/wkt:timestamp_go_proto",
    ],
)

go_binary(
    name = "main",
    srcs = ["main.go"],
    importpath = "github.com/cripplet/vangogh/core/main",
    data = [
        "//lib/api/proto/testdata:testdata",
        "//lib/core/template:templates",
    ],
    deps = [
        ":render",
        "//lib/api/proto:base_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
