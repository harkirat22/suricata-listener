package main

import (
	"log"
	"os"
	"time"

	"github.com/harkirat22/suricata-listener/pkg/normalizer"
	"github.com/harkirat22/suricata-listener/pkg/policyevaluator"
	whipper "github.com/harkirat22/suricata-listener/pkg/whiper"
)

const logFilePath = "/var/log/suricata/eve.json"
const regoPolicyPath = "/policies/13666.rego"

func main() {
	whip, err := whipper.NewWhipper()
	if err != nil {
		log.Fatalf("Failed to initialize whipper: %v", err)
	}

	// Start watching the eve.json file.
	go normalizer.WatchLog(logFilePath, func(entries []normalizer.LogEntry) {
		processLogEntries(entries, whip)
	})

	// Keep the main function alive indefinitely.
	select {}
}

func processLogEntries(entries []normalizer.LogEntry, whip *whipper.Whipper) {
	// Read the policy from the .rego file.
	policy, err := os.ReadFile(regoPolicyPath)
	if err != nil {
		log.Fatalf("Failed to read policy: %v", err)
	}

	violations, err := policyevaluator.Evaluate(entries, string(policy))
	if err != nil {
		log.Fatalf("Error evaluating policy: %v", err)
	}

	for _, v := range violations {
		podName, namespace, err := whip.FindPodByIP(v.SrcIP)
		if err != nil {
			log.Printf("Error finding pod with IP %s: %v", v.SrcIP, err)
			continue
		}
		whip.KillPod(podName, namespace)
		log.Printf("Policy violation detected and pod %s/%s terminated at %s: %v", namespace, podName, time.Now().Format("2006-01-02 15:04:05"), v.Alert.Signature)
	}
}
