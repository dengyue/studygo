package split

import (
	"strings"
)

//Split string
//example:
//abc,b ==> [abc,b]
func Split(str, sep string) []string {
	var ret []string
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
