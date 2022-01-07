package utils

// MockData represents the date used when mocking functionality with gomock.
type MockData struct {
	Calls      int
	Args       []interface{}
	SetObjs    map[int]interface{}
	ReturnObjs []interface{}
}
