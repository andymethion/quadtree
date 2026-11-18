package quadtree

import "testing"

// Test le cas où le quadtree contient plusieurs cases sans eau
func TestMultipleLands(t *testing.T) {
	quadtrees := []Quadtree{
		{
			width: 6, height: 4, root: &node{topLeftX: 0, topLeftY: 0, width: 6, height: 4,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, width: 3, height: 2, isLeaf: true, content: 5},
				topRightNode: &node{topLeftX: 3, topLeftY: 0, width: 3, height: 2,
					topLeftNode:     &node{topLeftX: 3, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 5},
					topRightNode:    &node{topLeftX: 3, topLeftY: 0, width: 2, height: 1, isLeaf: true, content: 5},
					bottomLeftNode:  &node{topLeftX: 3, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 1, width: 2, height: 1, isLeaf: true, content: 5},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, width: 3, height: 2,
					topLeftNode: &node{topLeftX: 0, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 0},
					topRightNode: &node{topLeftX: 1, topLeftY: 2, width: 2, height: 1,
						topLeftNode:     &node{topLeftX: 1, topLeftY: 2, width: 1, height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 2, topLeftY: 2, width: 1, height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 1, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 2, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 5},
					},
					bottomLeftNode: &node{topLeftX: 0, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 5},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 3, width: 2, height: 1,
						topLeftNode:     &node{topLeftX: 1, topLeftY: 3, width: 1, height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 2, topLeftY: 3, width: 1, height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 1, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 5},
						bottomRightNode: &node{topLeftX: 2, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 0},
					},
				},
				bottomRightNode: &node{topLeftX: 3, topLeftY: 2, width: 3, height: 2, isLeaf: true, content: 5},
			},
		},
		{
			width: 4, height: 4, root: &node{topLeftX: 0, topLeftY: 0, width: 4, height: 4,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 2,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 5},
					topRightNode:    &node{topLeftX: 1, topLeftY: 0, width: 1, height: 1, isLeaf: true, content: 5},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 5},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
				},
				topRightNode: &node{topLeftX: 2, topLeftY: 0, width: 2, height: 2, isLeaf: true, content: 5},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, width: 2, height: 2,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 5},
					topRightNode:    &node{topLeftX: 1, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 0},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 5},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 5},
				},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 2, width: 2, height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 3, topLeftY: 2, width: 1, height: 1, isLeaf: true, content: 0},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 5},
				},
			},
		},
	}

	results := [][]int{{3, 1}, {1, 1}}

	for i, quadtree := range quadtrees {
		result := results[i]

		topLeftX, topLeftY := quadtree.FindFirstLandNodeCoordinates()
		resultX, resultY := result[0], result[1]

		if topLeftX != resultX || topLeftY != resultY {
			t.Fatalf("%s\nPour ce quadtree le première terre à retourner était aux coordonnées (%d, %d) mais vous avez retourné (%d, %d)", quadtree.GetAsText(""), resultX, resultY, topLeftX, topLeftY)
		}
	}
}

