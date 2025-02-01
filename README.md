# article-generator

A CLI tool for generaing blog article templates. 

## Features

![article-generator](https://github.com/user-attachments/assets/afd48141-eec9-4c7d-b819-94c734f49b25)

- Interactive CLI Interface using bubbletea
- Generate markdown files with frontmatter
- Creates directory strucrture for image of blog article
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

