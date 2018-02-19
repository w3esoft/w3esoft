package tool

import (
	"path/filepath"
	"strings"
	"regexp"
	"os"
	"io/ioutil"
)

const (
	PB_EXTENSION = 2
	PB_ALLPATH   = 3
	PB_ANYPATH   = 4
	PB_PATH      = 5
)

type PathBuild struct {
	Index int
	Value string
}

func ignore(test string, bExcludes [][]*PathBuild) bool {

	var _ignore func(test string, bExclude []*PathBuild) bool;
	_ignore = func(test string, bExclude []*PathBuild) bool {
		if len(bExclude) == 0 {
			return false;
		}
		var bpath = bExclude[0:1][0]
		bExclude = bExclude[1:]
		var LL = len(bExclude)
		var reSplit = regexp.MustCompile("[\\/\\\\]")
		if bpath.Index == PB_PATH {
			b := strings.Index(filepath.Clean(test), filepath.Clean(bpath.Value)) == 0
			if (LL > 0 && b) {
				var s string = ""
				if len(test) > len(bpath.Value) {
					s = (test)[len(bpath.Value)+1:]
				}
				b = _ignore(s, bExclude)
			}
			return b
		} else if bpath.Index == PB_ANYPATH {
			if LL == 0 {
				return true;
			}
			ps := reSplit.Split(test, -1)
			if !(len(ps) > 1) {
				return true;
			}
			ps = ps[1:]
			_ignore(strings.Join(ps, "/"), bExclude)
		} else if bpath.Index == PB_ALLPATH {
			if LL == 0 {
				return true;
			}

			ps := reSplit.Split(test, -1)
			var psr []string = []string{};
			var a = ""
			var b = ""
			for _, p := range ps {
				a = a + b + p;
				b = "/"
				psr = append([]string{a}, psr...)
			}
			var c = false;
			for _, p := range psr {
				var s = ""
				if len(test) > len(p) {
					s = (test)[len(p)+1:]
				}
				c = _ignore(s, bExclude);
				if (c) {
					break
				}
			}
			return c;
		} else if bpath.Index == PB_EXTENSION {
			if strings.HasSuffix(test, bpath.Value) {
				return true;
			}
			return false;
		}
		return false;
	}
	var c = false;
	for _, bbE := range bExcludes {
		c = _ignore(test, bbE);
		if (c) {
			break;
		}
	}
	return c;
}

func Find(ROOTDIR string, _files []string, _excludes []string) []string{
	ROOTDIR = filepath.Clean(ROOTDIR)
	bFiles := createBuild(_files);
	bExcludes := createBuild(_excludes);
	var files = []string{};
	for _, bFile := range bFiles {
		var fn func(rootDir string, bFile []*PathBuild, bExcludes [][]*PathBuild);
		fn = func(parentDir string, bFile []*PathBuild, bExcludes [][]*PathBuild) {
			if len(bFile) == 0 {
				return;
			}
			var rPath = []string{};
			var bpath = bFile[0:1][0]
			bFile = bFile[1:]
			var LL = len(bFile)
			if bpath.Index == PB_PATH {
				rpath := filepath.Join(parentDir, bpath.Value)
				rFile := filepath.Join(ROOTDIR, rpath)
				if !ignore(rpath, bExcludes) {
					if _, err := os.Stat(rFile); err == nil {
						rPath = append(rPath, rpath)
					}
				}
			} else if bpath.Index == PB_ALLPATH {
				var fnRecursive func(p string);
				fnRecursive = func(parentDir string) {
					rFile := filepath.Join(ROOTDIR, parentDir)
					f, err := os.Stat(rFile);
					if err == nil {
						if f.IsDir() {
							if !ignore(parentDir, bExcludes) {
								if LL > 0 {
									rPath = append(rPath, parentDir)

								}
								fileInfoArray, e := ioutil.ReadDir(rFile);
								if e == nil {
									for _, fileInfo := range fileInfoArray {
										rF := filepath.Join(parentDir, fileInfo.Name())
										fnRecursive(rF);
									}
								}
							}
						} else {
							if LL == 0 {
								if !ignore(parentDir, bExcludes) {
									rPath = append(rPath, parentDir)
								}
							}
						}
					}

				}
				fnRecursive(parentDir);
				//
				//if _, err := os.Stat(rFile); os.IsNotExist(err) {
				//	files =append(files,rpath)
				//}

			} else if bpath.Index == PB_ANYPATH {

			} else if bpath.Index == PB_EXTENSION {
				rFile := filepath.Join(ROOTDIR, parentDir)
				f, err := os.Stat(rFile);
				if err == nil {
					if f.IsDir() {
						fileInfoArray, e := ioutil.ReadDir(rFile);
						if e == nil {
							for _, fileInfo := range fileInfoArray {
								if (!fileInfo.IsDir()) {
									var name = fileInfo.Name()
									if (strings.HasSuffix(name, bpath.Value) ) {
										rF := filepath.Join(parentDir, fileInfo.Name())
										if !ignore(rF, bExcludes) {

											rPath = append(rPath, rF)
										}
									}
								}
							}
						}
					}
				}

			}
			if (LL > 0) {
				for _, file := range rPath {
					fn(file, bFile, bExcludes);
				}
			} else {
				for _, file := range rPath {
					files = append(files, file)
				}
			}
		}
		fn("", bFile, bExcludes);
	}
	return  files
}


func createBuild(files []string) [][]*PathBuild {
	itemX := [][]*PathBuild{}
	for _, file := range files {
		paths := strings.Split(file, "/")
		itemY := []*PathBuild{}
		var ii *PathBuild = nil
		for _, p := range paths {
			var reRecursiveextension = regexp.MustCompile("^(\\*\\*)\\.")
			var reExtension = regexp.MustCompile("^(\\*)\\.")
			var reAllpath = regexp.MustCompile("^(\\*\\*)$")
			var reAnypath = regexp.MustCompile("^(\\*)$")
			if reRecursiveextension.MatchString(p) {
				ii = nil
				itemY = append(itemY, &PathBuild{
					Index: PB_ALLPATH,
				})
				v := reRecursiveextension.ReplaceAllString(p, "")
				itemY = append(itemY, &PathBuild{
					Index: PB_EXTENSION,
					Value: v,
				})
			} else if reExtension.MatchString(p) {
				ii = nil
				itemY = append(itemY, &PathBuild{
					Index: PB_EXTENSION,
					Value: reExtension.ReplaceAllString(p, ""),
				})
			} else if (reAllpath.MatchString(p)) {
				ii = nil
				itemY = append(itemY, &PathBuild{
					Index: PB_ALLPATH,
				})
			} else if (reAnypath.MatchString(p)) {
				ii = nil
				itemY = append(itemY, &PathBuild{
					Index: PB_ANYPATH,
				})
			} else {
				if ii == nil {
					PP := &PathBuild{
						Index: PB_PATH,
						Value: p,
					}
					itemY = append(itemY, PP)
					ii = PP
				} else {
					ii.Value = ii.Value + "/" + p
				}
			}
		}
		itemX = append(itemX, itemY)
	}
	return itemX

}
