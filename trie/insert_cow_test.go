package trie

import "testing"

func Test_InsertCOW(t *testing.T) {
	t.Parallel()

	t.Run("new node must be added", func(t *testing.T) {
		trie0 := NewTrie()
		trie1 := trie0.InsertCOW("a", 1)

		node1, found1 := find(trie1, "a")

		if found1 != true {
			t.Errorf("Key \"%s\" is expected to be \"%v\", got \"%v\"", "a", true, found1)
		}

		if found1 && node1.Value != 1 {
			t.Errorf("Key \"%s\" is expected to have value \"%v\", got \"%v\"", "a", 1, node1.Value)
		}

		node0, found0 := find(trie0, "a")

		if found0 != false || node0 != nil {
			t.Errorf("Key \"%s\" is expected to be \"%v\", got \"%v\"", "a", false, found0)
		}
	})

	t.Run("new node must be added and retain old nodes", func(t *testing.T) {
		trie0 := NewTrie()
		trie0.Insert("aa", 1)
		trie0.Insert("ab", 2)
		trie0.Insert("bc", 3)

		trie1 := trie0.InsertCOW("ba", 4)

		testSteps := []struct {
			name         string
			tree         *Trie
			key          string
			expectedVal  int
			expectedTerm bool
		}{
			{
				name:         "find in trie0 and must not find \"a\"",
				tree:         trie0,
				key:          "a",
				expectedVal:  0,
				expectedTerm: false,
			},
			{
				name:         "find in trie0 and must find key \"aa\"",
				tree:         trie0,
				key:          "aa",
				expectedVal:  1,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must not find key \"a\"",
				tree:         trie1,
				key:          "a",
				expectedVal:  0,
				expectedTerm: false,
			},
			{
				name:         "find in trie1 and must find key \"aa\"",
				tree:         trie1,
				key:          "aa",
				expectedVal:  1,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must find key \"ab\"",
				tree:         trie1,
				key:          "ab",
				expectedVal:  2,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must find \"bc\"",
				tree:         trie1,
				key:          "bc",
				expectedVal:  3,
				expectedTerm: true,
			},
			{
				name:         "find in trie1 and must find \"ba\"",
				tree:         trie1,
				key:          "ba",
				expectedVal:  4,
				expectedTerm: true,
			},
			{
				name:         "find in trie0 and must not find \"ba\"",
				tree:         trie0,
				key:          "ba",
				expectedVal:  0,
				expectedTerm: false,
			},
		}

		for _, steps := range testSteps {
			node, found := find(steps.tree, steps.key)
			if found != steps.expectedTerm {
				t.Errorf("Step '%s' expected to be %v, got %v", steps.name, steps.expectedTerm, found)
			}

			if found && node.Value != steps.expectedVal {
				t.Errorf("Step '%s' expected to be %v, got %v", steps.name, steps.expectedVal, node.Value)
			}
		}
	})
}
