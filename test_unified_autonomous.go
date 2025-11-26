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
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                                ║")
	fmt.Println("║     🌳 ECHO9LLAMA UNIFIED AUTONOMOUS AGENT TEST 🌳            ║")
	fmt.Println("║                                                                ║")
	fmt.Println("║     Fully Autonomous Wisdom-Cultivating Deep Tree Echo AGI     ║")
	fmt.Println("║                                                                ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Get API keys from environment
	anthropicKey := os.Getenv("ANTHROPIC_API_KEY")
	openrouterKey := os.Getenv("OPENROUTER_API_KEY")

	if anthropicKey == "" {
		fmt.Println("⚠️  Warning: ANTHROPIC_API_KEY not set")
	}
	if openrouterKey == "" {
		fmt.Println("⚠️  Warning: OPENROUTER_API_KEY not set")
	}

	// Create unified autonomous agent
	fmt.Println("🔧 Creating unified autonomous agent...")
	agent, err := core.NewUnifiedAutonomousAgent(anthropicKey, openrouterKey)
	if err != nil {
		fmt.Printf("❌ Failed to create agent: %v\n", err)
		os.Exit(1)
	}

	// Start the agent
	fmt.Println("🚀 Starting autonomous operation...")
	if err := agent.Start(); err != nil {
		fmt.Printf("❌ Failed to start agent: %v\n", err)
		os.Exit(1)
	}

	// Set up graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Run for demonstration
	fmt.Println("\n📊 Agent is now running autonomously...")
	fmt.Println("   Press Ctrl+C to stop\n")

	// Periodic status updates
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-sigChan:
			fmt.Println("\n\n🛑 Shutdown signal received...")
			if err := agent.Stop(); err != nil {
				fmt.Printf("⚠️  Error stopping agent: %v\n", err)
			}
			
			// Print final metrics
			fmt.Println("\n📊 Final Agent Metrics:")
			metrics := agent.GetMetrics()
			fmt.Printf("   Uptime: %.0f seconds\n", metrics["uptime"])
			fmt.Printf("   Autonomous Cycles: %d\n", metrics["autonomous_cycles"])
			fmt.Printf("   Thoughts Generated: %d\n", metrics["thoughts_generated"])
			fmt.Printf("   Goals Achieved: %d\n", metrics["goals_achieved"])
			fmt.Printf("   Final State: %s\n", metrics["wake_state"])
			
			fmt.Println("\n✅ Agent stopped successfully")
			return

		case <-ticker.C:
			// Print periodic status
			metrics := agent.GetMetrics()
			fmt.Printf("\n📈 Status Update - Uptime: %.0fs | Cycles: %d | Thoughts: %d | Goals: %d/%d\n",
				metrics["uptime"],
				metrics["autonomous_cycles"],
				metrics["thoughts_generated"],
				metrics["goals_achieved"],
				metrics["active_goals"])
		}
	}
}
