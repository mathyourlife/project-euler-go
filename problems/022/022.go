/*
Using names.txt, a 46K text file containing over five-thousand first
names, begin by sorting it into alphabetical order. Then working out
the alphabetical value for each name, multiply this value by its
alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN,
which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the
list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?

Note: this should be executed in the same directory as "names.txt"

*/

package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"os"
	"strings"
)


func load_names() []string {

  dir, err := os.Getwd()
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
	f := path.Join(dir, "names.txt")
	content, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	content_string := string(content)

	quoted_names := strings.Split(content_string, ",")

	names := []string{}
	for _, quoted_name := range quoted_names {
		names = append(names, quoted_name[1:len(quoted_name)-1])
	}

	return names
}

func swap(names []string, a int, b int) {
	names[a], names[b] = names[b], names[a]

}

func main() {
	names := load_names()

	for {
		swaped := false
		for i := 0; i < len(names) - 1; i++ {
			if names[i] > names[i+1] {
				swaped = true
				swap(names, i, i+1)
			}
		}
		if !swaped {
			break
		}
	}

	total := 0
	for i, name := range names {
		sum := 0
		for _, n := range []byte(name) {
			sum += int(n) - 64
		}
		total += (i+1) * sum
	}
	fmt.Println(total)
}