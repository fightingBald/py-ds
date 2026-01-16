package student_ds_tutorial

import (
	"testing"

	"github.com/fightingBald/py-ds/student_ds_tutorial/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sampleStudents() []common.Student {
	return []common.Student{
		{
			ID:       "s1",
			Name:     "Ada Lovelace",
			Class:    "A",
			Grade:    1,
			Enrolled: true,
			Scores:   map[string]int{"math": 95, "eng": 88},
			Tags:     []string{"prefect", "robot"},
			Profile: &common.Profile{
				Phone:   "111",
				Address: "Street 1",
				Emergency: &common.Contact{
					Name:  "Ada Mom",
					Phone: "999",
				},
			},
		},
		{
			ID:       "s2",
			Name:     "Bob",
			Class:    "A",
			Grade:    1,
			Enrolled: false,
			Scores:   map[string]int{"math": 60, "eng": 70},
			Tags:     []string{"late"},
			Profile: &common.Profile{
				Phone: "222",
			},
		},
		{
			ID:       "s3",
			Name:     "Cora",
			Class:    "B",
			Grade:    2,
			Enrolled: true,
			Scores:   map[string]int{"math": 78, "eng": 92},
			Tags:     []string{"robot", "mentor"},
			Profile:  nil,
		},
		{
			ID:       "s4",
			Name:     "Dan",
			Class:    "B",
			Grade:    2,
			Enrolled: true,
			Scores:   nil,
			Tags:     nil,
			Profile: &common.Profile{
				Phone:     "333",
				Emergency: nil,
			},
		},
	}
}

func ids(students []common.Student) []string {
	out := make([]string, 0, len(students))
	for _, s := range students {
		out = append(out, s.ID)
	}
	return out
}

func TestNormalizeName(t *testing.T) {
	assert.Equal(t, "Ada Lovelace", NormalizeName("  Ada   Lovelace  "))
	assert.Equal(t, "", NormalizeName("   "))
	assert.Equal(t, "Ada", NormalizeName("Ada"))
}

func TestMakeStudentEmail(t *testing.T) {
	email, ok := MakeStudentEmail("  Ada   Lovelace ", "Example.COM ")
	assert.True(t, ok)
	assert.Equal(t, "ada.lovelace@example.com", email)

	email, ok = MakeStudentEmail("", "example.com")
	assert.False(t, ok)
	assert.Equal(t, "", email)

	email, ok = MakeStudentEmail("Ada", "   ")
	assert.False(t, ok)
	assert.Equal(t, "", email)
}

func TestParseTagLine(t *testing.T) {
	assert.Nil(t, ParseTagLine("   "))
	assert.Equal(t, []string{"robot", "mentor", "late", "prefect"}, ParseTagLine(" robot, mentor; late  ,  prefect  "))
}

func TestBuildDisplayName(t *testing.T) {
	assert.Equal(t, "Ada (s1)", BuildDisplayName(common.Student{Name: "Ada", ID: "s1"}))
	assert.Equal(t, "Ada", BuildDisplayName(common.Student{Name: "Ada"}))
	assert.Equal(t, "s1", BuildDisplayName(common.Student{ID: "s1"}))
	assert.Equal(t, "", BuildDisplayName(common.Student{}))
}

func TestHasNamePrefix(t *testing.T) {
	s := common.Student{Name: "Ada Lovelace"}
	assert.True(t, HasNamePrefix(s, "ad"))
	assert.True(t, HasNamePrefix(s, "ADA"))
	assert.False(t, HasNamePrefix(s, "lo"))
}

func TestSplitFullName(t *testing.T) {
	first, last := SplitFullName("  Ada   Lovelace ")
	assert.Equal(t, "Ada", first)
	assert.Equal(t, "Lovelace", last)

	first, last = SplitFullName("Ada")
	assert.Equal(t, "Ada", first)
	assert.Equal(t, "", last)

	first, last = SplitFullName("  Ada Byron Lovelace ")
	assert.Equal(t, "Ada", first)
	assert.Equal(t, "Byron Lovelace", last)

	first, last = SplitFullName("   ")
	assert.Equal(t, "", first)
	assert.Equal(t, "", last)
}

