/*
Copyright (C) MarckTomack <marcktomack@tutanota.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/marcktomack/aloneMP/app"
)

var (
	dir           = flag.Bool("dir", false, "Directory with audio files")
	knowExtension = [4]string{".mp3", ".wav", ".flac", ".ogg"}
)

func getFiles(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(err)
		}
		if !info.IsDir() && contains(knowExtension, filepath.Ext(path)) {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	return files
}

func contains(s [4]string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		u, _ := user.Current()
		musicDir := fmt.Sprintf("%s/Music", u.HomeDir)
		f, err := os.Stat(musicDir)
		if f.IsDir() {
			files := getFiles(musicDir)
			if len(files) != 0 {
				os.Chdir(musicDir)
				app.Run(files)
			} else {
				log.Fatalln("No audio file found")
			}
		} else if err != nil {
			log.Fatalln(err)
		} else {
			log.Fatalln("Music directory does not exist")
		}
	}
	if *dir {
		userDir := os.Args[2]
		f, err := os.Stat(userDir)
		if f.IsDir() {
			files := getFiles(userDir)
			if len(files) != 0 {
				os.Chdir(userDir)
				app.Run(files)
			} else {
				log.Fatalln("No audio file found")
			}
		} else if err != nil {
			log.Fatalln(err)
		} else {
			log.Fatalf("Directory %s does not exist", userDir)
		}

	}
}
