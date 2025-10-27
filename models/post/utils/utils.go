package utils

import "bifrost/models/post/shared"

func MakeLocalizedString(lang string, text string) *shared.LocalizedString {
	if text == "" {
		return nil
	}
	ls := shared.LocalizedString{lang: text}
	return &ls
}
