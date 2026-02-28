# Whack — Implementation Roadmap

## 1. Project Overview

Whack is a 100% procedurally generated arena combat game built in Go 1.24+ with the Ebiten v2 game engine. Inspired by Super Smash Brothers, Street Fighter, and Mortal Kombat, every element — fighters, arenas, movesets, music, SFX, and visual effects — is generated at runtime from a deterministic seed with zero embedded assets. As the W-Series successor to the V-Series (`venture`, `vania`, `violence`, `velocity`), Whack exceeds its predecessors through frame-data-driven combat, hitbox/hurtbox generation, procedural combo trees, ring-out/KO physics, rollback-style netcode adapted for high-latency (Tor/onion) play, and five-genre theming that changes not just aesthetics but gameplay mechanics — giving each of fantasy, sci-fi, horror, cyberpunk, and post-apocalyptic a distinctly different fighting experience.

---

## 2. Core Architecture

**ECS:** Entities are uint64 IDs, but component storage is densely indexed: each component type maintains a packed slice plus an entity→dense-index map (sparse-set style). Components (`HitboxComponent`, `FighterStatsComponent`, `PositionComponent`, etc.) are pure data structs in these slices, and systems (`PhysicsSystem`, `CombatSystem`, `NetworkSyncSystem`) contain all logic and iterate over the dense component arrays.

**PCG Interface (inherited from V-Series):**
```go
type Generator interface {
    Generate(seed int64, params GenerationParams) (interface{}, error)
    Validate() error
}
// GenerationParams.GenreID ∈ {fantasy, sci-fi, horror, cyberpunk, post-apocalyptic}
```

**V-Series reuse:** `pkg/procgen` (noise, name gen, stat tables), `pkg/audio` (oscillators, envelopes, motif system), `pkg/rendering` (sprite pipeline, palette, post-processing), `pkg/network` (TCP server, prediction, delta compression), `pkg/engine` (ECS core, event bus, scheduler).

**Whack-specific additions:** `pkg/combat` (frame data, hitbox resolution, knockback), `pkg/fighter` (moveset trees, combo logic), `pkg/arena` (platform layout, hazard placement), `pkg/tournament` (bracket, spectator state).

**Binary:** Single `cmd/client` + `cmd/server` binary, no runtime asset directories.

---

## 3. Implementation Phases

### Phase 1 — Foundation & ECS Core (4 weeks)
*Establish project scaffold, ECS, PCG interfaces, genre system.*

1. Initialize Go module; scaffold `cmd/client`, `cmd/server`, `pkg/` layout mirroring V-Series.
   - **AC:** `go build ./...` succeeds; no embedded assets in repo.
2. Port/adapt ECS core from V-Series `pkg/engine`; add `HitboxComponent`, `HurtboxComponent`, `FighterStatsComponent`, `KnockbackComponent`.
   - **AC:** ECS handles 100 entities at 60 Hz with <1 ms update time in benchmark.
3. Implement `GenerationParams` with `GenreID` validation for all five genres; wire into generator interface.
   - **AC:** All five `GenreID` values pass `Validate()`; invalid genre returns typed error.
4. Stub `FighterGenerator`, `ArenaGenerator`, `MovesetGenerator` implementing the PCG interface.
   - **AC:** Each stub returns deterministic output for the same seed across 1000 invocations.

### Phase 2 — PCG Core: Fighters, Arenas, Movesets (6 weeks)
*Implement procedural content generators with full genre differentiation.*

5. `FighterGenerator`: body proportions, stat distribution (speed/weight/jump/reach), visual palette per genre.
   - **AC:** Two fighters with same seed+genre are byte-identical; different genres produce measurably different stat distributions (p < 0.01).
6. `ArenaGenerator`: platform layout via constrained Poisson sampling, hazard placement, blast-zone sizing, genre-specific aesthetics.
   - **AC:** No two seeds produce identical platform configurations; all genres produce collision-valid arenas.
7. `MovesetGenerator`: combo tree (BFS from normal → special → super), frame data (startup/active/recovery frames), hitbox shapes per move, genre-specific special move types.
   - **AC:** Every generated moveset has ≥8 normals, ≥4 specials, ≥1 super; no move has 0 active frames.
