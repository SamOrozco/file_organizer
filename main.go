package main

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Organizer interface {
	Organize(file os.FileInfo, rootLocation string) error
}

var SeparatorString = string(filepath.Separator)

var (
	OrganizeByType        = false
	OrganizeByTypeAndDate = false
	OrganizeByDateAndType = false
	ExcludePattern        = ""
	FolderSuffix          = ""
	SampleMillis          = 500
	rootCmd               = &cobra.Command{
		Use:   "fo",
		Short: "File Organizer",
		Long:  `Watcher a directory and organizes your files. Defaults to organizing by type`,
		Run: func(cmd *cobra.Command, args []string) {
			watchFile(args[0], ExcludePattern, FolderSuffix, SampleMillis, OrganizeByType, OrganizeByTypeAndDate, OrganizeByDateAndType)
		},
	}
)

func main() {
	rootCmd.PersistentFlags().BoolVarP(&OrganizeByType, "organize-by-type", "t", false, "this will organize a parent folder into sub folders based on file extension")
	rootCmd.PersistentFlags().BoolVarP(&OrganizeByTypeAndDate, "organize-by-type-and-date", "l", false, "this will organize a parent folder into sub folders of types and then sub-sub folders of date added")
	rootCmd.PersistentFlags().BoolVarP(&OrganizeByDateAndType, "organize-by-date-and-type", "d", false, "this will organize a parent folder into sub folders of dates added and then sub-sub folders of type")
	rootCmd.PersistentFlags().IntVarP(&SampleMillis, "sample-millis", "m", 1000, "how often the watcher will check for new files and move them")
	rootCmd.PersistentFlags().StringVarP(&ExcludePattern, "exclude-pattern", "x", "", "if set will exclude file from being organized if it matches regex pattern")
	rootCmd.PersistentFlags().StringVarP(&FolderSuffix, "type-folder-suffix", "f", "_files", "for the type created folders this will be the suffix")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func watchFile(location, excludePattern, typeFolderSuffix string, sampleMillis int, organizeByType, organizeByTypeAndDate, organizeByDateAndType bool) {
	// create watching dir if not exists
	if err := CreateDirIfNotExists(location); err != nil {
		panic(err)
	}

	// create our exclude predicate
	excludePredicate := createExcludePredicate(excludePattern)

	// setup organizer
	var organizer Organizer
	if organizeByType {
		organizer = &TypeOrganizer{
			FolderSuffix: typeFolderSuffix,
		}
	} else if organizeByTypeAndDate {
		organizer = &TypeDateOrganizer{
			FolderSuffix: typeFolderSuffix,
		}
	} else if organizeByDateAndType {
		organizer = &DateTypeOrganizer{
			FolderSuffix: typeFolderSuffix}
	}

	// watch files
	for {
		files, err := ioutil.ReadDir(location)
		if err != nil {
			panic(err)
		}

		// look through files
		for i := range files {
			currentFileInfo := files[i]
			// exclude file
			if excludePredicate(currentFileInfo.Name()) || currentFileInfo.IsDir() {
				continue
			}

			// organize file
			if err = organizer.Organize(currentFileInfo, location); err != nil {
				log.Error(err)
			}
		}

		// wait for next sample
		<-time.After(time.Millisecond * time.Duration(sampleMillis))
	}

}

func createExcludePredicate(excludePattern string) func(string) bool {
	if len(excludePattern) > 0 {
		regex, _ := regexp.Compile(excludePattern)
		return func(val string) bool {
			return regex.Match([]byte(val))
		}
	}
	// no op - no excluding
	return func(val string) bool {
		return false
	}
}
