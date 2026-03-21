# Project Overview

Whack is a 100% procedurally generated arena combat game built with Go 1.24+ and Ebitengine v2.9.8. Inspired by Super Smash Brothers, Street Fighter, and Mortal Kombat, the game generates every element—fighters, arenas, movesets, music, SFX, and visual effects—at runtime from a deterministic seed with zero embedded assets. As part of the opd-ai "W-Series" game suite, Whack delivers frame-data-driven combat, hitbox/hurtbox systems, procedural combo trees, ring-out/KO physics, and rollback-style netcode designed for high-latency (200–5000ms) network conditions including Tor/onion routing.

The project targets competitive fighting game enthusiasts who appreciate procedural variety. Every match can feature unique fighters and arenas while maintaining deterministic reproducibility—the same seed produces identical content across all platforms. The five-genre theming system (fantasy, sci-fi, horror, cyberpunk, post-apocalyptic) changes not just aesthetics but gameplay mechanics, giving each genre a distinctly different fighting experience.

The single-binary philosophy is paramount: the compiled client and server contain all content generation logic. No PNG, WAV, OGG, JSON, TTF, or other asset files should ever be added to the repository. All graphics are synthesized via procedural algorithms, all audio via oscillators and envelopes, all levels via parameterized generators.

## Sibling Repository Context

Whack is one of 8 sibling repositories in the opd-ai procedural game suite. All repos share architectural patterns, conventions, and will eventually share library packages. Code written for Whack should be structured to enable future extraction into shared libraries.

| Repo | Genre | Description |
|------|-------|-------------|
| `opd-ai/venture` | Co-op action-RPG | Top-down co-op adventure with procedural dungeons |
| `opd-ai/vania` | Metroidvania platformer | Side-scrolling exploration with ability gating |
| `opd-ai/velocity` | Galaga-like shooter | Vertical shooter with procedural enemy waves |
| `opd-ai/violence` | Raycasting FPS | First-person shooter with modding/scripting support |
| `opd-ai/way` | Battle-cart racer | Procedural tracks and combat racing |
| `opd-ai/wyrm` | First-person survival RPG | Survival mechanics with procedural world |
| `opd-ai/where` | Wilderness survival | Outdoor survival simulation |
| `opd-ai/whack` | Arena battle game | Platform fighter with frame-data combat (this repo) |

## Technical Stack

- **Primary Language**: Go 1.24.13
- **Game Framework**: Ebitengine v2.9.8 — 2D game engine with cross-platform + WASM support
- **Configuration**: Viper v1.21.0 — YAML configuration management
- **Testing**: Go standard `testing` package, table-driven tests, benchmarks
- **Build**: `go build ./...` for all packages; `go build -o whack-client ./cmd/client` and `go build -o whack-server ./cmd/server` for individual binaries

### Key Indirect Dependencies
- `ebitengine/gomobile` — Mobile platform support
- `ebitengine/purego` — Pure Go foreign function interface
- `golang.org/x/sync` — Extended synchronization primitives
- `golang.org/x/sys` — System-level operations
- `golang.org/x/text` — Text processing utilities

## Project Structure

Whack uses a **velocity-style** layout: `cmd/` + `pkg/` with Viper configuration.

```
cmd/
  client/           # Game client (Ebitengine window) - main.go
  server/           # Authoritative game server - main.go
pkg/
  arena/            # Platform layout, hazard placement, blast zones
  audio/            # Oscillators, envelopes, motif system, audio playback
  combat/           # CombatSystem, KnockbackSystem, frame data resolution
  config/           # Viper-based configuration loading
  engine/           # ECS core (World, Entity, ComponentStore), EventBus, Scheduler
  fighter/          # Fighter entity, ComboTree, moveset logic
  network/          # TCP server, client, NetworkSyncSystem, input frames
  physics/          # PhysicsSystem with gravity and collision
  procgen/          # Generator interface and all content generators
  rendering/        # SpriteRenderer, Palette, PostProcessor
  tournament/       # Bracket generation, spectator state
config.yaml         # Default configuration file
```

---

## ⚠️ CRITICAL: Complete Feature Integration (Zero Dangling Features)

