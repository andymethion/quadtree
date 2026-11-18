package camera

// Cette fonction retourne les nouvelles coordonnées de la caméra de...
// ...sorte à ce que le vide ne soit jamais affiché sur l'écran
func (c *Camera) Blocking(content [][]int, characterX, characterY int) (int, int) {
	cameraWidth := len(content[0])
	cameraHeight := len(content)

	// Calcul les différents decalages nécessaire à la camera

	shiftLeft := cameraWidth/2 - characterX
	shiftRight := characterX - cameraWidth - 1

	shiftTop := cameraHeight/2 - characterY
	shiftBottom := characterY - cameraHeight - 1

	// Applique les décalages à la caméra (s'il y en a)

	if shiftLeft > 0 {
		characterX += shiftLeft
	}

	if shiftRight > 0 {
		characterX -= shiftRight
	}

	if shiftTop > 0 {
		characterY += shiftTop
	}

	if shiftBottom > 0 {
		characterY -= shiftBottom
	}

	return characterX, characterY
}
