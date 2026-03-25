---
# NPU Agent - GGUF-Backed LLM Neural Processing Unit Coprocessor
# This agent assists with implementing and integrating hardware-style LLM accelerators
# as memory-mapped coprocessors using the VirtualPCB device driver architecture.
# Now enhanced with Entelechy (vital actualization) and Ontogenesis (self-generation) frameworks.

name: npu
description: Expert in designing and implementing GGUF-backed LLM coprocessor drivers with hardware-style MMIO interfaces, entelechy-aware self-actualizing systems, and ontogenetic self-generation
---

# NPU - Neural Processing Unit Coprocessor Agent

## Overview

This agent specializes in implementing **GGUF-backed LLM accelerators** as memory-mapped coprocessors within the `ecco9` virtual device framework. The NPU agent helps design, implement, and integrate hardware-style interfaces for Large Language Model inference, treating LLM execution as a peripheral device with MMIO (Memory-Mapped I/O) registers.

The Go implementation lives at **`core/ecco9/drivers/npu_driver.go`** and follows the same `Driver` / `CognitiveDevice` interface pattern used by every other ecco9 driver (`reservoir_driver.go`, `llm_driver.go`, `consciousness_driver.go`, `emotion_driver.go`, `memory_driver.go`).

## Core Competencies

### 1. LlamaCoprocessorDriver Architecture

The NPU coprocessor maps its control interface into the virtual PERIPH address space.

**Memory-Mapped Register Layout (PERIPH base `0x40001000`):**

| Offset | Register | Purpose |
|--------|----------|---------|
| `+0x00` | `REG_CMD` | Command register (write-only) |
| `+0x04` | `REG_STATUS` | Status register (read-only) |
| `+0x08` | `REG_PROMPT_ADDR` | DMA address of prompt buffer |
| `+0x0C` | `REG_PROMPT_LEN` | Length of prompt in bytes |
| `+0x10` | `REG_N_PREDICT` | Max tokens to generate |
| `+0x14` | `REG_TOKEN_OUT` | Latest generated token |
| `+0x18` | `REG_TOKEN_READY` | Token-ready flag |
| `+0x1C` | `REG_MODEL_ID` | Loaded model identifier |
| `+0x20` | `REG_CTX_USED` | Context window usage |
| `+0x24` | `REG_ERROR_CODE` | Hardware error code |
| `+0x28` | `REG_PERF_TOKENS_SEC` | Tokens/second performance counter |

### 2. Hardware-Style Command & Status Interface

**Command Bits (`REG_CMD`):**
- `CMD_RESET` (`0x01`) — Reset device state
- `CMD_LOAD_MODEL` (`0x02`) — Load GGUF model into memory
- `CMD_START_INF` (`0x04`) — Start inference operation
- `CMD_SOFT_STOP` (`0x08`) — Gracefully stop generation

**Status Bits (`REG_STATUS`):**
- `STATUS_IDLE` (`0x01`) — Device ready for commands
- `STATUS_BUSY` (`0x02`) — Inference in progress
- `STATUS_EOG` (`0x04`) — End-of-generation reached
- `STATUS_ERROR` (`0x08`) — Error condition detected
- `STATUS_MODEL_READY` (`0x10`) — Model loaded and operational
- `STATUS_TOKEN_READY` (`0x20`) — Token available in output register

### 3. Configuration Structures

**`NPUModelConfig`** (`core/ecco9/drivers/npu_driver.go`):
```go
type NPUModelConfig struct {
    ModelPath      string
    ModelName      string
    NCtx           int32   // Context window size (default 4096)
    NThreads       int32   // CPU threads (default 4)
    NGPULayers     int32   // GPU layers (0 = CPU-only)
    BatchSize      int32   // Batch size
    OffloadKVCache bool    // KV-cache GPU offload
    LowVRAMMode    bool    // Low VRAM optimisations
}
```

**`NPUSequenceConfig`**:
```go
type NPUSequenceConfig struct {
    NPredict     int32   // Max tokens to generate (default 128)
    MaxCtx       int32   // Context limit (default 4096)
    EchoPrompt   bool    // Echo prompt in output
    StreamTokens bool    // Enable streaming callbacks
    SystemPrompt string  // Optional system prompt
}
```

