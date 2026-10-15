package onlineshopping
import "testing"
func TestOrderRetentionAndUnavailableCart(t *testing.T) {
 s := GetInstance(); u := NewUser("retention", "User", "u@example.com", "password"); s.RegisterUser(u)
 p := NewProduct("available", "Meal", "Item", 10, 3); unavailable := NewProduct("empty", "Empty", "Item", 10, 0)
 s.AddProduct(p); s.AddProduct(unavailable)
 purchase := func() *Order { c:=NewShoppingCart(); c.AddItem(p,1); o,e:=s.PlaceOrder(u,c,&CreditCardPayment{}); if e!=nil {t.Fatal(e)}; return o }
 first:=purchase(); second:=purchase()
 if first.ID==second.ID || s.Orders[first.ID]!=first || s.Orders[second.ID]!=second {t.Fatal("order overwritten")}
 c:=NewShoppingCart(); c.AddItem(p,1); c.AddItem(unavailable,1)
 if o,e:=s.PlaceOrder(u,c,&CreditCardPayment{}); e==nil || o!=nil {t.Fatal("partial cart silently accepted")}
 if len(c.GetItems())!=2 || !p.IsAvailable(1) {t.Fatal("rejection lost cart or stock")}
}
