package common

// Student 表示内存中的学生信息。
type Student struct {
	ID       string
	Name     string
	Class    string
	Grade    int
	Enrolled bool
	Scores   map[string]int
	Tags     []string
	Profile  *Profile
}

type Profile struct {
	Phone     string
	Address   string
	Emergency *Contact
}

type Contact struct {
	Name  string
	Phone string
}
