package math

func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func GCDReminderRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDReminderRecursive(b, a % b)
}

func GCDReminder(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}

	return a
}
