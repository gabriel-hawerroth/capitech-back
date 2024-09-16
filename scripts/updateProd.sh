ssh root@15.229.18.114 "pm2 delete api.capitech"
ssh root@15.229.18.114 "rm -rf /home/ubuntu/workspace/capitech_api/capitech_api"

scp -r capitech_api root@15.229.18.114:/home/ubuntu/workspace/capitech_api/

ssh root@15.229.18.114 "/root/start_capitech_api.sh"

rm capitech_api
