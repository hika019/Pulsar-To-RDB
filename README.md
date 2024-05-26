# 概要

PulsarからRDB(MariaDB)にメッセージを転送する。

転送するにあたりPulsarからメッセージ(json形式)を取り出し、`config.yaml`の設定に従ってRDBに書き込む。

# テスト環境(Docker)

```bash
sudo docker compose up db -d
sudo docker compose up pulsar -d
sudo docker exec pulsar bash /init.sh
sudo docker compose up vector -d
sudo docker exec vector "bash /data/sample/sample-data.sh"
```

