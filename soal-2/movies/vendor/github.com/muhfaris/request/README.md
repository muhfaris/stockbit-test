# Request
## Deskripsi
Library yang memudahkan kita untuk membuat http request lebih mudah.

## Fitur
- Request dengan method Get, Post, Delete, Patch.
- Sistem retry, untuk melakukan request ulang jika terjadi error. Kamu bisa set total retry pada field `Retry` dan kamu juga bisa melakukan setting jeda waktu antar request yang satu dengan yang lain.
- Parse response data ke struct / map[string].
- custom user agent

## Install
Untuk menggunakan paket request, Anda harus menginstall Go dan setup Go workspace.
- Install paket, jalankan perintah berikut
`go get github.com/muhfaris/request`

- Import ke dalam kode:
`import "github.com/muhfaris/request"`

## Get Request dan parse response data ke struct
```
package main

import (
	"log"

	"github.com/muhfaris/request"
)

type QuoteModel struct {
	Anime     string
	Character string
	Quote     string
}

func main() {
	var quoteModel QuoteModel
	config := &request.Config{URL: "https://animechan.vercel.app/api/random"}
	response := request.Get(config).Parse(&quoteModel)
	if response.Error != nil {
		log.Printf("error get quote anime, %v", response.Error)
		return
	}

	log.Println("Quote:>")
	log.Println(quoteModel.Anime)
	log.Println(quoteModel.Character)
	log.Println(quoteModel.Quote)
}
```

## Post request dengan retry
```
import "github.com/muhfaris/request"
func main(){
    resp, err := request.Post(
        &request.Config{
            "URL": "https://facebook.com/v1/api/profile",
            "Method": "POST",
            "Retry": 1, 
            "Delay": 10 * time.Seconds,
        },
    )
}
```

## Post application/json
```

package main

import (
	"encoding/json"
	"log"

	"github.com/muhfaris/request"
)

type User struct {
	Name string
	Job  string
}

func main() {
	user := User{
		Name: "faris",
		Job:  "software developer",
	}

	raw, err := json.Marshal(user)
	if err != nil {
		log.Printf("error marshal the user data, %v", err)
		return
	}

	config := &request.Config{
		URL:  "https://reqres.in/api/users",
		Body: raw,
	}

	resp := request.Post(config)
	if resp.Error != nil {
		log.Printf("error create new post user, %v", resp.Error)
		return
	}

	log.Println("Resp")
	log.Println(string(resp.Body))
}
```
