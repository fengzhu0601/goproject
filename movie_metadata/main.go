package main

import (
	"movie_metadata/common"
	"movie_metadata/utils"
)

func main() {

	common.RenameFile(utils.DIR)
	common.CreateDir(utils.DIR)
	common.GenMoviesNfo(utils.DIR)

	//utils.MoveFilesInDirectory(utils.DIR)

	//utils.GetMovieInfo(utils.DIR+"/"+"STARS-236"+"/", "STARS-236", "//avmoo.online/cn/movie/7194ccaf2da7387f")
}
