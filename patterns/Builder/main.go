package main

import "example.com/builder"

func main() {
	GamingBuilder := builder.GetBuilder("GamingPC")
	OfficeBuilder := builder.GetBuilder("OfficePC")

	Dir := builder.NewDir(GamingBuilder)
	GamingPC := Dir.CreatePC()
	GamingPC.PrintConfig()

	Dir.SetBuilder(OfficeBuilder)
	OfficePC := Dir.CreatePC()
	OfficePC.PrintConfig()
}
