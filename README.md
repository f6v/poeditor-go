# poeditor
CLI tool to update translations from https://poeditor.com/

## Usage

* Clone repository
* `go get "github.com/codegangsta/cli"`
* `go install`
* Add poeditor.json to the root of your project with language mapping(see sample in config_test.json)
* run: `poeditor update --path="{dir}"` where dir is your project directory
