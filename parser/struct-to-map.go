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
				//fmt.Println("------------" + m[k] + "---------------")
				found[k]=1
				maps[m[k]] = make(map[string]string)
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
				maps[m[k]][preComment] = name
				//fmt.Println(name + "=" + preComment)
			}
		}
	}

	return maps;
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

	// map of maps: e.g. m[cups_edge_e]={0:CUPS_EDGE_TOP, 1:CUPS_EDGE_RIGHT etc}
	//var maps map[string]map[int]string

	//maps = map[string]map[int]string

	var maps = readLines("./cups-enums.h", m)
	for k := range maps{
		for kk := range maps[k]{
			fmt.Println(k + " -->   " + kk + " = " + maps[k][kk])
		}
	}
}
