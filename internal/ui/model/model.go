package model

import (
	"fmt"
	"strings"
	"hexatui/internal/hexatui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	// Улучшенная цветовая схема
	primaryColor = lipgloss.Color("#00FF9F")      // Яркий неоновый зеленый для пользователя
	secondaryColor = lipgloss.Color("#00B8FF")    // Яркий голубой для ассистента
	accentColor = lipgloss.Color("#B39DDB")       // Лавандовый для акцентов вместо розового
	bgColor = lipgloss.Color("#1A1B26")          // Темный фон
	subtleColor = lipgloss.Color("#A9B1D6")      // Приглушенный цвет для второстепенного текста

	// Обновленные стили с уменьшенными отступами
	userStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		PaddingLeft(1)

	assistantStyle = lipgloss.NewStyle().
		Foreground(secondaryColor).
		PaddingLeft(1)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true).
		PaddingLeft(1)

	inputBoxStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(accentColor).
		Padding(0, 1)

	titleStyle = lipgloss.NewStyle().
		Foreground(accentColor).
		Bold(true).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(accentColor).
		Padding(0, 1).
		Align(lipgloss.Center)

	chatBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(subtleColor).
		Padding(1, 1)

	statusBarStyle = lipgloss.NewStyle().
		Foreground(subtleColor).
		Background(bgColor).
		PaddingLeft(1).
		PaddingRight(1).
		Bold(true)

	helpStyle = lipgloss.NewStyle().
		Foreground(subtleColor).
		Background(bgColor).
		Align(lipgloss.Center).
		PaddingTop(0).
		PaddingBottom(0)

	timestampStyle = lipgloss.NewStyle().
		Foreground(subtleColor).
		Italic(true).
		PaddingLeft(1)

	separatorStyle = lipgloss.NewStyle().
		Foreground(subtleColor).
		Align(lipgloss.Center)
)

type Model struct {
	messages  []hexatui.Message
	input     string
	client    *hexatui.Client
	err       error
	waiting   bool
	width     int
	height    int
}

func InitialModel() Model {
	client := hexatui.NewClient("ikGOYpcb1HZ1jURHJdcX656IXnOwjvrA")

	return Model{
		messages: []hexatui.Message{},
		client:   client,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	var b strings.Builder

	// Заголовок
	title := titleStyle.Width(m.width).Render("✧ HexaTui ✧")
	b.WriteString(title + "\n")

	// Разделитель
	b.WriteString(separatorStyle.Width(m.width).Render("════════════════") + "\n")

	// Область чата
	var chatContent strings.Builder
	for _, msg := range m.messages {
		switch msg.Role {
		case "user":
			content := wordwrap.String(msg.Content, m.width-10)
			fmt.Fprintf(&chatContent, "%s %s\n%s\n\n", 
				userStyle.Render("◉ Вы"),
				timestampStyle.Render("•"),
				userStyle.Render(content))
		case "assistant":
			content := wordwrap.String(msg.Content, m.width-10)
			fmt.Fprintf(&chatContent, "%s %s\n%s\n\n",
				assistantStyle.Render("◉ HexaTui"),
				timestampStyle.Render("•"),
				assistantStyle.Render(content))
		}
	}

	chatBox := chatBoxStyle.Width(m.width - 4).Render(chatContent.String())
	b.WriteString(chatBox + "\n")

	// Статус бар с иконками
	var status string
	if m.err != nil {
		status = errorStyle.Render("✗ Ошибка: " + m.err.Error())
	} else if m.waiting {
		status = statusBarStyle.Render("⟳ HexaTui печатает...")
	}
	b.WriteString(statusBarStyle.Width(m.width).Render(status) + "\n")

	// Поле ввода
	prompt := fmt.Sprintf("❯ %s", m.input)
	b.WriteString(inputBoxStyle.Width(m.width - 4).Render(prompt) + "\n")

	// Подсказки с иконками
	help := []string{
		"⌃-c: выход",
		"⌃-u: очистить строку",
		"⏎: отправить",
		"⌫: удалить",
	}
	b.WriteString(helpStyle.Width(m.width).Render(strings.Join(help, " │ ")))

	return b.String()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "ctrl+u": // Добавляем очистку строки
			m.input = ""
			return m, nil

		case "enter":
			if m.input == "" {
				return m, nil
			}

			userMsg := hexatui.Message{
				Role:    "user",
				Content: m.input,
			}
			m.messages = append(m.messages, userMsg)
			m.input = ""
			m.waiting = true
			m.err = nil

			return m, func() tea.Msg {
				response, err := m.client.Chat(m.messages)
				if err != nil {
					return errMsg{err}
				}
				return responseMsg(response)
			}

		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}

		default:
			if !m.waiting {
				m.input += msg.String()
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case responseMsg:
		m.waiting = false
		assistantMsg := hexatui.Message{
			Role:    "assistant", 
			Content: string(msg),
		}
		m.messages = append(m.messages, assistantMsg)

	case errMsg:
		m.err = msg.error
		m.waiting = false
	}

	return m, nil
}

type responseMsg string

type errMsg struct{ error }

func (e errMsg) Error() string { return e.error.Error() }