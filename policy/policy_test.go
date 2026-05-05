package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 77, Capacity: 103, Latency: 16, Risk: 6, Weight: 5}
	if got := Score(signal); got != 193 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 71, Capacity: 84, Latency: 14, Risk: 9, Weight: 4}
	if got := Score(signal); got != 149 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 98, Capacity: 71, Latency: 21, Risk: 16, Weight: 5}
	if got := Score(signal); got != 133 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
