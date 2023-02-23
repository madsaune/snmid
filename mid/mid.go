package mid

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cheggaaa/pb/v3"
)

const MIDServerPackageBaseURL = "https://install.service-now.com/glide/distribution/builds/package"

// MIDDownloader handles parsing and downloading of ServiceNow MID Server installer packages
type MIDDownloader struct {
	// Which platform the installer is meant for (windows, linux)
	OperatingSystem string

	// The installer can either be an msi or zip file
	FileType string

	BuildStamp BuildStamp
}

type BuildStamp struct {
	Stamp string
	Year  string
	Month string
	Day   string
}

// New creates a new instance of MIDDownloader
func New(operatingsystem, filetype, buildstamp string) *MIDDownloader {
	md := &MIDDownloader{
		OperatingSystem: operatingsystem,
		FileType:        filetype,
	}
	md.parse(buildstamp)

	return md
}

// Download fetches the installer file from ServiceNow
func (md *MIDDownloader) Download(output string) error {
	u := md.URL()
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%d: failed to download file", res.StatusCode)
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}

	bar := pb.Simple.Start64(res.ContentLength)
	barReader := bar.NewProxyReader(res.Body)

	_, err = io.Copy(f, barReader)
	if err != nil {
		f.Close()
		return err
	}
	f.Close()

	bar.Finish()

	return nil
}

// URL prints the installer download URL
func (md *MIDDownloader) URL() string {
	var endpoint string
	var fileprefix string
	switch md.OperatingSystem {
	case "windows":
		fileprefix = "mid-windows-installer"
		endpoint = "/app-signed/mid-windows-installer"
	default:
		fileprefix = "mid"
		endpoint = "/mid"
	}

	baseURL := MIDServerPackageBaseURL + endpoint
	return fmt.Sprintf(
		"%s/%s/%s/%s/%s.%s.%s.x86-64.%s",
		baseURL,
		md.BuildStamp.Year,
		md.BuildStamp.Month,
		md.BuildStamp.Day,
		fileprefix,
		md.BuildStamp.Stamp,
		md.OperatingSystem,
		md.FileType,
	)
}

func (md *MIDDownloader) parse(bs string) {
	s := strings.Split(bs, "_")
	d := s[3]
	ds := strings.Split(d, "-")

	md.BuildStamp.Stamp = bs
	md.BuildStamp.Year = ds[2]
	md.BuildStamp.Month = ds[0]
	md.BuildStamp.Day = ds[1]
}
