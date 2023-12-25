package kirmath

import "testing"

func TestAdd(t *testing.T) {
	// @COOL instead of commenting out the whole test
	t.SkipNow()
	result := Add(1, 2)
	if result != 3 {
		t.Error("Adding 1 and 2 doesn't produce 3")
	} else {
		t.Log("Result is:", result)
	}
}

func TestAddWithTables(t *testing.T) {
	// @COOL instead of commenting out the whole test
	t.SkipNow()
	testCases := []struct {
		a      int
		b      int
		result int
	}{
		{1, 2, 3},
		{12, 30, 42},
		{100, -1, 99},
	}
	for _, testCase := range testCases {
		result := Add(testCase.a, testCase.b)
		if result != testCase.result {
			t.Errorf("Adding %d and %d doesn't produce %d, instead it produces %d",
				testCase.a, testCase.b, testCase.result, result)

		}

	}
}

func TestAddWithSubTest(t *testing.T) {
	// @COOL instead of commenting out the whole test
	t.SkipNow()
	testCases := []struct {
		name   string
		a      int
		b      int
		result int
	}{
		{"OneDigit", 1, 2, 3},
		{"TwoDigits", 12, 30, 42},
		{"ThreeDigits", 100, -1, 99},
	}
	for _, testCase := range testCases {
		// @COOL
		t.Run(testCase.name, func(t *testing.T) {
			result := Add(testCase.a, testCase.b)
			if result != testCase.result {
				t.Errorf("Adding %d and %d doesn't produce %d, instead it produces %d",
					testCase.a, testCase.b, testCase.result,
					result)

			}
		})
	}
}
