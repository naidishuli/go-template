package tests

type MockMatcher struct {
	matchFun func(any) bool
}

func (u MockMatcher) Matches(a any) bool {
	return u.matchFun(a)
}

func (u MockMatcher) String() string {
	return ""
}

func MockMatch(matchFun func(a any) bool) *MockMatcher {
	return &MockMatcher{matchFun}
}
