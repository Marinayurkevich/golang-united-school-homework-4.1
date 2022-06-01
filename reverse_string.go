package reverse_string

func ReverseString(input string) (output string) {
	line:=[]rune(input)
	var output []rune
	for i:=len(line)-1;i>=0;i--{
	output=append(output,line[i])
	}
	return string(output)
}
