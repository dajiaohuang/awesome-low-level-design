package snakeandladdergame
import ("testing";"sync";"time")
func TestConcurrentGameCreation(t *testing.T){
 const count=8
 var callers, games sync.WaitGroup
 managers:=make(chan *GameManager,count)
 for i:=0;i<count;i++ { callers.Add(1);go func(){defer callers.Done();gm:=GetGameManager();managers<-gm;gm.StartNewGame(&games,[]string{"Alice","Bob"})}() }
 callers.Wait(); close(managers)
 gm:=GetGameManager();for got:=range managers{if got!=gm{t.Fatal("different managers")}}
 if len(gm.Games)!=count {t.Fatal("lost session")}
 done:=make(chan struct{});go func(){games.Wait();close(done)}()
 select {case <-done:case <-time.After(45*time.Second):t.Fatal("bounded real game completion timeout")}
}
