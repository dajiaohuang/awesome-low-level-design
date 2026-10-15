package trafficsignalsystem
import("testing";"time")
func TestCoordinatedCycleAndEmergency(t *testing.T){
 tc:=GetTrafficController();a:=NewRoad("A","First");b:=NewRoad("B","Second");a.SetTrafficLight(NewTrafficLight("A",10,5,10));b.SetTrafficLight(NewTrafficLight("B",10,5,10));tc.AddRoad(a);tc.AddRoad(b);tc.StartTrafficControl()
 defer func(){tc.mu.Lock();tc.generation++;tc.timer.Stop();tc.mu.Unlock()}()
 snapshot:=func()(Signal,Signal){tc.mu.Lock();defer tc.mu.Unlock();return a.TrafficLight.CurrentSignal,b.TrafficLight.CurrentSignal}
 deadline:=time.Now().Add(time.Second);for{first,_:=snapshot();if first==Green{break};if time.Now().After(deadline){t.Fatal("first road never advanced")};time.Sleep(time.Millisecond)}
 tc.HandleEmergency("B");first,second:=snapshot();if first!=Red||second!=Green{t.Fatal("override did not red-light conflicting road",first,second)}
 seenA,seenB:=false,true;deadline=time.Now().Add(250*time.Millisecond);for time.Now().Before(deadline){first,second=snapshot();if first!=Red && second!=Red{t.Fatal("conflicting live phases",first,second)};seenA=seenA||first==Green;seenB=seenB||second==Green;time.Sleep(time.Millisecond)}
 if !seenA||!seenB{t.Fatal("normal coordinated cycle did not continue")}
}
