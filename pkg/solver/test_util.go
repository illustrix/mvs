package solver

import "testing"

type BoardCheckCase struct {
	Pos      Vec2
	Expected bool
}

func parseTestCases(s string) []*BoardCheckCase {
	lines := trimBoardString(s)
	cases := make([]*BoardCheckCase, 0)
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '1':
				cases = append(cases, &BoardCheckCase{Vec2{x, y}, true})
			case '0':
				cases = append(cases, &BoardCheckCase{Vec2{x, y}, false})
			}
		}
	}
	return cases
}

func testBoardCheck(t *testing.T, rule Rule, b Board, c []*BoardCheckCase) {
	for _, tc := range c {
		if rule.Check(b, tc.Pos) != tc.Expected {
			t.Errorf("Check failed for (%d, %d) expected: %v actual: %v", tc.Pos[0], tc.Pos[1], tc.Expected, !tc.Expected)
			t.FailNow()
		}
	}
}
