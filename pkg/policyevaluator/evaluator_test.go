package policyevaluator_test

import (
	"github.com/harkirat22/suricata-listener/pkg/normalizer"
	"github.com/harkirat22/suricata-listener/pkg/policyevaluator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Policyevaluator", func() {
	Context("When evaluating a policy", func() {
		It("Should evaluate correctly", func() {
			logEntries := []normalizer.LogEntry{
				{
					Type:      "alert",
					Timestamp: "2023-09-20T12:34:56.789",
					SrcIP:     "192.168.0.1",
					Alert: normalizer.Alert{
						Signature: "Suspicious connection to port 13666",
					},
				},
				// You can add more entries as needed for diverse test cases
			}

			policy := `
			package suricata

			default alert := {}

			# Trigger an alert if the Signature field contains the specific string "Suspicious connection to port 13666".
			alert := {"src_ip": inp.SrcIP} {
				inp := input[_]
				inp.Type == "alert"
				contains(inp.Alert.Signature, "Suspicious connection to port 13666")
			}
			`

			results, err := policyevaluator.Evaluate(logEntries, policy)
			Expect(err).ToNot(HaveOccurred())

			// You might want to adjust this based on your policy logic and expected results.
			Expect(results).To(HaveLen(1))
			Expect(results[0].SrcIP).To(Equal("192.168.0.1"))
			// ... more assertions based on the expected policy output ...
		})
	})
})