8. `ItemGenerator`: weapon/pickup types, stat modifiers, visual flavor, spawn weight table per genre.
   - **AC:** Item pool differs visually and statistically across genres with same seed.

### Phase 3 — Combat Engine & Physics (5 weeks)
*Frame-data-driven combat, hitbox resolution, knockback, KO conditions.*

9. `CombatSystem`: hitbox vs. hurtbox intersection per frame, damage accumulation, hitstun calculation from frame data.
   - **AC:** Hitbox resolution runs in <0.5 ms for 4-player match; no false positives in 10k random frame tests.
10. `KnockbackSystem`: vector-based knockback with damage-scaling (SSB-style percentage), DI (directional influence) input, tumble state transitions.
    - **AC:** Knockback at 150% damage launches fighter beyond blast zone; DI deflects trajectory by ≥15°.
11. `PhysicsSystem`: platform collision, fast-fall, wall/ceiling interaction, ledge-grab detection.
    - **AC:** Fighter never clips through a generated platform at any tested velocity.
12. KO/ring-out detection, stock/life system, match timer, round state machine.
    - **AC:** Match correctly ends after last stock lost; sudden-death triggers at timer=0 with tied stocks.

### Phase 4 — Multiplayer & Netcode (5 weeks)
*Authoritative server, input sync, prediction, lag compensation for Tor-range latency.*

13. Authoritative TCP server: input packet framing, server-side combat simulation, delta-compressed state broadcast.
    - **AC:** Server simulates 4-player match at 60 Hz; state divergence between two clients <1 frame after reconciliation.
14. Client-side prediction: optimistic input application, server reconciliation rollback (up to 300 frames buffer).
    - **AC:** At 200 ms RTT, visual stutter <2 frames; at 2000 ms RTT, game remains playable with correct final state.
15. Entity interpolation for remote fighters; lag compensation for hitbox validation (server rewinds up to 5000 ms).
    - **AC:** Hit registration accuracy ≥95% at 1000 ms simulated latency in automated test harness.
16. Spectator mode: read-only state stream, no input authority; tournament bracket generator (single/double elimination).
    - **AC:** Spectator client connects mid-match and renders correct state within 3 seconds; bracket validates for 4–64 players.

### Phase 5 — Audio, Rendering & Post-Processing (4 weeks)
*Procedural audio pipeline, runtime sprite generation, genre visual palettes.*

17. `FighterRenderer`: runtime sprite generation via signed-distance-field composition; genre-driven palette (e.g., horror desaturated greens, cyberpunk neon).
    - **AC:** Sprite generation <50 ms per fighter; all five genres produce visually distinct fighter silhouettes.
18. `ArenaRenderer`: tile/platform rendering from noise maps; parallax background layers; genre hazard VFX particles.
    - **AC:** Arena renders at ≥60 FPS on reference hardware with 4 fighters active.
19. `AudioSystem`: port V-Series oscillator/envelope engine; add fight-intensity scaling (BPM increases with stock loss), procedural announcer (text-to-waveform), hit/impact SFX variety via genre pitch/timbre rules.
    - **AC:** Audio latency <20 ms; announcer produces distinct phonemes for "KO", "FIGHT", "WINNER"; SFX vary across genres.
20. Post-processing passes: screen shake on KO, genre-specific filters (scanlines for sci-fi, blood-tint for horror, glow for cyberpunk).
    - **AC:** Post-processing adds <3 ms per frame on reference hardware.

### Phase 6 — Polish, Balance & Release (3 weeks)
*Balance pass, mode completeness, single binary packaging.*

21. Automated balance validation: run 10k simulated matches per fighter/arena combination; flag stat outliers >2σ from mean win rate.
    - **AC:** No generated fighter wins >65% of matches vs. random opponents across 10k games.
22. 1v1, FFA (up to 4 players), team battle (2v2) mode support; mode selection in server config.
    - **AC:** All three modes complete without desync across 100 test sessions.
23. Single binary build: `go build -tags netgo -ldflags "-extldflags -static"` for client and server; verify no runtime asset loading.
    - **AC:** Binary runs on clean system with no game assets present; `ldd` shows statically linked.

---

## 4. PCG Systems Inventory

