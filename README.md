
```bash
sudo docker compose up db -d
sudo docker compose up pulsar -d
sudo docker exec pulsar bash /init.sh
sudo docker compose up vector -d
sudo docker exec vector bash /data/sample/sample-data.sh
```

