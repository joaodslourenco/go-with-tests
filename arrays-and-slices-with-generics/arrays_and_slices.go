package arraysandslices

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }

	return Reduce(numbers, add, 0)
}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}

	}

	return Reduce(numbers, sumTail, []int{})
}
