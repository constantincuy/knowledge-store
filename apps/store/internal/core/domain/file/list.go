package file

type List map[string]File

func NewList(files []File) (List, error) {
	list := make(List, len(files))
	for _, f := range files {
		list[string(f.Path)] = f
	}

	return list, nil
}
