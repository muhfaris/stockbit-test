## The Question
### Please refactor the code below to make it more concise, efficient and readable with good logic flow.
```
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}
```

## Answer
### Solution: 
I use `strings` library from package go [https://pkg.go.dev/strings](https://pkg.go.dev/strings). The package have more functions, this case some functions help me to achive the goal. The list functions like below:

- `HasPrefix` is checking the string, whether have x in begin.
- `TrimPrefix` is to remove x from beginning the string. 
- `HasSuffix` is checking the string, whether has x in last. 
- `TrimSuffix` is removing x from last the string.

### How to test the code
You can run via terminal: `go run find_string.go "(ayo main bareng)"`

or You can run the test file `go test -v ./...`
