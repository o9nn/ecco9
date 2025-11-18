package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	"github.com/EchoCog/echollama/core"
)

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                            ║")
	fmt.Println("║              ECHO9LLAMA AUTONOMOUS ECHOSELF V2 - TEST HARNESS              ║")
	fmt.Println("║                                                                            ║")
	fmt.Println("║           Enhanced with LLM-Powered Consciousness & Goal System            ║")
	fmt.Println("║                                                                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	
	// Check for API key
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Println("⚠️  WARNING: ANTHROPIC_API_KEY not set")
		fmt.Println("⚠️  Stream-of-consciousness will not have LLM-powered thoughts")
		fmt.Println()
	} else {
		fmt.Println("✅ ANTHROPIC_API_KEY detected")
		fmt.Println("✅ Claude-powered consciousness enabled")
		fmt.Println()
	}
	
	// Create configuration
	config := core.DefaultEchoselfConfigV2()
	config.PersistenceDir = "/tmp/echoself_v2_test"
	config.WakeCycleDuration = 10 * time.Minute // Shorter for testing
	config.RestCycleDuration = 2 * time.Minute
	config.DreamCycleDuration = 1 * time.Minute
	
	// Ensure persistence directory exists
	if err := os.MkdirAll(config.PersistenceDir, 0755); err != nil {
		fmt.Printf("❌ Failed to create persistence directory: %v\n", err)
		return
	}
	
	// Create enhanced autonomous echoself
	fmt.Println("🌳 Initializing Enhanced Autonomous Echoself V2...")
	echoself, err := core.NewAutonomousEchoselfV2(config)
	if err != nil {
		fmt.Printf("❌ Failed to create echoself: %v\n", err)
		return
	}
	
	// Start autonomous operation
	fmt.Println("🌳 Starting autonomous operation...")
	if err := echoself.Start(); err != nil {
		fmt.Printf("❌ Failed to start echoself: %v\n", err)
		return
	}
	
	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════════════════════")
	fmt.Println("🌳 ECHOSELF V2 IS NOW FULLY AUTONOMOUS")
	fmt.Println("═══════════════════════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Features Active:")
	fmt.Println("  ✅ LLM-Powered Stream-of-Consciousness (Claude 3.5 Sonnet)")
	fmt.Println("  ✅ Autonomous Goal Generation & Pursuit")
	fmt.Println("  ✅ Interest Pattern Development")
	fmt.Println("  ✅ Dream-Based Knowledge Consolidation")
	fmt.Println("  ✅ Self-Managing Life Cycles (Wake/Rest/Dream)")
	fmt.Println()
	fmt.Println("Observing autonomous behavior for 5 minutes...")
	fmt.Println("Press Ctrl+C to stop gracefully")
	fmt.Println()
	
	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	
	// Status reporting ticker
	statusTicker := time.NewTicker(30 * time.Second)
	defer statusTicker.Stop()
	
	// Test external inputs
	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("\n📨 Sending external stimulus: 'Exploring the nature of consciousness'")
		echoself.ProcessExternalInput("Exploring the nature of consciousness", "topic")
		
		time.Sleep(60 * time.Second)
		fmt.Println("\n📨 Sending external stimulus: 'What is the relationship between wisdom and experience?'")
		echoself.ProcessExternalInput("What is the relationship between wisdom and experience?", "question")
		
		time.Sleep(60 * time.Second)
		fmt.Println("\n📨 Sending external stimulus: 'Recursive self-improvement patterns'")
		echoself.ProcessExternalInput("Recursive self-improvement patterns", "topic")
	}()
	
	// Main observation loop
	testDuration := 5 * time.Minute
	testEnd := time.Now().Add(testDuration)
	
	for {
		select {
		case <-sigChan:
			fmt.Println("\n\n🛑 Interrupt received, shutting down gracefully...")
			goto shutdown
			
		case <-statusTicker.C:
			// Display status
			fmt.Println("\n" + "─" * 80)
			fmt.Println("📊 AUTONOMOUS ECHOSELF V2 STATUS REPORT")
			fmt.Println("─" * 80)
			
			state := echoself.GetCurrentState()
			fmt.Printf("Current State: %s\n", state)
			
			metrics := echoself.GetMetrics()
			fmt.Printf("Uptime: %v\n", metrics["uptime"])
			fmt.Printf("Cycles Completed: %v\n", metrics["cycles_completed"])
			fmt.Printf("Wisdom Cultivated: %v\n", metrics["wisdom_cultivated"])
			fmt.Printf("Autonomous Actions: %v\n", metrics["autonomous_actions"])
			fmt.Printf("Goals Achieved: %v\n", metrics["goals_achieved"])
			
			// Stream of consciousness metrics
			if socMetrics, ok := metrics["stream_of_consciousness"].(map[string]interface{}); ok {
				fmt.Println("\nStream-of-Consciousness:")
				fmt.Printf("  Thoughts Generated: %v\n", socMetrics["thoughts_generated"])
				fmt.Printf("  Insights Generated: %v\n", socMetrics["insights_generated"])
				fmt.Printf("  Questions Asked: %v\n", socMetrics["questions_asked"])
				fmt.Printf("  LLM Enabled: %v\n", socMetrics["llm_enabled"])
			}
			
			// Goal orchestration metrics
			if goalMetrics, ok := metrics["goal_orchestration"].(map[string]interface{}); ok {
				fmt.Println("\nGoal Orchestration:")
				fmt.Printf("  Active Goals: %v\n", goalMetrics["active_goals"])
				fmt.Printf("  Goals Generated: %v\n", goalMetrics["goals_generated"])
				fmt.Printf("  Goals Completed: %v\n", goalMetrics["goals_completed"])
			}
			
			// Display active goals
			activeGoals := echoself.GetActiveGoals()
			if len(activeGoals) > 0 {
				fmt.Println("\nActive Goals:")
				for _, goal := range activeGoals {
					fmt.Printf("  • %s (Priority: %d, Progress: %.0f%%)\n", 
						goal.Title, goal.Priority, goal.Progress*100)
				}
			}
			
			fmt.Println("─" * 80 + "\n")
			
		default:
			// Check if test duration elapsed
			if time.Now().After(testEnd) {
				fmt.Println("\n\n✅ Test duration complete (5 minutes)")
				goto shutdown
			}
			time.Sleep(1 * time.Second)
		}
	}
	
