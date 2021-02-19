package main

import (
	"github.com/evanw/esbuild/pkg/api"
)

type pleasePlugin struct {
	importConfig map[string]string
}


func

func main() {
	plugin := new(pleasePlugin)

	api.Build(api.BuildOptions{
		EntryPoints: []string{"input.js"},
		Outfile:     "output.js",
		Bundle:      true,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
		Plugins:           []api.Plugin{
			{Name: "please", Setup: func(build api.PluginBuild) {
				return plugin
			}},
		},
		Watch:             nil,
	})
}