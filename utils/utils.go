package utils

import (
	"bufio"
	"io"
)

// RunesReadMultipleSectionReaders ...
func GetRunesReadMultipleSectionReaders(sectionReaders ...*io.SectionReader) []int {
	chs := make([]chan int, len(sectionReaders))
	for i := range chs {
		chs[i] = make(chan int)
	}
	out := make([]int, len(sectionReaders))

	for i := range sectionReaders {	
		go getRunesReadSectionReader(i, sectionReaders[i], chs[i])
	}	

	for i := range sectionReaders {
		out[i] = <-chs[i]
	}

	
	return out

}

func getRunesReadSectionReader(sectionNo int, inp *io.SectionReader, ch chan int) {	
	runesRead := 0
	br := bufio.NewReader(inp)
	
	for {
		_, _, err := br.ReadRune()
		if err != nil {
			if err == io.EOF {				
				break
			}
			panic(err)
		}
		runesRead++
	}
	
	ch <- runesRead
	close(ch)
}
