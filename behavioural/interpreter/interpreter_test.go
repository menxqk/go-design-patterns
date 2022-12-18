package interpreter

import "testing"

func TestCalculate(t *testing.T) {
	tempOperation := "3 4 sum 2 sub"
	correctResult := 5
	res, err := Calculate(tempOperation)

	if err != nil {
		t.Error(err)
	}

	if res != correctResult {
		t.Errorf("expected result = %d, got %d", correctResult, res)
	}
}
