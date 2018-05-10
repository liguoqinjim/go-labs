package main

import (
	"github.com/beevik/etree"
	"os"
)

func main() {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	doc.CreateProcInst("xml-stylesheet", `type="text/xsl" href="style.xsl"`)

	config := doc.CreateElement("config")

	fullName := config.CreateElement("FullName")
	fullName.CreateCharData("Grace R. Emlin")

	company := config.CreateElement("Company")
	company.CreateCharData("Example Inc.")

	email1 := config.CreateElement("Email")
	email1.CreateAttr("where", "home")
	addr1 := email1.CreateElement("Addr")
	addr1.CreateCharData("gre@example.com")

	email2 := config.CreateElement("Email")
	email2.CreateAttr("where", "work")
	addr2 := email2.CreateElement("Addr")
	addr2.CreateCharData("gre@work.com")

	group := config.CreateElement("Group")
	value1 := group.CreateElement("Value")
	value1.CreateCharData("Friends")
	value2 := group.CreateElement("Value")
	value2.CreateCharData("Squash")

	city := config.CreateElement("City")
	city.CreateCharData("Hanga Roa")

	state := config.CreateElement("State")
	state.CreateCharData("Easter Island")

	doc.Indent(4)
	doc.WriteTo(os.Stdout)
}
