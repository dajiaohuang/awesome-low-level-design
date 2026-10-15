package composite

import "testing"

func TestRemoveRemovesRequestedSameNamedChild(t *testing.T) {
	folder := NewFolder("parent")
	first := NewFile("duplicate", 10)
	requested := NewFile("duplicate", 20)
	folder.Add(first)
	folder.Add(requested)

	folder.Remove(requested)

	if got := folder.GetSize(); got != 10 {
		t.Fatalf("remaining size = %d, want 10 after removing requested child", got)
	}
	if len(folder.children) != 1 || folder.children[0] != first {
		t.Fatalf("remaining child = %#v, want first child", folder.children)
	}
}
