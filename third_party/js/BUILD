package(default_visibility = ["PUBLIC"])

subinclude("//build_defs")

node_module(
    name = "lru-cache",
    out = "lru-cache-7.8.1.tgz",
    version = "7.8.1",
)

node_module(
    name = "parse-yarn-lock",
    out = "parse-yarn-lock-0.28.4.tgz",
    version = "0.28.4",
)

node_module(
    name = "semver",
    out = "semver-7.3.6.tgz",
    version = "7.3.6",
    deps = [
        ":lru-cache",
    ],
)
