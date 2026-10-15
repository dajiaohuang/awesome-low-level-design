package hotelmanagement
import("testing";"time")
func TestCancelBookedRoom(t *testing.T){
 s:=GetHotelManagementSystem();g:=&Guest{ID:"G1"};r:=NewRoom("R1",RoomTypeSingle,100);s.AddGuest(g);s.AddRoom(r)
 b,e:=s.BookRoom(g,r,time.Now(),time.Now().Add(24*time.Hour));if e!=nil{t.Fatal(e)}
 if e=s.CancelReservation(b.ID);e!=nil{t.Fatal(e)}
 if r.GetStatus()!=RoomStatusAvailable{t.Fatal("canceled room not available")}
 if _,exists:=s.reservations[b.ID];exists{t.Fatal("canceled registry entry retained")}
}
