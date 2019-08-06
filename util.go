package proxy

import (
	"os"
	"strings"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetListenAddress() string {
	port := GetEnv("PORT", "1338")
	return ":" + port
}

func GetProxyURL(proxyConditionRaw string) string {
	proxyCondition := strings.ToUpper(proxyConditionRaw)

	a_condition_url := os.Getenv("A_CONDITION_URL")
	b_condition_url := os.Getenv("B_CONDITION_URL")
	default_condition_url := os.Getenv("DEFAULT_CONDITION_URL")

	// check for the condition url
	switch proxyCondition {
	case "A":
		return a_condition_url
	case "B":
		return b_condition_url
	default:
		return default_condition_url
	}
}
