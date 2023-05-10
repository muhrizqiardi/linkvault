# LinkVault

A self-hosted bookmarking & archiving app inspired by Zotero, Raindrop.io, Pocket, and Wayback Machine. 

Still work-in-progress.

## Features

- Save links as bookmarks
- Save an archive of the saved link
- Categorize links using tags
- Separate saved links using folders

## How to get this project up and running

This project have two parts: the **front end** and the **back end**. 

To be able to run this project, make sure that you have installed these first:

- Node, NPM, and the package manager [PNPM](https://pnpm.io/)
- Go programming language

First, clone this repository, then `cd` into the directory where the repository was cloned. 

After that, start running the project by doing these:

### Running the front end

1. Build the front end by running this:

  ```bash
  pnpm -F web build
  ```
2. Finally run the front end by running this:
  ```bash
  pnpm -F web start
  ```

### Running the back end 

1. Make sure that you're on the root directory of the project, then run these to build the project:

  ```bash
  cd server
  go build
  ```
