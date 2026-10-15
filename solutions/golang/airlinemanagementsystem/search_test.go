package airlinemanagementsystem
import ("testing"; "time")
func TestAddedFlightSearch(t *testing.T) {
 s:=NewAirlineManagementSystem(); date:=time.Date(2026,10,20,12,0,0,0,time.UTC)
 f:=NewFlight("F001","New York","London",date,date.Add(time.Hour)); s.AddFlight(f)
 results:=s.SearchFlights("New York","London",date)
 if len(results)!=1 || results[0]!=f {t.Fatal("new catalog flight not found")}
}
