package components

func PrdimgHashPath(hash string) string {
	return hash[0:2] + "/" + hash[2:4]
}
