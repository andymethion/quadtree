package camera

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
)

// Zoom en supprimant des lignes et colonnes dans floor.content
// La fonction s'assure que la taille de la map soit au minimum de 1x1
func (c *Camera) zoomIn(f *floor.Floor) {
	content := f.GetContent()

	// Arrête la fonction si la taille de la map est de 3x3
	if len(content) == 3 && len(content[0]) == 3 {
		return
	}

	// On supprime la dernière ligne
	content = content[:len(content)-1]

	// On supprime la dernière colonne (équivaut à supprimer le dernier élément de chaque ligne)
	for i := range content {
		content[i] = content[i][:len(content[i])-1]
	}

	// On retire 1 au nombre de lignes et colonnes dans la configuration
	configuration.Global.NumTileY -= 1
	configuration.Global.NumTileX -= 1

	// On agrandit la taille de l'écran
	configuration.Global.ScreenHeight -= configuration.Global.TileSize
	configuration.Global.ScreenWidth -= configuration.Global.TileSize

	// On replace correctement la caméra au centre
	configuration.Global.ScreenCenterTileX = len(content[0]) / 2
	configuration.Global.ScreenCenterTileY = len(content) / 2

	// On modifie floor.content
	f.SetContent(content)
}

// Dézoom en ajoutant des lignes et colonnes dans floor.content
func (c *Camera) zoomOut(f *floor.Floor) {
	content := f.GetContent()

	// On ajoute une ligne à la fin
	content = append(content, make([]int, len(content[0])))

	// On ajoute une colonne à la fin (équivaut à ajouter un élément à la fin de chaque ligne)
	for i := range content {
		content[i] = append(content[i], 0)
	}

	// On ajoute 1 au nombre de lignes et colonnes dans la configuration
	configuration.Global.NumTileY += 1
	configuration.Global.NumTileX += 1

	// On rétrécit la taille de l'écran
	configuration.Global.ScreenHeight += configuration.Global.TileSize
	configuration.Global.ScreenWidth += configuration.Global.TileSize

	// On replace correctement la caméra au centre
	configuration.Global.ScreenCenterTileX = len(content[0]) / 2
	configuration.Global.ScreenCenterTileY = len(content) / 2

	// On modifie floor.content
	f.SetContent(content)
}

// Fonction principale de la feature zoom
func (c *Camera) Zoom(f *floor.Floor) {
	// Zoom lorsque la touche i est préssée
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		c.zoomIn(f)
	}

	// Dézoom lorsque la touche j est préssée
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		c.zoomOut(f)
	}
}