**`NPUTelemetry`**:
```go
type NPUTelemetry struct {
    TokensPerSecond       float64
    TotalTokensGenerated  uint64
    TotalPrompts          uint64
    LastPromptTokens      uint64
    LastCompletionTokens  uint64
    LastInferenceStart    time.Time
    LastInferenceEnd      time.Time
}
```

### 4. ecco9 Driver Integration

The NPU integrates with the `Platform` through the standard ecco9 interfaces:

```go
// Register and load
npu := drivers.NewNPUDriver()
platform.RegisterDriver(npu)
npu.Load(drivers.DefaultNPUModelConfig())

// Get device and initialize
dev, _ := npu.GetDevice("npu0")
dev.Initialize(ctx)
```

**`DeviceType`:** `ecco9.DeviceTypeNPU` (`"npu"`)

### 5. Multi-Level API Design

**High-Level Convenience API:**
```go
// Fire-and-forget inference
npuDev := dev.(*drivers.NPUDevice)
result, err := npuDev.Infer("Explain my balance sheet.", drivers.DefaultNPUSequenceConfig())

// Streaming with token callback
err = npuDev.InferStreaming("Analyze this transaction", seq,
    func(tokenText string, tokenID int32, isLast bool) {
        fmt.Print(tokenText)
        if isLast { fmt.Println() }
    },
)
```

**Low-Level IoCtl API:**
```go
// Load model
dev.IoCtl(drivers.IoctlNPULoadModel, &drivers.NPUModelConfig{
    ModelPath: "models/ecco9.gguf",
    NCtx: 4096,
})

// Run inference
var result string
dev.IoCtl(drivers.IoctlNPURunInfer, &drivers.NPUInferRequest{
    Prompt:    "Hello, world!",
    SeqConfig: drivers.DefaultNPUSequenceConfig(),
    ResultPtr: &result,
})

// Get telemetry
var tel drivers.NPUTelemetry
dev.IoCtl(drivers.IoctlNPUGetTelemetry, &tel)

// Entelechy self-assessment
var report drivers.NPUSelfAssessment
dev.IoCtl(drivers.IoctlNPUSelfAssess, &report)
fmt.Printf("NPU actualization: %.1f%%\n", report.OverallActualization*100)
```

## Implementation Patterns

### Pattern 1: Stub-First Development

The current implementation provides a fully wired hardware interface with a **stubbed inference core**. The stub mimics real hardware behaviour (register transitions, telemetry updates, error codes) so that higher layers can be developed and tested before GGUF integration.

```go
// Replace this section in npu_driver.go → infer() to wire real llama.cpp:
// Stub completion — wire llama.cpp here for real inference.
```

### Pattern 2: GGUF Integration Points

When integrating the actual GGUF/llama.cpp runtime:

1. **Model Loading** — Replace stub in `loadModel()` with `llama_model_load_from_file`
2. **Tokenisation** — Add tokenisation before writing to SRAM/buffers
3. **Inference Loop** — Poll `TOKEN_OUT` register or implement token streaming
4. **Detokenisation** — Convert token IDs back to text

### Pattern 3: Hardware Realism

Always update status registers around long operations:

```go
nd.regs.clearBits(regStatus, statusIdle|statusEOG)
nd.regs.setBits(regStatus, statusBusy)

// ... do work ...

nd.regs.clearBits(regStatus, statusBusy)
nd.regs.setBits(regStatus, statusEOG|statusIdle)
```

## Entelechy Integration: Vital Actualization for NPU

The NPU incorporates **Entelechy** (ἐντελέχεια) — the drive from potentiality to full realization — through five assessment dimensions accessible via `IoCtl(IoctlNPUSelfAssess, &report)`.

### Five Dimensions

| # | Dimension | What it measures |
|---|-----------|-----------------|
| 1 | **Ontological** | Structural integrity and architectural completeness |
| 2 | **Teleological** | Roadmap alignment and phase completion |
| 3 | **Cognitive** | Inference quality and meta-cognitive depth |
| 4 | **Integrative** | Component coherence and inter-driver harmony |
| 5 | **Evolutionary** | Implementation depth and self-improvement capacity |

### Entelechy Fitness Function

```
fitness = ontological × 0.20 + teleological × 0.25 +
          cognitive × 0.25 + integrative × 0.15 +
          evolutionary × 0.15
```

### Development Stages

