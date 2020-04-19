resource "pingaccess_hsm_provider" "test" {
  class_name    = "com.pingidentity.pa.hsm.cloudhsm.plugin.AwsCloudHsmProvider"
  name          = "demo"
  configuration = <<EOF
  {
      "user": "bob",
      "password": "top_secret",
      "partition": "p1"
  }
  EOF
}
