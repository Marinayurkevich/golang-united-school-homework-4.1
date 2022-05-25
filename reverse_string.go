package reverse_string

func ReverseString(input string) (output string) {
	line:=[]rune(input)
	for i:=len(line)-1;i>=0;i--{
	output=append(output,line[i])
	}
	return output
}