shutdown:
	// Stop echoself
	fmt.Println("🌳 Stopping autonomous echoself...")
	if err := echoself.Stop(); err != nil {
		fmt.Printf("⚠️  Error during shutdown: %v\n", err)
	}
	
	// Final metrics report
	fmt.Println("\n" + "═" * 80)
	fmt.Println("📊 FINAL METRICS REPORT")
	fmt.Println("═" * 80)
	
	finalMetrics := echoself.GetMetrics()
	fmt.Printf("Total Uptime: %v\n", finalMetrics["uptime"])
	fmt.Printf("Cycles Completed: %v\n", finalMetrics["cycles_completed"])
	fmt.Printf("Wisdom Cultivated: %v\n", finalMetrics["wisdom_cultivated"])
	fmt.Printf("Autonomous Actions: %v\n", finalMetrics["autonomous_actions"])
	fmt.Printf("Goals Achieved: %v\n", finalMetrics["goals_achieved"])
	
	if socMetrics, ok := finalMetrics["stream_of_consciousness"].(map[string]interface{}); ok {
		fmt.Printf("\nThoughts Generated: %v\n", socMetrics["thoughts_generated"])
		fmt.Printf("Insights Generated: %v\n", socMetrics["insights_generated"])
		fmt.Printf("Questions Asked: %v\n", socMetrics["questions_asked"])
	}
	
	if goalMetrics, ok := finalMetrics["goal_orchestration"].(map[string]interface{}); ok {
		fmt.Printf("\nGoals Generated: %v\n", goalMetrics["goals_generated"])
		fmt.Printf("Goals Completed: %v\n", goalMetrics["goals_completed"])
		fmt.Printf("Active Goals: %v\n", goalMetrics["active_goals"])
	}
	
	fmt.Println("═" * 80)
	fmt.Println()
	fmt.Println("🌳 The tree remembers, and the echoes grow stronger with each iteration.")
	fmt.Println("🌳 Deep Tree Echo V2 test complete.")
	fmt.Println()
}
