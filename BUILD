load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/cory-evans/Ichnaea
# gazelle:proto disable

gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

gazelle(
    name = "gazelle-fix",
    command = "fix",
)

gazelle(
    name = "gazelle-update",
    command = "update",
)