**This is the single most important rule for this codebase.** Every feature, system, component, generator, and integration MUST be fully wired into the runtime. Dangling features are a maintenance burden, a source of frustration, and actively degrade code quality.

### The Dangling Feature Problem

In complex procedural game codebases, it is extremely common for features to be:
1. **Defined but never instantiated** — A system struct exists but is never created in `main()` or system registration
2. **Instantiated but never integrated** — A system runs but its output is never consumed by other systems
3. **Partially integrated** — A system works for one genre/theme but silently no-ops for others
4. **Tested in isolation but broken in context** — Unit tests pass but the system was never wired into the game loop

### Current Integration Gaps (Known Issues)

Based on codebase analysis, the following systems are currently **skeleton implementations** that require full integration:

| System/Generator | File | Status | Integration Path Needed |
|------------------|------|--------|------------------------|
| `CombatSystem` | `pkg/combat/combat.go` | Skeleton | Register with Scheduler in client, implement hitbox/hurtbox intersection |
| `KnockbackSystem` | `pkg/combat/combat.go` | Skeleton | Register with Scheduler, consume CombatSystem damage events |
| `PhysicsSystem` | `pkg/physics/physics.go` | Skeleton | Register with Scheduler, integrate with arena platforms |
| `AudioSystem` | `pkg/audio/audio.go` | Skeleton | Initialize in client, wire to EventBus for hit/KO sounds |
| `SpriteRenderer` | `pkg/rendering/rendering.go` | Skeleton | Call in Draw(), feed with generator outputs |
| `PostProcessor` | `pkg/rendering/rendering.go` | Skeleton | Apply in Draw() after sprite rendering |
| `FighterGenerator` | `pkg/procgen/fighter.go` | Returns static data | Use `math/rand` with seed, vary by genre |
| `ArenaGenerator` | `pkg/procgen/arena.go` | Returns static data | Implement Poisson disk sampling, genre-specific layouts |
| `MovesetGenerator` | `pkg/procgen/moveset.go` | Returns empty data | Implement BFS combo tree, frame data allocation |
| `HitboxGenerator` | `pkg/procgen/hitbox.go` | Returns empty data | Generate hitbox shapes per move |
| `ItemGenerator` | `pkg/procgen/item.go` | Returns empty data | Implement spawn tables, stat modifiers |
| `MusicGenerator` | `pkg/procgen/music.go` | Returns empty data | Generate motifs, wire to AudioSystem |
| `SFXGenerator` | `pkg/procgen/sfx.go` | Returns empty data | Generate hit/KO samples |
| `AnnouncerGenerator` | `pkg/procgen/announcer.go` | Returns empty data | Phoneme synthesis for fight calls |
| `ArenaVisualGenerator` | `pkg/procgen/visual.go` | Returns empty data | Tile palettes, background layers |
| `FighterVisualGenerator` | `pkg/procgen/visual.go` | Returns empty data | Sprite layers, costume palettes |
| `NetworkSyncSystem` | `pkg/network/network.go` | Skeleton | Server-side state broadcast, client reconciliation |
| `TournamentGenerator` | `pkg/tournament/tournament.go` | Ignores seed | Use seed for deterministic bracket ordering |

### Mandatory Checks Before Adding or Modifying Any Feature

**Before writing ANY new code, verify the full integration chain:**

1. **Definition → Instantiation**: Is the struct/system created at runtime? Trace from `main()` through system registration.
2. **Instantiation → Registration**: Is the system registered with the Scheduler via `scheduler.Register()`? Check `cmd/client/main.go`.
3. **Registration → Update Loop**: Does the system's `Update()` method get called each frame via `scheduler.Update()`?
4. **Update → Output**: Does the system produce outputs (components, events, state changes) that other systems consume?
5. **Output → Consumer**: Is there at least one other system that reads this system's output?
6. **Consumer → Player Effect**: Does the chain ultimately produce something visible, audible, or mechanically felt by the player?

If ANY link in this chain is missing, the feature is dangling. **Do not submit dangling features.**

### Specific Anti-Patterns to Reject

