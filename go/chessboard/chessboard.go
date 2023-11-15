package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	currentFile := cb[file]
	
	for _, v := range currentFile {
		if v {
			count++
		}
	}
	return count
	//panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	
	if rank <= 0 || rank > len(cb) {
		return 0
	}
	
	for _, v := range cb {
		if v[rank - 1] {
			count++
		}
	}
	return count
	//panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	
	count := 0
	
	for _,v := range cb {
		count += len(v)
	}
	
	return count
	//panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	
	count := 0
	
	for _, files := range cb {
		for _, space := range files {
			if space {
				count++
			}
		}
	}
	
	return count
	//panic("Please implement CountOccupied()")
}
