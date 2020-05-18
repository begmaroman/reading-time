package readingtime

import (
	"fmt"
	"math"
	"time"
)

const (
	// defaultWordsPerMinute is the default number of words per minute an average reader can read.
	defaultWordsPerMinute = 200
)

var (
	// defaultWordBoundFunc is the default WordBoundFunc implementation.
	defaultWordBoundFunc WordBoundFunc = func(b byte) bool {
		return b == ' ' || b == '\n' || b == '\r' || b == '\t'
	}
)

// WordBoundFunc is a function that returns a boolean value depending on if a character is considered as a word bound
type WordBoundFunc func(b byte) bool

// Option is the function that modifies the given options.
type Option func(opts *Options)

// WordsPerMinute is the function that modifies the number of words per minute.
func WordsPerMinute(wordsPerMinute int) Option {
	return func(opts *Options) {
		opts.WordsPerMinute = wordsPerMinute
	}
}

// Options contains options to configure time reading estimation logic.
type Options struct {
	// WordsPerMinute is the number of words per minute an average reader can read
	// Default: 200.
	WordsPerMinute int

	// WordBound is a WordBoundFunc representation.
	// Default: spaces, new lines and tabulations.
	WordBound WordBoundFunc
}

// Result contains result of an estimated time to read an article.
type Result struct {
	Text     string
	Duration time.Duration
	Words    int
}

// Estimate estimates how long the given text will take to read using the given configuration.
func Estimate(text string, opts ...Option) *Result {
	if len(text) == 0 {
		return &Result{
			Text:     "0 min read",
			Duration: 0,
			Words:    0,
		}
	}

	// Init options with default values.
	op := Options{
		WordsPerMinute: defaultWordsPerMinute,
		WordBound:      defaultWordBoundFunc,
	}

	// Apply custom configuration.
	for _, opt := range opts {
		opt(&op)
	}

	words := 0
	start := 0
	end := len(text) - 1

	// Fetch bounds.
	for op.WordBound(text[start]) {
		start++
	}
	for op.WordBound(text[end]) {
		end--
	}

	// Calculate the number of words.
	for i := start; i <= end; {
		for i <= end && !op.WordBound(text[i]) {
			i++
		}

		words++

		for i <= end && op.WordBound(text[i]) {
			i++
		}
	}

	// Reading time stats.
	minutes := math.Ceil(float64(words) / float64(op.WordsPerMinute))
	duration := time.Duration(math.Ceil(minutes) * float64(time.Minute))

	return &Result{
		Text:     fmt.Sprintf("%d min read", int(minutes)),
		Duration: duration,
		Words:    words,
	}
}
