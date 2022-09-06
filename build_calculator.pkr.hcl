variables {
  REPOSITORY = ""
  USERNAME   = ""
  PASSWORD   = ""
}

packer {
  required_plugins {
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
  }
}

source "docker" "go" {
  image  = "golang:1.19"
  commit = "true"
  changes = [
    "WORKDIR /calculator",
    "EXPOSE 8000",
    "ENTRYPOINT  [\"go\", \"run\", \".\"]"
  ]
}

build {
  name = "go-calculator"

  sources = ["source.docker.go"]

  provisioner "file" {
    source      = "./src"
    destination = "/calculator"
  }

  post-processors {
    post-processor "docker-tag" {
    repository = "${var.REPOSITORY}"
    tags       = ["latest"]
  }

  post-processor "docker-push" {
    login          = true
    login_username = "${var.USERNAME}"
    login_password = "${var.PASSWORD}"
  }

}