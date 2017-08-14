package page

import (
	"net/http"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	Description string
	OgTitle		string
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
			return true
		}
	}
	return false
}

func isOgTitle(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "property" && attr.Val == "og:title" {
			return true
		}
	}
	return false
}

func f(n *html.Node, page *Page) {
	if n.Type == html.ElementNode && n.Data == "title" {
		page.Title = n.FirstChild.Data
	}
	if isDescription(n.Attr) {
		for _, attr := range n.Attr {
			// キーがcontentであるアトリビュートの値を格納
			if attr.Key == "content" {
				page.Description = attr.Val
			}
		}
	}
	if isOgTitle(n.Attr) {
		for _, attr := range n.Attr {
			if attr.Key == "content" {
				page.OgTitle = attr.Val
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c, page)
	}
}

func Get(url string) (*Page, error) {
	var resPage Page

	// HTTPリクエスト
	resp, err := http.Get(url)
	if err != nil {
		// HTTPリクエストエラー
		return nil, err
	}
	defer resp.Body.Close()

	// タイトルとディスクリプションの抽出
	doc, err := html.Parse(resp.Body)
	if err != nil {
		// 抽出時エラー
		return nil, err
	}

	// Page型に格納
	f(doc, &resPage)

	return &resPage, nil
}