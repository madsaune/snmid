package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/madsaune/snmid/mid"
)

type CLIArgs struct {
	URLOnly    bool
	Platform   string
	FileType   string
	BuildStamp string
	Output     string
}

func main() {
	var opts CLIArgs

	flag.Usage = func() {
		fmt.Println("Usage: snmid <buildstamp> [flags]")
		fmt.Println("")
		fmt.Println("  --url-only\tonly print the download url")
		fmt.Println("  -p\t\ttarget platform (windows, linux)")
		fmt.Println("  -t\t\tinstaller filetype (msi, zip)")
		fmt.Println("  -o\t\toutput file")
	}
	flag.BoolVar(&opts.URLOnly, "url-only", false, "only print the download url")
	flag.StringVar(&opts.Platform, "p", "windows", "target platform (windows, linux)")
	flag.StringVar(&opts.FileType, "t", "msi", "installer filetype (msi, zip)")
	flag.StringVar(&opts.Output, "o", "", "output file")
	opts.BuildStamp = flag.Arg(0)
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	md := mid.New(opts.Platform, opts.FileType, opts.BuildStamp)

	if opts.URLOnly {
		fmt.Println(md.URL())
		os.Exit(0)
	}

	err := md.Download(opts.Output)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

// "sandiego-12-22-2021__patch9a-hotfix1-01-31-2023_02-01-2023_1625"
