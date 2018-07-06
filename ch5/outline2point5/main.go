package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// forEachNode calls the functions pre(x) and post(x) for each node x in
// the tree rooted at n. Both functions are optional. pre is called
// before the children are visited (preorder) and post is called after
// (postorder)
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func makeNestFns(depth int) (start, end func(n *html.Node)) {
	// note: unlike in some languages, function declarations only work
	// at the topmost level, so we *must* use function expressions in
	// function bodies. This is covered in section 5.6
	// e.g. `func start(n *html.Node) { ... }` is a syntax error here
	start = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	end = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	return start, end
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	startElement, endElement := makeNestFns(2)
	forEachNode(doc, startElement, endElement)
}
