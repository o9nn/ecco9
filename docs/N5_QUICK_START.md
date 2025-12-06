# Quick Start: Iteration N+5 Components

## Overview

Iteration N+5 adds meta-cognitive awareness and advanced reasoning to Deep Tree Echo. This guide helps you quickly integrate and use the new components.

## Installation

```bash
# Clone the repository
git clone https://github.com/o9nn/ecco9.git
cd ecco9

# Checkout the N+5 implementation
git checkout copilot/implement-next-steps

# Build the project
go build ./...
```

## Quick Usage Examples

### 1. Using the Integrated N5 System (Recommended)

```go
package main

import (
    "fmt"
    "github.com/EchoCog/echollama/core/deeptreeecho"
)

func main() {
    // Create and start the integrated N5 system
    n5 := deeptreeecho.NewN5CognitiveSystem()
    n5.Start()
    defer n5.Stop()
    
    // Process a task with full meta-cognition
    result := n5.ProcessWithMetaCognition("How can I improve my learning efficiency?")
    
    fmt.Printf("Strategy Used: %s\n", result.StrategyUsed)
    fmt.Printf("Wisdom Applied: %d entries\n", result.WisdomApplied)
    fmt.Printf("Meta-Insight: %s\n", result.MetaInsight)
    fmt.Printf("Success: %v\n", result.Success)
    
    // Reason about a complex problem
    reasoningResult := n5.ReasonAboutProblem(
        "Design an optimal study schedule",
        deeptreeecho.ProblemPlanning,
    )
    
    fmt.Printf("Sub-problems identified: %d\n", reasoningResult.SubProblems)
    fmt.Printf("Causal conclusion: %s\n", reasoningResult.CausalConclusion)
    fmt.Printf("Final answer: %s\n", reasoningResult.Answer)
    fmt.Printf("Confidence: %.2f\n", reasoningResult.Confidence)
    
    // Apply wisdom to a context
    wisdomResult := n5.ApplyWisdomToContext("learning new programming languages")
    
    fmt.Printf("Wisdom found: %d\n", wisdomResult.WisdomFound)
    fmt.Printf("Applied: %s\n", wisdomResult.WisdomApplied)
    fmt.Printf("Effectiveness: %.2f\n", wisdomResult.Effectiveness)
    
    // Get comprehensive system status
    status := n5.GetSystemStatus()
    fmt.Printf("System Status: %+v\n", status)
}
```

### 2. Using Individual Components

#### Meta-Cognitive Monitor

```go
import "github.com/EchoCog/echollama/core/deeptreeecho"

// Create monitor
monitor := deeptreeecho.NewMetaCognitiveMonitor()

// Track a cognitive process
processID := monitor.StartProcess("complex_reasoning", deeptreeecho.ProcessReasoning)
monitor.UpdateProcess(processID, 0.5, deeptreeecho.StatusActive)
monitor.UpdateProcess(processID, 1.0, deeptreeecho.StatusCompleted)

// Record and assess a decision
decisionID := monitor.RecordDecision(
    "choosing learning strategy",
    "spaced repetition",
    "evidence shows it's most effective for long-term retention",
    []string{"spaced repetition", "massed practice", "interleaving"},
    0.85,
)

outcome := &deeptreeecho.DecisionOutcome{
    Success:         true,
    ActualBenefit:   0.9,
    ExpectedBenefit: 0.8,
}
monitor.AssessDecisionQuality(decisionID, outcome)

// Select optimal strategy for a task
strategy := monitor.SelectStrategy(
    deeptreeecho.ProcessProblemSolving,
    map[string]interface{}{
        "accuracy_required": true,
        "speed_required":    false,
    },
)
fmt.Printf("Selected strategy: %s\n", strategy.Name)

// Generate meta-thoughts
thought := monitor.GenerateMetaThought(
    "learning_process",
    "I'm thinking about how I learn",
    0, // depth 0
)
fmt.Printf("Meta-insight: %s\n", thought.Insight)

// Get self-awareness metrics
awareness := monitor.GetSelfAwareness()
fmt.Printf("Awareness: %+v\n", awareness)
```

#### Advanced Reasoning Engine

