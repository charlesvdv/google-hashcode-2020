package main

import (
	"bufio"
	"os"
	"sort"
)

type Library struct {
	ID            int
	BookCount     int
	SignupTime    int
	MaxBookPerDay int
	BookIDs       []int
}

type LibraryOutput struct {
	ID           int
	BooksScanned []int
}

type Process struct {
	Days                  int
	BookScores            []int
	NonProcessedLibraries []Library
	PendingSignupLib      *Library
	SignedUpLibrary       []Library
	BookProcessed         []bool
	SignUpLibrary         []int
	BooksScannedByLib     map[int][]int
}

func NewProcess(days int, booksScores []int, libraries []Library) Process {
	booksScannedByLib := map[int][]int{}

	libSorting := LibrarySorting{
		libraries: libraries,
	}
	sort.Sort(libSorting)
	libraries = libSorting.libraries
	for _, lib := range libraries {
		orderBooksInLibrary(lib, booksScores)
		booksScannedByLib[lib.ID] = []int{}
	}

	// TODO: book processed
	return Process{
		Days:                  days,
		BookScores:            booksScores,
		NonProcessedLibraries: libraries,
		PendingSignupLib:      nil,
		SignedUpLibrary:       []Library{},
		BookProcessed:         make([]bool, len(booksScores)),
		SignUpLibrary:         []int{},
		BooksScannedByLib:     booksScannedByLib,
	}
}

func (proc *Process) Calculate() ([]int, map[int][]int) {
	for day := 0; day < proc.Days; day++ {
		// fmt.Printf("%v %v %v\n", day, len(proc.NonProcessedLibraries), (proc.SignedUpLibrary))
		proc.keepLibrarySignupGoing()

		for _, library := range proc.SignedUpLibrary {
			proc.registerBook(library)
		}
	}
	return proc.SignUpLibrary, proc.BooksScannedByLib
}

func (proc *Process) keepLibrarySignupGoing() {
	if proc.PendingSignupLib != nil {
		if proc.PendingSignupLib.SignupTime == 0 {
			proc.SignUpLibrary = append(proc.SignUpLibrary, proc.PendingSignupLib.ID)
			proc.SignedUpLibrary = append(proc.SignedUpLibrary, *proc.PendingSignupLib)
			proc.PendingSignupLib = nil
		}
	}
	if proc.PendingSignupLib == nil && len(proc.NonProcessedLibraries) != 0 {
		proc.PendingSignupLib = &proc.NonProcessedLibraries[0]
		proc.NonProcessedLibraries = proc.NonProcessedLibraries[1:]
	}
	if proc.PendingSignupLib != nil {
		proc.PendingSignupLib.SignupTime -= 1
	}
}

func (proc *Process) registerBook(library Library) {
	bookRegistered := []int{}
	bookPassed := 0
	for bookIndex, book := range library.BookIDs {
		if len(bookRegistered) > library.BookCount {
			break
		}
		if proc.BookProcessed[book] {
			continue
		}
		bookRegistered = append(bookRegistered, book)
		bookPassed = bookIndex
		proc.BookProcessed[book] = true
	}
	library.BookIDs = library.BookIDs[bookPassed+1:]
	proc.BooksScannedByLib[library.ID] = append(proc.BooksScannedByLib[library.ID], bookRegistered...)
}

func main() {
	days, books, libraries := parseInput(bufio.NewReader(os.Stdin))

	process := NewProcess(days, books, libraries)
	libRegistered, booksByLib := process.Calculate()

	format(libRegistered, booksByLib)
}
