package starter

import (
	"fmt"
	"forex/library/times"
	"forex/systems"
	"time"
)

type App struct {
	ModuleName   string
	ModuleID     int
	InUseService []string
	DebugMode    bool

	// basic setting
	RootPath         string
	MinimumGoVersion string

	FileEncryptKey string

	// upload setting
	UploadPath      string
	UploadSizeLimit int
	UploadFileTypes []string

	// download setting
	DownloadPath      string
	DownloadSizeLimit int
	DownloadFileTypes []string

	ThumbNailPath string
	ThumbNailSize int

	// file location setting
	FileLocationShiftInterval int // Hourly
}

func (m *App) Builder(c *Content) error {
	var start = []int{0, 0, 0, 0}
	var routineFunc = func() error {
		var err error
		var now = time.Now().Local()
		var upath = systems.ReplaceSplit(fmt.Sprintf("%s%s", m.RootPath, m.UploadPath))
		var dpath = systems.ReplaceSplit(fmt.Sprintf("%s%s", m.RootPath, m.DownloadPath))
		var name = ""
		_, err = systems.MustOpen(name, upath)
		Assert(err)
		_, err = systems.MustOpen(name, dpath)
		Assert(err)
		_, err = systems.MustOpen(name, dpath+time.Date(now.Year(), now.Month(), now.Day(), start[0], start[1], start[2], start[3], time.Local).Format("2006-01-02")+systems.GetSplit())
		Assert(err)
		_, err = systems.MustOpen(name, upath+time.Date(now.Year(), now.Month(), now.Day(), start[0], start[1], start[2], start[3], time.Local).Format("2006-01-02")+systems.GetSplit())
		Assert(err)
		_, err = systems.MustOpen(name, dpath+time.Date(now.Year(), now.Month(), now.Day(), start[0], start[1], start[2], start[3], time.Local).Format("2006-01-02")+systems.GetSplit()+m.ThumbNailPath)
		Assert(err)
		return err
	}
	go times.Routine(start, 10, 24, routineFunc)

	return nil
}

func (m *App) Starter(c *Content) error {
	return nil
}

func (m *App) Router(s *Server) {

}
