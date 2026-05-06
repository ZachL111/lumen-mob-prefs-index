package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 77, Capacity: 103, Latency: 16, Risk: 6, Weight: 5}, wantScore: 193, wantDecision: "accept"},
		{name: "case_2", signal: Signal{Demand: 71, Capacity: 84, Latency: 14, Risk: 9, Weight: 4}, wantScore: 149, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 98, Capacity: 71, Latency: 21, Risk: 16, Weight: 5}, wantScore: 133, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
