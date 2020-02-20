package main

import (
	"bufio"
	"os"
	"sort"

	"github.com/jinzhu/copier"
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
		// // fmt.Printf("%v %v %v\n", day, len(proc.NonProcessedLibraries), (proc.SignedUpLibrary))
		proc.keepLibrarySignupGoing(proc.Days - day)

		for _, library := range proc.SignedUpLibrary {
			proc.registerBook(library)
		}
	}
	return proc.SignUpLibrary, proc.BooksScannedByLib
}

func (proc *Process) keepLibrarySignupGoing(daysLeft int) {
	// fmt.Printf("pending: %v\n", proc.PendingSignupLib)
	if proc.PendingSignupLib != nil {
		if proc.PendingSignupLib.SignupTime == 0 {
			proc.SignUpLibrary = append(proc.SignUpLibrary, proc.PendingSignupLib.ID)
			proc.SignedUpLibrary = append(proc.SignedUpLibrary, *proc.PendingSignupLib)
			proc.PendingSignupLib = nil
		}
	}
	if proc.PendingSignupLib == nil && len(proc.NonProcessedLibraries) != 0 {
		// fmt.Printf("before %v\n", proc.NonProcessedLibraries)
		bestLibIndex := proc.pickBestLibrary(daysLeft)
		// fmt.Printf("index: %v\n", bestLibIndex)
		pendingSignup := Library{}
		//proc.PendingSignupLib = &proc.NonProcessedLibraries[bestLibIndex]
		copier.Copy(&pendingSignup, &proc.NonProcessedLibraries[bestLibIndex])
		proc.PendingSignupLib = &pendingSignup
		// fmt.Printf("before delete: %v\n", proc.NonProcessedLibraries)
		// fmt.Printf("pending signup: %v\n", proc.PendingSignupLib)
		proc.NonProcessedLibraries = append(proc.NonProcessedLibraries[0:bestLibIndex], proc.NonProcessedLibraries[bestLibIndex+1:]...)
		// fmt.Printf("after delete: %v\n", proc.NonProcessedLibraries)
		// fmt.Printf("pending signup: %v\n", proc.PendingSignupLib)
	}
	if proc.PendingSignupLib != nil {
		proc.PendingSignupLib.SignupTime -= 1
	}
}

func (proc *Process) pickBestLibrary(daysLeft int) int {
	bestLib := 0
	bestScore := 0
	for libraryIndex := range proc.NonProcessedLibraries {
		library := proc.NonProcessedLibraries[libraryIndex]
		daysLeftWithoutSignup := daysLeft - library.SignupTime
		maxPossibleBooks := daysLeftWithoutSignup * library.BookCount
		score := 0
		bookTaken := 0
		for bookIndex := 0; bookIndex < len(library.BookIDs); bookIndex++ {
			if bookTaken > maxPossibleBooks {
				break
			}
			bookID := library.BookIDs[bookIndex]
			if proc.BookProcessed[bookID] {
				continue
			}
			score += proc.BookScores[bookID]
			bookTaken++
		}
		// fmt.Printf("%v\n", library)
		score = score / library.SignupTime
		if score > bestScore {
			bestLib = libraryIndex
			bestScore = score
		}
	}

	return bestLib
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
	// fmt.Printf("%v\n", libraries)

	process := NewProcess(days, books, libraries)
	libRegistered, booksByLib := process.Calculate()

	format(libRegistered, booksByLib)
}
