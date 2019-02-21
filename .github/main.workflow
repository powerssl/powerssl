workflow "Build powerctl" {
  on = "release"
  resolves = [
    "release darwin/amd64",
    "release windows/amd64",
    "release linux/amd64",
  ]
}

action "release darwin/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "darwin"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerctl"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release windows/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "windows"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerctl"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release linux/amd64" {
  uses = "./build/action/"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    POWERSSL_COMPONENT = "powerctl"
  }
  secrets = ["GITHUB_TOKEN"]
}
