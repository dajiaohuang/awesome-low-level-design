package airlinemanagementsystem
import("testing";"time")
func TestSeatReservationLifecycle(t *testing.T){
 s:=NewAirlineManagementSystem();date:=time.Now();f:=NewFlight("seat-flight","A","B",date,date.Add(time.Hour));p:=NewPassenger("P","Passenger","p@example.com","123");s.AddFlight(f);seat:=NewSeat("25A",SeatTypeEconomy)
 b:=s.BookFlight(f,p,seat,100);if b==nil||seat.GetStatus()!=SeatStatusReserved{t.Fatal("not reserved")}
 if s.BookFlight(f,p,seat,100)!=nil{t.Fatal("double confirmed")}
 s.CancelBooking(b.BookingNumber);if seat.GetStatus()!=SeatStatusAvailable{t.Fatal("not released")}
 if s.BookFlight(f,p,seat,100)==nil{t.Fatal("not rebookable")}
 s.CancelBooking(b.BookingNumber);if seat.GetStatus()!=SeatStatusReserved{t.Fatal("stale cancellation released new booking")}
}
