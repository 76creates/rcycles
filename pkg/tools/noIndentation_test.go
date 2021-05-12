package tools

import "testing"

func TestNoIndent(t *testing.T) {
	test1 := `
		Testing
	Test1
		Test test 1

		Test2
	`
	expecterResult1 := `	Testing
Test1
	Test test 1

	Test2
`
	result1 := DeIndent(test1)
	if expecterResult1 != result1 {
		t.Log("testing didnt yield expected result")
		t.Errorf("\nExpected:\n%+v\n\nActual:\n%+v", expecterResult1, result1)
	}

	test2 := `Testing
	Test1
			Test test 1
	`
	expecterResult2 := `Testing
Test1
		Test test 1
`
	result2 := DeIndent(test2)
	if expecterResult2 != result2 {
		t.Log("testing didnt yield expected result")
		t.Errorf("\nExpected:\n%+v\n\nActual:\n%+v", expecterResult2, result2)
	}

	// first line indented with space, second with tabs
	test3 := `
  Test1
		Test test 1
	`
	expecterResult3 := test3
	result3 := DeIndent(test3)
	if expecterResult3 != result3 {
		t.Log("expecter deindent to error out and produce output same as input")
		t.Errorf("\nExpected:\n%+v\n\nActual:\n%+v", expecterResult3, result3)
	}

	test4 := `Test1
Test test 1
`
	expecterResult4 := test4
	result4 := DeIndent(test4)
	if expecterResult4 != result4 {
		t.Log("expecter deindent to error out and produce output same as input")
		t.Errorf("\nExpected:\n%+v\n\nActual:\n%+v", expecterResult4, result4)
	}
}