```go
import "github.com/EchoCog/echollama/core/deeptreeecho"

engine := deeptreeecho.NewAdvancedReasoningEngine()

// Build a reasoning chain
chainID := engine.StartReasoningChain("Find the best approach to learn Go")

engine.AddReasoningStep(chainID, deeptreeecho.StepDeduction,
    "Go emphasizes simplicity and readability",
    "Simple languages are easier to learn",
    "Therefore, focus on core concepts first",
    0.9,
)

engine.AddReasoningStep(chainID, deeptreeecho.StepInduction,
    "Successful Go developers use projects to learn",
    "Learning by doing is effective",
    "Build small projects while learning",
    0.8,
)

engine.CompleteReasoningChain(chainID, 
    "Best approach: Learn core concepts, then build small projects", 
    0.85,
)

// Decompose a complex problem
problem := engine.DecomposeProblem(
    "Build a distributed web application",
    deeptreeecho.ProblemDesign,
)
fmt.Printf("Sub-problems: %d\n", len(problem.SubProblems))

// Perform causal reasoning
inference := engine.PerformCausalReasoning(
    "Regular practice improves coding skills",
    []string{
        "Daily coding shows improvement",
        "Improvement correlates with practice frequency",
        "Controlled studies confirm the effect",
    },
)
fmt.Printf("Causal conclusion: %s\n", inference.Conclusion)

// Generate counterfactual scenarios
counterfactual := engine.GenerateCounterfactual(
    "Currently learning through tutorials",
    "If I learned through building real projects instead",
)
fmt.Printf("Alternative outcome: %s\n", counterfactual.PredictedOutcome)

// Chain-of-thought reasoning
thoughtChain := engine.ChainOfThought("What's the fastest way to become proficient in Go?")
fmt.Printf("Answer: %s\n", thoughtChain.Answer)
fmt.Printf("Reasoning steps: %d\n", len(thoughtChain.Thoughts))
```

#### Wisdom Application Engine

```go
import "github.com/EchoCog/echollama/core/deeptreeecho"

wisdomEngine := deeptreeecho.NewWisdomApplicationEngine()

// Find relevant wisdom
matches := wisdomEngine.FindRelevantWisdom("learning from mistakes", 3)

for _, match := range matches {
    fmt.Printf("Wisdom: %s\n", match.Wisdom.Content)
    fmt.Printf("Relevance: %.2f\n", match.RelevanceScore)
    fmt.Printf("Match type: %s\n", match.MatchType)
    fmt.Printf("Explanation: %s\n\n", match.Explanation)
}

// Apply wisdom
if len(matches) > 0 {
    application := wisdomEngine.ApplyWisdom(
        matches[0].WisdomID,
        "trying to debug a complex issue",
    )
    
    fmt.Printf("Effectiveness: %.2f\n", application.Effectiveness)
    fmt.Printf("Outcome: %s\n", application.Outcome)
    
    // Provide feedback
    wisdomEngine.ProvideFeedback(
        application.ID,
        0.9,
        "Very helpful for systematic debugging",
    )
}

// Synthesize new wisdom
if len(matches) >= 2 {
    sourceIDs := []string{matches[0].WisdomID, matches[1].WisdomID}
    synthesized := wisdomEngine.SynthesizeWisdom(sourceIDs)
    
    fmt.Printf("Synthesized wisdom: %s\n", synthesized.Content)
    fmt.Printf("Confidence: %.2f\n", synthesized.Confidence)
}

// Get metrics
metrics := wisdomEngine.GetWisdomMetrics()
fmt.Printf("Total wisdom entries: %v\n", metrics["total_wisdom"])
fmt.Printf("Success rate: %.2f\n", metrics["success_rate"])
```

#### Interest Evolution Engine

