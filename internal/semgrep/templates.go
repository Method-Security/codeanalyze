package semgrep

import (
	"fmt"
)

var semgrepTemplatesMap = map[string]string{
	"secrets": "--config p/secrets --config %s/generic/secrets/",
}

// CreateSemgrepConfigValue creates a semgrep config value based on the provided template and rules directory. If the
// template is not valid, an error is returned.
func CreateSemgrepConfigValue(template string, rulesDir string) (string, error) {
	if templateValue, exists := semgrepTemplatesMap[template]; exists {
		return fmt.Sprintf(templateValue, rulesDir), nil
	}

	return "", fmt.Errorf("invalid template value '%s', must be one of %v", template, semgrepTemplatesMap)
}
