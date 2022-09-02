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
      repository = "gabrielys/go-calculator"
      tags       = ["latest"]
    }
  }

}