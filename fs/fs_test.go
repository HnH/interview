package fs

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilesystemSuite struct {
	fs []Inode

	suite.Suite
}

func TestFilesystemSuite(t *testing.T) {
	suite.Run(t, new(FilesystemSuite))
}

func (suite *FilesystemSuite) SetupSuite() {
	suite.fs = []Inode{
		directory{
			name: "/first",
			children: []Inode{
				directory{
					name: "/first/inner",
					children: []Inode{
						directory{
							name: "/first/inner/directory",
							children: []Inode{
								directory{
									name: "/first/inner/directory/sub",
									children: []Inode{
										file{
											name: "/first/inner/directory/sub/aerosmith-dream-on.mp3",
											size: 1024 * 1024 * 6,
										},
									},
								},
								file{
									name: "/first/inner/directory/aerosmith-dream-on.mp3",
									size: 1024 * 1024 * 4,
								},
							},
						},
						directory{
							name: "/first/inner/music",
							children: []Inode{
								file{
									name: "/first/inner/music/metallica-nothing-else-matters.mp3",
									size: 1024 * 1024 * 6,
								},
								file{
									name: "/first/inner/music/metallica-unforgiven.mp3",
									size: 1024 * 1024 * 5,
								},
								file{
									name: "/first/inner/music/metallica-enter-sandman.mp3",
									size: 1024*1024*5 - 1,
								},
							},
						},
						file{
							name: "/first/inner/linking-park-nothing-else-matters.mp3",
							size: 1024 * 1024 * 7,
						},
					},
				},
				directory{
					name: "/first/second",
					children: []Inode{
						directory{
							name:     "/first/second/third",
							children: []Inode{},
						},
					},
				},
			},
		},
		directory{
			name: "/downloads",
			children: []Inode{
				file{
					name: "/downloads/lecture.mp3",
					size: 1024 * 1024 * 15,
				},
				file{
					name: "/downloads/lecture.pdf",
					size: 1024 * 1024 * 15,
				},
				file{
					name: "/downloads/lecture.ppt",
					size: 1024 * 1024 * 15,
				},
				file{
					name: "/downloads/lecture.mp33",
					size: 1024 * 1024 * 15,
				},
			},
		},
		directory{
			name: "/documents",
			children: []Inode{
				file{
					name: "/documents/book.pdf",
					size: 1024 * 1024 * 15,
				},
				file{
					name: "/documents/book2.fb2",
					size: 1024 * 1024 * 15,
				},
			},
		},
		directory{
			name:     "/home",
			children: nil,
		},
		file{
			name: "/the-cranberries-zombie.mp3",
			size: 1024 * 1024 * 6,
		},
		file{
			name: "/scorpions-white-dove.mp3",
			size: 1024 * 1024 * 4,
		},
	}
}

func (suite *FilesystemSuite) TestFilter() {
	var (
		output   = filter(suite.fs)
		expected = []string{
			"/first/inner/directory/sub/aerosmith-dream-on.mp3",
			"/first/inner/music/metallica-nothing-else-matters.mp3",
			"/first/inner/music/metallica-unforgiven.mp3",
			"/first/inner/linking-park-nothing-else-matters.mp3",
			"/downloads/lecture.mp3",
			"/the-cranberries-zombie.mp3",
		}
	)

	suite.Require().NotNil(output)
	suite.Require().Equal(len(expected), len(output))
	//suite.Require().Equal(expected, output)

	var contains bool
	for _, f := range expected {
		contains = false
		for _, got := range output {
			if got == f {
				contains = true
			}
		}

		if !contains {
			suite.T().Errorf("file '%s' not found in output", f)
		}
	}
}
