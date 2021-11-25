# Please with esbuild

These rules attempt to use esbuild to treat Javascript bundling more like a traditional compiler. 
By that I mean, there's some analogy to distinct compile and link stages: 

## Compiling
Similar to `go tool compile`, esbuild has a concept of resolution that happens before the load phase. This is similar 
to how the go compiler resolves import paths to `.a` files. I have written a special `please` resolver. This takes in 
a list of known modules, that we can derrive from the direct dependencies of the rule. This method looks as such:

```golang
func(args api.OnResolveArgs) (api.OnResolveResult, error) {
  # opts.Modules here are the list of known dependencies of this modules we're "compiling"
  if path, ok := opts.Modules[args.Path]; ok { 
    return api.OnResolveResult{
      Path:      path,
      Namespace: "please",
    }, nil
  }
  return api.OnResolveResult{}, nil
}
```
 
 At this point, esbuild has tagged this `require()` as being part of the `please` namespace. This means that the 
 please plugin will handle this going forward. We've essentially resolved the `require()` to a filepath based on the 
 `opts.Module` mapping. 
 
 The `node_module()` and `js_library()` rules will use this to resolve thier `require()`s to the correct paths. 
 
 ## Linking
 Similar to `go tool link`, esbuild has a `load` phase. This will be used by the `js_binary()` and `js_test()` targets
 to produce a single `bundle.js`. At this point, we have injected some metadata into the `require()` calls in the previous
 "link" stage. We simply have to read this out and provide the correct data back to esbuild. The code looks like this:
 
 ```golang
func(args api.OnLoadArgs) (api.OnLoadResult, error) {
  # args.Path is set by the resolver above for us
  path := filepath.Join(wd, args.Path)
  data, err := ioutil.ReadFile(path)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to load %v: %v\n", args.Path, err)
    os.Exit(1)
  }

  contents := string(data)
  return api.OnLoadResult{
    Contents: &contents,
  }, nil
}
```
 
 
# Considerations

* Can node modules provide resources, or require runtime data? 


