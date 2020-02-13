load("@io_bazel_rules_go//go:def.bzl", "go_test")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "testdata",
    srcs = glob([
        "example.textpb",
    ]),
)

go_test(
    name = "data_validation_test",
    srcs = ["data_validation_test.go"],
    importpath = "github.com/cripplet/vangogh/api/proto/testdata",
    data = [":testdata"],
    deps = [
        "//lib/api/proto:base_go_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
