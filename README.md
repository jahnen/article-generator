
![스크린샷 2025-02-01 오후 9 14 43](https://github.com/user-attachments/assets/10e1a102-7bc0-4852-9f30-33244ec0edfa)

# article-generator

A CLI tool for generaing blog article templates. 

## Features

- Interactive CLI Interface.
- Generate markdown files with frontmatter.
- Creates directory strucrture for image of blog article.

![article-generator](https://github.com/user-attachments/assets/afd48141-eec9-4c7d-b819-94c734f49b25)

Will create this file, which you can start with.

```title.md
---
title: "Title"
description: "Something special"
tags: ["AI", "GO", "Blog"]
pubDate: "Feb 01 2025"
heroImage: "/blog/250201/thumbnail.webp"
---
```

## Install

```sh
go install github.com/jahnen/article-generator@latest
```

## How to use

1. modify your `package.json` (recommend)

```json
{
  "scripts": {
    "write": "go run github.com/jahnen/article-generator@latest",
    "astro": "astro",
    // ... other scripts ...
  }
}
```

then run the command in your blog directory

```sh
npm run write
```

or 

2. run the command in your blog directory: 

```sh
go run github.com/jahnen/article-generator@latest
```

## Acknowledgement

- generated gopher image from [gopherkon](https://github.com/quasilyte/gopherkon)
- go, go mono font from [go.googlesource.com](https://go.googlesource.com/image/+archive/master/font/gofont/ttfs.tar.gz)
- CLI built with [bubbletea](https://github.com/charmbracelet/bubbletea)
