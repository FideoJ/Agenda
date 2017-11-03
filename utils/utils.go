package utils

import (
	"time"

	"../err"
	"../logger"
	"github.com/spf13/cobra"
)

const TimeLayout = "2006-01-02 15:04"

func GetNonEmptyString(cmd *cobra.Command, key string) string {
	value, _ := cmd.Flags().GetString(key)
	if value == "" {
		logger.FatalIf(err.RequireNonEmpty(key))
	}

	return value
}

func GetNonEmptyStringSlice(cmd *cobra.Command, key string) []string {
	value, _ := cmd.Flags().GetStringSlice(key)
	if len(value) == 0 {
		logger.FatalIf(err.RequireNonEmpty(key))
	}

	return value
}

func Overlapped(s1, e1, s2, e2 time.Time) bool {
	return !s2.Before(s1) && s2.Before(e1) ||
		e2.After(s1) && !e2.After(e1) ||
		s2.Before(s1) && e2.After(e1)
}
