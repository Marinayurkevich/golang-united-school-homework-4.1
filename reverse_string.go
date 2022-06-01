package reverse_string
import "strings"
func ReverseString(input string) (output string) {
	runes:=[]rune(input)
	var line_runes []rune
	for i:=len(runes)-1;i>=0;i--{
	line_runes=append(line_runes,runes[i])
	}	
	return string(line_runes)
	}
