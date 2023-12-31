package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mirzakhany/chapar/components"
)

type RequestItem struct {
	Name string
	Type string
}

func createMainWindow(app fyne.App) fyne.Window {
	myWindow := app.NewWindow("API Testing Tool")

	// Combine sidebar and main content
	splitLayout := container.NewHSplit(newSideBar(), newMainContent(myWindow))
	splitLayout.Offset = 0.18 // Sidebar takes 20% of the window

	appTabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Collections", theme.FolderIcon(), splitLayout),
		container.NewTabItemWithIcon("Environments", theme.GridIcon(), widget.NewLabel("Environments")),
	)

	appTabs.SetTabLocation(container.TabLocationLeading)
	// Set the content of the window
	myWindow.SetContent(container.NewBorder(newNavBar(myWindow), nil, nil, nil, appTabs))
	myWindow.Resize(fyne.NewSize(800, 600)) // Set window size

	return myWindow
}

func showCreateEnvDialog(win fyne.Window) {
	newEnvEntry := widget.NewEntry()
	newEnvEntry.SetPlaceHolder("Enter new environment name")

	dialog.ShowForm("Create New Environment", "Save", "Cancel", []*widget.FormItem{
		{Text: "Environment Name", Widget: newEnvEntry},
	}, func(b bool) {
		if b {
			newEnvName := newEnvEntry.Text
			fmt.Println("New Environment Created:", newEnvName)
			// Add logic to handle the creation of the new environment
		}
	}, win)
}

func newMainContent(win fyne.Window) *fyne.Container {
	// Tabs container
	tabs := container.NewDocTabs(createTab(win, 1))
	i := 1
	tabs.CreateTab = func() *container.TabItem {
		i++
		return createTab(win, i)
	}

	// Layout for the add button and the tabs
	return container.NewBorder(nil, nil, nil, nil, tabs)
}

func createTab(win fyne.Window, tabNum int) *container.TabItem {
	requestContainer := newRequestContainer(win)
	responseEntry := widget.NewMultiLineEntry()
	responseEntry.Disable()
	content := container.NewBorder(requestContainer, nil, nil, nil, responseEntry)
	return container.NewTabItem(fmt.Sprintf("Request %d", tabNum), content)
}

func newNavBar(win fyne.Window) *fyne.Container {
	// Navigation Bar Setup
	logo := widget.NewLabel("App Logo") // Replace with an actual logo widget

	// Environment Dropdown Setup
	envLabel := widget.NewLabel("Environment:")
	envOptions := []string{"Development", "Staging", "Production", "Create New..."}
	envDropdown := widget.NewSelect(envOptions, func(value string) {
		if value == "Create New..." {
			showCreateEnvDialog(win)
		} else {
			fmt.Println("Environment Selected:", value)
			// Add logic to handle environment change
		}
	})
	envDropdown.SetSelectedIndex(0) // Set default environment

	return container.NewHBox(
		logo,
		layout.NewSpacer(), // This pushes the dropdown to the right
		envLabel,
		envDropdown,
	)
}

func newSideBar() *fyne.Container {
	addBtn := widget.NewButton("Add", func() {})
	addBtn.Importance = widget.LowImportance

	importBtn := widget.NewButton("Import", func() {})
	importBtn.Importance = widget.LowImportance

	toolbar := container.NewHBox(
		layout.NewSpacer(),
		addBtn,
		importBtn,
	)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search ...")
	searchEntry.ActionItem = widget.NewIcon(theme.SearchIcon())

	requests := []RequestItem{
		{"Request 1", "GET"},
		{"Request 2", "POST"},
		// Add more requests here
	}

	// List of requests
	requestList := widget.NewList(
		func() int { return len(requests) },
		func() fyne.CanvasObject { return components.NewRequestListObject("", "") },
		func(index widget.ListItemID, item fyne.CanvasObject) {
			i := item.(*components.RequestListObject)
			i.SetType(requests[index].Type)
			i.SetName(requests[index].Name)
		},
	)

	br := container.NewBorder(searchEntry, nil, nil, nil, requestList)
	cn := container.New(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, br)
	return cn
}

func newRequestContainer(win fyne.Window) *fyne.Container {
	// Request Config Setup
	protocolSelect := widget.NewSelect([]string{"HTTP/S", "GRPC"}, func(value string) {})
	protocolSelect.SetSelectedIndex(0)

	requestName := components.NewEditableLabel("Request Name")
	nameConfig := container.NewBorder(nil, widget.NewSeparator(), protocolSelect, nil, requestName)

	methodDropdown := widget.NewSelect([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION", "HEAD"}, func(value string) {})
	methodDropdown.SetSelectedIndex(0)

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter request URL")
	sendButton := widget.NewButton("Send", func() {
		// Logic to handle the API request goes here
	})

	requestConfig := container.NewBorder(nameConfig, nil, methodDropdown, sendButton, urlEntry)
	return requestConfig
}
