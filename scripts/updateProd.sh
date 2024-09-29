ssh root@147.79.81.13 "pm2 delete api.capitech"
ssh root@147.79.81.13 "rm -rf /home/ubuntu/workspace/capitech_api/capitech_api"

scp -r capitech_api root@147.79.81.13:/home/ubuntu/workspace/capitech_api/

ssh root@147.79.81.13 "/root/start_capitech_api.sh"

rm capitech_api