```go
// ❌ BAD: System defined but never added to the scheduler
type WeatherSystem struct { ... }
func (w *WeatherSystem) Update(world *engine.World, dt float64) { ... }
// ...but never called: scheduler.Register(&WeatherSystem{})

// ✅ GOOD: System defined, instantiated, registered, and consuming/producing
weather := NewWeatherSystem(seed)
scheduler.Register(weather)
// AND other systems react to weather state via EventBus or World queries
```

```go
// ❌ BAD: Generator implements interface but returns static/empty data
func (g *FighterGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
    return &FighterData{Name: "Fighter", Speed: 1.0, ...}, nil  // Ignores seed!
}

// ✅ GOOD: Generator uses seed for deterministic variation
func (g *FighterGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
    rng := rand.New(rand.NewSource(seed))
    baseSpeed := genreSpeedBase[params.GenreID]
    return &FighterData{
        Name:  generateName(rng, params.GenreID),
        Speed: baseSpeed + rng.Float64()*0.4 - 0.2,
        // ...
    }, nil
}
```

```go
// ❌ BAD: Event emitted but no handler exists
eventBus.Publish(engine.Event{Type: "fighter.hit", Data: hitData})
// No eventBus.Subscribe("fighter.hit", ...) anywhere in the codebase

// ✅ GOOD: Event has both emitter and handler
// In CombatSystem:
eventBus.Publish(engine.Event{Type: "fighter.hit", Data: hitData})
// In AudioSystem init:
eventBus.Subscribe("fighter.hit", func(e engine.Event) {
    audio.PlayHitSound(e.Data.(*HitData).Damage)
})
```

```go
// ❌ BAD: World created but never passed to systems that need it
world := engine.NewWorld()
_ = world  // Unused!

// ✅ GOOD: World flows through the entire system
world := engine.NewWorld()
scheduler.Update(world, dt)  // Systems receive world in Update()
```

### Integration Verification Checklist (run before every PR)

```bash
# Every constructor should have at least one non-test caller
grep -rn 'func New' --include='*.go' pkg/ | grep -v _test.go

# All TODOs should be tracked
grep -rn 'TODO\|FIXME\|HACK\|XXX\|Skeleton' --include='*.go' pkg/

# Check that generators use the seed parameter
grep -rn '_ = seed' --include='*.go' pkg/

# Verify no empty method bodies
grep -A2 'func.*Update.*{$' --include='*.go' pkg/ | grep -B1 '^}'
```

---

## Networking Best Practices (MANDATORY for all Go network code)

### Interface-Only Network Types (Hard Constraint)

When declaring network variables, ALWAYS use interface types. This is a **non-negotiable project rule**.

| ❌ Never Use (Concrete Type) | ✅ Always Use (Interface Type) |
|------------------------------|-------------------------------|
| `*net.UDPAddr` | `net.Addr` |
| `*net.IPAddr` | `net.Addr` |
| `*net.TCPAddr` | `net.Addr` |
| `*net.UDPConn` | `net.PacketConn` |
| `*net.TCPConn` | `net.Conn` |
| `*net.TCPListener` | `net.Listener` |

```go
// ✅ GOOD: Interface types in pkg/network/network.go
type Server struct {
    Address  string
    listener net.Listener  // Interface type ✓
}

type Client struct {
    ServerAddress string
    conn          net.Conn  // Interface type ✓
}

// ❌ BAD: Concrete types
type Server struct {
    listener *net.TCPListener  // Concrete type ✗
}
```

**Never use type assertions to access concrete methods:**

```go
// ❌ BAD: Type assertion to concrete type
if tcpConn, ok := conn.(*net.TCPConn); ok {
    tcpConn.SetKeepAlive(true)
}

// ✅ GOOD: Use interface methods or wrap functionality
conn.SetDeadline(time.Now().Add(timeout))
```

### High-Latency Network Design (200–5000ms)

Whack's multiplayer networking MUST function correctly under **200–5000ms round-trip latency**. This supports Tor/onion routing, satellite internet, and intercontinental connections.

#### Mandatory Design Principles

1. **Client-Side Prediction**: The client simulates game state locally and reconciles with server authoritative state when it arrives. Never block the game loop waiting for a server response.

2. **State Interpolation / Extrapolation**: Remote entity positions must be interpolated between known states. When packets are delayed beyond the interpolation window, extrapolate using last-known velocity.

