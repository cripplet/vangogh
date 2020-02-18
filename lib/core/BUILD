load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary",)

go_library(
    name = "render",
    srcs = ["render.go"],
    importpath = "github.com/cripplet/vangogh/core/render",
    deps = [
        ":render_post",
        ":render_util",
        ":type",
        "//lib/api:api",
        "//lib/api/proto:base_go_proto",
        "//lib/core/proto:core_extension_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        # Import the any.Any type from here instead of //ptypes/any.
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
        "//lib/api:api",
        "//lib/api:api_util",
        "//lib/api/proto:base_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)

go_library(
    name = "type",
    srcs = ["type.go"],
    importpath = "github.com/cripplet/vangogh/core/type",
    deps = [
        "//lib/api/proto:base_go_proto",
    ],
)

go_library(
    name = "render_util",
    srcs = ["render_util.go"],
    importpath = "github.com/cripplet/vangogh/core/render_util",
    deps = [
        "//lib/api/proto:base_go_proto",
        "//lib/core/proto:core_extension_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@io_bazel_rules_go//proto/wkt:any_go_proto",
        "@io_bazel_rules_go//proto/wkt:timestamp_go_proto",
    ],
)

go_library(
    name = "render_post",
    srcs = ["render_post.go"],
    importpath = "github.com/cripplet/vangogh/core/render_post",
    deps = [
        ":render_util",
        ":type",
        "//lib/api:api",
        "//lib/api/proto:base_go_proto",
    ],
)
