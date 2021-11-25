#!/usr/bin/node
let bar = require("example/bar/bar")
const semver = require("semver")

function foo() {
    return bar.bar()
}


console.log(foo())
console.log(semver.clean('  =v1.2.3   '))