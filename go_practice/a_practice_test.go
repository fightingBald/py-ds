package go_practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDSPlayground(t *testing.T) {
	nums := []int{1, 2, 3}
	users := []UserProfile{
		{ID: 1, Name: "Alice", Score: 90},
	}
	tags := []string{"go", "go", "ds"}

	ds := NewDSPlayground(nums, users, tags)
	require.NotNil(t, ds)
	assert.Equal(t, nums, ds.Numbers)
	assert.Equal(t, users, ds.Users)
	assert.Len(t, ds.Tags, 2)
	_, hasGo := ds.Tags["go"]
	_, hasDS := ds.Tags["ds"]
	assert.True(t, hasGo)
	assert.True(t, hasDS)
}

func TestAppendUniqueNumber(t *testing.T) {
	ds := &DSPlayground{Numbers: []int{1, 2}}

	AppendUniqueNumber(ds, 3)
	assert.Equal(t, []int{1, 2, 3}, ds.Numbers)

	AppendUniqueNumber(ds, 2)
	assert.Equal(t, []int{1, 2, 3}, ds.Numbers)
}

func TestRemoveNumberAt(t *testing.T) {
	ds := &DSPlayground{Numbers: []int{10, 20, 30}}

	RemoveNumberAt(ds, 1)
	assert.Equal(t, []int{10, 30}, ds.Numbers)

	RemoveNumberAt(ds, -1)
	assert.Equal(t, []int{10, 30}, ds.Numbers)

	RemoveNumberAt(ds, 5)
	assert.Equal(t, []int{10, 30}, ds.Numbers)
}

func TestFilterNumbers(t *testing.T) {
	ds := &DSPlayground{Numbers: []int{-5, 0, 5, 10, 15}}

	filtered := FilterNumbers(ds, 0, 10)
	assert.Equal(t, []int{0, 5, 10}, filtered)
	assert.Equal(t, []int{-5, 0, 5, 10, 15}, ds.Numbers)
}

func TestMapNumbersToString(t *testing.T) {
	ds := DSPlayground{Numbers: []int{1, 2, 3}}

	got := MapNumbersToString(ds)
	assert.Equal(t, []string{"#1", "#2", "#3"}, got)
}

func TestUpsertUserScore(t *testing.T) {
	ds := &DSPlayground{
		Users: []UserProfile{
			{ID: 1, Name: "Alice", Score: 90},
		},
	}

	UpsertUserScore(ds, 1, "Alice", 95)
	require.Len(t, ds.Users, 1)
	assert.Equal(t, UserProfile{ID: 1, Name: "Alice", Score: 95}, ds.Users[0])

	UpsertUserScore(ds, 2, "Bob", 80)
	require.Len(t, ds.Users, 2)
	assert.Equal(t, UserProfile{ID: 2, Name: "Bob", Score: 80}, ds.Users[1])
}

func TestSortUsersByScoreThenID(t *testing.T) {
	ds := &DSPlayground{
		Users: []UserProfile{
			{ID: 3, Name: "Carol", Score: 95},
			{ID: 2, Name: "Bob", Score: 95},
			{ID: 5, Name: "Eve", Score: 80},
			{ID: 4, Name: "Dan", Score: 95},
			{ID: 1, Name: "Alice", Score: 80},
		},
	}

	SortUsersByScoreThenID(ds)
	assert.Equal(t, []UserProfile{
		{ID: 2, Name: "Bob", Score: 95},
		{ID: 3, Name: "Carol", Score: 95},
		{ID: 4, Name: "Dan", Score: 95},
		{ID: 1, Name: "Alice", Score: 80},
		{ID: 5, Name: "Eve", Score: 80},
	}, ds.Users)
}

func TestDeleteUserByNamePrefix(t *testing.T) {
	ds := &DSPlayground{
		Users: []UserProfile{
			{ID: 1, Name: "Alice", Score: 90},
			{ID: 2, Name: "Bob", Score: 85},
			{ID: 3, Name: "Alfred", Score: 88},
			{ID: 4, Name: "Ann", Score: 92},
			{ID: 5, Name: "Eve", Score: 70},
		},
	}

	DeleteUserByNamePrefix(ds, "Al")
	assert.Equal(t, []UserProfile{
		{ID: 2, Name: "Bob", Score: 85},
		{ID: 4, Name: "Ann", Score: 92},
		{ID: 5, Name: "Eve", Score: 70},
	}, ds.Users)
}

func TestFilterUsersByScore(t *testing.T) {
	ds := DSPlayground{
		Users: []UserProfile{
			{ID: 1, Name: "Alice", Score: 40},
			{ID: 2, Name: "Bob", Score: 70},
			{ID: 3, Name: "Carol", Score: 85},
			{ID: 4, Name: "Dave", Score: 95},
		},
	}

	got := FilterUsersByScore(ds, 70, 90)
	assert.Equal(t, []UserProfile{
		{ID: 2, Name: "Bob", Score: 70},
		{ID: 3, Name: "Carol", Score: 85},
	}, got)
	assert.Equal(t, 4, len(ds.Users))
}

func TestCollectUserNames(t *testing.T) {
	ds := DSPlayground{
		Users: []UserProfile{
			{ID: 1, Name: "alice", Score: 40},
			{ID: 2, Name: "bob", Score: 70},
			{ID: 3, Name: "Alice", Score: 85},
			{ID: 4, Name: "bob", Score: 95},
		},
	}

	got := CollectUserNames(&ds)
	assert.Equal(t, []string{"Alice", "alice", "bob"}, got)
}

func TestMergeTags(t *testing.T) {
	ds := &DSPlayground{
		Tags: map[string]struct{}{"go": {}, "ds": {}},
	}

	added := MergeTags(ds, []string{"go", "lab", "concurrency"})
	assert.Equal(t, 2, added)
	assert.Len(t, ds.Tags, 4)
}

func TestIndexUsersByScore(t *testing.T) {
	ds := DSPlayground{
		Users: []UserProfile{
			{ID: 1, Name: "Alice", Score: 100},
			{ID: 2, Name: "Bob", Score: 90},
			{ID: 3, Name: "Carol", Score: 100},
		},
	}

	got := IndexUsersByScore(ds)
	require.Len(t, got, 2)
	assert.ElementsMatch(t, []int{1, 3}, got[100])
	assert.ElementsMatch(t, []int{2}, got[90])
}

func TestSplitTags(t *testing.T) {
	ds := DSPlayground{
		Tags: map[string]struct{}{
			"odd":   {},
			"even":  {},
			"xx":    {},
			"y":     {},
			"three": {},
		},
	}

	even, odd := SplitTags(ds)
	assert.ElementsMatch(t, []string{"even", "xx"}, even)
	assert.ElementsMatch(t, []string{"odd", "y", "three"}, odd)
}