3. **Jitter Buffers**: Incoming state updates must be buffered and played back at a consistent rate. Design for ±500ms jitter tolerance minimum.

4. **Idempotent Messages**: Every network message must be safe to process multiple times. Use frame numbers for deduplication.

5. **No Synchronous RPC in Game Loops**: Never issue a blocking network call inside `Update()` or `Draw()`. All network I/O must be asynchronous.

6. **Graceful Degradation**: At 5000ms latency the game must remain playable, not just connected. Increase prediction windows and hide latency with animations.

7. **Timeout Tolerance**: Connection timeouts must be ≥10 seconds. Disconnect detection must use heartbeat absence over a sliding window (≥3 missed heartbeats).

```go
// ❌ BAD: Tight timeout that drops players on high-latency connections
conn.SetReadDeadline(time.Now().Add(1 * time.Second))

// ✅ GOOD: Generous timeout for high-latency environments
conn.SetReadDeadline(time.Now().Add(10 * time.Second))

// ❌ BAD: Blocking receive in game loop
func (g *Game) Update() error {
    state := g.receiveState()  // Blocks until data arrives
    g.world = state
    return nil
}

// ✅ GOOD: Async receive with interpolation
func (g *Game) Update() error {
    select {
    case state := <-g.stateChannel:
        g.interpolator.PushServerState(state)
    default:
        // No new state — continue with prediction
    }
    g.world = g.interpolator.GetInterpolatedState(time.Now())
    return nil
}
```

#### Input Frame Format

The `InputFrame` struct in `pkg/network/network.go` uses an efficient 8-byte format:

```go
type InputFrame struct {
    Frame   uint32  // Frame index (timestamp derived from frame at 60 Hz)
    Buttons uint16  // Button bitmask
    StickX  int8    // Analog stick X (-128 to 127)
    StickY  int8    // Analog stick Y (-128 to 127)
}
```

#### Rollback Requirements

Per ROADMAP.md Phase 4:
- Client predicts up to 300 frames (5 seconds at 60 Hz)
- Server rewinds up to 5000ms for hit registration validation
- State snapshot ring-buffer: 300 frames × ~4 KB ≈ 1.2 MB per client
- At >2000ms latency, switch to interpolation-only mode

---

## Code Assistance Guidelines

### 1. Deterministic Procedural Generation

All content generation MUST be deterministic and seed-based. Given the same seed and genre, the game MUST produce identical output across all platforms and runs.

```go
// ✅ GOOD: Explicit seed-based RNG, never global
func (g *FighterGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
    rng := rand.New(rand.NewSource(seed))
    speed := baseSpeed + rng.Float64()*variance
    // ...
}

// ❌ BAD: Global rand (non-deterministic, not thread-safe)
func (g *FighterGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
    speed := baseSpeed + rand.Float64()*variance  // Uses global rand!
}

// ❌ BAD: Time-based seeding in generation code
rng := rand.New(rand.NewSource(time.Now().UnixNano()))

// ✅ GOOD: Derived seeds for sub-generators (deterministic hierarchy)
func (g *FighterGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
    statsSeed := seed ^ 0x53544154    // "STAT"
    visualSeed := seed ^ 0x56495355  // "VISU"
    movesSeed := seed ^ 0x4D4F5645   // "MOVE"
    
    statsRNG := rand.New(rand.NewSource(statsSeed))
    visualRNG := rand.New(rand.NewSource(visualSeed))
    movesRNG := rand.New(rand.NewSource(movesSeed))
    // ...
}
```

### 2. ECS Architecture Discipline

Whack uses a sparse-set backed ECS with generics:

- **Entities** are `uint64` IDs (type alias `engine.Entity`)
- **Components** are pure data structs stored in typed `ComponentStore[T]`
- **Systems** implement `System` interface with `Update(w *World, dt float64)`
- **World** owns all component stores and entity allocation

