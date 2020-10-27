package fs

type file struct {
	name string
	size int
}

func (f file) GetType() InodeObjectType {
	return InodeTypeFile
}

func (f file) GetName() string {
	return f.name
}

func (f file) GetSize() int {
	return f.size
}

func (f file) GetChildren() []Inode {
	return nil
}
