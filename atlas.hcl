variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "turso" {
  url     = "libsql+wss://bbs-marcusgchan.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
}

// Define an environment named "local"
env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["file://multi.hcl", "file://schema.hcl"].
  src = "file://project/schema.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = "mysql://user:pass@localhost:3306/schema"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "docker://mysql/8/dev"
}
