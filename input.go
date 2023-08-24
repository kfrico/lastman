package main

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Input struct {
	lastBulletTime time.Time
}

func (i *Input) Update(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.ship.x -= g.cfg.ShipSpeedFactor
		if g.ship.x < -float64(g.ship.width)/2 {
			g.ship.x = -float64(g.ship.width) / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.ship.x += g.cfg.ShipSpeedFactor
		if g.ship.x > float64(g.cfg.ScreenWidth)-float64(g.ship.width)/2 {
			g.ship.x = float64(g.cfg.ScreenWidth) - float64(g.ship.width)/2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.ship.y -= g.cfg.ShipSpeedFactor
		if g.ship.y < -float64(g.ship.height)/2 {
			g.ship.y = -float64(g.ship.height) / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.ship.y += g.cfg.ShipSpeedFactor
		if g.ship.y > float64(g.cfg.ScreenHeight)-float64(g.ship.height)/2 {
			g.ship.y = float64(g.cfg.ScreenHeight) - float64(g.ship.height)/2
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if len(g.bullets) < g.cfg.MaxBulletNum &&
			time.Now().Sub(i.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
			bullet := NewBullet(g.cfg, g.ship)
			g.addBullet(bullet)

			i.lastBulletTime = time.Now()
		}
	}
}
