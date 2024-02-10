package file

import "github.com/constantincuy/knowledgestore/internal/core/domain/common"

type File struct {
	Id       common.Id
	Path     Path
	Provider Provider
	Created  Created
	Updated  Updated
}
