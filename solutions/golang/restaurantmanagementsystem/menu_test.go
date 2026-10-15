package restaurantmanagementsystem
import "testing"
func TestMenuSnapshotSurvivesRemoval(t *testing.T) {
 r := GetRestaurantInstance()
 first := &MenuItem{ID:1}; second := &MenuItem{ID:2}
 r.AddMenuItem(first); r.AddMenuItem(second)
 menu := r.GetMenu(); r.RemoveMenuItem(first)
 if len(menu) != 2 || menu[0].ID != 1 || menu[1].ID != 2 { t.Fatal("returned menu changed after removal") }
}
