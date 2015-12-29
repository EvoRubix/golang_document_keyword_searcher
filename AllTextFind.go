package main

import (
	"path/filepath"
	"os"
	"fmt"
	"os/exec"
	"bufio"
	//"log"
)

func transform(path string){
	cmd := exec.Command("echo", path)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	fmt.Printf("%s\n", err)
}

func readtext(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}


func main() {
	//Variable declaration of myWalkFunc equals the output func which has the parameters (path, which is  string; etc)
	myWalkFunc := func(path string, info os.FileInfo, err error) error {
		//Package.name handles parameters
		fmt.Printf("%s\n", path)
		//Invokes next func
		transform(path)
		return nil
	}
	//Variable declaration of err to Package.name handing a hardcoded path
	err := filepath.Walk("C:/Users/peter_000/Desktop/testpath", myWalkFunc) 
	fmt.Printf("%s\n", err) 
	
}