package bd

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase[W any, I any] struct {
	want  W
	input I
}

func TestCheckPalindrome(t *testing.T) {
	t.Run("check bd ok", func(t *testing.T) {
		err := CheckConnection()
		assert.Equal(t, 1, err)
	})

	t.Run("force disconnect ok", func(t *testing.T) {
		err := ConnectDB().Disconnect(context.TODO())
		assert.Equal(t, nil, err)
	})

	t.Run("Test palindrome fro multiple in testtable", func(t *testing.T) {
		var testsCases = []struct {
			input string
			want  bool
		}{
			{"", false},
			{"a", true},
			{"aa", true},
			{"ab", false},
			{"kayak", true},
			{"detartrated", true},
			{"AVA", true},
			{"dsaasd", true},
		}

		for _, test := range testsCases {
			if got := CheckPalindrome(test.input); got != test.want {
				t.Errorf("CheckPalindrome(%q) = %v, want %v", test.input, got, test.want)
			}
		}
	})

	t.Run("should fail for empty string", func(t *testing.T) {
		testCase := TestCase[bool, string]{
			input: "",
			want:  false,
		}
		if got := CheckPalindrome(testCase.input); got != bool(testCase.want) {
			t.Errorf("CheckPalindrome(%q) = %v, want %v", testCase.input, got, testCase.want)
		}
	})

	t.Run("dsaasd is palindrome OK", func(t *testing.T) {
		testCase := TestCase[bool, string]{
			input: "dsaasd",
			want:  true,
		}
		if got := CheckPalindrome(testCase.input); got != bool(testCase.want) {
			t.Errorf("CheckPalindrome(%q) = %v, want %v", testCase.input, got, testCase.want)
		}
	})

	t.Run("test service OK search string", func(t *testing.T) {
		testCase := TestCase[bool, string]{
			input: "dsaasd",
			want:  true,
		}
		res, err := SearchProductsByDescriptionBrand(testCase.input, "s")

		if err != nil {
			t.Errorf("SearchProductsByDescriptionBrand(%s) = %v, want %v", testCase.input, err, nil)
		}

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("test service OK id", func(t *testing.T) {
		testCase := TestCase[bool, string]{
			input: "weñxoab",
			want:  true,
		}
		res, err := SearchProductsByDescriptionBrand(testCase.input, "id")

		if err != nil {
			t.Errorf("SearchProductsByDescriptionBrand(%s) = %v, want %v", testCase.input, err, nil)
		}

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}