func TestNormalizeStudentID(t *testing.T) {
	assert.Equal(t, "001", NormalizeStudentID("  STU-001-TMP "))
	assert.Equal(t, "abc", NormalizeStudentID("stu-abc"))
	assert.Equal(t, "123", NormalizeStudentID("  123  "))
	assert.Equal(t, "xyz", NormalizeStudentID("STU-XYZ"))
	assert.Equal(t, "777", NormalizeStudentID("777-tmp"))
}

func TestIsSchoolEmail(t *testing.T) {
	assert.True(t, IsSchoolEmail("Ada@Example.Com", " example.com "))
	assert.False(t, IsSchoolEmail("Ada@Example.Com", "other.com"))
	assert.False(t, IsSchoolEmail("@example.com", "example.com"))
	assert.False(t, IsSchoolEmail("ada@example.com", ""))
	assert.False(t, IsSchoolEmail("", "example.com"))
}

func TestNameContainsKeyword(t *testing.T) {
	s := common.Student{Name: "Ada Lovelace"}
	assert.True(t, NameContainsKeyword(s, "love"))
	assert.True(t, NameContainsKeyword(s, "ADA"))
	assert.False(t, NameContainsKeyword(s, "bob"))
	assert.False(t, NameContainsKeyword(s, "   "))
}

func TestNameToSlug(t *testing.T) {
	assert.Equal(t, "ada-lovelace", NameToSlug("  Ada   Lovelace "))
	assert.Equal(t, "ada", NameToSlug("Ada"))
	assert.Equal(t, "", NameToSlug("   "))
}

func TestSplitSubjects(t *testing.T) {
	assert.Nil(t, SplitSubjects("   "))
	assert.Equal(t, []string{"math", "eng", "bio"}, SplitSubjects(" math| eng |bio| "))
}

func TestFilterEnrolled(t *testing.T) {
	assert.Nil(t, FilterEnrolled(nil))

	students := sampleStudents()
	got := FilterEnrolled(students)
	assert.Equal(t, []string{"s1", "s3", "s4"}, ids(got))
}

func TestNames(t *testing.T) {
	assert.Nil(t, Names(nil))

	students := sampleStudents()
	assert.Equal(t, []string{"Ada Lovelace", "Bob", "Cora", "Dan"}, Names(students))
}

func TestFindByID(t *testing.T) {
	students := sampleStudents()

	got, ok := FindByID(students, "s1")
	require.True(t, ok)
	require.NotNil(t, got)

	got.Name = "AdaX"
	assert.Equal(t, "AdaX", students[0].Name)

	got, ok = FindByID(students, "missing")
	assert.False(t, ok)
	assert.Nil(t, got)
}

func TestAddStudent(t *testing.T) {
	var students []common.Student
	students = AddStudent(students, common.Student{ID: "s1"})
	assert.Equal(t, []string{"s1"}, ids(students))

	students = AddStudent(students, common.Student{ID: "s2"})
	assert.Equal(t, []string{"s1", "s2"}, ids(students))
}

func TestUpdateStudentName(t *testing.T) {
	students := sampleStudents()
	ok := UpdateStudentName(students, "s2", "  Bobby  ")
	assert.True(t, ok)
	assert.Equal(t, "Bobby", students[1].Name)

	ok = UpdateStudentName(students, "missing", "New")
	assert.False(t, ok)

	ok = UpdateStudentName(students, "s1", "   ")
	assert.False(t, ok)
}

func TestRemoveByID(t *testing.T) {
	assert.Nil(t, RemoveByID(nil, "s1"))

	students := sampleStudents()
	after := RemoveByID(students, "s2")
	assert.Equal(t, []string{"s1", "s3", "s4"}, ids(after))

	unchanged := RemoveByID(students, "missing")
	assert.Equal(t, ids(students), ids(unchanged))
}

