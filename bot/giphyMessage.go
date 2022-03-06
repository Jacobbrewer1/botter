package bot

import (
	"fmt"
	"github.com/Jacobbrewer1/botter/api"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
)

func stickerMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if !message.isEmpty() {
		log.Println("Running search sticker")
		var sticker api.Search
		sticker, err := api.SearchQuery(api.StickerText, message.query)
		if err != nil {
			return err
		}
		if resp := api.DecodeResponse(sticker.Meta.Status); resp.IsOK() {
			if sticker.Pagination.TotalCount > 0 {
				if _, err := sendMessage(s, m.ChannelID, chooseRandomSearchGif(sticker), poweredByGiphyResponse); err != nil {
					return err
				}
			} else {
				log.Println("Response from giphy API was nil")
				if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(stickerCommand.secondResponse, message.query)); err != nil {
					return err
				}
				return nil
			}
		} else {
			if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(apiResponseCodeText, message.query, resp.Code)); err != nil {
				return err
			}
			return nil
		}
	} else {
		log.Println("Running trending sticker")
		var sticker api.Trending
		sticker, err := api.TrendingSearch(api.StickerText)
		if err != nil {
			return err
		}
		if resp := api.DecodeResponse(sticker.Meta.Status); resp.IsOK() {
			if sticker.Pagination.TotalCount > 0 {
				if _, err := sendMessage(s, m.ChannelID, chooseRandomTrendingGif(sticker), poweredByGiphyResponse); err != nil {
					return err
				}
			} else {
				log.Println("Response from giphy API was nil")
				if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(stickerCommand.secondResponse, gifTrendingText)); err != nil {
					return err
				}
				return nil
			}
		} else {
			if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(apiResponseCodeText, gifTrendingText, resp.Code)); err != nil {
				return err
			}
			return nil
		}
	}
	if err := deleteMessage(s, m); err != nil {
		return err
	}
	return nil
}

func gifMessage(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if !message.isEmpty() {
		log.Println("Running search gif")
		var gif api.Search
		gif, err := api.SearchQuery(api.GifText, message.query)
		if err != nil {
			return err
		}
		if resp := api.DecodeResponse(gif.Meta.Status); resp.IsOK() {
			if gif.Pagination.TotalCount > 0 {
				if _, err := sendMessage(s, m.ChannelID, chooseRandomSearchGif(gif), poweredByGiphyResponse); err != nil {
					return err
				}
			} else {
				log.Println("Response from giphy API was nil")
				if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(gifCommand.secondResponse, message.query)); err != nil {
					return err
				}
				return nil
			}
		} else {
			if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(apiResponseCodeText, message.query, resp.Code)); err != nil {
				return err
			}
			return nil
		}
	} else {
		log.Println("Running trending gif")
		var gif api.Trending
		gif, err := api.TrendingSearch(api.GifText)
		if err != nil {
			return err
		}
		if resp := api.DecodeResponse(gif.Meta.Status); resp.IsOK() {
			if gif.Pagination.TotalCount > 0 {
				if _, err := sendMessage(s, m.ChannelID, chooseRandomTrendingGif(gif), poweredByGiphyResponse); err != nil {
					return err
				}
			} else {
				log.Println("Response from giphy API was nil")
				if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(gifCommand.secondResponse, gifTrendingText)); err != nil {
					return err
				}
				return nil
			}
		} else {
			if _, err := sendMessage(s, m.ChannelID, fmt.Sprintf(apiResponseCodeText, gifTrendingText, resp.Code)); err != nil {
				return err
			}
			return nil
		}
	}
	if err := deleteMessage(s, m); err != nil {
		return err
	}
	return nil
}

func chooseRandomSearchGif(gifs api.Search) string {
	if len(gifs.Data) == 1 {
		return gifs.Data[0].URL
	}
	rndInt := rand.Intn(len(gifs.Data))
	return gifs.Data[rndInt].URL
}

func chooseRandomTrendingGif(gifs api.Trending) string {
	if len(gifs.Data) == 1 {
		return gifs.Data[0].URL
	}
	rndInt := rand.Intn(len(gifs.Data) - 1)
	return gifs.Data[rndInt].URL
}
