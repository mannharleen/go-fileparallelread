package fileparallelread

import (
	"reflect"
	"testing"

	"github.com/mannharleen/goParallelRead/utils"
)

func TestGetSectionReadersUsingMultipleHandle(t *testing.T) {
	type args struct {
		filename  string
		splitInto int
	}
	tests := []struct {
		name    string
		args    args
		want    int     //NoFileHandles
		want1   int     //NoSectionReaders
		want2   []int64 //SectionSizes
		want3   []int   //RunesRead
		wantErr bool
	}{
		{
			"test1: MultipleHandle 2.csv into 2",
			args{filename: "..\\testdata\\2.csv", splitInto: 2},
			2,
			2,
			[]int64{0, 67, 134},
			[]int{67, 67},
			false,
		},
		{
			"test2: MultipleHandle 2.csv into 3",
			args{filename: "..\\testdata\\2.csv", splitInto: 3},
			3,
			3,
			[]int64{0, 67, 134, 134},
			[]int{67, 67, 0},
			false,
		},
		{
			"test3: MultipleHandle 500.csv into 3",
			args{filename: "..\\testdata\\500.csv", splitInto: 3},
			3,
			3,
			[]int64{0, 9902, 19811, 29602},
			[]int{9902, 9909, 9791},
			false,
		},
		{
			"test4: MultipleHandle xx.csv into 3 (error)",
			args{filename: "..\\testdata\\xx.csv", splitInto: 3},
			3,
			3,
			[]int64{0, 9902, 19811, 29602},
			[]int{9902, 9909, 9791},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fReaders, sectionReaders, sectionSizes, err := GetSectionReadersUsingMultipleHandle(tt.args.filename, tt.args.splitInto, '\n')
			got, got1, got2 := len(fReaders), len(sectionReaders), sectionSizes
			var got3 []int

			if err == nil {
				got3 = utils.GetRunesReadMultipleSectionReaders(sectionReaders...)
			}
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetSectionReadersUsingMultipleHandle() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSectionReadersUsingMultipleHandle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetSectionReadersUsingMultipleHandle() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetSectionReadersUsingMultipleHandle() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("GetSectionReadersUsingMultipleHandle() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func TestGetSectionReadersUsingSingleHandle(t *testing.T) {
	type args struct {
		filename  string
		splitInto int
	}
	tests := []struct {
		name    string
		args    args
		want    int     //NoFileHandles
		want1   int     //NoSectionReaders
		want2   []int64 //SectionSizes
		want3   []int   //RunesRead
		wantErr bool
	}{
		{
			"test1: SingleHandle 2.csv into 2",
			args{filename: "..\\testdata\\2.csv", splitInto: 2},
			1,
			2,
			[]int64{0, 67, 134},
			[]int{67, 67},
			false,
		},
		{
			"test2: SingleHandle 2.csv into 3",
			args{filename: "..\\testdata\\2.csv", splitInto: 3},
			1,
			3,
			[]int64{0, 67, 134, 134},
			[]int{67, 67, 0},
			false,
		},
		{
			"test3: SingleHandle 500.csv into 3",
			args{filename: "..\\testdata\\500.csv", splitInto: 3},
			1,
			3,
			[]int64{0, 9902, 19811, 29602},
			[]int{9902, 9909, 9791},
			false,
		},
		{
			"test4: SingleHandle xx.csv into 3 (error)",
			args{filename: "..\\testdata\\xx.csv", splitInto: 3},
			1,
			3,
			[]int64{0, 9902, 19811, 29602},
			[]int{9902, 9909, 9791},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fReaders, sectionReaders, sectionSizes, err := GetSectionReadersUsingSingleHandle(tt.args.filename, tt.args.splitInto, '\n')
			got, got1, got2 := len(fReaders), len(sectionReaders), sectionSizes
			got3 := utils.GetRunesReadMultipleSectionReaders(sectionReaders...)

			// if (err != nil) != tt.wantErr {
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetSectionReadersUsingsingleHandle() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSectionReadersUsingsingleHandle() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetSectionReadersUsingsingleHandle() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GetSectionReadersUsingsingleHandle() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("GetSectionReadersUsingsingleHandle() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
