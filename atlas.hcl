variable "env" {
    type    = string
    default = "./.env"
}

locals {
    env = {
        for line in split("\n", file(var.env)): split("=", line)[0] => regex("=(.*)", line)[0]
        if !startswith(line, "#") && length(split("=", line)) > 1
    }
}

env "dev" {
    url     = local.env["DB_URL"]
    src     = "file://database/schema.sql"
    dev     = "docker://mysql/8/dev"
}


env "turso" {
    url = local.env["TEST"]
    exclude = ["_litestream*"]
    src = "file://database/schema.sql"
    dev  = "sqlite://dev?mode=memory"
}
