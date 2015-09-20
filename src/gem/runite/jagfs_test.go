package runite

import (
	"strconv"
	"testing"
)

func TestJagFS(t *testing.T) {
	// paths seem to be relative to source file..
	// need to either create a smaller set of test data and put it in this dir,
	// or find a better way to reference the main game data
	basePath := "../../../data"
	idxBase := basePath + "/main_file_cache.idx"
	idxFiles := make([]string, 5)
	for i := 0; i < len(idxFiles); i++ {
		idxFiles[i] = idxBase + strconv.Itoa(i)
	}

	fs, err := UnpackJagFSFiles(basePath+"/main_file_cache.dat", idxFiles)
	if err != nil {
		t.Fatalf("Couldn't unpack file system: %v", err)
	}

	if fs.IndexCount() != 5 {
		t.Errorf("Invalid index count %v", fs.IndexCount())
	}

/*
	for x := 0; x < fs.IndexCount(); x++ {
		idx, err := fs.Index(x)
		if err != nil {
			t.Fatalf("Couldn't load index %v: %v", x, err)
		}

		t.Logf("File count: %v", idx.FileCount())

		for i := 0; i < idx.FileCount(); i++ {
			data, err := idx.File(i)
			if err != nil {
				t.Errorf("Couldn't load index %v file %v: %v", x, i, err)
			}

			ioutil.WriteFile(basePath+"/dump/"+strconv.Itoa(x)+"/"+strconv.Itoa(i), data, 777)
		}
	}
*/
}
