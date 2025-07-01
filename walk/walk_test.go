package walk

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with non-string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name:          "nested fields",
			Input:         Person{"Chris", Profile{33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name:          "pointers to things",
			Input:         &Person{"Chris", Profile{33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{33, "London"},
				{34, "Dhaka"},
			},
			ExpectedCalls: []string{"London", "Dhaka"},
		},
		{
			Name: "arrays",
			Input: []Profile{
				{33, "London"},
				{34, "Dhaka"},
			},
			ExpectedCalls: []string{"London", "Dhaka"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
			}
		})
	}
	t.Run("with map", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Tokyo"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Tokyo"}

		walk(aChannel, func(s string) {
			got = append(got, s)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Paris"}
		}

		var got []string
		want := []string{"Berlin", "Paris"}

		walk(aFunction, func(s string) {
			got = append(got, s)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	if !slices.Contains(haystack, needle) {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
