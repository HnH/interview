package fs

type InodeObjectType int

const (
	InodeTypeFile InodeObjectType = iota + 1
	InodeTypeDirectory
)

type Inode interface {
	GetType() InodeObjectType
	GetName() string
	GetSize() int
	GetChildren() []Inode
}

func filter(in ...Inode) []string {
	return []string{}
}
