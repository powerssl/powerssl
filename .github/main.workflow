workflow "release powerctl" {
  on = "release"
  resolves = [
    "release powerctl darwin/amd64",
    "release powerctl windows/amd64",
    "release powerctl linux/amd64",
  ]
}

action "release powerctl darwin/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "darwin"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerctl"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release powerctl windows/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "windows"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerctl"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release powerctl linux/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerctl"
  }
  secrets = ["GITHUB_TOKEN"]
}

workflow "release powerutil" {
  on = "release"
  resolves = [
    "release powerutil darwin/amd64",
    "release powerutil windows/amd64",
    "release powerutil linux/amd64",
  ]
}

action "release powerutil darwin/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "darwin"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerutil"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release powerutil windows/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "windows"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerutil"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release powerutil linux/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerutil"
  }
  secrets = ["GITHUB_TOKEN"]
}
