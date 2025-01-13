package file

import (
	"io"
	"os"
	"path/filepath"
	"sort"
)

func Copy(fromLocation string, toLocation string) error {
	var err error

	original, err := os.Open(fromLocation)
	if err != nil {
		return err
	}

	new, err := os.Create(toLocation)
	if err != nil {
		original.Close()
		return err
	}

	_, err = io.Copy(new, original)

	original.Close()
	new.Close()
	return err
}

type modTimeSorter []os.FileInfo

func (m modTimeSorter) Len() int {
	return len(m)
}

func (m modTimeSorter) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m modTimeSorter) Less(i, j int) bool {
	return m[i].ModTime().Before(m[j].ModTime())
}

// ListByModTime returns a list of files in the specified directory
// The files are returned in the order they were modified
func ListByModTime(dirname string) ([]string, error) {
	de, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	var mts modTimeSorter
	for _, d := range de {
		fi, _ := d.Info()
		mts = append(mts, fi)
	}

	sort.Sort(mts)

	var files []string
	for _, f := range mts {
		files = append(files, filepath.Join(dirname, f.Name()))
	}

	return files, nil

}
