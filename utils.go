package main

import "strconv"

func _UnescapeUnicodeCharactersInJSON(data string) ([]byte, error) {
	quoted := `"` + data + `"`
	unescaped, err := strconv.Unquote(quoted)
	if err != nil {
		return nil, err
	}

	return []byte(unescaped), nil
}
