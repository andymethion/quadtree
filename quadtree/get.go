package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	// Pour chaque pixel de contentHolder
	for contentHolderY := 0; contentHolderY < len(contentHolder); contentHolderY++ {
		for contentHolderX := 0; contentHolderX < len(contentHolder[contentHolderY]); contentHolderX++ {

			// On récupère la position du noeud concerné
			nodeX := topLeftX + contentHolderX
			nodeY := topLeftY + contentHolderY

			// On modifie le pixel de contentHolder en fonction du noeud concerné
			contentHolder[contentHolderY][contentHolderX] = q.root.getContentFromNode(nodeX, nodeY)
		}
	}
}

// Cette fonction permet d'obtenir la valeur d'un noeud à partir de ses coordonnées
func (n *node) getContentFromNode(nodeX int, nodeY int) int {
	// On s'assure que le noeud n'est pas vide
	if n == nil {
		return -1
	}

	// On s'assure que les coordonnées du noeud ne dépassent pas
	firstCondition := n.topLeftX <= nodeX && n.topLeftX+n.width > nodeX
	secondCondition := n.topLeftY <= nodeY && n.topLeftY+n.height > nodeY

	if !firstCondition || !secondCondition {
		return -1
	}

	// On retourne le contenu du noeud lorsqu'il est une feuille
	if n.isLeaf {
		return n.content
	}

	// On récupère la moitié de la largeur et de la hauteur du noeud
	halfNodeWidth := n.width / 2
	halfNodeHeight := n.height / 2

	// On récupère la médianne de la largeur et de la hauteur du noeud
	midPointX := n.topLeftX + halfNodeWidth
	midPointY := n.topLeftY + halfNodeHeight

	if nodeY < midPointY {
		// Partie haute du noeud
		if nodeX < midPointX {
			// Partie gauche du noeud
			return n.topLeftNode.getContentFromNode(nodeX, nodeY)
		} else {
			// Partie droite du noeud
			return n.topRightNode.getContentFromNode(nodeX, nodeY)
		}
	} else {
		// Partie basse du noeud
		if nodeX < midPointX {
			// Partie gauche du noeud
			return n.bottomLeftNode.getContentFromNode(nodeX, nodeY)
		} else {
			// Partie droite du noeud
			return n.bottomRightNode.getContentFromNode(nodeX, nodeY)
		}
	}
}
