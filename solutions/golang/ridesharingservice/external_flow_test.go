package ridesharingservice_test
import ("testing";"time";service "github.com/ashishps1/awesome-low-level-design/solutions/golang/ridesharingservice")
func TestExternalDriverRequestFlow(t *testing.T){
 s:=service.GetRideService(); location:=&service.Location{Latitude:1,Longitude:1}
 p:=&service.Passenger{ID:1,Name:"Passenger",Location:location}; d:=&service.Driver{ID:1,Name:"Driver",Location:location,Status:service.Available}
 s.AddPassenger(p);s.AddDriver(d);s.RequestRide(p,location,&service.Location{Latitude:2,Longitude:2})
 select {case ride:=<-s.PendingRideRequests():s.AcceptRide(d,ride);s.StartRide(ride);s.CompleteRide(ride);if ride.Status!=service.Completed || d.Status!=service.Available {t.Fatal("driver flow failed")};case <-time.After(time.Second):t.Fatal("request unavailable")}
}
