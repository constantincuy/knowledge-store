package file

import "time"

type Filesystem struct {
	files List
}

func (f Filesystem) Sync(files List) (ChangeList, error) {
	created := make(List)
	updated := make(List)
	for key, file := range files {
		old, ex := f.files[key]
		if !ex {
			created[key] = file
		} else {
			t := time.Time(old.Updated)
			if t.Before(time.Time(file.Updated)) {
				file.Id = old.Id
				updated[key] = file
			}
		}
	}

	deleted := make(List)
	for key, file := range f.files {
		_, ex := files[key]
		if !ex {
			deleted[key] = file
		}
	}

	return ChangeList{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}, nil
}

func NewFilesystem(files List) Filesystem {
	return Filesystem{files}
}
