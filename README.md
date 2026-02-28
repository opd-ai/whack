# Whack

Procedurally generated arena combat game built with Go and Ebitengine.

## Directory Structure

```
cmd/
  client/          Game client (Ebitengine window)
  server/          Authoritative game server
pkg/
  arena/           Platform layout and hazard placement
  audio/           Oscillators, envelopes, motif system
  combat/          Frame data, hitbox resolution, knockback
  config/          Configuration loading (Viper)
  engine/          ECS core, event bus, scheduler
  fighter/         Moveset trees, combo logic
  network/         TCP server, prediction, delta compression
  physics/         Collision, gravity, movement
  procgen/         PCG interface and all generators
  rendering/       Sprite pipeline, palette, post-processing
  tournament/      Bracket generation, spectator state
config.yaml        Default configuration file
```

## Requirements

Building the client requires platform graphics libraries:

- **Linux**: `libx11-dev libxrandr-dev libxcursor-dev libxi-dev libxinerama-dev libxxf86vm-dev libgl1-mesa-dev libasound2-dev`
- **macOS / Windows**: No additional dependencies

## Build

```sh
go build ./...
```

Build individual binaries:

```sh
go build -o whack-client ./cmd/client
go build -o whack-server ./cmd/server
```

## Run

Start the client:

```sh
./whack-client
```

Start the server:

```sh
./whack-server
```

## Configuration

Edit `config.yaml` in the working directory. Defaults are used if the file is absent.

| Section  | Key              | Default         |
|----------|------------------|-----------------|
| game     | title            | Whack           |
| game     | width            | 800             |
| game     | height           | 600             |
| game     | tps              | 60              |
| game     | seed             | 0               |
| game     | genre            | fantasy         |
| server   | address          | :7000           |
| server   | max_players      | 4               |
| server   | tick_rate        | 60              |
| client   | server_address   | localhost:7000  |
| match    | stocks           | 3               |
| match    | time_limit       | 300             |
| match    | mode             | ffa             |

Valid genres: `fantasy`, `sci-fi`, `horror`, `cyberpunk`, `post-apocalyptic`.

## Dependencies

- [Ebitengine v2](https://ebitengine.org/) — game engine
- [Viper](https://github.com/spf13/viper) — configuration management

## License

MIT
