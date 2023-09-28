<h1 align="center">💻 Hackertea | Hackernews TUI</h1>
<p align="center">A minimal application for browsing Hacker News on the terminal</p>

<p align="center">

<a style="text-decoration: none" href="go.mod">
<img src="https://img.shields.io/github/go-mod/go-version/KarolosLykos/hackertea?style=flat-square" alt="Go version">
</a>

<a href="https://codecov.io/gh/KarolosLykos/hackertea" style="text-decoration: none">
<img src="https://img.shields.io/codecov/c/github/KarolosLykos/hackertea?color=magenta&style=flat-square&token=JQTnUwF7AK"alt="Downloads">
</a>

<a style="text-decoration: none" href="https://github.com/KarolosLykos/hackertea/actions?query=workflow%3AGo+branch%3Amain">
<img src="https://img.shields.io/github/actions/workflow/status/KarolosLykos/hackertea/go.yml?style=flat-square" alt="Build Status">
</a>

<a style="text-decoration: none" href="https://github.com/KarolosLykos/hackertea/actions?query=workflow%3ACodeQL+branch%3Amain">
<img src="https://img.shields.io/github/actions/workflow/status/KarolosLykos/hackertea/codeql-analysis.yml?style=flat-square&label=CodeQL" alt="Build Status">
</a>

<a style="text-decoration: none" href="https://github.com/KarolosLykos/hackertea">
<img src="https://img.shields.io/github/languages/top/KarolosLykos/hackertea?style=flat-square" alt="Build Status">
</a>

<br />
<a style="text-decoration: none" href="https://github.com/KarolosLykos/hackertea/stargazers">
<img src="https://img.shields.io/github/stars/KarolosLykos/hackertea.svg?style=flat-square" alt="Stars">
</a>
<a style="text-decoration: none" href="https://github.com/KarolosLykos/hackertea/fork">
<img src="https://img.shields.io/github/forks/KarolosLykos/hackertea.svg?style=flat-square" alt="Forks">
</a>
<a style="text-decoration: none" href="https://github.com/KarolosLykos/hackertea/issues">
<img src="https://img.shields.io/github/issues/KarolosLykos/hackertea.svg?style=flat-square" alt="Issues">
</a>

-----
A command-line interface (CLI) tool that allows users to browse the Top, New, and Best stories on Hacker News. The tool includes a minimalist text-based user interface (TUI) that is developed using Bubble Tea, Lip Gloss, and Bubble libraries.

<img alt="Welcome to Hachertea" src="demo.gif" width="1920"/>

## Features


- Read Top, New and Best stories.
- Fetch stories concurrently. (You can set the number of workers in the config file)
- In-memory thread-safe cache for caching news.
- A shiny UI to gaze your eyes upon.
    - Tabs
    - Separate pagination for each tab
    - Fetch next pages
    - Vim-like movements

## Libraries used


* [Bubbletea](https://github.com/charmbracelet/bubbles): The fun, functional and stateful way to build terminal apps.
* [Bubbles](https://github.com/charmbracelet/bubbles): Common Bubble Tea components such as text inputs, viewports, spinners and so on
* [Lip Gloss](https://github.com/charmbracelet/lipgloss): Style, format and layout tools for terminal applications


## Styling


The default theme is already loaded by default, but the good news is that you have the option to add any theme of your choice!
Simply take a look at the "config-example.yaml" file to see the available options.

<img alt="Welcome to Hachertea" src="examples/demo.gif" width="1920"/>

## Roadmap


- [ ] Add more screens
    - [ ] Add Comments screen
    - [ ] Add User profile screen
    - [ ] Add Ask HN screen
    - [ ] Add Jobs screen
- [ ] Add Changelog
- [ ] Add additional styling options w/ Examples
- [ ] Multi-language Support

## Contributing


Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

You can open issues for bugs you've found or features you think are missing. You can also submit pull requests to this repository. To get started, take a look at [CONTRIBUTING.md](CONTRIBUTING.md)