| Generator | Algorithm/Approach | Genre Effect |
|---|---|---|
| `FighterGenerator` | Seeded stat sampling from weighted distributions; body-shape via SDF composition | Fantasy: high magic/low speed; Sci-fi: +30% speed, energy visuals; Horror: high damage, grotesque shapes; Cyberpunk: balanced + neon palette; Post-apoc: heavy/tanky, weathered look |
| `ArenaGenerator` | Constrained Poisson disk for platform centres; hazard placement via noise threshold | Fantasy: floating islands, spike pits; Sci-fi: moving platforms, laser hazards; Horror: falling floor tiles, fog; Cyberpunk: vertical neon towers, electric rails; Post-apoc: crumbling platforms, fire pits |
| `MovesetGenerator` | BFS combo tree with frame-budget allocation; special-move type selected by `GenreID` | Fantasy: projectile spells, summons; Sci-fi: energy beams, teleport; Horror: blood/tendrils, grab-heavy; Cyberpunk: dash cancel chains, EMP bursts; Post-apoc: weapon swings, explosive throws |
| `HitboxGenerator` | Capsule/rect fitting from limb length distribution; active frames from attack speed stat | All genres share shape logic; size/reach scales with genre stat bias |
| `ItemGenerator` | Spawn table with per-genre weight biases; SDF item sprites; stat mod ranges by item tier | Fantasy: potions, scrolls; Sci-fi: ammo packs, shields; Horror: cursed relics; Cyberpunk: mods/chips; Post-apoc: scrap weapons |
| `ArenaVisualGenerator` | Layered noise-based tile palette; parallax background from genre seed | Each genre: distinct colour palette, background narrative elements, hazard visual style |
| `FighterVisualGenerator` | SDF layering for costume/body; palette swap driven by `GenreID` | Fantasy: robes/armour; Sci-fi: suits/visors; Horror: gore/flesh; Cyberpunk: implants/neon; Post-apoc: rags/rust |
| `MusicGenerator` | Seeded motif (4–8 notes) + genre-BPM base (fantasy 90, sci-fi 130, horror 70, cyberpunk 150, post-apoc 110); intensity scales BPM ±30% | Each genre uses distinct waveform mix (fantasy: strings/pipes; sci-fi: synth; horror: dissonant pads; cyberpunk: bass/arps; post-apoc: distorted guitar) |
| `SFXGenerator` | Hit sounds: noise burst + genre pitch shift; KO: genre-specific impact sample via additive synthesis | Sci-fi: +30% pitch; Horror: -30% pitch + vibrato; Cyberpunk: +40% pitch + hard clipping; Fantasy: reverb; Post-apoc: low rumble |
| `AnnouncerGenerator` | Phoneme waveform synthesis from text; modulated by genre timbre profile | Fantasy: booming herald; Sci-fi: synthetic; Horror: whispering; Cyberpunk: vocoder; Post-apoc: raspy/distorted |

---

## 5. Multiplayer Design

