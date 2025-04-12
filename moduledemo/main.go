package main

import "moduledemo/author"

func main() {
	authorInstance := author.NewAuthor("John Doe", "john.doe@example.com")

	authorInstance.WriteChapter("Chapter 1", "Lorem ipsum dolor sit amet.")
}
