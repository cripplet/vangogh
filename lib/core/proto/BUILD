load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility=["//visibility:public"])

proto_library(
    name = "core_extension_proto",
    srcs = ["core_extension.proto"],
    deps = [
        "//lib/api/proto:base_proto",
    ],
)

go_proto_library(
    name = "core_extension_go_proto",
    importpath = "github.com/cripplet/vangogh/core/proto",
    proto = ":core_extension_proto",
    deps = [
        "//lib/api/proto:base_go_proto",
    ],
)
