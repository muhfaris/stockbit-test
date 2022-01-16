## The Question 4
### Logic Test
Anagram adalah istilah dimana suatu string jika dibolak balik ordernya maka akan sama eg. 'aku' dan
'kua' adalah Anagram, 'aku' dan 'aka' bukan Anagram.

Dibawah ini ada array berisi sederetan Strings.
`['kita', 'atik', 'tika', 'aku', 'kia', 'makan', 'kua']`

Silahkan kelompokkan/group kata-kata di dalamnya sesuai dengan kelompok Anagramnya,

### Expected Outputs
``` 
[
    ["kita", "atik", "tika"],
    ["aku", "kua"],
    ["makan"],
    ["kia"]
]
```

## Answer
### Solution
Loop data anagrams (array data), then compare the anagram with the list of anagrams. If the anagram finds another anagram with the same length, validate the letter. whether they were identic or not.

The identic anagram will be grouped and stored in the array.

###  Test the code

You can run application with this command:

`go run anagram.go kita atik tika aku kia makan kua`

or 

you can run the test file

`go test -v ./...`


result:
```

[
  [
   "kita",
   "atik",
   "tika"
  ],
  [
   "aku",
   "kua"
  ],
  [
   "kia"
  ],
  [
   "makan"
  ]
 ]

```
