/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ini

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"fmt"
)

var (
	emptyRunes = []rune{}
)

func isSep(b []rune) bool {
	if len(b) == 0 {
		return false
	}

	switch b[0] {
	case '[', ']':
		return true
	default:
		return false
	}
}

var (
	openBrace  = []rune("[")
	closeBrace = []rune("]")
)

func newSepToken(b []rune) (Token, int, error) {
	tok := Token{}

	switch b[0] {
	case '[':
		tok = newToken(TokenSep, openBrace, NoneType)
	case ']':
		tok = newToken(TokenSep, closeBrace, NoneType)
	default:
		return tok, 0, NewParseError(fmt.Sprintf("unexpected sep type, %v", b[0]))
	}
	return tok, 1, nil
}
