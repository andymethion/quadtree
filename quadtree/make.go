package quadtree

import (
	"math"
)

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {

	// Appel de la fonction qui crée le cadrée
	q = makeFromArrayRecuQuad(floorContent)

	// Implémentation des TopLeftX et TopLeftY sur ce quadtree
	makeFromArrayRecuTopLeft(q.root)

	return q
}

// Fonction récursive qui prend une liste de liste en paramètre et renvoie un quadtree avec les TopLeft pas mis à jour
func makeFromArrayRecuQuad(floorContent [][]int) (q Quadtree) {

	// Obtention des constantes pour la taille du tableau
	H := len(floorContent)
	W := 0
	if H > 0 {
		W = len(floorContent[0])
	}

	CW := int(math.Ceil(float64(W) / 2))
	CH := int(math.Ceil(float64(H) / 2))

	FW := int(math.Floor(float64(W) / 2))
	FH := int(math.Floor(float64(H) / 2))

	// Première condition d'arrêt
	if W <= 1 && H <= 1 {
		q.width = W
		q.height = H

		// prise en compte cas où la feuille a une width ou height égale à 0
		content2 := 0
		if W != 0 && H != 0 {
			content2 = floorContent[0][0]
		}

		// creation du noeud
		q.root = &node{topLeftX: 0, topLeftY: 0, width: W, height: H, isLeaf: true, content: content2}
		return q
	} else {

		// Vérification du floorContent pour voir si une réduction est possible
		ContentIsAllSame := true
		FirstContent := floorContent[0][0]

		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if floorContent[i][j] != FirstContent {
					ContentIsAllSame = false
					break
				}
			}
			if !ContentIsAllSame {
				break
			}
		}

		// Récursivité pour construire le quadtree
		if !ContentIsAllSame {

			topL, topR, bottomL, bottomR := makeQuadtreeCorner(H, W, FH, FW, floorContent)

			// création d'un nouveau noeud du quadtree avec la récursivité pour poursuivre le quadtree
			q.width = W
			q.height = H
			q.root = &node{topLeftX: 0, topLeftY: 0, width: W, height: H,
				topLeftNode:     makeFromArrayRecuQuad(topL).root,
				topRightNode:    makeFromArrayRecuQuad(topR).root,
				bottomLeftNode:  makeFromArrayRecuQuad(bottomL).root,
				bottomRightNode: makeFromArrayRecuQuad(bottomR).root,
			}
			// mise a jour du width et height de tout les noeuds
			q.root.topLeftNode.width, q.root.topLeftNode.height = FW, FH
			q.root.topRightNode.width, q.root.topRightNode.height = CW, FH
			q.root.bottomLeftNode.width, q.root.bottomLeftNode.height = FW, CH
			q.root.bottomRightNode.width, q.root.bottomRightNode.height = CW, CH

			return q
		} else {
			// Réduction

			q.width = W
			q.height = H

			q.root = &node{topLeftX: 0, topLeftY: 0, width: W, height: H, isLeaf: true, content: FirstContent}
			return q
		}
	}
}

// Fonction récursive qui prend un noeud root et modifie toutes les branches jusqu'à la feuille et met à jour le TopLeftX et TopLeftY
func makeFromArrayRecuTopLeft(root *node) {
	// condition d'arret
	if !root.isLeaf {
		W := root.width
		H := root.height
		// mise a jour
		root.topLeftNode.topLeftX, root.topLeftNode.topLeftY = root.topLeftX, root.topLeftY
		root.topRightNode.topLeftX, root.topRightNode.topLeftY = root.topLeftX+W/2, root.topLeftY
		root.bottomLeftNode.topLeftX, root.bottomLeftNode.topLeftY = root.topLeftX, root.topLeftY+H/2
		root.bottomRightNode.topLeftX, root.bottomRightNode.topLeftY = root.topLeftX+W/2, root.topLeftY+H/2
		// récursivité
		makeFromArrayRecuTopLeft(root.topLeftNode)
		makeFromArrayRecuTopLeft(root.topRightNode)
		makeFromArrayRecuTopLeft(root.bottomLeftNode)
		makeFromArrayRecuTopLeft(root.bottomRightNode)
	}
}

// fonction qui permet de créé les coins du quadtree a partir d'un floorcontent et de sa HEIGHT et WIDTH
func makeQuadtreeCorner(H, W, FH, FW int, floorContent [][]int) (topL, topR, bottomL, bottomR [][]int) {

	// création des listes de listes représentant les coins du floorContent
	topL = [][]int{}
	topR = [][]int{}
	bottomL = [][]int{}
	bottomR = [][]int{}

	// boucle pour parcourir la height du floorContent
	for i := 0; i < H; i++ {
		// mise a jour de la hauteur des nouvelles listes
		empty := []int{}
		up := false

		if i < FH {
			// vérifie si on est en haut du floorContent
			up = true
			topL = append(topL, empty)
			topR = append(topR, empty)
		} else {
			bottomL = append(bottomL, empty)
			bottomR = append(bottomR, empty)
		}

		// Boucle pour parcourir la width du floorContent
		for j := 0; j < W; j++ {
			// vérifie si on est sur un côté gauche du floorContent
			left := false
			if j < FW {
				left = true
			}
			// Ajout des éléments dans les bonnes listes en fonction de leur position dans le floorContent
			if up {
				if left {
					topL[i] = append(topL[i], floorContent[i][j])
				} else {
					topR[i] = append(topR[i], floorContent[i][j])
				}
			} else {
				if left {
					bottomL[i-FH] = append(bottomL[i-FH], floorContent[i][j])
				} else {
					bottomR[i-FH] = append(bottomR[i-FH], floorContent[i][j])
				}
			}
		}
	}

	return topL, topR, bottomL, bottomR
}
