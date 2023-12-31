# Starlight
An offline-first todo app with the option to save your data to an online data store


## Considerations
As of 2024, January, this app is to be written in either SolidJS, or Svelte. While we could use either Angular, React, or Vue, I want to make this using one of the two upcoming starlet frameworks of 2023, 2024. Using either Svelte or SolidJS will allow us to reduce the default size of the application. As this will be a desktop app running on the operating system's web view library, we will use the rust Tauri package/library/crate to render the JavaScript framework as if it were a native desktop application.

## Design
Version 1.0 will feature two views – a table of tasks with sub-tasks, and a list of documentation nodes. The User Interface (UI) should be minimal, but easy to learn and navigate.

### Tasks

The project screen will feature Tasks and Sub-tasks. Tasks can feature sub-tasks of their own if there need to be more feature complexity. It should be a table of tasks that are movable by user input, can be manually sorted by user input, can be edited and be deleted, and can feature sub-tasks. There should be as minimal complexity when creating a project, adding tasks to a project, and editing a project.

### Documentation

The goal of the Documentation view is to include the project documents with the project itself, rather than in a separate text file. Documentation will likely be rendered into Markup textareas that are used to explain tasks and sub-tasks. This Documentation can be used either for project development, version 1.0 and onward documentation, or for application documentation.

The Documentation view should be in a list of titles. These titles can be clicked-on to expand the textarea and view/render the text information contained inside the list item. This text information will be about what the task and sub-tasks are for, their feature impact on the application, and their actions made during the data flow on user interaction.

### Database storage

Data will be saved to an SQLite database and will be stored localy. This data should have the option of being encrypted. If possible, the primary database used will be SQLCipher, an encryption-first SQLite fork.