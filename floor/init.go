package floor

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	case RandomFloor:
		f.quadtreeContent = quadtree.MakeFromArray(CreateRandomFloor(configuration.Global.RandomSizeX, configuration.Global.RandomSizeY))
	}
}

// readFloorFromFile lit un fichier texte et le convertit en une matrice d'entiers.
// Chaque ligne du fichier devient une ligne de la matrice, et chaque caractère devient un entier.
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// Ouvre le fichier spécifié
	file, err := os.Open(fileName)

	// S'assure qu'aucune erreur s'est produite lors de l'ouverture du fichier
	if err != nil {
		log.Fatal(err)
	}

	// S'assure que le fichier sera fermé à la fin de la fonction
	defer file.Close()

	// Crée un scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	// Parcourt chaque ligne du fichier
	for scanner.Scan() {
		floorLine := []int{}

		// Parcourt chaque caractère de la ligne
		for _, charRune := range scanner.Text() {
			// Convertit le caractère en entier
			charInt, err := strconv.Atoi(string(charRune))

			// S'assure qu'aucune erreur s'est produite la conversion du caractère en entier
			if err != nil {
				log.Fatal(err)
			}

			// Ajoute l'entier à la ligne courante
			floorLine = append(floorLine, charInt)
		}

		// Ajoute la ligne complète à la matrice
		floorContent = append(floorContent, floorLine)
	}

	// Vérifie s'il y a eu une erreur lors de la lecture du fichier
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Retourne la matrice complète
	return floorContent
}
