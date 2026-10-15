package elevatorsystem
import ("testing";"time")
func TestRequestVisitsPickupBeforeDestination(t *testing.T){
 e:=NewElevator(1,2);e.AddRequest(NewRequest(3,2));e.Run();defer e.Stop()
 deadline:=time.NewTimer(8*time.Second);defer deadline.Stop();tick:=time.NewTicker(5*time.Millisecond);defer tick.Stop()
 seenPickup:=false
 for {select {case <-deadline.C:t.Fatal("pickup/destination timeout");case <-tick.C:
  floor:=e.getCurrentFloor();if floor==3 {seenPickup=true};if floor==2&&seenPickup{return}
 }}
}
