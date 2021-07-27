package palindrome

type Pal struct {

}

func New() *Pal {
	return &Pal{}
}

func (pal *Pal) IsPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return pal.IsPalindrome(s[1 : len(s)-1])
	}
	return false

}

func (pal *Pal) RepeatCharacters(s string) string {

	return  ""
}
