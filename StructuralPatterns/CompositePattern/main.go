package main

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{name: "Folder1"}
	folder2 := &folder{name: "Folder2"}

	folder1.add(file1)

	folder2.add(file2)
	folder1.add(folder2)

	folder1.add(file3)

	folder1.search("Marigold")

}