func TestInsertAt(t *testing.T) {
	students := sampleStudents()
	atStart := InsertAt(students, 0, common.Student{ID: "s0", Name: "Zero"})
	assert.Equal(t, []string{"s0", "s1", "s2", "s3", "s4"}, ids(atStart))

	students = sampleStudents()
	atMiddle := InsertAt(students, 2, common.Student{ID: "sx", Name: "Middle"})
	assert.Equal(t, []string{"s1", "s2", "sx", "s3", "s4"}, ids(atMiddle))

	students = sampleStudents()
	atEnd := InsertAt(students, len(students), common.Student{ID: "se", Name: "End"})
	assert.Equal(t, []string{"s1", "s2", "s3", "s4", "se"}, ids(atEnd))

	students = sampleStudents()
	outOfRange := InsertAt(students, -1, common.Student{ID: "bad"})
	assert.Equal(t, ids(students), ids(outOfRange))

	outOfRange = InsertAt(students, len(students)+1, common.Student{ID: "bad"})
	assert.Equal(t, ids(students), ids(outOfRange))
}

func TestCloneStudents(t *testing.T) {
	assert.Nil(t, CloneStudents(nil))

	students := sampleStudents()
	cloned := CloneStudents(students)
	assert.Equal(t, ids(students), ids(cloned))

	if len(students) > 0 {
		assert.NotEqual(t, &students[0], &cloned[0])
	}

	cloned[0].Name = "Changed"
	assert.Equal(t, "Ada Lovelace", students[0].Name)
}

func TestDedupByID(t *testing.T) {
	students := []common.Student{
		{ID: "s1", Name: "A"},
		{ID: "s2", Name: "B"},
		{ID: "s1", Name: "A2"},
	}

	out := DedupByID(students)
	assert.Equal(t, []string{"s1", "s2"}, ids(out))
	assert.Equal(t, "A", out[0].Name)
}

func TestSortByName(t *testing.T) {
	students := []common.Student{
		{ID: "s2", Name: "bob"},
		{ID: "s1", Name: "Ada"},
		{ID: "s3", Name: "ada"},
		{ID: "s4", Name: "cora"},
	}

	sorted := SortByName(students)
	assert.Equal(t, []string{"s1", "s3", "s2", "s4"}, ids(sorted))
	assert.Equal(t, []string{"s2", "s1", "s3", "s4"}, ids(students))
}

func TestContainsID(t *testing.T) {
	students := sampleStudents()
	assert.True(t, ContainsID(students, "s1"))
	assert.False(t, ContainsID(students, "sx"))
	assert.False(t, ContainsID(nil, "s1"))
}

func TestDeleteAt(t *testing.T) {
	assert.Nil(t, DeleteAt(nil, 0))

	students := sampleStudents()
	after := DeleteAt(students, 1)
	assert.Equal(t, []string{"s1", "s3", "s4"}, ids(after))

	unchanged := DeleteAt(students, -1)
	assert.Equal(t, ids(students), ids(unchanged))

	unchanged = DeleteAt(students, len(students))
	assert.Equal(t, ids(students), ids(unchanged))
}

func TestReverseIDs(t *testing.T) {
	assert.Nil(t, ReverseIDs(nil))

	input := []string{"s1", "s2", "s3"}
	reversed := ReverseIDs(input)
	assert.Equal(t, []string{"s3", "s2", "s1"}, reversed)
	assert.Equal(t, []string{"s1", "s2", "s3"}, input)
}

func TestCompactIDs(t *testing.T) {
	assert.Nil(t, CompactIDs(nil))

	input := []string{"a", "a", "b", "b", "b", "c"}
	out := CompactIDs(input)
	assert.Equal(t, []string{"a", "b", "c"}, out)
	assert.Equal(t, []string{"a", "a", "b", "b", "b", "c"}, input)
}

func TestIndexByID(t *testing.T) {
	index := IndexByID(nil)
	require.NotNil(t, index)
	assert.Len(t, index, 0)

	students := sampleStudents()
	index = IndexByID(students)
	assert.Equal(t, "Ada Lovelace", index["s1"].Name)
	assert.Len(t, index, 4)

	index["s1"].Name = "AdaX"
	assert.Equal(t, "AdaX", students[0].Name)
}

func TestGetFromIndex(t *testing.T) {
	got, ok := GetFromIndex(nil, "s1")
	assert.False(t, ok)
	assert.Nil(t, got)

	students := sampleStudents()
	index := IndexByID(students)
	got, ok = GetFromIndex(index, "s1")
	assert.True(t, ok)
	assert.Equal(t, "Ada Lovelace", got.Name)
}