```go
// Component: Pure data, no logic
type PositionComponent struct {
    X, Y float64
}

// System: All logic, operates on World
type PhysicsSystem struct {
    Gravity float64
}

func (s *PhysicsSystem) Update(w *engine.World, dt float64) {
    positions := w.Positions.All()
    knockbacks := w.Knockbacks.All()
    entities := w.Positions.Entities()
    
    for i, pos := range positions {
        if kb, ok := w.Knockbacks.Get(entities[i]); ok {
            pos.X += kb.VX * dt
            pos.Y += kb.VY * dt
            pos.Y += s.Gravity * dt
            w.Positions.Set(entities[i], pos)
        }
    }
}
```

**Key patterns:**
- Never store entity references directly; use `engine.Entity` IDs
- Systems declare dependencies implicitly through which `ComponentStore`s they access
- Use `World.NewEntity()` for entity creation
- Use `ComponentStore.Set()` / `Get()` / `All()` / `Entities()` for component access

### 3. Event Bus Usage

The `EventBus` in `pkg/engine/event.go` provides publish/subscribe:

```go
// Subscribe to events
eventBus.Subscribe("fighter.hit", func(e engine.Event) {
    data := e.Data.(*HitEventData)
    // Handle hit
})

// Publish events
eventBus.Publish(engine.Event{
    Type: "fighter.hit",
    Data: &HitEventData{Attacker: attackerID, Target: targetID, Damage: damage},
})
```

**Standard event types (establish consistency):**
- `fighter.hit` — Hitbox connected with hurtbox
- `fighter.ko` — Fighter crossed blast zone
- `match.start` — Match begins
- `match.end` — Match ends (winner determined)
- `stock.lost` — Player lost a stock

### 4. Genre System

The five genres in `pkg/procgen/generator.go` must affect ALL procedural systems:

```go
const (
    GenreFantasy       GenreID = "fantasy"
    GenreSciFi         GenreID = "sci-fi"
    GenreHorror        GenreID = "horror"
    GenreCyberpunk     GenreID = "cyberpunk"
    GenrePostApocalyptic GenreID = "post-apocalyptic"
)
```

Per ROADMAP.md, genre differences include:
- **Stats**: Fantasy has high magic/low speed; Sci-fi has +30% speed; Horror has high damage; Post-apoc is tanky
- **Arenas**: Fantasy floating islands; Sci-fi moving platforms/lasers; Horror falling tiles/fog; Cyberpunk neon/electric; Post-apoc crumbling/fire
- **Moves**: Fantasy spells/summons; Sci-fi beams/teleport; Horror grabs/lifesteal; Cyberpunk dash-cancel; Post-apoc explosives
- **Audio**: Fantasy 90 BPM orchestral; Sci-fi 130 BPM synth; Horror 70 BPM dissonant; Cyberpunk 150 BPM bass; Post-apoc 110 BPM industrial
- **Visuals**: Fantasy warm gold; Sci-fi cool blue/scanlines; Horror desaturated/vignette; Cyberpunk neon/chromatic; Post-apoc muted orange

### 5. Performance Requirements

- Target 60 FPS on mid-range hardware (Intel i5, 8 GB RAM)
- Client memory budget: <500MB
- Sprite generation: <50ms per fighter
- Hitbox resolution: <0.5ms for 4-player match
- Post-processing: <3ms per frame
- Audio latency: <20ms

Use benchmarks for hot paths:
```go
func BenchmarkCombatSystem(b *testing.B) {
    world := setupTestWorld(100) // 100 entities
    sys := &CombatSystem{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        sys.Update(world, 1.0/60.0)
    }
}
```

### 6. Zero External Assets

ALL content is generated at runtime:
- **Graphics**: SDF composition, noise functions, pixel manipulation
- **Audio**: Oscillators (Sine, Square, Sawtooth, Triangle, Noise) + ADSR envelopes
- **Levels**: Poisson disk sampling, noise thresholds, parameterized platforms
- **Music**: Motif sequences with genre-specific BPM and waveform mix
- **SFX**: Noise bursts with genre pitch/timbre shifts

**Never add asset files** (PNG, WAV, OGG, JSON, TTF, GLB) to the repository.

Validate with:
```bash
find . -path ./.git -prune -o -type f \( -name "*.png" -o -name "*.wav" -o -name "*.ogg" -o -name "*.mp3" -o -name "*.json" -o -name "*.ttf" -o -name "*.glb" \) -print
# Should return empty

git grep -n "//go:embed"
# Should return no matches
```

