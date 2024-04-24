load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/hika019/Pulsar-To-RDB.git
gazelle(name = "gazelle")

go_library(
    name = "Pulsar-To-RDB_git_lib",
    srcs = ["main.go"],
    importpath = "github.com/hika019/Pulsar-To-RDB.git",
    visibility = ["//visibility:private"],
    deps = ["//config"],
)

go_binary(
    name = "main",
    embed = [":Pulsar-To-RDB_git_lib"],
    visibility = ["//visibility:public"],
)
