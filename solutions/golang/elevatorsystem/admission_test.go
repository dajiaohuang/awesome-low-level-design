package elevatorsystem
import "testing"
func TestControllerReportsFullQueue(t *testing.T){
 e:=NewElevator(1,1); c:=&ElevatorController{elevators:[]*Elevator{e}}
 if !c.RequestElevator(1,2) {t.Fatal("first request rejected")}
 if c.RequestElevator(1,3) {t.Fatal("full queue reported accepted")}
}
