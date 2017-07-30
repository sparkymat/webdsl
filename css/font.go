package css

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css/font"
)

func FontFace(name string, defaultSource string, sources map[string]string, weight font.Weight, style font.Style) RuleSet {
	sourceMaps := []string{}

	for sourceType, sourceUrl := range sources {
		sourceMaps = append(sourceMaps, fmt.Sprintf("url('%v') format('%v')", sourceUrl, sourceType))
	}

	return For(FontFaceSelector).Set(
		FontFamily(name),
		Src(fmt.Sprintf("url('%v')", defaultSource)),
		Src(strings.Join(sourceMaps, ", ")),
		FontWeight(weight),
		FontStyle(style),
	)
}
