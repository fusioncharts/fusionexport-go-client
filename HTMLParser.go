package FusionExport

import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
	"io"
)

type HTMLParser struct {
	Data      []byte
	tokenizer *html.Tokenizer
}

type Node struct {
	attrs map[string]string
}

func (hp *HTMLParser) GetElemetsByTagName(tagName string) ([]Node, error) {
	if hp.Data == nil {
		return nil, errors.New("no data found")
	}

	hp.tokenizer = html.NewTokenizer(bytes.NewReader(hp.Data))
	nodes, err := hp.iterateAndFilterNodes(tagName)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return nodes, nil
}

func (hp *HTMLParser) iterateAndFilterNodes(tagName string) ([]Node, error) {
	nodes := make([]Node, 0)

	for {
		tt := hp.tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			return nodes, hp.tokenizer.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			tn, _ := hp.tokenizer.TagName()

			if string(tn) == tagName {
				attrs := hp.findAttributes()
				node := Node{attrs: attrs}
				nodes = append(nodes, node)
			}
		}
	}
}

func (hp *HTMLParser) findAttributes() map[string]string {
	attrs := make(map[string]string)

	for {
		key, val, moreAttr := hp.tokenizer.TagAttr()

		if string(key) != "" {
			attrs[string(key)] = string(val)
		}

		if !moreAttr {
			break
		}
	}

	return attrs
}

func (n *Node) GetAttribute(attrName string) string {
	for k, v := range n.attrs {
		if k == attrName {
			return v
		}
	}

	return ""
}
