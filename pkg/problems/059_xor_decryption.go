package problems

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type XORDecryption struct {
	cipherFile string
}

func (p *XORDecryption) ID() int {
	return 59
}

func (p *XORDecryption) Text() string {
	return `Each character on a computer is assigned a unique
code and the preferred standard is ASCII (American Standard
Code for Information Interchange). For example, uppercase
A = 65, asterisk (*) = 42, and lowercase k = 107.

A modern encryption method is to take a text file, convert
the bytes to ASCII, then XOR each byte with a given value,
taken from a secret key. The advantage with the XOR function
is that using the same encryption key on the cipher text,
restores the plain text; for example, 65 XOR 42 = 107,
then 107 XOR 42 = 65.

For unbreakable encryption, the key is the same length as
the plain text message, and the key is made up of random
bytes. The user would keep the encrypted message and the
encryption key in different locations, and without both
"halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users,
so the modified method is to use a password as a key. If
the password is shorter than the message, which is likely,
the key is repeated cyclically throughout the message.
The balance for this method is using a sufficiently long
password key for security, but short enough to be memorable.

Your task has been made easy, as the encryption key
consists of three lower case characters. Using cipher.txt,
a file containing the encrypted ASCII codes, and the
knowledge that the plain text must contain common English
words, decrypt the message and find the sum of the ASCII
values in the original text.
`
}

func (p *XORDecryption) Solve() (string, error) {
	file, err := os.Open(p.cipherFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fb, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	content := make([]byte, 0, len(fb))
	for _, d := range strings.Split(string(bytes.TrimSpace(fb)), ",") {
		dd, err := strconv.ParseInt(d, 0, 8)
		if err != nil {
			return "", err
		}
		content = append(content, byte(dd))
	}

	found := false
	var plaintext []byte
	for a := byte(97); a < byte(123); a++ {
		for b := byte(97); b < byte(123); b++ {
			for c := byte(97); c < byte(123); c++ {
				key := []byte{a, b, c}
				plaintext = make([]byte, 0, len(content))
				for i, by := range content {
					plaintext = append(plaintext, by^key[i%3])
				}
				if bytes.Contains(plaintext, []byte(" the ")) &&
					bytes.Contains(plaintext, []byte(" of ")) {
					found = true
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	sum := int64(0)
	for _, b := range plaintext {
		sum += int64(b)
	}
	return fmt.Sprintf("%d", sum), nil

}
