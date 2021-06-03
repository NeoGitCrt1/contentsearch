package main

import (
	"bufio"
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type void struct{}
var reg *regexp.Regexp
var ch chan string = make(chan string, 200)
var wg *sync.WaitGroup = &sync.WaitGroup{}
var pos int
func main() {
	rootPath := flag.String("root", "C:/work/f", "search root")
	outfile := flag.String("out", "C:/work/tmp/r.txt", "search result output file fullname")
	regExp := flag.String("reg", `MailMapper\.getMailBody - ==> Parameters:\s(?P<name>.*)\(String\)`, "regex for search line")
	flag.IntVar(&pos, "p", 1, "index of the capture group, -1 for whole line")
	flag.Parse()
	reg = regexp.MustCompile(*regExp)
	
	log.Printf("start from %s", *rootPath)
	
	go out(outfile)
	
	iterate(rootPath, scan, wg)
	wg.Wait()
	close(ch)

}
func out(outfile *string) {
	fi, _ := os.Create(*outfile)
	defer fi.Close()
	var member void
	set := make(map[string]void)
	for f := range ch {
		_, exists := set[f]
		if !exists {
			set[f] = member
			fi.WriteString(f)
			fi.WriteString("\r\n")
		//	log.Println(f)
		}
	}
}

func scan(fn string) {
	defer wg.Done()
	logFile, _ := os.Open(fn)
	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		line := scanner.Text()
		found := reg.FindStringSubmatch(line)
		if ( found != nil ) {
			if (pos == -1) {
				ch <- line
			} else {
				ch <- found[pos]
			}
		}
		
	}
}


func iterate(path *string, searchMailId func(fn string), wg *sync.WaitGroup) {
	filepath.WalkDir(*path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
			return err
		}
		if d.IsDir() {
			log.Printf("------------------------------------ %s", path)
			return nil
		}
		if filepath.Ext(d.Name()) == ".zip" {
			return nil
		}
		log.Printf(">> %s", d.Name())
		wg.Add(1)
		go searchMailId(path)
		return nil
	})

}
