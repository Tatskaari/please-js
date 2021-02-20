package main

import (
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/thought-machine/go-flags"
)


var opts = struct {
	Usage string

	Modules map[string]string `short:"m" long:"module" description:"Module mapping"`
	EntryPoints []string `short:"e" long:"entry_point"`
	Out string `short:"o" long:"out"`

	Link struct {
	} `command:"link" alias:"c" description:"Compile the entry_points, redirecting requires for the provided modules"`
}{
	Usage: `
esbuild provides a wrapper around esbuild, using plugins to perform a more traditional "compile" and "link" workflow 
around bundling. 
`,
}


var wd, wdErr = os.Getwd()
var plugin = api.Plugin{
	Name:  "please",
	Setup: func(build api.PluginBuild) {
		build.OnResolve(api.OnResolveOptions{Filter: `[^\.].*`},
			func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				var path = filepath.Join(args.ResolveDir, args.Path)
				if p, ok := opts.Modules[args.Path]; ok {
					path = filepath.Join(wd, p)
				}
				return api.OnResolveResult{
					Path: path,
				}, nil
			})
	},
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	if wdErr != nil {
		panic(wdErr)
	}

	result := api.Build(api.BuildOptions{
		EntryPoints: opts.EntryPoints,
		Outfile:     opts.Out,
		Bundle:      true,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
	})
	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}