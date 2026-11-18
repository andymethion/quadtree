package floor

import (
	"os"
	"strconv"
)

// Fonction qui créé ou réecris un fichier Save.txt qui contient un array de la map aleatoire
func SaveRandom(newfloor [][]int) {
	// Ouverture ou création du fichier
	file, err := os.OpenFile("../floor-files/Save.txt", os.O_CREATE|os.O_WRONLY, 0600)
	// Verification de l'erreur
	if err != nil {
		panic(err)
	}

	defer file.Close() // on ferme automatiquement le fichier à la fin de notre programme

	for y := 0; y < len(newfloor); y++ {
		// création d'un string vide pour chaque nouvelle ligne de l'array
		line := ""

		// ajout de chaque element de la ligne Y de l'array dans le string line
		for x := 0; x < len(newfloor); x++ {
			line += strconv.Itoa(newfloor[y][x])
		}
		_, err = file.WriteString(line + "\n") // écrire dans le fichier la ligne Y de l'array newfloor
		if err != nil {
			panic(err)
		}
	}
}