**Input Synchronization:** Clients send timestamped input frames (8-byte: frame#, buttons, stick X/Y) at 60 Hz. Server buffers 1–3 frames and simulates authoritatively. Input delay is adaptive: min 2 frames at <50 ms, scales to 8 frames at >500 ms.

**Rollback Netcode (high-latency adaptation):** Client predicts up to 300 frames; on server reconciliation mismatch, replays from divergence point. State snapshot ring-buffer (300 frames × ~4 KB = ~1.2 MB per client). At >2000 ms latency (Tor), switches to interpolation-only mode with deferred input confirmation.

**Hitbox Validation:** Server rewinds entity state up to 5000 ms (300 frames at 60 Hz) to validate hit registration from lagging clients. Prevents phantom hits while allowing late-arriving inputs.

**Knockback Reconciliation:** Knockback vectors are server-authoritative. Client shows predicted trajectory; corrects smoothly over 3 frames on mismatch to avoid visual pop.

**Spectator State:** Dedicated read-only stream at 20 Hz snapshot rate with entity interpolation. Spectators receive full ECS state diff; no input packets sent.

**Tournament Brackets:** `TournamentGenerator` produces single/double-elimination brackets for 4–64 players deterministically from seed. Bracket state replicated to all participants; bye rounds auto-resolved.

---

## 6. Genre Differentiation Matrix

| Dimension | Fantasy | Sci-Fi | Horror | Cyberpunk | Post-Apocalyptic |
|---|---|---|---|---|---|
| Arena Theme | Floating castles, lava pits, enchanted forests | Space stations, laser grids, moving platforms | Haunted mansions, spike floors, fog | Neon city rooftops, electric rails, holo-barriers | Desert ruins, fire pits, crumbling floors |
| Fighter Style | Mages/knights, high magic damage, slow | Troopers/androids, fast projectiles, teleport | Monsters/cultists, grabs, lifesteal, slow | Hackers/cyborgs, dash-cancel combos, EMP | Scavengers/mutants, weapon-heavy, tanky |
| Special Moves | Spell projectiles, area summons, shields | Energy beams, teleport blink, drone assists | Tentacle grabs, blood mist, fear-stun | Dash cancel, EMP burst, system hack debuff | Explosive throw, chain whip, radiation AOE |
| Hazards | Dragon fly-bys, cursed floor tiles | Laser sweeps, gravity flips, teleporter pads | Falling ceiling, shadow enemies spawn | Power surge, auto-turret spawns | Acid rain, ground collapse |
| KO Visual | Golden explosion, angelic choir sting | Energy disintegration, sci-fi klaxon | Gore dissolve, horror sting | Pixel glitch, vocoder KO call | Dust collapse, distorted crunch |
| Music Style | Orchestral fantasy (90 BPM, strings) | Synthwave (130 BPM, lead synth) | Dissonant horror (70 BPM, pads) | Drum & bass (150 BPM, bass/arps) | Industrial rock (110 BPM, distorted guitar) |
| Visual Palette | Warm gold/green, soft lighting | Cool blue/white, bloom + scanlines | Desaturated grey-green, vignette | Neon pink/cyan, chromatic aberration | Muted orange/brown, dust particles |

---

## 7. Success Criteria

| Indicator | Target | Measurement Method |
|---|---|---|
| Zero embedded assets | 0 image/audio/data files in repo | `find . -path ./.git -prune -o -type f \( -name "*.png" -o -name "*.wav" -o -name "*.ogg" -o -name "*.mp3" -o -name "*.json" -o -name "*.ttf" -o -name "*.glb" \)` returns empty and `git grep -n "//go:embed"` returns no matches |
| Deterministic PCG | Identical output for same seed+genre | 1000-run hash comparison in CI |
| Combat frame accuracy | Hitbox resolution error 0 frames at 60 Hz local | Automated frame-step test harness |
| Netcode at 200 ms RTT | <2 frame visual stutter | Network simulator test with tc netem |
| Netcode at 2000 ms RTT | Correct final game state, playable | Automated session test with 2 s simulated delay |
| Hit registration at 1000 ms | ≥95% accuracy | Server-rewind validation test vs. ground truth |
| Genre differentiation | All 5 genres produce statistically distinct fighter stats | ANOVA p < 0.01 across genre stat distributions |
| Arena validity | 0 platform-clip events | Physics fuzz test: 10k random positions/velocities |
| Balance | No fighter >65% win rate | 10k simulated match tournament per seed |
| Binary size | Single static binary ≤50 MB | `ls -lh` on release build |
| Render performance | ≥60 FPS, 4 players, all effects | Profiler on reference hardware (Intel i5, 8 GB RAM) |

---

## 8. Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Rollback netcode incompatible with 5000 ms Tor latency | High | High | Implement interpolation-only fallback mode; test with simulated Tor latency from day 1 of Phase 4 |
| PCG fighter balance impossible to guarantee | Medium | High | Automated balance CI gate; stat-clamping in `FighterGenerator`; re-seed on outlier detection |
| Ebiten audio latency exceeds 20 ms target | Medium | Medium | Evaluate `oto` buffer sizing; fallback to larger buffer with visual compensation |
| Hitbox rewind at 5000 ms consumes excessive memory | Medium | Medium | Cap rewind buffer at 300 frames (5000 ms × 60 Hz ÷ 1000, ~1.2 MB per client); beyond cap, server rules in favour of defender |
| Five-genre PCG scope bloat delays core combat | Medium | High | Genres share base generator logic; genre effect is a parameter overlay, not a separate code path |
| Platform clipping at high knockback velocities | Low | High | Sub-step physics integration (CCD) for velocities >platform-thickness per frame |
| Single binary size exceeds target with all PCG | Low | Low | Profile with `go tool nm`; tree-shake unused genre modules via build tags |
