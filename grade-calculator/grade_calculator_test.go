package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetRemainingLetterGrades(t *testing.T) {
	cases := []struct {
		score    int
		expected string
	}{
		{70, "C"},
		{60, "D"},
		{59, "F"},
	}

	for _, tc := range cases {
		t.Run(tc.expected, func(t *testing.T) {
			calculator := NewGradeCalculator()
			calculator.AddGrade("assignment", tc.score, Assignment)
			calculator.AddGrade("exam", tc.score, Exam)
			calculator.AddGrade("essay", tc.score, Essay)

			actual := calculator.GetFinalGrade()
			if actual != tc.expected {
				t.Errorf("Expected %s; got %s", tc.expected, actual)
			}
		})
	}
}

func TestGradeTypeString(t *testing.T) {
	cases := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, tc := range cases {
		actual := tc.gradeType.String()
		if actual != tc.expected {
			t.Errorf("Expected %s; got %s", tc.expected, actual)
		}
	}
}
