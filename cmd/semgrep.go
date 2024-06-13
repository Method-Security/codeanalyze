package cmd

import (
	"fmt"

	"github.com/Method-Security/codeanalyze/internal/semgrep"
	"github.com/spf13/cobra"
)

func (a *CodeAnalyze) InitSastCommand() {
	a.SemgrepCmd = &cobra.Command{
		Use:   "semgrep",
		Short: "Run semgrep against code directory",
		Long:  `Run semgrep against code directory`,
		Run: func(cmd *cobra.Command, args []string) {
			// Validate targets flag
			target, err := cmd.Flags().GetString("target")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate the configType flag
			configType, err := cmd.Flags().GetString("config-type")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}
			validTypes := map[string]bool{"direct": true, "template": true}
			if _, valid := validTypes[configType]; !valid {
				err = fmt.Errorf("invalid config type '%s', must be one of 'direct', 'template'", configType)
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate the localRulesDir flag
			localRulesDir, err := cmd.Flags().GetString("local-rules-dir")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate the configValue flag
			configValue, err := cmd.Flags().GetString("config-value")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}
			var finalConfigValue string
			if configType == "template" {
				// Attempt to get the template string
				finalConfigValue, err = semgrep.CreateSemgrepConfigValue(configValue, localRulesDir)
				if err != nil {
					errorMessage := err.Error()
					a.OutputSignal.ErrorMessage = &errorMessage
					a.OutputSignal.Status = 1
					return
				}
			} else {
				finalConfigValue = configValue
			}

			report, err := semgrep.ExecuteSemgrep(cmd.Context(), target, finalConfigValue)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report

		},
	}

	a.SemgrepCmd.Flags().String("target", "", "Local folder or file code target to scan")
	a.SemgrepCmd.Flags().String("config-type", "", "SAST config type (direct|template), direct to write custom config string (e.g. --config p/secrets), template to use a pre defined built-in and custom rule set combo")
	a.SemgrepCmd.Flags().String("config-value", "", "SAST config value, either a string to be passed directly to semgrep CLI or a template value (e.g. secrets)")
	a.SemgrepCmd.Flags().String("local-rules-dir", "/opt/method/codeanalyze/var/conf/resources/semgrep/", "Absolute path to local semgrep rules directory")

	a.RootCmd.AddCommand(a.SemgrepCmd)
}
