package arraysandslices

import "testing"

func assertCorrectMessage(t testing.TB, got int, want int, array []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, array)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)

	})
}