### 7. Error Handling

```go
// ✅ GOOD: Return errors with context
func (g *ArenaGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
    if err := params.Validate(); err != nil {
        return nil, fmt.Errorf("arena generation: %w", err)
    }
    // ...
}

// ❌ BAD: Panic in library code
func (g *ArenaGenerator) Generate(seed int64, params GenerationParams) *ArenaData {
    if params.GenreID == "" {
        panic("empty genre")  // Never panic in game logic
    }
}

// ✅ GOOD: Graceful fallback in systems
func (s *RenderSystem) Update(w *engine.World, dt float64) {
    sprite, err := s.generator.Generate(s.seed, s.params)
    if err != nil {
        log.Printf("sprite generation failed: %v, using fallback", err)
        sprite = s.fallbackSprite
    }
}
```

Panics are acceptable ONLY in `main()` for unrecoverable startup failures.

---

## Cross-Repository Code Sharing Patterns

### Shared Pattern Catalog

When implementing features, follow these patterns so code can be extracted into shared packages later:

| Pattern | Package | Used By |
|---------|---------|---------|
| ECS core (World, Entity, ComponentStore, System) | `pkg/engine/` | All repos |
| PCG interface and GenerationParams | `pkg/procgen/` | All repos |
| Seed derivation utilities | Inline or `pkg/seed/` | All repos |
| Sprite/tile generation | `pkg/rendering/` | All repos |
| Audio synthesis (Oscillator, Envelope) | `pkg/audio/` | All repos |
| Input handling | `pkg/input/` | All repos |
| Configuration (Viper) | `pkg/config/` | violence, velocity, whack |
| Networking | `pkg/network/` | venture, violence, whack |
| Physics | `pkg/physics/` | vania, venture, whack, way |

### Guidelines for Shareable Code

1. **Keep dependencies minimal**: Shared packages should depend only on stdlib + Ebiten
2. **Use interfaces at boundaries**: Define interfaces for game-specific behavior
3. **Parameterize, don't specialize**: Generators accept parameters for any genre
4. **Same naming conventions**: If `venture` uses `pkg/engine/World`, whack uses `pkg/engine/World` with same signatures
5. **Identical System interface**: `Update(w *World, dt float64)` is universal

### When Adding a Feature That Exists in a Sibling Repo

1. Check the sibling repo's implementation first
2. Use the same package structure and naming conventions
3. Match the interface signatures for future extraction
4. Document divergences in ROADMAP.md

---

## Configuration

Whack uses Viper for YAML configuration. See `pkg/config/config.go`:

```yaml
# config.yaml
game:
  title: "Whack"
  width: 800
  height: 600
  tps: 60
  seed: 0           # 0 = random seed at startup
  genre: "fantasy"  # fantasy, sci-fi, horror, cyberpunk, post-apocalyptic

server:
  address: ":7000"
  max_players: 4
  tick_rate: 60

client:
  server_address: "localhost:7000"

match:
  stocks: 3
  time_limit: 300   # seconds
  mode: "ffa"       # ffa, 1v1, 2v2
```

Access via `config.Load()` which returns `*Config` with defaults applied.

---

## Quality Standards

### Testing Requirements

- **Coverage**: ≥40% per package (≥30% for Ebiten-dependent packages)
- **Table-driven tests** for all generators and systems
- **Benchmarks** for hot-path code (combat resolution, physics, rendering)
- **Race detection**: All tests must pass under `go test -race ./...`
- **Determinism tests**: Same seed must produce identical output

```go
func TestFighterGeneratorDeterminism(t *testing.T) {
    gen := &FighterGenerator{}
    params := GenerationParams{GenreID: GenreFantasy}
    
    result1, _ := gen.Generate(12345, params)
    result2, _ := gen.Generate(12345, params)
    
    if !reflect.DeepEqual(result1, result2) {
        t.Error("same seed produced different results")
    }
}
```

### Build Verification

```bash
# Build all packages
go build ./...

# Run tests with race detection
go test -race ./...

# Verify no embedded assets
find . -path ./.git -prune -o -type f \( -name "*.png" -o -name "*.wav" \) -print
```

