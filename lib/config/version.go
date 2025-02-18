package config

import "log"

// Version Set using compile flags
var Version string

func init() {
	if Version == "" {
		log.Fatalln("Version not set using compile flags. Use -ldflags \"-X 'duckdb-version-manager/config.Version=1.0.0'\" to set the version.")
	}
}
