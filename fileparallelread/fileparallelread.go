package fileparallelread

import (
	"bufio"
	"io"
	"os"
)

// GetSectionReadersUsingMultipleHandle ...
func GetSectionReadersUsingMultipleHandle(filename string, splitInto int, delim rune) ([]*os.File, []*io.SectionReader, []int64, error) {
	var fReaders []*os.File
	for i := 0; i < splitInto; i++ {
		f, err := os.Open(filename)
		if err != nil {
			return nil, nil, nil, err
		}
		fReaders = append(fReaders, f)
	}

	fInfo, _ := fReaders[0].Stat()
	
	sectionSizes, err := getSectionSizes(fReaders[0], fInfo.Size(), splitInto, delim)
	if err != nil {
		return nil, nil, nil, err
	}	

	sectionReaders := make([]*io.SectionReader, splitInto, splitInto)

	for i := 0; i < splitInto; i++ {
		sectionReaders[i] = io.NewSectionReader(fReaders[i], sectionSizes[i], sectionSizes[i+1]-sectionSizes[i])
	}

	return fReaders, sectionReaders, sectionSizes, nil
}

//GetSectionReadersUsingSingleHandle ...
func GetSectionReadersUsingSingleHandle(filename string, splitInto int, delim rune) ([]*os.File, []*io.SectionReader, []int64, error) {
	var fReaders []*os.File
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}
	
	fReaders = append(fReaders, f)

	fInfo, err := f.Stat()
	if err != nil {
		return nil, nil, nil, err
	}
	
	sectionSizes, err := getSectionSizes(fReaders[0], fInfo.Size(), splitInto, delim)
	if err != nil {
		return nil, nil, nil, err
	}
	
	sectionReaders := make([]*io.SectionReader, splitInto, splitInto)
	for i := 0; i < splitInto; i++ {
		sectionReaders[i] = io.NewSectionReader(f, sectionSizes[i], sectionSizes[i+1]-sectionSizes[i])
	}

	return fReaders, sectionReaders, sectionSizes, nil
}

func getSectionSizes(inp *os.File, size int64, splitInto int, delim rune) ([]int64, error) {
	sectionSizes := make([]int64, splitInto+1)
	sectionSizes[0] = 0
	for i := 0; i < splitInto; i++ {
		x := (int64(i) + 1) * size / int64(splitInto) - 4 // - 4 to goback by 1 rune to cover if \n = exactly at mid of file
		for {
			inp.Seek(x, 0)
			br := bufio.NewReader(inp)
			r, size, err := br.ReadRune()
			if err != nil {
				if err == io.EOF {
					sectionSizes[i+1] = x
					break
				}
				return nil, err
			}
			if r == '\n' {
				x += int64(size)
				break
			} else {
				x += int64(size)
			}
		}
		sectionSizes[i+1] = x
	}
	return sectionSizes, nil
}
