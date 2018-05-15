package main

import (
    "fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func isNumber(s string) bool{
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func readLines(path string, m map[string]string) map[string]map[string]string{
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	  scanner.Split(bufio.ScanLines) 

	var found = make(map[string]int)
	for k := range m{
		found[k] = 0
	}
	
	var maps = make(map[string]map[string]string)
	for scanner.Scan() {
		for k := range m {
			if strings.Contains(scanner.Text(), m[k]){
				fmt.Println("------------" + k + "---------------")
				found[k]=1
				maps[k] = make(map[string]string)
			}
			if strings.Contains(scanner.Text(), "}"){
				found[k]=0
			}
			if strings.Contains(scanner.Text(), "=") && found[k]>0{
				var name = strings.Split(scanner.Text(), "=")[0]
				name = strings.Replace(name, " ", "", -1)
				name = strings.Replace(name, "\t", "", -1)
				name = strings.Replace(name, "\n", "", -1)
				var rest = strings.Split(scanner.Text(), "=")[1]
				var preComment = strings.Split(rest, "/")[0]
				preComment = strings.Replace(preComment, ",", "", -1)
				preComment = strings.Replace(preComment, " ", "", -1)
				preComment = strings.Replace(preComment, "\t", "", -1)
				preComment = strings.Replace(preComment, "\n", "", -1)
				maps[k][preComment] = name
				fmt.Println(name + "=" + preComment)
			}
		}
	}

	return maps;
  }

  func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
	  return err
	}
	defer file.Close()
  
	w := bufio.NewWriter(file)
	for _, line := range lines {
	  fmt.Fprintln(w, line)
	}
	return w.Flush()
  }

func main() {
	// init main map
	var m map[string]string
	
	m = make(map[string]string)

	m["AdvanceMedia"] = "cups_adv_e"
	m["cupsColorSpace"] = "cups_cspace_e"
	m["CutMedia"] = "cups_cut_e"
	m["LeadingEdge"] = "cups_edge_e"
	m["Jog"] = "cups_jog_e"
	m["cupsColorOrder"] = "cups_order_e"
	m["Orientation"] = "cups_orient_e"

	var maps = readLines("./cups-enums.h", m)

	var lines = make([]string, 0)
	// prepare file 'header'
	lines = append(lines,"package cups_itf")
	lines = append(lines,"\n")
	lines = append(lines,"var Maps = map[string]map[string]string {")

	for k := range maps{
		lines = append(lines, "\"" + k + "\": {")
		for kk := range maps[k]{
			lines = append(lines, "\""  + kk + "\": " + "\"" + maps[k][kk] + "\",")
		}
		lines = append(lines, "},")
	}
	lines = append(lines,"}")
	lines = append(lines, "func main(){\n}")

	writeLines(lines, "../cups_itf/cups_itf.go")

	for i := range lines{
		fmt.Println(lines[i])
	}

	
}
