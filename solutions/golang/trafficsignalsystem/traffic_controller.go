package trafficsignalsystem

import (
 "sort"
 "sync"
 "time"
)

type TrafficController struct {
 roads map[string]*Road
 mu sync.Mutex
 order []string
 activeIndex int
 phase Signal
 timer *time.Timer
 generation uint64
}

var instance *TrafficController
var once sync.Once

func GetTrafficController() *TrafficController {
 once.Do(func(){instance=&TrafficController{roads:make(map[string]*Road)}})
 return instance
}

func (tc *TrafficController) AddRoad(road *Road) {
 tc.mu.Lock()
 defer tc.mu.Unlock()
 tc.roads[road.ID]=road
}

func (tc *TrafficController) prepareCycleLocked() {
 if tc.timer!=nil {tc.timer.Stop()}
 tc.generation++
 tc.order=make([]string,0,len(tc.roads))
 for id,road:=range tc.roads {
  tc.order=append(tc.order,id)
  road.TrafficLight.ChangeSignal(Red)
 }
 sort.Strings(tc.order)
}

func (tc *TrafficController) StartTrafficControl() {
 tc.mu.Lock()
 defer tc.mu.Unlock()
 tc.prepareCycleLocked()
 if len(tc.order)==0 {return}
 tc.activeIndex=0
 tc.setPhaseLocked(Red)
}

func (tc *TrafficController) setPhaseLocked(phase Signal) {
 tc.phase=phase
 light:=tc.roads[tc.order[tc.activeIndex]].TrafficLight
 light.ChangeSignal(phase)
 var duration int
 switch phase {
 case Red: duration=light.RedDuration
 case Green: duration=light.GreenDuration
 case Yellow: duration=light.YellowDuration
 }
 generation:=tc.generation
 tc.timer=time.AfterFunc(time.Duration(duration)*time.Millisecond,func(){tc.advance(generation)})
}

func (tc *TrafficController) advance(generation uint64) {
 tc.mu.Lock()
 defer tc.mu.Unlock()
 // Stop cannot retract a callback already waiting for the controller lock.
 if generation!=tc.generation {return}
 switch tc.phase {
 case Red: tc.setPhaseLocked(Green)
 case Green: tc.setPhaseLocked(Yellow)
 case Yellow:
  tc.roads[tc.order[tc.activeIndex]].TrafficLight.ChangeSignal(Red)
  tc.activeIndex=(tc.activeIndex+1)%len(tc.order)
  tc.setPhaseLocked(Red)
 }
}

func (tc *TrafficController) HandleEmergency(roadID string) {
 tc.mu.Lock()
 defer tc.mu.Unlock()
 if _,exists:=tc.roads[roadID];!exists {return}
 tc.prepareCycleLocked()
 for index,id:=range tc.order {
  if id==roadID {tc.activeIndex=index;break}
 }
 tc.setPhaseLocked(Green)
}
