package floor

import (
	"math/rand"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Function pour créé un Array remplie d'entier aléatoire qui vas servir a créé une map completement aléatoire
func CreateRandomFloor(tailleX int, tailleY int) (newFloor [][]int) {

	// remplissage aléatoire de l'array de taille tailleX en largeur et tailleY en hauteur
	for y := 0; y < tailleY; y++ {
		newFloor = append(newFloor, []int{})
		for x := 0; x < tailleX; x++ {
			newFloor[y] = append(newFloor[y], rand.Intn(6))
		}
	}

	// appel de la fonction de sauvegarde si l'option est activé
	if configuration.Global.EnableRandomSave {
		SaveRandom(newFloor)
	}

	return newFloor
}
