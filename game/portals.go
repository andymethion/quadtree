package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Fonction permettant de placer les portails si la touche "p" est préssée.
// Elle utilise le booléan firstPortalToSet pour déterminer quel est le portail à poser
// Le booléan s'inverse à chaque portail posé.
// Le portail est placé aux coordonnées du joueur
func (g *Game) placePortals() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		if g.setFirstPortal {
			g.firstPortal.X = g.character.X
			g.firstPortal.Y = g.character.Y
		} else {
			g.secondPortal.X = g.character.X
			g.secondPortal.Y = g.character.Y
		}
		g.setFirstPortal = !g.setFirstPortal
		g.justUsePortal = true
	}
}

// Téléporte le joueur d'un portail à un autre
func (g *Game) teleportToPortal(portal Portal) {
	g.character.X = portal.X
	g.character.Y = portal.Y
	g.justUsePortal = true

	// Met à jour le jeu pour s'adapter à la téléportation
	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.character.X, g.character.Y)
}

// Téléporte le joueur au second portail lorsqu'il marche sur le premier portail
func (g *Game) teleportToSecondPortal() {
	if g.character.X == g.firstPortal.X && g.character.Y == g.firstPortal.Y && !g.justUsePortal && !g.character.GetMoving() {
		g.teleportToPortal(g.secondPortal)
	}
}

// Téléporte le joueur au premier portail lorsqu'il marche sur le second portail
func (g *Game) teleportToFirstPortal() {
	if g.character.X == g.secondPortal.X && g.character.Y == g.secondPortal.Y && !g.justUsePortal && !g.character.GetMoving() {
		g.teleportToPortal(g.firstPortal)
	}
}

func (g *Game) Portals() {
	// Lorsque nous posons ou spawnons sur un portail il est necéssaire que justUsePortal soit sur true pour ne pas...
	// ... que nous soyons téléporter à l'infini entre les deux portails
	// Cette condition remet justUsePortal à false lorsque nous quittons un portail pour pouvoir à nouveau les utiliser
	if g.character.GetMoving() {
		g.justUsePortal = false
	}

	g.placePortals()

	// S'assure que les deux portails sont posés pour pouvoir se téléporter
	if g.firstPortal.X != -1 && g.firstPortal.Y != -1 && g.secondPortal.X != -1 && g.secondPortal.Y != -1 {
		g.teleportToSecondPortal()
		g.teleportToFirstPortal()
	}
}
