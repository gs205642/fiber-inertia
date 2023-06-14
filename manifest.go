package inertia

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Manifest map[string]struct {
	Css            []string `json:"css"`
	File           string   `json:"file"`
	IsEntry        bool     `json:"isEntry"`
	Imports        []string `json:"imports"`
	DynamicImports []string `json:"dynamicImports"`
	IsDynamicEntry bool     `json:"isDynamicEntry"`
	Src            string   `json:"src"`
}

func manifest(buildDirectory ...string) Manifest {
	path := "public/build"
	if len(buildDirectory) > 0 {
		path = buildDirectory[0]
	}
	path = filepath.Join(path, "manifest.json")

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var manifest Manifest
	err = json.Unmarshal(data, &manifest)
	if err != nil {
		panic(err)
	}
	return manifest
}
