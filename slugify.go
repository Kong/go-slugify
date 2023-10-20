package slugify

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

const (
	defaultToLower                = true
	defaultWordSeparator          = "-"
	defaultInvalidCharReplacement = "-"
)

type slugifier struct {
	toLower                bool   // convert to lower or not
	wordSeparator          string // separator char between words
	invalidCharReplacement string // replacement for illegal chars
	reInValidChar          *regexp.Regexp
	reDupSeparatorChar     *regexp.Regexp
}

// NewSlugifier creates a new slugifier. Defaults to lowercasing and using a dash ("-") as the word separator and
// as the invalid character replacement.
func NewSlugifier() slugifier {
	s := slugifier{
		toLower:                defaultToLower,
		invalidCharReplacement: defaultInvalidCharReplacement,
	}
	s.WordSeparator(defaultWordSeparator)
	return s
}

// ToLower sets the flag indicating if the slugified result should be lowercased.
func (s *slugifier) ToLower(toLower bool) {
	s.toLower = toLower
}

// WordSeparator sets the word separator character to use. Defaults to a dash ("-"). The word separator is used to
// replace whitespace. Leading and trailing word separators are trimmed. Multiple successive word separators are
// replaced with a single word separator.
func (s *slugifier) WordSeparator(wordSeparator string) error {
	s.wordSeparator = wordSeparator
	separatorForRe := regexp.QuoteMeta(s.wordSeparator)
	var err error
	if s.reInValidChar, err = regexp.Compile(fmt.Sprintf("[^%sa-zA-Z0-9]", separatorForRe)); err != nil {
		return err
	}
	if s.reDupSeparatorChar, err = regexp.Compile(fmt.Sprintf("%s{2,}", separatorForRe)); err != nil {
		return err
	}
	return nil
}

// InvalidCharReplacement sets the character to use to replace invalid characters (anything not a-z, A-Z, 0-9, the
// word separator, or the InvalidCharReplacement). Defaults to a dash ("-"). Leading and trailing
// InvalidCharReplacements are trimmed. Multiple successive InvalidCharReplacements are NOT replaced with a single
// InvalidCharReplacement.
func (s *slugifier) InvalidCharReplacement(invalidCharReplacement string) {
	s.invalidCharReplacement = invalidCharReplacement
}

// Version return version
func Version() string {
	return "0.2.0"
}

// Slugify implements making a pretty slug from the given text.
// e.g. Slugify("kožušček hello world") => "kozuscek-hello-world"
func (s *slugifier) Slugify(txt string) string {
	txt = unidecode.Unidecode(txt)
	txt = strings.Join(strings.Fields(txt), s.wordSeparator)
	txt = s.reInValidChar.ReplaceAllString(txt, s.invalidCharReplacement)
	txt = s.reDupSeparatorChar.ReplaceAllString(txt, s.wordSeparator)

	// trim leading and trailing word separators and invalidCharReplacements
	for len(txt) > 0 &&
		(string(txt[0]) == s.wordSeparator || string(txt[0]) == s.invalidCharReplacement) {
		txt = txt[1:]
	}
	for len(txt) > 0 &&
		(string(txt[len(txt)-1]) == s.wordSeparator || string(txt[len(txt)-1]) == s.invalidCharReplacement) {
		txt = txt[:len(txt)-1]
	}

	if s.toLower {
		txt = strings.ToLower(txt)
	}
	return txt
}
