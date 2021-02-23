package main

import (
	"html/template"
	"time"

	"github.com/gomarkdown/markdown"
)

type Homepage struct {
	ID          int       `json:"id"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Seo         struct {
		ID              int    `json:"id"`
		MetaTitle       string `json:"metaTitle"`
		MetaDescription string `json:"metaDescription"`
		ShareImage      struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			AlternativeText string `json:"alternativeText"`
			Caption         string `json:"caption"`
			Width           int    `json:"width"`
			Height          int    `json:"height"`
			Formats         struct {
				Thumbnail struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"thumbnail"`
				Large struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"large"`
				Medium struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"medium"`
				Small struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"small"`
			} `json:"formats"`
			Hash             string      `json:"hash"`
			Ext              string      `json:"ext"`
			Mime             string      `json:"mime"`
			Size             float64     `json:"size"`
			URL              string      `json:"url"`
			PreviewURL       interface{} `json:"previewUrl"`
			Provider         string      `json:"provider"`
			ProviderMetadata interface{} `json:"provider_metadata"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
		} `json:"shareImage"`
	} `json:"seo"`
	Hero struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"hero"`
}

type Global struct {
	ID         int       `json:"id"`
	SiteName   string    `json:"siteName"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Title      string    `json:"title"`
	Language   string    `json:"language"`
	SiteURL    string    `json:"siteURL"`
	SiteRepo   string    `json:"siteRepo"`
	Email      string    `json:"email"`
	Github     string    `json:"github"`
	Twitter    string    `json:"twitter"`
	Facebook   string    `json:"facebook"`
	Youtube    string    `json:"youtube"`
	Linkedin   string    `json:"linkedin"`
	Twitch     string    `json:"twitch"`
	DefaultSeo struct {
		ID              int    `json:"id"`
		MetaTitle       string `json:"metaTitle"`
		MetaDescription string `json:"metaDescription"`
		ShareImage      struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			AlternativeText string `json:"alternativeText"`
			Caption         string `json:"caption"`
			Width           int    `json:"width"`
			Height          int    `json:"height"`
			Formats         struct {
				Thumbnail struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"thumbnail"`
				Large struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"large"`
				Medium struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"medium"`
				Small struct {
					Name   string      `json:"name"`
					Hash   string      `json:"hash"`
					Ext    string      `json:"ext"`
					Mime   string      `json:"mime"`
					Width  int         `json:"width"`
					Height int         `json:"height"`
					Size   float64     `json:"size"`
					Path   interface{} `json:"path"`
					URL    string      `json:"url"`
				} `json:"small"`
			} `json:"formats"`
			Hash             string      `json:"hash"`
			Ext              string      `json:"ext"`
			Mime             string      `json:"mime"`
			Size             float64     `json:"size"`
			URL              string      `json:"url"`
			PreviewURL       interface{} `json:"previewUrl"`
			Provider         string      `json:"provider"`
			ProviderMetadata interface{} `json:"provider_metadata"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
		} `json:"shareImage"`
	} `json:"defaultSeo"`
	Favicon struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		AlternativeText string `json:"alternativeText"`
		Caption         string `json:"caption"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
		Formats         struct {
			Thumbnail struct {
				Name   string      `json:"name"`
				Hash   string      `json:"hash"`
				Ext    string      `json:"ext"`
				Mime   string      `json:"mime"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Size   float64     `json:"size"`
				Path   interface{} `json:"path"`
				URL    string      `json:"url"`
			} `json:"thumbnail"`
			Large struct {
				Name   string      `json:"name"`
				Hash   string      `json:"hash"`
				Ext    string      `json:"ext"`
				Mime   string      `json:"mime"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Size   float64     `json:"size"`
				Path   interface{} `json:"path"`
				URL    string      `json:"url"`
			} `json:"large"`
			Medium struct {
				Name   string      `json:"name"`
				Hash   string      `json:"hash"`
				Ext    string      `json:"ext"`
				Mime   string      `json:"mime"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Size   float64     `json:"size"`
				Path   interface{} `json:"path"`
				URL    string      `json:"url"`
			} `json:"medium"`
			Small struct {
				Name   string      `json:"name"`
				Hash   string      `json:"hash"`
				Ext    string      `json:"ext"`
				Mime   string      `json:"mime"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Size   float64     `json:"size"`
				Path   interface{} `json:"path"`
				URL    string      `json:"url"`
			} `json:"small"`
		} `json:"formats"`
		Hash             string      `json:"hash"`
		Ext              string      `json:"ext"`
		Mime             string      `json:"mime"`
		Size             float64     `json:"size"`
		URL              string      `json:"url"`
		PreviewURL       interface{} `json:"previewUrl"`
		Provider         string      `json:"provider"`
		ProviderMetadata interface{} `json:"provider_metadata"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
	} `json:"favicon"`
}

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Slug        string    `json:"slug"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Category    struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		Slug        string      `json:"slug"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		Description string      `json:"description"`
		Image       interface{} `json:"image"`
	} `json:"category"`
	Section string `json:"section"`
	Image   Image  `json:"image"`
}

type ShortArticle struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Slug        string    `json:"slug"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Category    int       `json:"category"`
	Section     string    `json:"section"`
	Image       Image     `json:"image"`
}
type Image struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	AlternativeText string `json:"alternativeText"`
	Caption         string `json:"caption"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Formats         struct {
		Thumbnail struct {
			Name   string      `json:"name"`
			Hash   string      `json:"hash"`
			Ext    string      `json:"ext"`
			Mime   string      `json:"mime"`
			Width  int         `json:"width"`
			Height int         `json:"height"`
			Size   float64     `json:"size"`
			Path   interface{} `json:"path"`
			URL    string      `json:"url"`
		} `json:"thumbnail"`
		Large struct {
			Name   string      `json:"name"`
			Hash   string      `json:"hash"`
			Ext    string      `json:"ext"`
			Mime   string      `json:"mime"`
			Width  int         `json:"width"`
			Height int         `json:"height"`
			Size   float64     `json:"size"`
			Path   interface{} `json:"path"`
			URL    string      `json:"url"`
		} `json:"large"`
		Medium struct {
			Name   string      `json:"name"`
			Hash   string      `json:"hash"`
			Ext    string      `json:"ext"`
			Mime   string      `json:"mime"`
			Width  int         `json:"width"`
			Height int         `json:"height"`
			Size   float64     `json:"size"`
			Path   interface{} `json:"path"`
			URL    string      `json:"url"`
		} `json:"medium"`
		Small struct {
			Name   string      `json:"name"`
			Hash   string      `json:"hash"`
			Ext    string      `json:"ext"`
			Mime   string      `json:"mime"`
			Width  int         `json:"width"`
			Height int         `json:"height"`
			Size   float64     `json:"size"`
			Path   interface{} `json:"path"`
			URL    string      `json:"url"`
		} `json:"small"`
	} `json:"formats"`
	Hash             string      `json:"hash"`
	Ext              string      `json:"ext"`
	Mime             string      `json:"mime"`
	Size             float64     `json:"size"`
	URL              string      `json:"url"`
	PreviewURL       interface{} `json:"previewUrl"`
	Provider         string      `json:"provider"`
	ProviderMetadata interface{} `json:"provider_metadata"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

func (a Article) Markdown() template.HTML {

	return template.HTML(markdown.ToHTML([]byte(a.Content), nil, nil))
}

type Category struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Description string         `json:"description"`
	Image       Image          `json:"image"`
	Articles    []ShortArticle `json:"articles"`
}
