package readingtime_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	readingtime "github.com/begmaroman/reading-time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func generateText(words int) string {
	var b []byte
	for i := 0; i < words; i++ {
		wordLength := seededRand.Intn(len(charset)) + 1
		for j := 0; j < wordLength; j++ {
			b = append(b, charset[seededRand.Intn(len(charset))])
		}
		b = append(b, ' ')
	}
	return string(b)
}

func TestEstimate(t *testing.T) {
	testTable := []struct {
		testName    string
		text        string
		opts        []readingtime.Option
		expectedRes *readingtime.Result
	}{
		{
			testName: "should handle less than 1 minute text",
			text:     generateText(2),
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    2,
			},
		},
		{
			testName: "should handle less than 1 minute text",
			text:     generateText(50),
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    50,
			},
		},
		{
			testName: "should handle 1 minute text",
			text:     generateText(100),
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    100,
			},
		},
		{
			testName: "should handle 2 minutes text",
			text:     generateText(300),
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "2 min read",
				Duration: time.Minute * 2,
				Words:    300,
			},
		},
		{
			testName: "should handle a very long text",
			text:     generateText(500),
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "3 min read",
				Duration: time.Minute * 3,
				Words:    500,
			},
		},
		{
			testName: "should handle text containing multiple successive whitespaces",
			text:     "word  word    word",
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    3,
			},
		},
		{
			testName: "should handle text starting with whitespaces",
			text:     "   word word word",
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    3,
			},
		},
		{
			testName: "should handle text ending with whitespaces",
			text:     "word word word   ",
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    3,
			},
		},
		{
			testName: "should handle text containing links",
			text:     "word https://github.com/begmaroman word",
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    3,
			},
		},
		{
			testName: "should handle text containing markdown links",
			text:     "word [github](https://github.com/begmaroman) word",
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    3,
			},
		},
		{
			testName: "should handle text containing one word correctly",
			text:     "0",
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "1 min read",
				Duration: time.Minute,
				Words:    1,
			},
		},
		{
			testName: "should handle text containing a black hole",
			text:     "",
			opts:     nil,
			expectedRes: &readingtime.Result{
				Text:     "0 min read",
				Duration: 0,
				Words:    0,
			},
		},
		{
			testName: "should accept a custom word per minutes value",
			text:     generateText(200),
			opts: []readingtime.Option{
				readingtime.WordsPerMinute(100),
			},
			expectedRes: &readingtime.Result{
				Text:     "2 min read",
				Duration: time.Minute * 2,
				Words:    200,
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			actualRes := readingtime.Estimate(tt.text, tt.opts...)
			require.Equal(t, tt.expectedRes, actualRes)
		})
	}
}
