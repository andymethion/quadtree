package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	case RandomFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX
			if absX < 0 {
				absX = -absX
			}
			absY := topLeftY
			if absY < 0 {
				absY = -absY
			}
			f.content[y][x] = ((x + absX%2) + (y + absY%2)) % 2
		}
	}
}

// updateFromFileFloor met à jour le contenu visible du sol (f.content)
// à partir du contenu complet (f.fullContent), en fonction de la position donnée
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
	// Parcourir chaque pixel du contenu visible
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			// Calculer les coordonnées correspondantes dans le contenu complet
			fullY := topLeftY + y
			fullX := topLeftX + x

			// Vérifier si les coordonnées sont dans les limites du contenu complet
			if fullY < len(f.fullContent) && fullY >= 0 && fullX < len(f.fullContent[0]) && fullX >= 0 {
				f.content[y][x] = f.fullContent[fullY][fullX]
			} else {
				f.content[y][x] = -1
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
