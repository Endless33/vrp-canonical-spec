package vrptest

import "testing"

type AuthorityRaceDecision string

const (
	AuthorityAcceptedWinner       AuthorityRaceDecision = "ACCEPTED_WINNER"
	AuthorityRejectedLowerEpoch   AuthorityRaceDecision = "REJECTED_LOWER_EPOCH"
	AuthorityRejectedLowerPriority AuthorityRaceDecision = "REJECTED_LOWER_PRIORITY"
)

type AuthorityRaceInput struct {
	MutationID string
	Authority  string
	Epoch      int
	Priority   int
}

type AuthorityRaceWinner struct {
	MutationID string
	Authority  string
	Epoch      int
	Priority   int
}

func SelectAuthorityWinner(inputs []AuthorityRaceInput) AuthorityRaceWinner {
	winner := inputs[0]

	for _, input := range inputs[1:] {
		if input.Epoch > winner.Epoch {
			winner = input
			continue
		}

		if input.Epoch == winner.Epoch && input.Priority < winner.Priority {
			winner = input
			continue
		}
	}

	return AuthorityRaceWinner{
		MutationID: winner.MutationID,
		Authority:  winner.Authority,
		Epoch:      winner.Epoch,
		Priority:   winner.Priority,
	}
}

func TestAuthorityRaceSelectsSingleCanonicalWinner(t *testing.T) {
	inputs := []AuthorityRaceInput{
		{
			MutationID: "payment-001",
			Authority:  "node-a",
			Epoch:      2,
			Priority:   90,
		},
		{
			MutationID: "payment-001",
			Authority:  "node-b",
			Epoch:      3,
			Priority:   50,
		},
		{
			MutationID: "payment-001",
			Authority:  "node-c",
			Epoch:      3,
			Priority:   10,
		},
	}

	winner := SelectAuthorityWinner(inputs)

	if winner.MutationID != "payment-001" {
		t.Fatalf("expected mutation payment-001, got %s", winner.MutationID)
	}

	if winner.Authority != "node-c" {
		t.Fatalf("expected node-c to win, got %s", winner.Authority)
	}

	if winner.Epoch != 3 {
		t.Fatalf("expected epoch 3, got %d", winner.Epoch)
	}

	if winner.Priority != 10 {
		t.Fatalf("expected priority 10, got %d", winner.Priority)
	}
}

func TestAuthorityRaceIsOrderIndependent(t *testing.T) {
	inputsA := []AuthorityRaceInput{
		{MutationID: "payment-001", Authority: "node-a", Epoch: 2, Priority: 90},
		{MutationID: "payment-001", Authority: "node-b", Epoch: 3, Priority: 50},
		{MutationID: "payment-001", Authority: "node-c", Epoch: 3, Priority: 10},
	}

	inputsB := []AuthorityRaceInput{
		{MutationID: "payment-001", Authority: "node-c", Epoch: 3, Priority: 10},
		{MutationID: "payment-001", Authority: "node-a", Epoch: 2, Priority: 90},
		{MutationID: "payment-001", Authority: "node-b", Epoch: 3, Priority: 50},
	}

	winnerA := SelectAuthorityWinner(inputsA)
	winnerB := SelectAuthorityWinner(inputsB)

	if winnerA != winnerB {
		t.Fatalf("expected same winner under reorder, got A=%+v B=%+v", winnerA, winnerB)
	}
}