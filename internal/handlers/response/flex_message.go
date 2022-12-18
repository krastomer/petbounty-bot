package response

import (
	"fmt"

	"github.com/krastomer/petbounty-bot/internal/repositories/bounty"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func MapBountyToFlexMessage(bounty bounty.Bounty) linebot.FlexContainer {
	return &linebot.BubbleContainer{
		Type:      linebot.FlexContainerTypeBubble,
		Direction: linebot.FlexBubbleDirectionTypeLTR,
		Hero: &linebot.ImageComponent{
			Type:        linebot.FlexComponentTypeImage,
			URL:         "https://ichef.bbci.co.uk/news/800/cpsprodpb/1124F/production/_119932207_indifferentcatgettyimages.png",
			Size:        linebot.FlexImageSizeTypeFull,
			AspectRatio: linebot.FlexImageAspectRatioType16to9,
			AspectMode:  linebot.FlexImageAspectModeTypeCover,
		},
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    bounty.Name,
					Size:    linebot.FlexTextSizeTypeXl,
					Gravity: linebot.FlexComponentGravityTypeCenter,
					Wrap:    true,
					Weight:  linebot.FlexTextWeightTypeBold,
				},
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Text:    fmt.Sprintf("Reward: %.2f Baht", bounty.Reward),
					Size:    linebot.FlexTextSizeTypeLg,
					Gravity: linebot.FlexComponentGravityTypeCenter,
					Color:   "#999999",
				},
				BulletComponent("Detail", bounty.Detail),
				BulletComponent("Address", bounty.Address),
				BulletComponent("Tel.", bounty.Telephone),
				BulletComponent("Status", string(bounty.Status)),
				&linebot.TextComponent{},
			},
			Spacing: linebot.FlexComponentSpacingTypeMd,
		},
	}
}

func BulletComponent(title string, detail string) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:    linebot.FlexComponentTypeBox,
		Layout:  linebot.FlexBoxLayoutTypeVertical,
		Spacing: linebot.FlexComponentSpacingTypeSm,
		Margin:  linebot.FlexComponentMarginTypeLg,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:    linebot.FlexComponentTypeBox,
				Layout:  linebot.FlexBoxLayoutTypeBaseline,
				Spacing: linebot.FlexComponentSpacingTypeSm,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  title,
						Size:  linebot.FlexTextSizeTypeSm,
						Color: "#AAAAAA",
						Flex:  linebot.IntPtr(1),
					},
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  detail,
						Flex:  linebot.IntPtr(4),
						Size:  linebot.FlexTextSizeTypeSm,
						Wrap:  true,
						Color: "#666666",
					},
				},
			},
		},
	}
}
