package quadtree

// FindFirstLandNodeCoordinates cherche les coordonnées du premier nœud terrestre valide
func (q Quadtree) FindFirstLandNodeCoordinates() (int, int) {
	return findFirstLandNodeCoordinatesRecursive(q.root)
}

// findFirstLandNodeCoordinatesRecursive cherche récursivement les coordonnées du premier nœud terrestre valide
func findFirstLandNodeCoordinatesRecursive(n *node) (int, int) {
	// Retourne le nœud s'il respecte les conditions
	if isValidLandNode(n) {
		return n.topLeftX, n.topLeftY
	}

	// Lance la récursivité pour chacun des nœuds enfants
	for _, child := range []*node{n.topLeftNode, n.topRightNode, n.bottomLeftNode, n.bottomRightNode} {
		if child != nil && (isValidLandNode(child) || !child.isLeaf) {
			return findFirstLandNodeCoordinatesRecursive(child)
		}
	}

	// Retourne -1, -1 si tous les nœuds sont de l'eau
	// Dans ce cas, le personnage apparaîtra aux coordonnées 0, 0 sur un bloc de terre
	return -1, -1
}

// isValidLandNode vérifie si le nœud est une feuille terrestre de taille 1x1
func isValidLandNode(n *node) bool {
	return n.isLeaf && n.content != 5 && n.width == 1 && n.height == 1
}

// SpawnLandNode fait apparaître de la terre dans tout quadtree
// Cette fonction est appelée uniquement lorsque tous les nœuds sont de l'eau
func (q Quadtree) SpawnLandNode() {
	currentNode := q.root

	for currentNode.topLeftNode != nil {
		currentNode = currentNode.topLeftNode
	}

	currentNode.content = 0
}
