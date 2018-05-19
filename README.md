# terraform-provider-examples
Repo containing custom terraform providers for digging &amp; highlights

terraform-http-stateserver => un (very simple) remote backend http pour partager et collaborer sur les ressources terraform de l'equipe

terraform-provider-flownet => un custom provider pour un outil proprietaire "flownet", permettant de manipuler sous forme de ressources terraform les objets proprietaires de cet outil "flownet"

terraform-provider-bosh => un custom provider pour bosh, permettant de manipuler sous forme de ressources terraform des stemcell, releases et deploiements Bosh

## Build ##

Voir la doc [officielle ](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)

Sinon :

- (set|export) GOOS=linux
- (set|export) GOARCH=amd64
- fab@gobox:/go_workspace/src/terraform-provider-bosh$ go build -o terraform-provider-bosh

## Usage ##

1. cloner le repo
2. Demarrer le backend http en local
3. Demarrer le fake serveur flownet en local
4. Adapter la configuration du provider et du backend http dans le fichier main.tf (principalement les IP, port et URI)
5. terraform init && terraform plan && terraform apply !
