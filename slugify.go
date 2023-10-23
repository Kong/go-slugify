package slugify

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

const (
	defaultWordSeparator          = "-"
	defaultInvalidCharReplacement = "-"
)

type Slugifier struct {
	skipToLower            bool    // convert to lower or not
	wordSeparator          *string // separator char between words
	invalidCharReplacement *string // replacement for illegal chars
	reInValidChar          *regexp.Regexp
	reDupSeparatorChar     *regexp.Regexp
	initialized            bool
}

// NewSlugifier creates a new slugifier. Defaults to lowercasing and using a dash ("-") as the word separator and
// as the invalid character replacement.
func NewSlugifier() Slugifier {
	s := Slugifier{}
	s.InvalidChar(defaultInvalidCharReplacement)
	s.WordSeparator(defaultWordSeparator)
	return s
}

// initialize initializes the slugifier. This is called automatically by Slugify() if it hasn't already been called.
func (s *Slugifier) initialize() {
	if s.invalidCharReplacement == nil {
		s.InvalidChar(defaultInvalidCharReplacement)
	}

	if s.wordSeparator == nil {
		ws := defaultWordSeparator
		s.wordSeparator = &ws
	}

	separatorForRe := regexp.QuoteMeta(*s.wordSeparator)
	s.reInValidChar = regexp.MustCompile(fmt.Sprintf("[^%sa-zA-Z0-9]", separatorForRe))
	if separatorForRe != "" {
		s.reDupSeparatorChar = regexp.MustCompile(fmt.Sprintf("%s{2,}", separatorForRe))
	} else {
		s.reDupSeparatorChar = nil
	}

	s.initialized = true
}

// ToLower sets the flag indicating if the slugified result should be lowercased.
// Returns the slugifier for easy chaning.
func (s *Slugifier) ToLower(toLower bool) *Slugifier {
	s.skipToLower = !toLower
	return s
}

// WordSeparator sets the word separator character to use. Defaults to a dash ("-"). The word separator is used to
// replace whitespace. Leading and trailing word separators are trimmed. Multiple successive word separators are
// replaced with a single word separator.
// Returns the slugifier for easy chaning.
func (s *Slugifier) WordSeparator(wordSeparator string) *Slugifier {
	s.wordSeparator = &wordSeparator
	s.reInValidChar = nil
	s.reDupSeparatorChar = nil
	s.initialized = false
	return s
}

// InvalidChar sets the character to use to replace invalid characters (anything not a-z, A-Z, 0-9, the
// word separator, or the InvalidChar). Defaults to a dash ("-"). Leading and trailing
// InvalidCharReplacements are trimmed. Multiple successive InvalidCharReplacements are NOT replaced with a single
// InvalidChar.
// Returns the slugifier for easy chaning.
func (s *Slugifier) InvalidChar(invalidCharReplacement string) *Slugifier {
	s.invalidCharReplacement = &invalidCharReplacement
	return s
}

// Version return version
func Version() string {
	return "0.2.0"
}

// Slugify implements making a pretty slug from the given text.
// e.g. Slugify("kožušček hello world") => "kozuscek-hello-world"
func (s *Slugifier) Slugify(txt string) string {
	if !s.initialized {
		s.initialize()
	}
	txt = unidecode.Unidecode(txt)
	txt = strings.Join(strings.Fields(txt), *s.wordSeparator)
	txt = s.reInValidChar.ReplaceAllString(txt, *s.invalidCharReplacement)
	if s.reDupSeparatorChar != nil {
		txt = s.reDupSeparatorChar.ReplaceAllString(txt, *s.wordSeparator)
	}

	// trim leading and trailing word separators and invalidCharReplacements
	for len(txt) > 0 &&
		(string(txt[0]) == *s.wordSeparator || string(txt[0]) == *s.invalidCharReplacement) {
		txt = txt[1:]
	}
	for len(txt) > 0 &&
		(string(txt[len(txt)-1]) == *s.wordSeparator || string(txt[len(txt)-1]) == *s.invalidCharReplacement) {
		txt = txt[:len(txt)-1]
	}

	if !s.skipToLower {
		txt = strings.ToLower(txt)
	}
	return txt
}
