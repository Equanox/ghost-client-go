package content

import (
	"html/template"
	"time"
)

type Post struct {
	Access  bool `json:"access"`
	Authors []struct {
		Bio             interface{} `json:"bio"`
		CoverImage      interface{} `json:"cover_image"`
		Facebook        interface{} `json:"facebook"`
		ID              string      `json:"id"`
		Location        interface{} `json:"location"`
		MetaDescription interface{} `json:"meta_description"`
		MetaTitle       interface{} `json:"meta_title"`
		Name            string      `json:"name"`
		ProfileImage    interface{} `json:"profile_image"`
		Slug            string      `json:"slug"`
		Twitter         interface{} `json:"twitter"`
		URL             string      `json:"url"`
		Website         interface{} `json:"website"`
	} `json:"authors"`
	CanonicalURL         interface{}   `json:"canonical_url"`
	CodeinjectionFoot    interface{}   `json:"codeinjection_foot"`
	CodeinjectionHead    interface{}   `json:"codeinjection_head"`
	CommentID            string        `json:"comment_id"`
	CreatedAt            time.Time     `json:"created_at"`
	CustomExcerpt        string        `json:"custom_excerpt"`
	CustomTemplate       interface{}   `json:"custom_template"`
	EmailRecipientFilter string        `json:"email_recipient_filter"`
	EmailSubject         interface{}   `json:"email_subject"`
	Excerpt              string        `json:"excerpt"`
	FeatureImage         interface{}   `json:"feature_image"`
	Featured             bool          `json:"featured"`
	HTML                 template.HTML `json:"html"`
	ID                   string        `json:"id"`
	MetaDescription      interface{}   `json:"meta_description"`
	MetaTitle            interface{}   `json:"meta_title"`
	OgDescription        interface{}   `json:"og_description"`
	OgImage              interface{}   `json:"og_image"`
	OgTitle              interface{}   `json:"og_title"`
	PrimaryAuthor        struct {
		Bio             interface{} `json:"bio"`
		CoverImage      interface{} `json:"cover_image"`
		Facebook        interface{} `json:"facebook"`
		ID              string      `json:"id"`
		Location        interface{} `json:"location"`
		MetaDescription interface{} `json:"meta_description"`
		MetaTitle       interface{} `json:"meta_title"`
		Name            string      `json:"name"`
		ProfileImage    interface{} `json:"profile_image"`
		Slug            string      `json:"slug"`
		Twitter         interface{} `json:"twitter"`
		URL             string      `json:"url"`
		Website         interface{} `json:"website"`
	} `json:"primary_author"`
	PrimaryTag struct {
		AccentColor        interface{} `json:"accent_color"`
		CanonicalURL       interface{} `json:"canonical_url"`
		CodeinjectionFoot  interface{} `json:"codeinjection_foot"`
		CodeinjectionHead  interface{} `json:"codeinjection_head"`
		Description        interface{} `json:"description"`
		FeatureImage       interface{} `json:"feature_image"`
		ID                 string      `json:"id"`
		MetaDescription    interface{} `json:"meta_description"`
		MetaTitle          interface{} `json:"meta_title"`
		Name               string      `json:"name"`
		OgDescription      interface{} `json:"og_description"`
		OgImage            interface{} `json:"og_image"`
		OgTitle            interface{} `json:"og_title"`
		Slug               string      `json:"slug"`
		TwitterDescription interface{} `json:"twitter_description"`
		TwitterImage       interface{} `json:"twitter_image"`
		TwitterTitle       interface{} `json:"twitter_title"`
		URL                string      `json:"url"`
		Visibility         string      `json:"visibility"`
	} `json:"primary_tag"`
	PublishedAt            time.Time `json:"published_at"`
	ReadingTime            int       `json:"reading_time"`
	SendEmailWhenPublished bool      `json:"send_email_when_published"`
	Slug                   string    `json:"slug"`
	Tags                   []struct {
		AccentColor        interface{} `json:"accent_color"`
		CanonicalURL       interface{} `json:"canonical_url"`
		CodeinjectionFoot  interface{} `json:"codeinjection_foot"`
		CodeinjectionHead  interface{} `json:"codeinjection_head"`
		Description        interface{} `json:"description"`
		FeatureImage       interface{} `json:"feature_image"`
		ID                 string      `json:"id"`
		MetaDescription    interface{} `json:"meta_description"`
		MetaTitle          interface{} `json:"meta_title"`
		Name               string      `json:"name"`
		OgDescription      interface{} `json:"og_description"`
		OgImage            interface{} `json:"og_image"`
		OgTitle            interface{} `json:"og_title"`
		Slug               string      `json:"slug"`
		TwitterDescription interface{} `json:"twitter_description"`
		TwitterImage       interface{} `json:"twitter_image"`
		TwitterTitle       interface{} `json:"twitter_title"`
		URL                string      `json:"url"`
		Visibility         string      `json:"visibility"`
	} `json:"tags"`
	Title              string      `json:"title"`
	TwitterDescription interface{} `json:"twitter_description"`
	TwitterImage       interface{} `json:"twitter_image"`
	TwitterTitle       interface{} `json:"twitter_title"`
	UpdatedAt          time.Time   `json:"updated_at"`
	URL                string      `json:"url"`
	UUID               string      `json:"uuid"`
	Visibility         string      `json:"visibility"`
}
