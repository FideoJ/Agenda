package utils

import (
	"../err"
	"../logger"
	"github.com/spf13/cobra"
)

func GetNonEmptyString(cmd *cobra.Command, key string) string {
	value, _ := cmd.Flags().GetString(key)
	if value == "" {
		logger.FatalIf(err.RequireNonEmpty(key))
	}

	return value
}
