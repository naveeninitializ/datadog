package main

import (
	"github.com/initializ-buildpacks/datadog/helper"
	"github.com/paketo-buildpacks/libpak/sherpa"
)

func main() {
	sherpa.Execute(func() error {
		return sherpa.Helpers(map[string]sherpa.ExecD{
			"toggle": helper.Toggle{},
		})
	})
}
