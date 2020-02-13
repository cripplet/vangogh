load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility=["//visibility:public"])

proto_library(
    name = "base_proto",
    srcs = ["base.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "base_go_proto",
    importpath = "github.com/cripplet/vangogh/api/proto",
    proto = ":base_proto",
)