func TestUpsertIndex(t *testing.T) {
	var index map[string]*common.Student
	student := &common.Student{ID: "s1", Name: "Ada"}

	index, ok := UpsertIndex(index, student)
	assert.True(t, ok)
	require.NotNil(t, index)
	assert.Equal(t, "Ada", index["s1"].Name)

	index, ok = UpsertIndex(index, nil)
	assert.False(t, ok)

	index, ok = UpsertIndex(index, &common.Student{})
	assert.False(t, ok)
}

func TestGroupByClass(t *testing.T) {
	group := GroupByClass(nil)
	require.NotNil(t, group)
	assert.Len(t, group, 0)

	students := sampleStudents()
	group = GroupByClass(students)
	assert.Equal(t, []string{"s1", "s2"}, ids(group["A"]))
	assert.Equal(t, []string{"s3", "s4"}, ids(group["B"]))
}

func TestCountByGrade(t *testing.T) {
	counts := CountByGrade(nil)
	require.NotNil(t, counts)
	assert.Len(t, counts, 0)

	students := sampleStudents()
	counts = CountByGrade(students)
	assert.Equal(t, 2, counts[1])
	assert.Equal(t, 2, counts[2])
}

func TestMergeScoreTotals(t *testing.T) {
	dst := map[string]int{"math": 10}
	src := map[string]int{"math": 5, "eng": 7}

	out := MergeScoreTotals(dst, src)
	assert.Equal(t, map[string]int{"math": 15, "eng": 7}, out)

	out = MergeScoreTotals(nil, src)
	assert.Equal(t, map[string]int{"math": 5, "eng": 7}, out)

	out = MergeScoreTotals(dst, nil)
	assert.Equal(t, map[string]int{"math": 15, "eng": 7}, out)

	out = MergeScoreTotals(nil, nil)
	require.NotNil(t, out)
	assert.Len(t, out, 0)
}

func TestDeleteFromIndex(t *testing.T) {
	assert.False(t, DeleteFromIndex(nil, "s1"))

	index := IndexByID(sampleStudents())
	assert.True(t, DeleteFromIndex(index, "s1"))
	assert.False(t, DeleteFromIndex(index, "s1"))
}

func TestScoreOf(t *testing.T) {
	students := sampleStudents()

	score, ok := ScoreOf(students[0], "math")
	assert.True(t, ok)
	assert.Equal(t, 95, score)

	score, ok = ScoreOf(students[3], "math")
	assert.False(t, ok)
	assert.Equal(t, 0, score)

	score, ok = ScoreOf(students[0], "science")
	assert.False(t, ok)
	assert.Equal(t, 0, score)
}

func TestIDsFromIndex(t *testing.T) {
	assert.Nil(t, IDsFromIndex(nil))

	index := IndexByID(sampleStudents())
	assert.Equal(t, []string{"s1", "s2", "s3", "s4"}, IDsFromIndex(index))
}

func TestCloneScoreMap(t *testing.T) {
	assert.Nil(t, CloneScoreMap(nil))

	src := map[string]int{"math": 95}
	clone := CloneScoreMap(src)
	assert.Equal(t, src, clone)

	clone["math"] = 100
	assert.Equal(t, 95, src["math"])
}

func TestCopyScoreMap(t *testing.T) {
	src := map[string]int{"math": 90, "eng": 70}

	out := CopyScoreMap(nil, src)
	assert.Equal(t, map[string]int{"math": 90, "eng": 70}, out)

	dst := map[string]int{"math": 80}
	out = CopyScoreMap(dst, src)
	assert.Equal(t, map[string]int{"math": 90, "eng": 70}, out)
	assert.Equal(t, map[string]int{"math": 90, "eng": 70}, dst)

	out = CopyScoreMap(dst, nil)
	assert.Equal(t, map[string]int{"math": 90, "eng": 70}, out)
}

func TestDeleteLowScores(t *testing.T) {
	assert.Nil(t, DeleteLowScores(nil, 60))

	scores := map[string]int{"math": 50, "eng": 80, "bio": 60}
	out := DeleteLowScores(scores, 60)
	assert.Equal(t, map[string]int{"eng": 80, "bio": 60}, out)
}

func TestScoreSubjectsSorted(t *testing.T) {
	assert.Nil(t, ScoreSubjectsSorted(nil))

	scores := map[string]int{"eng": 80, "math": 90}
	assert.Equal(t, []string{"eng", "math"}, ScoreSubjectsSorted(scores))
}

