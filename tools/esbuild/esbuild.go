package main

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/thought-machine/go-flags"
	"io/ioutil"
	"os"
	"path/filepath"
)

var opts = struct {
	Usage string

	Modules     map[string]string `short:"m" long:"module" description:"Module mapping"`
	EntryPoints []string          `short:"e" long:"entry_point"`
	Out         string            `short:"o" long:"out"`

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
	Name: "please",
	Setup: func(build api.PluginBuild) {
		build.OnResolve(api.OnResolveOptions{Filter: `.*`},
			func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				if path, ok := opts.Modules[args.Path]; ok {
					return api.OnResolveResult{
						Path:      path,
						Namespace: "please",
					}, nil
				}
				return api.OnResolveResult{}, nil
			})
		build.OnLoad(api.OnLoadOptions{Namespace: "please", Filter: `.*`}, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
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
		Platform:    api.PlatformNode,
		Format:      api.FormatESModule,
		Plugins:     []api.Plugin{plugin},
	})
	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}
