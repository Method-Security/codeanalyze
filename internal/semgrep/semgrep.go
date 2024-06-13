package semgrep

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Metadata struct {
	Category           string   `json:"category" yaml:"category"`
	Confidence         string   `json:"confidence" yaml:"confidence"`
	CWE                []string `json:"cwe" yaml:"cwe"`
	CWE2021Top25       bool     `json:"cwe2021-top25" yaml:"cwe2021-top25"`
	CWE2022Top25       bool     `json:"cwe2022-top25" yaml:"cwe2022-top25"`
	Impact             string   `json:"impact" yaml:"impact"`
	License            string   `json:"license" yaml:"license"`
	Likelihood         string   `json:"likelihood" yaml:"likelihood"`
	OWASP              []string `json:"owasp" yaml:"owasp"`
	References         []string `json:"references" yaml:"references"`
	Subcategory        []string `json:"subcategory" yaml:"subcategory"`
	Technology         []string `json:"technology" yaml:"technology"`
	VulnerabilityClass []string `json:"vulnerability_class" yaml:"vulnerability_class"`
}

type Dev struct {
	Origin    string `json:"origin" yaml:"origin"`
	RID       string `json:"r_id" yaml:"r_id"`
	RuleID    string `json:"rule_id" yaml:"rule_id"`
	RVID      string `json:"rv_id" yaml:"rv_id"`
	URL       string `json:"url" yaml:"url"`
	VersionID string `json:"version_id" yaml:"version_id"`
}

type Extra struct {
	EngineKind      string             `json:"engine_kind" yaml:"engine_kind"`
	Fingerprint     string             `json:"fingerprint" yaml:"fingerprint"`
	IsIgnored       bool               `json:"is_ignored" yaml:"is_ignored"`
	Lines           string             `json:"lines" yaml:"lines"`
	Message         string             `json:"message" yaml:"message"`
	Metadata        Metadata           `json:"metadata" yaml:"metadata"`
	Metavars        map[string]Metavar `json:"metavars" yaml:"metavars"`
	Severity        string             `json:"severity" yaml:"severity"`
	ValidationState string             `json:"validation_state" yaml:"validation_state"`
	SemgrepDev      Dev                `json:"semgrep.dev" yaml:"semgrep.dev"`
	Shortlink       string             `json:"shortlink" yaml:"shortlink"`
	Source          string             `json:"source" yaml:"source"`
	SourceRuleURL   string             `json:"source-rule-url" yaml:"source-rule-url"`
}

type Metavar struct {
	AbstractContent string `json:"abstract_content" yaml:"abstract_content"`
	End             struct {
		Col    int `json:"col" yaml:"col"`
		Line   int `json:"line" yaml:"line"`
		Offset int `json:"offset" yaml:"offset"`
	} `json:"end" yaml:"end"`
	Start struct {
		Col    int `json:"col" yaml:"col"`
		Line   int `json:"line" yaml:"line"`
		Offset int `json:"offset" yaml:"offset"`
	} `json:"start" yaml:"start"`
}

type Result struct {
	CheckID string `json:"check_id" yaml:"check_id"`
	Path    string `json:"path" yaml:"path"`
	Start   struct {
		Line   int `json:"line" yaml:"line"`
		Col    int `json:"col" yaml:"col"`
		Offset int `json:"offset" yaml:"offset"`
	} `json:"start" yaml:"start"`
	End struct {
		Line   int `json:"line" yaml:"line"`
		Col    int `json:"col" yaml:"col"`
		Offset int `json:"offset" yaml:"offset"`
	} `json:"end" yaml:"end"`
	Extra Extra `json:"extra" yaml:"extra"`
}

type Error struct {
	Code    int    `json:"code" yaml:"code"`
	Level   string `json:"level" yaml:"level"`
	Message string `json:"message" yaml:"message"`
	Type    string `json:"type" yaml:"type"`
}

type Report struct {
	Errors  []Error  `json:"errors" yaml:"errors"`
	Results []Result `json:"results" yaml:"results"`
}

func ExecuteSemgrep(ctx context.Context, target string, configValue string) (Report, error) {
	// Prepare the semgrep arguments
	args := strings.Fields(configValue)
	args = append(args, target)
	args = append(args, "--json")
	args = append(args, "--quiet")

	cmd := exec.Command("semgrep", args...)

	// Create a pipe to capture the output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return Report{}, fmt.Errorf("failed to create stdout pipe: %w", err)
	}
	cmd.Stderr = cmd.Stdout

	// Start the command
	if err := cmd.Start(); err != nil {
		return Report{}, fmt.Errorf("failed to start command: %w", err)
	}

	// Read the command output
	var report Report
	decoder := json.NewDecoder(stdout)
	if err := decoder.Decode(&report); err != nil {
		return Report{}, fmt.Errorf("failed to decode json: %w", err)
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return Report{}, fmt.Errorf("error running semgrep: %w", err)
	}

	return report, nil
}
