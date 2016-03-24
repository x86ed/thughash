[![Coverage](http://gocover.io/_badge/github.com/x86ed/thughash?0)](http://gocover.io/github.com/x86ed/thughash)
[![Go Report Card](https://goreportcard.com/badge/github.com/x86ed/thughash)](https://goreportcard.com/report/github.com/x86ed/thughash) 

# thughash
Human readable hash library for those who like to keep their unique ID's 'real'.

### Installation

Assuming you already have a recent version of Go installed, pull down the code with `go get`:

```
go get github.com/x86ed/thughash
```

### Usage

```go
package main

import (

	"github.com/x86ed/thughash"
	"math/rand"
	"time"
	"fmt"

)

func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	var hash thughash.ThugHash
	hash.Generate(rand.Float64()*2147483647)
	
	fmt.Printf("Your random hash is %#v .\n", hash.MakeSlug())
	// Your Random hash is Mothafucka-turnt-holdin-it-down-1867 .

}

```

### To Do
 * add stdio functionality
 * make reconfigurable
 * add hashes of other lengths
 * more docs


### License

ThugHash is under the GNU GPL v3.0 license.