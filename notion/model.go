package notion

type Parent struct {
	Type   string `json:"-"`
	PageId string `json:"page_id"`
}

type Result struct {
	Id     string   `json:"id"`
	Object []Object `json:"results"`
}

type ChildPage struct {
	Title string `json:"title"`
}

type Object struct {
	Id        string    `json:"id"`
	Parent    Parent    `json:"parent"`
	Type      string    `json:"type"`
	ChildPage ChildPage `json:"child_page"`
}

type Content struct {
	Cont string `json:"content"`
}

type Text struct {
	Text Content `json:"text"`
}

type Title struct {
	Text []Text `json:"title"`
}

type Url struct {
	URL string `json:"url"`
}

type External struct {
	External Url `json:"external"`
}

type Body struct {
	Parent     Parent        `json:"parent"`
	Cover      External      `json:"cover"`
	Properties Title         `json:"properties"`
	Children   []interface{} `json:"children"`
}

type BodyNoChildren struct {
	Parent     Parent   `json:"parent"`
	Cover      External `json:"cover"`
	Properties Title    `json:"properties"`
}

type HeadingOneObject struct {
	Object  string  `json:"object"`
	Type    string  `json:"type"`
	HeadOne Heading `json:"heading_1"`
}

type HeadingTwoObject struct {
	Object  string  `json:"object"`
	Type    string  `json:"type"`
	HeadTwo Heading `json:"heading_2"`
}

type BulletedListObject struct {
	Object     string  `json:"object"`
	Type       string  `json:"type"`
	BulletList Heading `json:"bulleted_list_item"`
}

type Heading struct {
	RichText []RichText `json:"rich_text"`
}

type RichText struct {
	Type string  `json:"type"`
	Text Content `json:"text"`
}
