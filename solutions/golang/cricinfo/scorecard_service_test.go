package cricinfo

import "testing"

func TestCreateScorecardSetsMatchRelationship(t *testing.T) {
	match := &Match{ID: "match-1"}
	service := &ScorecardService{scorecards: make(map[string]*Scorecard)}
	id := service.CreateScorecard(match)
	if match.Scorecard == nil {
		t.Fatal("match scorecard relationship was not populated")
	}
	if match.Scorecard != service.GetScorecard(id) {
		t.Fatal("match relationship does not point to the registered scorecard")
	}
}
