# terraform-provider-examples
Repo containing custom terraform providers for digging &amp; highlights

terraform-http-stateserver => un (very simple) remote backend http pour partager et collaborer sur les ressources terraform de l'equipe

terraform-provider-flownet => un custom provider pour un outil proprietaire "flownet", permettant de manipuler sous forme de ressources terraform les objets proprietaires de cet outil "flownet"

## Usage ##

1. cloner le repo
2. Demarrer le backend http en local
3. Demarrer le fake serveur flownet en local
4. Adapter la configuration du provider et du backend http dans le fichier main.tf (principalement les IP, port et URI)
5. terraform init && terraform plan && terraform apply !
