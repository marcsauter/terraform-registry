[![Go Report Card](https://goreportcard.com/badge/github.com/marcsauter/terraform-registry)](https://goreportcard.com/report/github.com/marcsauter/terraform-registry)

# terraform-registry

for Artifactory acceptance testing set:
export ARTIFACTORY_BASE_URL=
export ARTIFACTORY_USERNAME=
export ARTIFACTORY_PASSWORD=


$ curl -o test.zip https://releases.hashicorp.com/terraform-provider-random/2.0.0/terraform-provider-random_2.0.0_linux_amd64.zip
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 3739k  100 3739k    0     0  4333k      0 --:--:-- --:--:-- --:--:-- 4333k
$ unzip -l test.zip
Archive:  test.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
 11468896  08-16-2018 01:55   terraform-provider-random_v2.0.0_x4
---------                     -------
 11468896                     1 file


##
curl http://localhost:8080/.well-known/terraform/json

curl -s ttp://localhost:8080/v1/providers/postfinance/aixboms/versions | jq

curl -s http://localhost:8080/v1/providers/postfinance/aixboms/1.1.8/download/linux/amd64 | jq