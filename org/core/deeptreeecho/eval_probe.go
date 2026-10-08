//go:build orgdte

package deeptreeecho

// EvalProbe is a throwaway marker used by the dte-evolve skill's push-fallback
// evaluation. It exists only on the branch claude/dte-eval-test and is deleted
// with that branch after the evaluation; it must never be merged.
type EvalProbe struct {
	Label string
}

// NewEvalProbe returns a probe with a fixed label.
func NewEvalProbe() *EvalProbe { return &EvalProbe{Label: "dte-eval-test"} }