// Test le cas où le quadtree ne contient qu'une seule case sans eau
func TestOneLand(t *testing.T) {
	quadtrees := []Quadtree{
		{
			width: 7, height: 2, root: &node{topLeftX: 0, topLeftY: 0, width: 7, height: 2,
				topLeftNode:    &node{topLeftX: 0, topLeftY: 0, width: 3, height: 1, isLeaf: true, content: 5},
				topRightNode:   &node{topLeftX: 3, topLeftY: 0, width: 4, height: 1, isLeaf: true, content: 5},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 1, width: 3, height: 1, isLeaf: true, content: 5},
				bottomRightNode: &node{topLeftX: 3, topLeftY: 1, width: 4, height: 1,
					topLeftNode:  &node{topLeftX: 3, topLeftY: 1, width: 2, height: 0, isLeaf: true, content: 0},
					topRightNode: &node{topLeftX: 5, topLeftY: 1, width: 2, height: 0, isLeaf: true, content: 0},
					bottomLeftNode: &node{topLeftX: 3, topLeftY: 1, width: 2, height: 1,
						topLeftNode:     &node{topLeftX: 3, topLeftY: 1, width: 1, height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 4, topLeftY: 1, width: 1, height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 3, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 4, topLeftY: 1, width: 1, height: 1, isLeaf: true, content: 5},
					},
					bottomRightNode: &node{topLeftX: 5, topLeftY: 1, width: 2, height: 1, isLeaf: true, content: 5},
				},
			},
		},
		{
			width: 16, height: 4, root: &node{topLeftX: 0, topLeftY: 0, width: 16, height: 4,
				topLeftNode:    &node{topLeftX: 0, topLeftY: 0, width: 8, height: 2, isLeaf: true, content: 5},
				topRightNode:   &node{topLeftX: 8, topLeftY: 0, width: 8, height: 2, isLeaf: true, content: 5},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, width: 8, height: 2, isLeaf: true, content: 5},
				bottomRightNode: &node{topLeftX: 8, topLeftY: 2, width: 8, height: 2,
					topLeftNode:    &node{topLeftX: 8, topLeftY: 2, width: 4, height: 1, isLeaf: true, content: 5},
					topRightNode:   &node{topLeftX: 12, topLeftY: 2, width: 4, height: 1, isLeaf: true, content: 5},
					bottomLeftNode: &node{topLeftX: 8, topLeftY: 3, width: 4, height: 1, isLeaf: true, content: 5},
					bottomRightNode: &node{topLeftX: 12, topLeftY: 3, width: 4, height: 1,
						topLeftNode:    &node{topLeftX: 12, topLeftY: 3, width: 2, height: 0, isLeaf: true, content: 0},
						topRightNode:   &node{topLeftX: 14, topLeftY: 3, width: 2, height: 0, isLeaf: true, content: 0},
						bottomLeftNode: &node{topLeftX: 12, topLeftY: 3, width: 2, height: 1, isLeaf: true, content: 5},
						bottomRightNode: &node{topLeftX: 14, topLeftY: 3, width: 2, height: 1,
							topLeftNode:     &node{topLeftX: 14, topLeftY: 3, width: 1, height: 0, isLeaf: true, content: 0},
							topRightNode:    &node{topLeftX: 15, topLeftY: 3, width: 1, height: 0, isLeaf: true, content: 0},
							bottomLeftNode:  &node{topLeftX: 14, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 5},
							bottomRightNode: &node{topLeftX: 15, topLeftY: 3, width: 1, height: 1, isLeaf: true, content: 0},
						},
					},
				},
			},
		},
	}

	results := [][]int{{3, 1}, {15, 3}}

	for i, quadtree := range quadtrees {
		result := results[i]

		topLeftX, topLeftY := quadtree.FindFirstLandNodeCoordinates()
		resultX, resultY := result[0], result[1]

		if topLeftX != resultX || topLeftY != resultY {
			t.Fatalf("%s\nPour ce quadtree le première terre à retourner était aux coordonnées (%d, %d) mais vous avez retourné (%d, %d)", quadtree.GetAsText(""), resultX, resultY, topLeftX, topLeftY)
		}
	}
}

// Test le cas où le quadtree est rempli d'eau
// Dans ce cas on est censé retourner (-1, -1)
func TestFullWater(t *testing.T) {
	quadtrees := []Quadtree{
		{width: 4, height: 4, root: &node{topLeftX: 0, topLeftY: 0, width: 4, height: 4, isLeaf: true, content: 5}},
		{width: 2, height: 4, root: &node{topLeftX: 0, topLeftY: 0, width: 2, height: 4, isLeaf: true, content: 5}},
	}

	for _, quadtree := range quadtrees {
		topLeftX, topLeftY := quadtree.FindFirstLandNodeCoordinates()

		if topLeftX != -1 || topLeftY != -1 {
			t.Fatalf("%s\nPour ce quadtree vous deviez retourner (-1, -1) puisqu'il n'y a aucune terre mais vous avez retourné (%d, %d)", quadtree.GetAsText(""), topLeftX, topLeftY)
		}
	}
}
