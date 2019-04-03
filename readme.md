# gitcli

Command line tool to pull GitLab repos.  
Before `go build`,do some work to remove *impl in vault,because `go-keychain` only support mac and ios.

## Basic Commands

    configure  Configure this app
    load       get repo or project relation from git 
    get        pull all directory under the path
    tree       tree all repo or project of current user
    ls         list all repo or project of current user
    reset      Remove user configuration of this app
    help, h    Shows a list of commands or help for one command

## Options

    --help, -h     show help
    --version, -v  print the version

## How to Use
[GITLAB API](https://docs.gitlab.com/ee/api/README.html#personal-access-tokens)

### get token 
    Open gitlab web,Settings->Access Tokens->`Add a personal access token`,
    then get a private token,use in http get param ,such as `private_token=XXXXXXX`

### configure

    $ ./gitcli configure
    or
    $ ./gitcli.exe configure
    Configure now, press Enter to use default.
    Enter the base url of your project [default:http://git.example.com]:
    Enter Private Token (visit http://git.example.com/profile/personal_access_tokens if you do not have):

### load

    $ ./gitcli load

### ls

    $ ./gitcli ls business
    business/backend/
    business/common/
    business/proto

### tree

    $ ./gitcli tree business
    business/
    business/backend/
    business/backend/ScheduleTask/
    business/backend/process/
    business/backend/services/
    business/backend/services/BP_Login
    business/backend/services/BP_WorkSheet
    business/common/
    business/common/utils
    business/proto


### get

    ./gitcli get business 
    ./gitcli get business -o /Users/wuzuoliang/Downloads 
    ./gitcli get business -o /Users/wuzuoliang/Downloads -e business/proto 
    ./gitcli get business -o /Users/wuzuoliang/Downloads -e business/proto -e business/common
