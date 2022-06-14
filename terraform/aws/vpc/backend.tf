terraform {
  cloud {
    organization = "homeserver"

    workspaces {
      name = "temporal"
    }
  }
}
