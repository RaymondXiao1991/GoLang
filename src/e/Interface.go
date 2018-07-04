package e

import (
	"io"
	"time"
)

/*
一个具体的类型可能实现了很多不相关的接口。考虑在一个组织出售数字文化产品比如音乐，电影和书籍的程序中可能定义了下列的具体类型：
Album
Book
Movie
Magazine
Podcast
TVEpisode
Track
*/

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}
type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}