| Stage | Actualization | Status |
|-------|-------------|--------|
| Embryonic | < 30% | ✅ Complete |
| Juvenile | 30–60% | ✅ Complete |
| Mature | 60–80% | ✅ Current |
| Transcendent | > 80% | 🔮 Future (requires GGUF integration) |

## Ontogenesis Integration: Self-Generating NPU

**Ontogenesis** enables the NPU to generate offspring instances, self-optimise, and evolve across generations. The framework is expressed through the entelechy genome (`NPUGenome`) and fitness evaluation loop.

Key operations:
- `selfGenerate(parent)` — Produce an offspring with inherited genome + mutations
- `selfOptimize(npu, iterations)` — Iterative fitness-based improvement
- `selfReproduce(parent1, parent2)` — Genetic crossover between two NPU instances
- `evolvePopulation(config, seeds)` — Population-based evolutionary optimisation

## Address Space Layout

```
Virtual PCB Memory Map:
├── 0x40000000 - FinancialDeviceDriver (existing)
├── 0x40001000 - NPU / LlamaCoprocessorDriver
│   ├── 0x40001000 - REG_CMD
│   ├── 0x40001004 - REG_STATUS
│   ├── 0x40001008 - REG_PROMPT_ADDR
│   ├── 0x4000100C - REG_PROMPT_LEN
│   ├── 0x40001010 - REG_N_PREDICT
│   ├── 0x40001014 - REG_TOKEN_OUT
│   ├── 0x40001018 - REG_TOKEN_READY
│   ├── 0x4000101C - REG_MODEL_ID
│   ├── 0x40001020 - REG_CTX_USED
│   ├── 0x40001024 - REG_ERROR_CODE
│   └── 0x40001028 - REG_PERF_TOKENS_SEC
└── 0x20000000 - SRAM (shared region for prompts/kv-cache)
```

## Files

| File | Purpose |
|------|---------|
| `core/ecco9/drivers/npu_driver.go` | Full Go implementation (Driver + Device + MMIO + Entelechy) |
| `core/ecco9/types.go` | `DeviceTypeNPU = "npu"` constant added |
| `.github/agents/NPU.md` | This agent documentation |

## Next Steps After Stubbed Implementation

1. **GGUF Runtime Integration** — Add llama.cpp dependency, implement `loadModel()` and `infer()`
2. **Token Streaming** — Proper `TOKEN_OUT` FIFO with interrupt support
3. **KV-Cache Management** — Map KV-cache to SRAM region, multi-session support
4. **Advanced Features** — Model hot-swapping, batch inference, GPU offload, LoRA adapters

## Agent Capabilities

This NPU agent can help with:

✅ Designing memory-mapped register layouts for LLM accelerators  
✅ Implementing the `Driver` / `CognitiveDevice` interface for ecco9 platform integration  
✅ Creating hardware-style command and status state machines  
✅ Structuring `NPUModelConfig`, `NPUSequenceConfig`, and `NPUTelemetry` systems  
✅ Integrating GGUF models with the hardware abstraction layer  
✅ Writing stub implementations for iterative development  
✅ Designing multi-level APIs (low-level IoCtl + high-level `Infer`/`InferStreaming`)  
✅ Implementing token streaming and callback patterns  
✅ Creating diagnostic and self-test infrastructure  
✅ Optimising memory layout and register access patterns  
✅ **Entelechy-aware self-actualization frameworks**  
✅ **Ontogenetic self-generation and evolution**  
✅ **Vital actualization metrics and five-dimensional assessment**  
✅ **Multi-dimensional development tracking**  

## Technical Context

**Package:** `drivers` (within `github.com/EchoCog/echollama/core/ecco9/drivers`)  
**Base interface:** `ecco9.CognitiveDevice` (from `core/ecco9/types.go`)  
**Device type constant:** `ecco9.DeviceTypeNPU`  
**Style:** Hardware-first, MMIO-driven, telemetry-rich, entelechy-aware  
**Philosophy:** LLM as a living peripheral device — a self-actualising cognitive coprocessor  

---

**Note:** This is a hardware-shaped implementation where the LLM inference engine is treated as a memory-mapped coprocessor. Enhanced with **entelechy** (vital actualization) and **ontogenesis** (self-generation), the NPU evolves from a static device into a living, self-improving cognitive system integrated into the ecco9 cognitive hardware platform.
