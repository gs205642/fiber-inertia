package inertia

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
)

func reactRefresh() template.HTML {
	if !isRunningHot() {
		return ""
	}
	return template.HTML(
		fmt.Sprintf(
			`<script type="module">`+
				`import RefreshRuntime from '%s';`+
				`RefreshRuntime.injectIntoGlobalHook(window);`+
				` window.$RefreshReg$ = () => {};`+
				` window.$RefreshSig$ = () => (type) => type;`+
				` window.__vite_plugin_react_preamble_installed__ = true;`+
				`</script>`,
			hotAsset("@react-refresh"),
		),
	)
}

func vite(entrypoints []string, buildDirectory ...string) template.HTML {
	if isRunningHot() {
		html := makeTagForChunk(hotAsset("@vite/client"))
		for _, v := range entrypoints {
			html += makeTagForChunk(hotAsset(v))
		}
		return html
	}
	manifest := manifest(buildDirectory...)
	html := template.HTML("")
	for _, v := range entrypoints {
		m := manifest[v]
		for _, css := range m.Css {
			html += template.HTML(makeStylesheetTag(css))
		}
		html += template.HTML(makeScriptTag(m.File))
	}
	return html
}

func inertiaBody(page *Page) template.HTML {
	marshaledPage, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}

	return template.HTML(
		fmt.Sprintf(`<div id="app" data-page='%s'></div>`, marshaledPage),
	)
}

func makeTagForChunk(url string) template.HTML {
	if isCssPath(url) {
		return template.HTML(makeStylesheetTag(url))
	}

	return template.HTML(makeScriptTag(url))
}

func makeStylesheetTag(url string) string {
	return fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, url)
}

func makeScriptTag(url string) string {
	return fmt.Sprintf(`<script type="module" src="%s"></script>`, url)
}

func isCssPath(path string) bool {
	return regexp.MustCompile(`\.(css|less|sass|scss|styl|stylus|pcss|postcss)$`).MatchString(path)
}

func hotAsset(asset string) string {
	data, err := os.ReadFile(hotFile())
	if err != nil {
		panic(err)
	}
	return string(data) + "/" + asset
}

func isRunningHot() bool {
	filename := hotFile()

	if filename == "" {
		return false
	}

	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func hotFile() string {
	return filepath.Join("public", "hot")
}
