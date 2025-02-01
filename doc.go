// Package article-generator is a CLI tool for generating blog article templates.
//
// It provides an interactive terminal interface using the Bubble Tea framework
// to collect article metadata such as title, description, and tags. The tool then
// generates a markdown file with proper frontmatter and creates necessary directories
// for blog posts.
//
// # Features
//   - Interactive CLI Interface
//   - Generate markdown files with frontmatter
//   - Creates directory structure for image of blog article
//
// # Installation
//
//	go install github.com/jahnen/article-generator@latest
//
// # Usage
//
// 1. Add to package.json (recommended):
//
//	{
//	  "scripts": {
//	    "write": "go run github.com/jahnen/article-generator@latest"
//	  }
//	}
//
// Then run:
//
//	npm run write
//
// 2. Or run directly:
//
//	go run github.com/jahnen/article-generator@latest
//
// # Generated File Structure
//
// The tool will generate a markdown file with this structure:
//
//	---
//	title: "Title"
//	description: "Something special"
//	tags: ["AI", "GO", "Blog"]
//	pubDate: "Feb 01 2025"
//	heroImage: "/blog/250201/thumbnail.webp"
//	---
//
// # Acknowledgements
//
// Built with:
//   - Bubble Tea Framework (github.com/charmbracelet/bubbletea)
//   - Gopher image from gopherkon
//   - Go and Go Mono fonts from go.googlesource.com
package main