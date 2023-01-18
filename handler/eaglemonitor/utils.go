package eaglemonitor

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func getSite(data *discordgo.MessageEmbed) string {
	for _, field := range data.Fields {
		if field.Name == "Site" {
			return strings.ToLower(field.Value)
		}
	}
	return ""

}

// func getSize(data *discordgo.MessageEmbed) string {
// 	for _, field := range data.Fields {
// 		if field.Name == "Size" {
// 			return field.Value
// 		}
// 	}
// 	return ""
// }

func getPid(data *discordgo.MessageEmbed) string {
	for _, field := range data.Fields {
		if field.Name == "PID" {
			return strings.ToLower(field.Value)
		}
	}
	return ""
}
