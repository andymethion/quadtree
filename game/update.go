package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))

	if configuration.Global.EnableCameraBlocking {
		cameraX, cameraY := g.camera.Blocking(g.floor.GetContent(), g.character.X, g.character.Y)
		g.camera.Update(cameraX, cameraY)
	} else {
		g.camera.Update(g.character.X, g.character.Y)
	}

	g.floor.Update(g.camera.X, g.camera.Y)

	if configuration.Global.EnableZoom {
		g.camera.Zoom(&g.floor)
	}

	if configuration.Global.EnablePortals {
		g.Portals()
	}

	return nil
}
