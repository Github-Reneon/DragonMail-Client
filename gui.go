package main

type model struct {
	choice   []string
	cursor   int
	selected map[int]struct{}
}
