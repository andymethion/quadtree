package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/character"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
)

type Portal struct {
	X, Y int
}

// Game est le type permettant de représenter les données du jeu.
// Aucun champs n'est exporté pour le moment.
//
// Les champs non exportés sont :
//   - camera : la représentation de la caméra
//   - floor : la représentation du terrain
//   - character : la représentation du personnage
//   - firstPortalToSet: permet de connaître l'ordre dans lequel nous devons modifier les coordonnées des portails
//   - justUsePortal: permet de déterminer s'il on vient d'utiliser un portail (pour empecher le personnage de faire du ping pong entre les deux portails)
type Game struct {
	camera         camera.Camera
	floor          floor.Floor
	character      character.Character
	firstPortal    Portal
	secondPortal   Portal
	setFirstPortal bool
	justUsePortal  bool
}
