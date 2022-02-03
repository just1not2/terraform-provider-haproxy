resource "haproxy_frontend" "my_frontend" {
  name    = "myhaproxyfrontend"
  backend = "myhaproxybackend"
}
