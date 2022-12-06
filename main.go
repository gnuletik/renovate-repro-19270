package app

import "golang.org/x/text/language"

func Func() {
	matcher := language.NewMatcher([]language.Tag{
		language.English, // The first language is used as fallback.
		language.MustParse("en-AU"),
		language.Danish,
		language.Chinese,
	})
	language.MatchStrings(matcher, "fr")
}
