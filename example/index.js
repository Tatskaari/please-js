let bar = require("./bar/bar")
// const semver = require("semver")

function foo() {
    return bar.bar()
}


console.log(foo())
// console.log(semver.minVersion('>=1.0.0'))