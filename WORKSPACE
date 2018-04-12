http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.2/rules_go-0.10.2.tar.gz",
    sha256 = "4b2c61795ac2eefcb28f3eb8e1cb2d8fb3c2eafa0f6712473bc5f93728f38758",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.10.1/bazel-gazelle-0.10.1.tar.gz",
    sha256 = "d03625db67e9fb0905bbd206fa97e32ae9da894fe234a493e7517fd25faec914",
)

# git_repository(
#     name = "org_pubref_rules_protobuf",
#     remote = "https://github.com/pubref/rules_protobuf",
#     tag = "v0.8.1",
# )

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_repository", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

new_http_archive(
    name = "com_github_googleapis_googleapis",
    sha256 = "63555639d21dd0041a2e7e02d0af57b60da0b8dfccbf638fba8a930ceb555176",
    url = "https://github.com/googleapis/googleapis/archive/6277786c1352a6d7be1e5a5d697394e1802f6608.zip",
    strip_prefix = "googleapis-6277786c1352a6d7be1e5a5d697394e1802f6608",
    build_file_content = """
package(default_visibility = ["//visibility:public"])

proto_library(
    name = "annotations_proto",
    srcs = [
        "google/api/annotations.proto",
        "google/api/http.proto",
    ],
    deps = ["@com_google_protobuf//:descriptor_proto",],
)

proto_library(
    name = "status_proto",
    srcs = ["google/rpc/status.proto"],
    deps = ["@com_google_protobuf//:any_proto"],
)
    """,
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "925541529c1fa6821df4e44ce2723319eb2be768",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    commit = "07f5e79768022f9a3265235f0db4ac8c3f675fec",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
)

go_repository(
    name = "in_gopkg_urfave_cli_v2",
    commit = "d3ae77c26ac8db90639677e4831a728d33c36111",
    importpath = "gopkg.in/urfave/cli.v2",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "35de2414665fc36f56b72d982c5af480d86de5ab",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "1e2570b1b19ade82d8dbb31bba4e65e9f9ef5b34",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_net",
    commit = "b68f30494add4df6bd8ef5e82803f308e7f7c59c",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "com_github_go_kit_kit",
    commit = "ca4112baa34cb55091301bdc13b1420a122b1b9e",
    importpath = "github.com/go-kit/kit",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    commit = "390ab7935ee28ec6b286364bba9b4dd6410cb3d5",
    importpath = "github.com/go-logfmt/logfmt",
)

go_repository(
    name = "com_github_go_stack_stack",
    commit = "259ab82a6cad3992b4e21ff5cac294ccb06474bc",
    importpath = "github.com/go-stack/stack",
)

go_repository(
    name = "com_github_jinzhu_gorm",
    commit = "6ed508ec6a4ecb3531899a69cbc746ccf65a4166",
    importpath = "github.com/jinzhu/gorm",
)

go_repository(
    name = "com_github_jinzhu_inflection",
    commit = "04140366298a54a039076d798123ffa108fff46c",
    importpath = "github.com/jinzhu/inflection",
)

go_repository(
    name = "com_github_kr_logfmt",
    commit = "b84e30acd515aadc4b783ad4ff83aff3299bdfe0",
    importpath = "github.com/kr/logfmt",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    commit = "6c771bb9887719704b210e87e934f08be014bdb1",
    importpath = "github.com/mattn/go-sqlite3",
)

go_repository(
    name = "com_github_oklog_oklog",
    commit = "e3ad1c411c27b4bc18c5facb9331820d141a5a54",
    importpath = "github.com/oklog/oklog",
)

go_repository(
    name = "com_github_oklog_run",
    commit = "4dadeb3030eda0273a12382bb2348ffc7c9d1a39",
    importpath = "github.com/oklog/run",
)
