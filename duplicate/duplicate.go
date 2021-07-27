package duplicate

import "strings"

func LongStringDuplicate(str string ) string {
	return longestStringNoDuplicate(repetition(str))
}

func repetition(st string) (stringList []string) {
	var part01, part02 string
	for i := 0; i < len(st); i++ {
		index := strings.Index(st[i+1:], string(st[i]))
		if index >= 0 {
			part01 = st[0 : index+1+i]
			part02 = st[index+1+i:]
			stringList = append(stringList, part01)
			st = part02
			i = -1
		}
	}
	stringList = append(stringList, part02)
	return
}

func longestStringNoDuplicate(stringList []string) (longest string) {
	var maxLength int = 0
	for _, data := range stringList {
		if len(data) > maxLength {
			longest = data
			maxLength = len(data)
		}
	}
	return
}

