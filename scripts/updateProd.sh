ssh ubuntu@hawetec "pm2 delete api.capitech"

scp -r capitech_api ubuntu@hawetec:/home/ubuntu/prd_projects/back/capitech/

ssh ubuntu@hawetec "/home/ubuntu/start_scripts/start_capitech_api.sh"

rm capitech_api
