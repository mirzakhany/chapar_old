package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type RequestItem struct {
	Name string
	Type string
}

func createMainWindow(app fyne.App) fyne.Window {
	myWindow := app.NewWindow("API Testing Tool")

	// Navigation Bar Setup
	logo := widget.NewLabel("App Logo") // Replace with an actual logo widget

	// Environment Dropdown Setup
	envLabel := widget.NewLabel("Environment:")
	envOptions := []string{"Development", "Staging", "Production", "Create New..."}
	envDropdown := widget.NewSelect(envOptions, func(value string) {
		if value == "Create New..." {
			showCreateEnvDialog(myWindow)
		} else {
			fmt.Println("Environment Selected:", value)
			// Add logic to handle environment change
		}
	})
	envDropdown.SetSelectedIndex(0) // Set default environment

	navBar := container.NewHBox(
		logo,
		layout.NewSpacer(), // This pushes the dropdown to the right
		envLabel,
		envDropdown,
	)

	// Sample data
	// requests := []RequestItem{
	// 	{"Request 1", "GET"},
	// 	{"Request 2", "POST"},
	// 	// Add more requests here
	// }

	// Function to create a widget for each request
	// createListItem := func(index int) fyne.CanvasObject {
	// 	return widget.NewLabel(requests[index].Name + " - " + requests[index].Type)
	// }

	// List of requests
	// requestList := widget.NewList(
	// 	func() int { return len(requests) },
	// 	func() fyne.CanvasObject { return widget.NewLabel("") },
	// 	func(index widget.ListItemID, item fyne.CanvasObject) {
	// 		item.(*widget.Label).SetText(requests[index].Type + " " + requests[index].Name)
	// 	},
	// )

	// Sidebar with search bar and list
	// sidebar := container.NewVBox(
	// 	newSearchEntry(),
	// 	requestList,
	// )

	// Function to create a new tab with request and response areas
	createTab := func(tabNum int) *container.TabItem {
		requestContainer := newRequestContainer()
		responseEntry := widget.NewMultiLineEntry()
		responseEntry.Disable()
		content := container.NewBorder(requestContainer, nil, nil, nil, responseEntry)
		return container.NewTabItem(fmt.Sprintf("Request %d", tabNum), content)
	}

	// Tabs container
	tabs := container.NewDocTabs(createTab(1))
	i := 1
	tabs.CreateTab = func() *container.TabItem {
		i++
		return createTab(i)
	}

	// Layout for the add button and the tabs
	mainContent := container.NewBorder(nil, nil, nil, nil, tabs)

	// Combine sidebar and main content
	splitLayout := container.NewHSplit(newSideBar(), mainContent)
	splitLayout.Offset = 0.18 // Sidebar takes 20% of the window

	mainContainer := container.NewHSplit(newActivityBar(), splitLayout)
	mainContainer.Offset = 0.04
	// Activity Bar
	// mainSplit := container.NewHSplit(newActivityBar(), splitLayout)
	// mainSplit.Offset = 0.04

	// Set the content of the window
	myWindow.SetContent(container.NewBorder(navBar, nil, nil, nil, mainContainer))
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

func newSearchEntry() *widget.Entry {
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search ...")
	searchEntry.ActionItem = widget.NewIcon(theme.SearchIcon())
	return searchEntry
}

func newSideBar() *fyne.Container {
	// toolbar := widget.NewToolbar(
	// 	widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
	// )

	toolbar := container.NewHBox(
		layout.NewSpacer(),
		widget.NewButton("Add", func() {}),
		widget.NewButton("Import", func() {}),
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
		func() fyne.CanvasObject { return NewRequestListObject("", "") },
		func(index widget.ListItemID, item fyne.CanvasObject) {
			i := item.(*RequestListObject)
			i.SetType(requests[index].Type)
			i.SetName(requests[index].Name)
		},
	)

	br := container.NewBorder(searchEntry, nil, nil, nil, requestList)
	cn := container.New(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, br)
	return cn
}

func newRequestContainer() *fyne.Container {
	methodDropdown := widget.NewSelect([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION", "HEAD", "-", "GRPC"}, func(value string) {})
	methodDropdown.SetSelectedIndex(0)

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter request URL")
	sendButton := widget.NewButton("Send", func() {
		// Logic to handle the API request goes here
	})

	requestConfig := container.NewBorder(nil, nil, methodDropdown, sendButton, urlEntry)
	return requestConfig
}

func newActivityBar() *fyne.Container {
	collectionsButton := widget.NewButtonWithIcon("", theme.FolderIcon(), func() {})
	environmentsButton := widget.NewButtonWithIcon("", theme.GridIcon(), func() {})

	activityBar := container.NewVBox(
		collectionsButton,
		environmentsButton,
	)

	return activityBar
}