func TestNewStudent(t *testing.T) {
	student := NewStudent("s9", "Eva")
	require.NotNil(t, student)
	assert.Equal(t, "s9", student.ID)
	assert.Equal(t, "Eva", student.Name)
	assert.False(t, student.Enrolled)
	assert.Nil(t, student.Scores)
	assert.Nil(t, student.Tags)
	assert.Nil(t, student.Profile)
}

func TestEnsureProfile(t *testing.T) {
	var nilStudent *common.Student
	assert.Nil(t, EnsureProfile(nilStudent))

	student := &common.Student{ID: "s1"}
	profile := EnsureProfile(student)
	require.NotNil(t, profile)
	require.NotNil(t, student.Profile)

	again := EnsureProfile(student)
	assert.Same(t, profile, again)
}

func TestUpdatePhone(t *testing.T) {
	var nilStudent *common.Student
	assert.False(t, UpdatePhone(nilStudent, "123"))

	student := &common.Student{ID: "s1"}
	assert.False(t, UpdatePhone(student, ""))
	assert.True(t, UpdatePhone(student, "123"))
	require.NotNil(t, student.Profile)
	assert.Equal(t, "123", student.Profile.Phone)
}

func TestAddScore(t *testing.T) {
	var nilStudent *common.Student
	assert.False(t, AddScore(nilStudent, "math", 90))

	student := &common.Student{ID: "s1"}
	assert.False(t, AddScore(student, "", 90))
	assert.True(t, AddScore(student, "math", 90))
	assert.Equal(t, 90, student.Scores["math"])

	assert.True(t, AddScore(student, "math", 95))
	assert.Equal(t, 95, student.Scores["math"])
}

func TestEmergencyName(t *testing.T) {
	assert.Equal(t, "", EmergencyName(nil))

	student := &common.Student{}
	assert.Equal(t, "", EmergencyName(student))

	student.Profile = &common.Profile{}
	assert.Equal(t, "", EmergencyName(student))

	student.Profile.Emergency = &common.Contact{Name: "Guardian"}
	assert.Equal(t, "Guardian", EmergencyName(student))
}

func TestParseGrade(t *testing.T) {
	grade, ok := ParseGrade(" 2 ")
	assert.True(t, ok)
	assert.Equal(t, 2, grade)

	grade, ok = ParseGrade("bad")
	assert.False(t, ok)
	assert.Equal(t, 0, grade)
}

func TestFormatGrade(t *testing.T) {
	assert.Equal(t, "3", FormatGrade(3))
}

func TestScoresToStringMap(t *testing.T) {
	assert.Nil(t, ScoresToStringMap(nil))

	out := ScoresToStringMap(map[string]int{"math": 95})
	assert.Equal(t, map[string]string{"math": "95"}, out)
}

func TestParseScores(t *testing.T) {
	assert.Nil(t, ParseScores(nil))

	out := ParseScores(map[string]string{"math": "95", "eng": " bad ", "bio": " 80 "})
	assert.Equal(t, map[string]int{"math": 95, "bio": 80}, out)
}

func TestIDsToCSV(t *testing.T) {
	assert.Equal(t, "", IDsToCSV(nil))
	assert.Equal(t, "", IDsToCSV([]string{}))
	assert.Equal(t, "s1,s2", IDsToCSV([]string{"s1", "s2"}))
}

func TestCSVToIDs(t *testing.T) {
	assert.Nil(t, CSVToIDs("   "))
	assert.Equal(t, []string{"s1", "s2", "s3"}, CSVToIDs(" s1, s2 ,, s3 "))
}

func TestTagSet(t *testing.T) {
	assert.Nil(t, TagSet(nil))

	set := TagSet([]string{"a", "b", "a"})
	assert.Equal(t, map[string]struct{}{"a": {}, "b": {}}, set)
}

