package tool

import (
	"path"
	"log"
	"strings"
	"regexp"
)

const (
	PB_RECURSIVEEXTENSION = 1
	PB_EXTENSION          = 2
	PB_ALLPATH            = 3
	PB_ANYPATH            = 4
	PB_PATH               = 5
)

type PathBuild struct {
	index int
	value string
}

func FindFiles(rootDir string, files []string, excludes []string) {
	rootDir = path.Clean(rootDir)
	bFiles := createBuild(files);
	bExcludes := createBuild(excludes);

	fn := func(rootDir string, ) {
		updated = make([]byte, len(current))
		copy(updated, current)
		//modify updated
		return
	}

	findFiles(rootDir,bFiles,bExcludes);
}
//func findFiles ( rootDir string, bFiles[][]*PathBuild ,bExcludes[][]*PathBuild)  {
//	for _,bFile := range  bFiles{
//		find(rootDir,bFile,bExcludes);
//	}
//}

func createBuild(files []string)[][]*PathBuild {
	itemX := [][]*PathBuild{}
	for _, file := range files {
		paths := strings.Split(file, "/")
		itemY := []*PathBuild{}
		var ii *PathBuild=nil
		for _, p := range paths {
			var reRecursiveextension =  regexp.MustCompile("^(\\*\\*)\\.")
			var reExtension = regexp.MustCompile("^(\\*)\\.")
			var reAllpath = regexp.MustCompile("^(\\*\\*)$")
			var reAnypath = regexp.MustCompile("^(\\*)$")
			if reRecursiveextension.MatchString(p) {
				ii =nil
				itemY = append(itemY, &PathBuild{
					index: PB_RECURSIVEEXTENSION,
					value: reRecursiveextension.ReplaceAllString(p, ""),
				})
			} else if reExtension.MatchString(p) {
				ii =nil
				itemY = append(itemY, &PathBuild{
					index: PB_EXTENSION,
					value: reExtension.ReplaceAllString(p, ""),
				})
			} else if (reAllpath.MatchString(p)) {
				ii =nil
				itemY = append(itemY, &PathBuild{
					index: PB_ALLPATH,
				})
			} else if (reAnypath.MatchString(p)) {
				ii =nil
				itemY = append(itemY, &PathBuild{
					index: PB_ANYPATH,
				})
			}else {
				if ii ==nil{
					PP:=&PathBuild{
						index: PB_PATH,
						value:p,
					}
					itemY = append(itemY,PP)
					ii=PP
				}else {
					ii.value =ii.value+"/"+ p
				}
			}
		}
		itemX = append(itemX, itemY)
	}
	return itemX

}
