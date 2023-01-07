package eaglemonitor

import "github.com/bwmarrin/discordgo"

func getStore(data *discordgo.MessageEmbed) string {
	for _, field := range data.Fields {
		if field.Name == "Store" {
			return field.Value
		}
	}
	return ""

}

func getSize(data *discordgo.MessageEmbed) string {
	for _, field := range data.Fields {
		if field.Name == "Size" {
			return field.Value
		}
	}
	return ""
}

func getPid(data *discordgo.MessageEmbed) string {
	for _, field := range data.Fields {
		if field.Name == "PID" {
			return field.Value
		}
	}
	return ""
}
