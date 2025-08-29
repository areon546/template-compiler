package dirs

func CleanPath(path string) string {
	if len(path) == 0 {
		path = "./"
	}

	if path[len(path)-1] != '/' {
		path += "/"
	}

	return path
}