```go
import "github.com/EchoCog/echollama/core/echobeats"

evolution := echobeats.NewInterestEvolutionEngine()

// Apply reinforcement learning
interest := &echobeats.Interest{
    ID:       "golang_learning",
    Name:     "Learning Go Programming",
    Strength: 0.6,
}

outcome := &echobeats.EngagementOutcome{
    InterestID:   interest.ID,
    Reward:       0.8,
    LearningGain: 0.7,
    Satisfaction: 0.85,
    Competence:   0.1,
    NoveltyValue: 0.5,
}

evolution.ApplyReinforcement(interest, outcome)
fmt.Printf("New interest strength: %.2f\n", interest.Strength)

// Discover interest clusters
interests := []*echobeats.Interest{interest /* add more */}
clusters := evolution.DiscoverClusters(interests)

for _, cluster := range clusters {
    fmt.Printf("Cluster: %s\n", cluster.CoreTopic)
    fmt.Printf("Members: %d\n", len(cluster.Members))
    fmt.Printf("Strength: %.2f\n", cluster.Strength)
}

// Generate emergent interests
parent1 := interest
parent2 := &echobeats.Interest{
    ID:       "systems_programming",
    Name:     "Systems Programming",
    Strength: 0.7,
}

emergent := evolution.GenerateEmergentInterest(parent1, parent2)
fmt.Printf("Emergent interest: %s\n", emergent.Name)

// Consolidate during rest
evolution.ConsolidateDuringRest(interests)
```

## Running Tests

```bash
# Run all N+5 tests
go test ./core/deeptreeecho -v -run TestN5

# Run specific component tests
go test ./core/deeptreeecho -v -run TestMetaCognitive
go test ./core/deeptreeecho -v -run TestAdvancedReasoning
go test ./core/deeptreeecho -v -run TestWisdomApplication

# Run integration tests
go test ./core/deeptreeecho -v -run TestN5Integration
```

## Integration with Existing Code

### With V5 Autonomous Agent

```go
// In your existing autonomous agent
import "github.com/EchoCog/echollama/core/deeptreeecho"

type MyAgent struct {
    // Existing fields...
    n5System *deeptreeecho.N5CognitiveSystem
}

func (a *MyAgent) Initialize() {
    a.n5System = deeptreeecho.NewN5CognitiveSystem()
    a.n5System.Start()
}

func (a *MyAgent) ProcessTask(task string) {
    // Use N+5 meta-cognition
    result := a.n5System.ProcessWithMetaCognition(task)
    
    // Apply the selected strategy
    switch result.StrategyUsed {
    case "Deliberate Reasoning":
        a.useDeliberateApproach(task)
    case "Intuitive Processing":
        a.useIntuitiveApproach(task)
    }
    
    // Incorporate meta-insights
    a.updateSelfModel(result.MetaInsight)
}
```

### With EchoDream Cycles

```go
func (d *DreamCycle) ConsolidateMemories() {
    // During dream/rest cycles, consolidate interests
    interests := d.getActiveInterests()
    d.n5System.interestEvolution.ConsolidateDuringRest(interests)
    
    // Synthesize wisdom from recent experiences
    recentWisdom := d.getRecentWisdomIDs()
    if len(recentWisdom) >= 2 {
        synthesized := d.n5System.wisdomApplication.SynthesizeWisdom(recentWisdom)
        d.storeNewWisdom(synthesized)
    }
}
```

## Performance Tips

1. **Interest Evolution**: Run clustering during low-activity periods (O(n²) complexity)
2. **Meta-Cognition**: Process tracking has minimal overhead, use freely
3. **Reasoning**: Chain-of-thought is intensive, use for important decisions
4. **Wisdom**: Caching reduces lookups, clear cache when wisdom changes significantly

## Next Steps

- Read `docs/ITERATION_N_PLUS_5_GUIDE.md` for comprehensive documentation
- Review `ITERATION_N_PLUS_5_SUMMARY.md` for architecture details
- Explore test examples in `core/deeptreeecho/n5_components_test.go`
- Check `ITERATION_N_PLUS_4_FINAL_REPORT.md` for context on previous iteration

## Troubleshooting

**Build taking too long?**
- Dependencies may be downloading. Use `go mod download` first.

**Tests timing out?**
- Some tests may take time for LLM-based operations. Increase timeout if needed.

**Import errors?**
- Ensure you're using Go 1.21+
- Run `go mod tidy` to clean up dependencies

## Support

For issues or questions:
- Check existing documentation in `docs/`
- Review test examples for usage patterns
- Refer to iteration summary for architecture details

---

**Quick Start Version**: 1.0  
**Date**: December 6, 2025  
**Iteration**: N+5
