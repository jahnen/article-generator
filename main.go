package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"	
	"github.com/charmbracelet/bubbletea"	
)

type model struct {
    title       textinput.Model
    description textinput.Model
    tags        textinput.Model 
    focused     int
    submitted   bool
    err         error
}

// Add Init method to implement tea.Model interface
func (m model) Init() tea.Cmd {
    return textinput.Blink
}

// Add View method to implement tea.Model interface
func (m model) View() string {
    var b strings.Builder

    b.WriteString("Create New Blog Article\n\n")

    b.WriteString("Title:\n")
    b.WriteString(m.title.View())
    b.WriteString("\n\n")

    b.WriteString("Description:\n")
    b.WriteString(m.description.View())
    b.WriteString("\n\n")

    b.WriteString("Tags (comma-separated):\n")
    b.WriteString(m.tags.View())
    b.WriteString("\n\n")

    button := "[ Submit ]"
    if m.focused == 3 {
        button = "[ » Submit « ]"
    }
    b.WriteString(button)

    b.WriteString("\n\n")
    b.WriteString("(esc to quit)")

    return b.String()
}

func initialModel() model {
    ti := textinput.New()
    ti.Placeholder = "Enter title"
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 50

    des := textinput.New()
    des.Placeholder = "Enter description"
    des.Width = 50

    tags := textinput.New()
    tags.Placeholder = "Enter tags (comma-separated)"
    tags.Width = 50

    return model{
        title:       ti,
        description: des,
        tags:        tags,
        focused:     0,
        submitted:   false,
    }
}

// Update the focused cases (remove cases 3 and 4)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "esc":
            return m, tea.Quit

        case "tab", "shift+tab", "enter", "up", "down":
            s := msg.String()

            if s == "enter" && m.focused == 3 {  // Changed from 5 to 3
                m.submitted = true
                return m, tea.Quit
            }

            if s == "up" || s == "shift+tab" {
                m.focused--
                if m.focused < 0 {
                    m.focused = 3
                }
            } else {
                m.focused++
                if m.focused > 3 {
                    m.focused = 0
                }
            }

            cmds = make([]tea.Cmd, 0)
            switch m.focused {
            case 0:
                cmds = append(cmds, m.title.Focus())
                m.description.Blur()
                m.tags.Blur()
            case 1:
                m.title.Blur()
                cmds = append(cmds, m.description.Focus())
                m.tags.Blur()
            case 2:
                m.title.Blur()
                m.description.Blur()
                cmds = append(cmds, m.tags.Focus())
            }

            return m, tea.Batch(cmds...)
        }
    }

    // Handle character input for each field
    var cmd tea.Cmd
    switch m.focused {
    case 0:
        m.title, cmd = m.title.Update(msg)
        cmds = append(cmds, cmd)
    case 1:
        m.description, cmd = m.description.Update(msg)
        cmds = append(cmds, cmd)
    case 2:
        m.tags, cmd = m.tags.Update(msg)
        cmds = append(cmds, cmd)
    }

    return m, tea.Batch(cmds...)
}

func main() {
    p := tea.NewProgram(initialModel())
    m, err := p.Run()
    if err != nil {
        fmt.Printf("Error running program: %v\n", err)
        os.Exit(1)
    }

    // After the program exits, if submitted, generate the article template
    if m.(model).submitted {
        finalModel := m.(model)
        dateStr := time.Now().Format("060102")
        currentDate := time.Now().Format("Jan 02 2006")
        
        // Use relative paths instead of absolute paths
        publicDir := "public/" + dateStr
        if err := os.MkdirAll(publicDir, 0755); err != nil {
            fmt.Printf("Error creating public directory: %v\n", err)
            os.Exit(1)
        }

        // Create content for markdown file
        // Process tags: split by comma, trim spaces, and wrap with quotes
        tagsList := strings.Split(finalModel.tags.Value(), ",")
        for i, tag := range tagsList {
            tagsList[i] = fmt.Sprintf(`"%s"`, strings.TrimSpace(tag))
        }
        formattedTags := strings.Join(tagsList, ", ")

        content := fmt.Sprintf(`---
title: "%s"
description: "%s"
tags: [%s]
pubDate: "%s"
heroImage: "/blog/%s/thumbnail.webp"
---

`, finalModel.title.Value(), 
           finalModel.description.Value(),
           formattedTags,
           currentDate,  // Use current date
           dateStr)     // Use dateStr for heroImage
        
        // Create blog content directory
        blogDir := "src/content/blog/"
        if err := os.MkdirAll(blogDir, 0755); err != nil {
            fmt.Printf("Error creating blog directory: %v\n", err)
            os.Exit(1)
        }

        // Create filename based on title
        filename := strings.ToLower(strings.ReplaceAll(finalModel.title.Value(), " ", "-")) + ".md"
        filepath := blogDir + filename

        // Write to file
        if err := os.WriteFile(filepath, []byte(content), 0644); err != nil {
            fmt.Printf("Error writing file: %v\n", err)
            os.Exit(1)
        }

        fmt.Printf("Article template generated at: %s\n", filepath)
        fmt.Printf("Public directory created at: %s\n", publicDir)
    }
}