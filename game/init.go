package game

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	g.character.Init(g.floor.GetQuadtreeContent(), g.floor.GetWidth(), g.floor.GetHeight())
	g.camera.Init(g.character.X, g.character.Y)

	// Définit les valeurs par défaut des portails
	// Veuillez consuler character/character.go pour en savoir plus

	g.firstPortal.X = -1
	g.firstPortal.Y = -1

	g.secondPortal.X = -1
	g.secondPortal.Y = -1

	g.setFirstPortal = true
	g.justUsePortal = true
}
