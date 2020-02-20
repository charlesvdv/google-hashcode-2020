package main

import (
	"sort"
)

type BookSorting struct {
	Scores []int
	Books  []int
}

func (s BookSorting) Len() int {
	return len(s.Books)
}

func (s BookSorting) Swap(i, j int) {
	s.Books[i], s.Books[j] = s.Books[j], s.Books[i]
}

func (s BookSorting) Less(i, j int) bool {
	return s.Scores[i] > s.Scores[j]
}

func orderBooksInLibrary(library Library, scores []int) {
	sorting := BookSorting{Scores: scores, Books: library.BookIDs}
	sort.Sort(sorting)
	library.BookIDs = sorting.Books
}

type LibrarySorting struct {
	libraries []Library
}

func (s LibrarySorting) Len() int {
	return len(s.libraries)
}

func (s LibrarySorting) Swap(i, j int) {
	s.libraries[i], s.libraries[j] = s.libraries[j], s.libraries[i]
}

func (s LibrarySorting) Less(i, j int) bool {
	return s.libraries[i].SignupTime < s.libraries[j].SignupTime
}
