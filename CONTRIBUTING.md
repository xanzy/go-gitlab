# How to Contribute

We want to make contributing to this project as easy as possible. 

## Reporting Issues and Enhancements

If you have an issue, please report it on the [issue tracker](https://github.com/xanzy/go-gitlab/issues)

If there is a documented feature within Gitlab 

## Contributing Code

Pull requests are always welcome. If adding code that needs to be tested, try to add tests if there aren't any.

We use [`gofumpt`](https://github.com/mvdan/gofumpt) to format this project.

### Setting up your local development environment to Contribute to `go-gitlab`

1. [Fork](https://github.com/xanzy/go-github/fork), then clone the repository.
    ```sh
    git clone https://github.com/<your-username>/go-gitlab.git
    # or via ssh
    git clone git@github.com:<your-username>/go-gitlab.git
    ```
1. Install dependencies:
    ```sh
    make setup
    ```
1. Make your changes on your feature branch
1. Run the tests and `gofumpt`
    ```sh
    make test && make fmt
    ```
1. Open up your pull request 

