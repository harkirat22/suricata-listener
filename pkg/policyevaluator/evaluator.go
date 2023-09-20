package policyevaluator

import (
	"context"

	"github.com/open-policy-agent/opa/rego"
	"github.com/your-repo-name/pkg/normalizer"
)

// Evaluate evaluates the log entries against the provided OPA policy.
func Evaluate(logEntries []normalizer.LogEntry, policy string) ([]normalizer.LogEntry, error) {
	ctx := context.Background()

	// Prepare OPA rego query
	r := rego.New(
		rego.Query("data.suricata.alert"),
		rego.Module("suricata_policy.rego", policy),
	)

	// Compile the module. Check for errors.
	query, err := r.PrepareForEval(ctx)
	if err != nil {
		return nil, err
	}

	var violations []normalizer.LogEntry
	for _, entry := range logEntries {
		// Here, we use the entry details for evaluation. Adjust as needed.
		results, err := query.Eval(ctx, rego.EvalInput(entry))
		if err != nil {
			return nil, err
		}
		// If the policy is violated, add to the violations list.
		if len(results) > 0 {
			violations = append(violations, entry)
		}
	}

	return violations, nil
}
