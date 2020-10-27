package fs

type directory struct {
	name     string
	children []Inode
}

func (dir directory) GetType() InodeObjectType {
	return InodeTypeDirectory
}

func (dir directory) GetName() string {
	return dir.name
}

func (dir directory) GetSize() int {
	return 0
}

func (dir directory) GetChildren() []Inode {
	return dir.children
}
