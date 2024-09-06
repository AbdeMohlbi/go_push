package main

import (
	"fmt"
	"os/exec"
	"os"
)

func main() {
	var commitMsg string ="++"
	dir,err:=os.Getwd()
	if err !=nil{
		fmt.Println("something wen wrong err: \t",err)
	}
	fmt.Println("current dir : ",dir)

	if len(os.Args) >1 {
		fmt.Println("passing ",os.Args[1]," as a commit message")
		commitMsg = os.Args[1]
	}else{
		fmt.Println("no args provided passing ++ as a commit message")
	}
	
	cmd := exec.Command("powershell", "-Command", "git commit -am \""+commitMsg+"\"; git push ")
	fmt.Println("passing following commands:\ngit commit -am \""+commitMsg+"\" \ngit push ")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Output:%s\n", output)
}