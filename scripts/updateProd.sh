ssh ubuntu@168.138.150.229 "pm2 delete api.capitech"
ssh ubuntu@168.138.150.229 "rm -rf /home/ubuntu/prd_projects/back/capitech/capitech_api"

scp -r capitech_api ubuntu@168.138.150.229:/home/ubuntu/prd_projects/back/capitech/

ssh ubuntu@168.138.150.229 "/home/ubuntu/start_scripts/start_capitech_api.sh"

rm capitech_api
