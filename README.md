# Please with esbuild

These rules attempt to use esbuild to treat Javascript bundling more like a traditional compiler. 
By that I mean, they's some analogy to compiling and linking: 

* Compiling: an npm module can be minified into a single `.js` file. In doing so, we provide a list of known dependencies of 
  this module, just like we would to a compiler. Any `require()` that isn't for one of these modules must be resolved at this stage. 
* Linking: we also minify, except this time all `require()` calls must be resolved, to produce a single `.bundle.js` file.

Though `esbuild` supports plugins that can implement a full featured bundler, I propose that instead we leave this to Please. If we 
wish to implement typescript, we'd implement similar rules that take in a `.ts` file and output a `.js` file by calling a typescript 
transpiler on it that is consumed by a `js_library()` rule. 

A similar approach can be taken to minify css or html, once we have a valid javascript module. 
 
# Considerations

* Can node modules provide resources, or require runtime data? 
* What about cases where we need multiple versions of the same module? I think this is solvable through some sort of import config file.

