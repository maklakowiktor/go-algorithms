package datastructures

import (
	"reflect"
	"testing"
)

func TestNewTrie(t *testing.T) {
	tests := []struct {
		name string
		want *Trie
	}{
		// TODO: Add test cases.
		{
			name: "Test empty NewTrie",
			want: &Trie{
				Root: NewTrieNode(),
			},
		},
		{
			name: "Word: Alla",
			want: &Trie{
				Root: &TrieNode{
					Children: map[rune]*TrieNode{
						'A': {
							Children: map[rune]*TrieNode{
								'l': {
									Children: map[rune]*TrieNode{
										'l': {
											Children: map[rune]*TrieNode{
												'a': {
													Children: map[rune]*TrieNode{
														'*': nil,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTrie("Alla"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTrieNode(t *testing.T) {
	tests := []struct {
		name string
		want *TrieNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTrieNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrieNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Insert(t1 *testing.T) {
	type fields struct {
		Root *TrieNode
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				Root: tt.fields.Root,
			}
			t.Insert(tt.args.word)
		})
	}
}

func TestTrie_Search(t1 *testing.T) {
	type fields struct {
		Root *TrieNode
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				Root: tt.fields.Root,
			}
			if got := t.Search(tt.args.word); got != tt.want {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