---

## Naming Conventions

- **Packages**: lowercase, single-word (`engine`, `procgen`, `audio`, `combat`)
- **Files**: snake_case (`combat_system.go`, `fighter_generator.go`)
- **Types**: PascalCase (`FighterGenerator`, `CombatSystem`, `HitboxComponent`)
- **Interfaces**: PascalCase, `-er` suffix for single-method (`Generator`, `System`)
- **Component types**: PascalCase + "Component" suffix (`PositionComponent`)
- **System types**: PascalCase + "System" suffix (`PhysicsSystem`)
- **Generator types**: PascalCase + "Generator" suffix (`FighterGenerator`)
- **Seeds**: Always `int64`, parameter named `seed`
- **Genres**: Always `GenreID` type, use constants (`GenreFantasy`, not `"fantasy"`)

---

## Frame Data and Combat Timing

Whack uses frame-data-driven combat similar to traditional fighting games. All combat timing is measured in frames at 60 FPS (16.67ms per frame).

### Move Frame Data Structure

Every move has three phases defined in `pkg/procgen/moveset.go`:

```go
type Move struct {
    Name     string
    Startup  int     // Frames before hitbox becomes active
    Active   int     // Frames hitbox is active (can hit)
    Recovery int     // Frames after active before next action
    Damage   float64 // Base damage on hit
}
```

**Frame advantage** = (Defender's hitstun) - (Attacker's recovery)
- Positive = Attacker can act first
- Negative = Defender can act first
- Zero = Both recover simultaneously

### Acceptance Criteria from ROADMAP.md

- Every generated moveset has ≥8 normals, ≥4 specials, ≥1 super
- No move has 0 active frames
- Hitbox resolution runs in <0.5ms for 4-player match
- Knockback at 150% damage launches beyond blast zone
- DI deflects trajectory by ≥15°

### Knockback Formula (SSB-style)

```
knockback = base_knockback + (damage * damage_scaling * weight_factor)
direction = hit_angle + DI_adjustment
```

Where:
- `base_knockback` comes from the move's frame data
- `damage_scaling` increases with accumulated percent (starts at 0%, caps at 999%)
- `weight_factor` = 1.0 / (fighter_weight) — lighter fighters fly further
- `DI_adjustment` = up to ±15° based on directional input during hitstun

---

## Blast Zones and KO Detection

The `BlastZone` struct in `pkg/arena/arena.go` defines four boundaries:

```go
type BlastZone struct {
    Left, Right, Top, Bottom float64
}
```

A fighter is KO'd when their center position crosses any boundary. The default blast zone extends 100 units past the visible screen on all sides.

**Stock system**: Each player has configurable stocks (default 3). Losing all stocks eliminates the player. Match ends when one player remains.

**Sudden death**: If timer expires with tied stocks, remaining players respawn at 300% damage for a final-stock sudden death.

---

## Audio Synthesis Pipeline

All audio in Whack is synthesized at runtime using the components in `pkg/audio/audio.go`.

### Waveform Types

```go
const (
    Sine     WaveformType = iota  // Pure tone, soft
    Square                        // Harsh, 8-bit style
    Sawtooth                      // Bright, buzzy
    Triangle                      // Softer than square
    Noise                         // White noise for percussion/SFX
)
```

### ADSR Envelope

Every sound uses an amplitude envelope:

```go
type Envelope struct {
    Attack  float64  // Time to reach peak amplitude (seconds)
    Decay   float64  // Time to drop from peak to sustain (seconds)
    Sustain float64  // Amplitude level during sustain phase (0.0-1.0)
    Release float64  // Time to fade to silence after note-off (seconds)
}
```

### Genre Audio Profiles

| Genre | Base BPM | Primary Waveforms | Character |
|-------|----------|-------------------|-----------|
| Fantasy | 90 | Sine, Triangle | Orchestral, strings/pipes |
| Sci-Fi | 130 | Square, Sawtooth | Synthwave, lead synths |
| Horror | 70 | Noise, Sawtooth | Dissonant pads, drones |
| Cyberpunk | 150 | Square, Sawtooth | DnB, heavy bass/arps |
| Post-Apocalyptic | 110 | Sawtooth, Noise | Industrial, distorted guitar |

