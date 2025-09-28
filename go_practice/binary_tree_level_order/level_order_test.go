package levelorder

import (
	"reflect"
	"testing"
)

func ptr(v int) *int { return &v }

func buildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &TreeNode{Val: *v}
		}
	}

	for i := 0; i < len(vals); i++ {
		if nodes[i] == nil {
			continue
		}

		left := 2*i + 1
		right := 2*i + 2

		if left < len(vals) {
			nodes[i].Left = nodes[left]
		}
		if right < len(vals) {
			nodes[i].Right = nodes[right]
		}
	}

	return nodes[0]
}

func TestLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		name string
		tree []*int
		want [][]int
	}{
		{
			name: "balanced",
			tree: []*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "single node",
			tree: []*int{ptr(1)},
			want: [][]int{{1}},
		},
		{
			name: "right skewed",
			tree: []*int{ptr(1), nil, ptr(2), nil, nil, nil, ptr(3)},
			want: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.tree)
			got := LevelOrder(root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("unexpected traversal. want %v, got %v", tc.want, got)
			}
		})
	}
}
