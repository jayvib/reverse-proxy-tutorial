package proxy

import (
	"os"

	"github.com/sirupsen/logrus"
)

func LogSetup() {
	a_condition_url := os.Getenv("A_CONDITION_URL")
	b_condition_url := os.Getenv("B_CONDITION_URL")
	default_condition_url := os.Getenv("DEFAULT_CONDITION_URL")

	logrus.Printf("Server will run on: %s\n", GetListenAddress())
	logrus.Printf("Redirecting to A url: %s\n", a_condition_url)
	logrus.Printf("Redirecting to B url: %s\n", b_condition_url)
	logrus.Printf("Redirecting to DEFAULT url: %s\n", default_condition_url)
}

func LogRequestPayload(payload RequestPayloadStruct, proxyUrl string) {
	logrus.Printf("proxy_condition: %s, proxy_url: %s\n", payload.ProxyCondition, proxyUrl)
}