### SFX Generation Rules

- **Hit sounds**: Noise burst + sine undertone, genre pitch shift
- **KO sounds**: Impact + reverb tail, genre-specific tonal signature
- **Announcer**: Phoneme synthesis with genre timbre (vocoder for cyberpunk, whisper for horror)

---

## Platform and Hazard System

Arenas contain platforms and optional hazards that affect gameplay.

### Platform Types

```go
type Platform struct {
    X, Y, W, H float64
    Moving     bool     // Platform moves on a path
}
```

Platforms support:
- **Solid collision**: Fighters stand on top, fall through from below
- **Moving paths**: Oscillate horizontally or vertically
- **Drop-through**: Player can press down to fall through

### Hazard System

```go
type Hazard struct {
    X, Y   float64
    Type   string   // Genre-specific hazard type
    Active bool     // Currently dangerous
}
```

Hazard types by genre (from ROADMAP.md):
- **Fantasy**: Dragon fly-bys, cursed floor tiles, spike pits
- **Sci-Fi**: Laser sweeps, gravity flips, teleporter pads
- **Horror**: Falling ceiling, shadow enemy spawns, fog
- **Cyberpunk**: Power surge, auto-turret spawns, electric rails
- **Post-Apocalyptic**: Acid rain, ground collapse, fire pits

Hazards activate on timers or proximity triggers. They deal damage and knockback equivalent to a medium attack.

---

## Combo Tree Structure

Fighters have procedurally generated combo trees in `pkg/fighter/fighter.go`:

```go
type ComboTree struct {
    Root *ComboNode
}

type ComboNode struct {
    MoveName string
    Children []*ComboNode  // Possible follow-ups
}
```

The `MovesetGenerator` builds combo trees using BFS:
1. Start with normal attacks as roots
2. Branch to special moves based on genre and stats
3. Cap chains at 8 moves to prevent infinites
4. Ensure frame advantage allows links (positive on-hit)

### Combo Notation Convention

Use fighting game notation for describing combos:
- `jab > jab > tilt > special` — Chain combo with links
- Numbers 1-9 for directional inputs (numpad notation)
- Example: `5A > 5A > 2B > 236C` = standing light, standing light, crouching medium, quarter-circle-forward heavy

---

## Spectator Mode

The tournament package provides spectator support:

```go
type SpectatorState struct {
    Frame    uint32
    Entities []SpectatorEntity
}

type SpectatorEntity struct {
    ID   uint64
    X, Y float64
}
```

Spectators:
- Receive read-only state stream at 20 Hz
- Cannot send input packets
- Interpolate entity positions between snapshots
- Connect mid-match and sync within 3 seconds

---

## ROADMAP.md Protocol

The ROADMAP.md file contains the implementation plan with phases, acceptance criteria, and success metrics. When contributing:

1. Check which phase the feature belongs to
2. Verify acceptance criteria (AC) before marking complete
3. Reference ROADMAP.md section numbers in commit messages
4. Update success criteria if implementation reveals new constraints

### Current Implementation Status

| Phase | Name | Status |
|-------|------|--------|
| 1 | Foundation & ECS Core | Partially complete |
| 2 | PCG Core: Fighters, Arenas, Movesets | Not started |
| 3 | Combat Engine & Physics | Not started |
| 4 | Multiplayer & Netcode | Not started |
| 5 | Audio, Rendering & Post-Processing | Not started |
| 6 | Polish, Balance & Release | Not started |

### Gap Identification Protocol

When Copilot identifies a potential gap:

1. **Note it immediately** in your response with severity assessment
2. **Suggest tracking** in ROADMAP.md Phase notes or new GAPS.md file
3. **Include location**: file path, line number, function name
4. **Propose fix**: Concrete implementation steps, not vague suggestions

Severity levels:
- **Critical**: Blocks core gameplay (e.g., combat system never runs)
- **High**: Degrades experience significantly (e.g., no audio on hits)
- **Medium**: Missing polish (e.g., no screen shake on KO)
- **Low**: Enhancement (e.g., additional post-processing effect)
