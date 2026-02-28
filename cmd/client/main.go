// Command client is the Whack game client.
package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/opd-ai/whack/pkg/config"
	"github.com/opd-ai/whack/pkg/engine"
)

// Game implements the ebiten.Game interface.
type Game struct {
	world     *engine.World
	scheduler *engine.Scheduler
	eventBus  *engine.EventBus
	cfg       *config.Config
}

// NewGame creates a new Game instance.
func NewGame(cfg *config.Config) *Game {
	return &Game{
		world:     engine.NewWorld(),
		scheduler: engine.NewScheduler(),
		eventBus:  engine.NewEventBus(),
		cfg:       cfg,
	}
}

// Update is called every tick.
func (g *Game) Update() error {
	g.scheduler.Update(g.world, 1.0/float64(g.cfg.Game.TPS))
	return nil
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 30, A: 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%s - %s", g.cfg.Game.Title, g.cfg.Game.Genre))
}

// Layout returns the game screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.cfg.Game.Width, g.cfg.Game.Height
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	game := NewGame(cfg)

	ebiten.SetWindowSize(cfg.Game.Width, cfg.Game.Height)
	ebiten.SetWindowTitle(cfg.Game.Title)
	ebiten.SetTPS(cfg.Game.TPS)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