func TestSetUnionIntersectDiff(t *testing.T) {
	a := map[string]struct{}{"a": {}, "b": {}}
	b := map[string]struct{}{"b": {}, "c": {}}

	assert.Nil(t, SetUnion(nil, nil))

	union := SetUnion(a, b)
	assert.Equal(t, map[string]struct{}{"a": {}, "b": {}, "c": {}}, union)

	inter := SetIntersect(a, b)
	assert.Equal(t, map[string]struct{}{"b": {}}, inter)

	diff := SetDiff(a, b)
	assert.Equal(t, map[string]struct{}{"a": {}}, diff)
}

func TestSetToSortedSlice(t *testing.T) {
	assert.Nil(t, SetToSortedSlice(nil))

	set := map[string]struct{}{"b": {}, "a": {}}
	assert.Equal(t, []string{"a", "b"}, SetToSortedSlice(set))
}

func TestCommonTags(t *testing.T) {
	a := common.Student{Tags: []string{"robot", "mentor", "late"}}
	b := common.Student{Tags: []string{"mentor", "prefect"}}
	assert.Equal(t, []string{"mentor"}, CommonTags(a, b))

	c := common.Student{Tags: nil}
	assert.Nil(t, CommonTags(a, c))
}

func TestUniqueTags(t *testing.T) {
	assert.Nil(t, UniqueTags(nil))

	tags := UniqueTags(sampleStudents())
	assert.Equal(t, []string{"late", "mentor", "prefect", "robot"}, tags)
}

func TestUniqueClasses(t *testing.T) {
	assert.Nil(t, UniqueClasses(nil))
	assert.Equal(t, []string{"A", "B"}, UniqueClasses(sampleStudents()))
}

func TestSortByScore(t *testing.T) {
	assert.Nil(t, SortByScore(nil, "math"))

	students := []common.Student{
		{ID: "s2", Scores: map[string]int{"math": 90}},
		{ID: "s1", Scores: map[string]int{"math": 90}},
		{ID: "s3", Scores: map[string]int{"math": 80}},
		{ID: "s4", Scores: nil},
	}

	sorted := SortByScore(students, "math")
	assert.Equal(t, []string{"s1", "s2", "s3", "s4"}, ids(sorted))
	assert.Equal(t, []string{"s2", "s1", "s3", "s4"}, ids(students))
}

func TestSortByGradeThenName(t *testing.T) {
	students := []common.Student{
		{ID: "s2", Name: "bob", Grade: 2},
		{ID: "s1", Name: "Ada", Grade: 2},
		{ID: "s3", Name: "ada", Grade: 2},
		{ID: "s4", Name: "cora", Grade: 1},
	}

	sorted := SortByGradeThenName(students)
	assert.Equal(t, []string{"s1", "s3", "s2", "s4"}, ids(sorted))
	assert.Equal(t, []string{"s2", "s1", "s3", "s4"}, ids(students))
}

func TestStableSortByClass(t *testing.T) {
	students := []common.Student{
		{ID: "s1", Class: "B"},
		{ID: "s2", Class: "A"},
		{ID: "s3", Class: "B"},
		{ID: "s4", Class: "A"},
	}

	sorted := StableSortByClass(students)
	assert.Equal(t, []string{"s2", "s4", "s1", "s3"}, ids(sorted))
	assert.Equal(t, []string{"s1", "s2", "s3", "s4"}, ids(students))
}

func TestBinarySearchByID(t *testing.T) {
	idx, ok := BinarySearchByID(nil, "s1")
	assert.False(t, ok)
	assert.Equal(t, 0, idx)

	students := []common.Student{{ID: "s1"}, {ID: "s2"}, {ID: "s4"}}
	idx, ok = BinarySearchByID(students, "s2")
	assert.True(t, ok)
	assert.Equal(t, 1, idx)

	idx, ok = BinarySearchByID(students, "s3")
	assert.False(t, ok)
	assert.Equal(t, 2, idx)
}

func TestTopStudentBySubject(t *testing.T) {
	best, ok := TopStudentBySubject(nil, "math")
	assert.False(t, ok)
	assert.Nil(t, best)

	students := sampleStudents()
	best, ok = TopStudentBySubject(students, "math")
	require.True(t, ok)
	require.NotNil(t, best)
	assert.Equal(t, "s1", best.ID)

	best.Name = "AdaX"
	assert.Equal(t, "AdaX", students[0].Name)

	best, ok = TopStudentBySubject(students, "science")
	assert.False(t, ok)
	assert.Nil(t, best)
}
