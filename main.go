package main

import (
	"log"

	"github.com/harkirat22/suricata-listener/pkg/normalizer"
	"github.com/harkirat22/suricata-listener/pkg/policyevaluator"
)

const logFilePath = "/path/to/fast.log"
const policy = `
package suricata

default alert = false

# Define your policy logic here.
# E.g., if entry.Details contains a specific string, trigger an alert.
alert {
	entry := input.Details
	entry == "specific string to watch"
}
`

func main() {
	entries, err := normalizer.Normalize(logFilePath)
	if err != nil {
		log.Fatalf("Error normalizing log: %v", err)
	}

	violations, err := policyevaluator.Evaluate(entries, policy)
	if err != nil {
		log.Fatalf("Error evaluating policy: %v", err)
	}

	for _, v := range violations {
		// Logic to kill the pod using the whipper package.
		// Example: whipper.KillPod("pod-name", "namespace")
		log.Printf("Policy violation: %v", v)
	}
}
