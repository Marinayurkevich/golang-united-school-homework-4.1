package reverse_string

func ReverseString(input string) (output string) {
	runes:=[]rune(input)
	var line_runes []rune
	for i:=len(runes)-1;i>=0;i--{
	line_runes=append(line_runes,runes[i])
	}
	for _,r:=range line_runes{
	output=append(output,line_runes[i])
	}
	return output
}
