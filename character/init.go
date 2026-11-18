package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init(quadtreeContent quadtree.Quadtree, floorWidth, floorHeight int) {
	c.animationStep = 1

	c.X = floorWidth / 2
	c.Y = floorHeight / 2

	// Empêche le joueur de spawner et de marcher sur l'eau
	if configuration.Global.BlockingWater {
		// Récupère les coordonnées d'un noeud qui n'est pas de l'eau
		topLeftX, topLeftY := quadtreeContent.FindFirstLandNodeCoordinates()

		// Si on trouve des coordonnées valides alors on fait spawn le personnage sur ce block
		// Sinon on remplace tout le quadtree qui est de l'eau par de la terre...
		// ... et les coordonnées du joueur ne sont pas modifiées
		if topLeftX != -1 && topLeftY != -1 {
			c.X, c.Y = topLeftX, topLeftY
		} else {
			quadtreeContent.SpawnLandNode()
		}
	}
}

func (c *Character) GetMoving() bool {
	return c.moving
}
