# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

It averages assignment, exam, and essay grades and applies weights of
50%, 35%, and 15%, respectively, to determine a final letter grade.

Runtime requirements:
Golang 1.24.4+

Setup:
```powershell
git clone https://github.com/riveraJoshuaUF/esep-grade-calculator.git
cd esep-grade-calculator
```

Usage:
Create a calculator with `NewGradeCalculator()`, add grades with
`AddGrade(name, grade, gradeType)`, and call `GetFinalGrade()`.
Grade types are `Assignment`, `Exam`, and `Essay`.

Run unit tests from the repository root:
```powershell
cd grade-calculator/
go test ./...
```

Run unit tests with code coverage from the same folder:
```powershell
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
```

The starter code currently fails the A and B tests.