package reverse_string

func ReverseString(input string) (output string) {
	str:=[]rune(input)
	var output[]rune
	for i:=len(str)-1;i>=0;i--{
	output=append(output,str[i])
	}
	return output
}
