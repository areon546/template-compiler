package compile

func OutputPath(internalPathToFile string) string {
	return directoryRoots["output"] + "/" + internalPathToFile
}

func ContentPath(internalPathToFile string) string {
	return directoryRoots["content"] + "/" + internalPathToFile
}
