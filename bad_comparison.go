package main

import "os"

func LoadConfiguration() string {
	if getConfigurationKey != nil { // HL
		// What we meant was getConfigurationKey().
		// But functions are first class, so this compiles fine.
		return "success!"
	} else {
		return "error!"
	}
}

func getConfigurationKey() (key string) {
	return os.Getenv("CONFIG_KEY")
}
