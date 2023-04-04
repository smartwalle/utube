package utube

import (
	"fmt"
	"net/url"
)

// GetPlaylistsParam https://developers.google.cn/youtube/v3/docs/playlists/list
type GetPlaylistsParam struct {
	// Required parameters
	Part Part

	// Filters (specify exactly one of the following parameters)
	ChannelId string
	Id        string
	Mine      int // authorized request 1 true, 2 false

	// Optional parameters
	HL                            string
	OnBehalfOfContentOwner        string // authorized request
	OnBehalfOfContentOwnerChannel string // authorized request

	MaxResults int
	PageToken  string
}

func (this GetPlaylistsParam) Values() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if len(this.ChannelId) > 0 {
		v.Add("channelId", this.ChannelId)
	}

	if len(this.Id) > 0 {
		v.Add("id", this.Id)
	}

	if this.Mine == 1 {
		v.Add("mine", "true")
	} else if (this.Mine) == 2 {
		v.Add("mine", "false")
	}

	if len(this.HL) > 0 {
		v.Add("hl", this.HL)
	}

	if len(this.OnBehalfOfContentOwner) > 0 {
		v.Add("onBehalfOfContentOwner", this.OnBehalfOfContentOwner)
	}

	if len(this.OnBehalfOfContentOwnerChannel) > 0 {
		v.Add("onBehalfOfContentOwnerChannel", this.OnBehalfOfContentOwnerChannel)
	}

	if this.MaxResults > 0 {
		v.Add("maxResults", fmt.Sprintf("%d", this.MaxResults))
	}

	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}

	return v
}

type Playlists struct {
	Kind          string      `json:"kind"`
	ETag          string      `json:"etag"`
	NextPageToken string      `json:"nextPageToken"`
	PrevPageToken string      `json:"prevPageToken"`
	RegionCode    string      `json:"regionCode"`
	PageInfo      *PageInfo   `json:"pageInfo"`
	Items         []*Playlist `json:"items,omitempty"`
}

type Playlist struct {
	Kind           string                  `json:"kind"`
	ETag           string                  `json:"etag"`
	Id             string                  `json:"id"`
	Snippet        *PlaylistSnippet        `json:"snippet,omitempty"`
	Status         *PlaylistSnippet        `json:"status"`
	ContentDetails *PlaylistContentDetails `json:"contentDetails"`
	Player         *Player                 `json:"player"`
	Localizations  *Localizations          `json:"localizations"`
}

type PlaylistSnippet struct {
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	ChannelId    string      `json:"channelId"`
	ChannelTitle string      `json:"channelTitle"`
	PublishedAt  string      `json:"publishedAt"`
	Thumbnails   *Thumbnails `json:"thumbnails"`
	Localized    *Localized  `json:"localized"`
}

type PlaylistStatus struct {
	PrivacyStatus string `json:"privacyStatus"`
}

type PlaylistContentDetails struct {
	ItemCount int `json:"itemCount"`
}

// GetPlaylistItemsParam https://developers.google.cn/youtube/v3/docs/playlistItems/list
type GetPlaylistItemsParam struct {
	// Required parameters
	Part Part

	// Filters (specify exactly one of the following parameters)
	Id         string
	PlaylistId string

	// Optional parameters
	VideoId                string
	OnBehalfOfContentOwner string // authorized request

	MaxResults int
	PageToken  string
}

func (this GetPlaylistItemsParam) Values() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if len(this.Id) > 0 {
		v.Add("id", this.Id)
	}

	if len(this.PlaylistId) > 0 {
		v.Add("playlistId", this.PlaylistId)
	}

	if len(this.VideoId) > 0 {
		v.Add("videoId", this.VideoId)
	}

	if len(this.OnBehalfOfContentOwner) > 0 {
		v.Add("onBehalfOfContentOwner", this.OnBehalfOfContentOwner)
	}

	if this.MaxResults > 0 {
		v.Add("maxResults", fmt.Sprintf("%d", this.MaxResults))
	}

	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}

	return v
}

type PlaylistItems struct {
	Kind          string          `json:"kind"`
	ETag          string          `json:"etag"`
	NextPageToken string          `json:"nextPageToken"`
	PrevPageToken string          `json:"prevPageToken"`
	RegionCode    string          `json:"regionCode"`
	PageInfo      *PageInfo       `json:"pageInfo"`
	Items         []*PlaylistItem `json:"items,omitempty"`
}

type PlaylistItem struct {
	Kind           string                      `json:"kind"`
	ETag           string                      `json:"etag"`
	Id             string                      `json:"id"`
	Snippet        *PlaylistItemSnippet        `json:"snippet,omitempty"`
	ContentDetails *PlaylistItemContentDetails `json:"contentDetails,omitempty"`
	Status         *PlaylistItemStatus         `json:"status,omitempty"`
}

type PlaylistItemSnippet struct {
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	ChannelId    string      `json:"channelId"`
	ChannelTitle string      `json:"channelTitle"`
	PublishedAt  string      `json:"publishedAt"`
	Thumbnails   *Thumbnails `json:"thumbnails"`
	Localized    *Localized  `json:"localized"`
	PlaylistId   string      `json:"playlistId"`
	Position     int         `json:"position"`
	ResourceId   *ResourceId `json:"resourceId"`
}

type PlaylistItemContentDetails struct {
	VideoId          string `json:"videoId"`
	VideoPublishedAt string `json:"videoPublishedAt"`
}

type PlaylistItemStatus struct {
	PrivacyStatus string `json:"privacyStatus"`
}
