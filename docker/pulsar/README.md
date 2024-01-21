# Pulsar setup

```bash
/pulsar/bin/pulsar-admin tenants create apache
/pulsar/bin/pulsar-admin namespaces create apache/pulsar
/pulsar/bin/pulsar-admin topics create-partitioned-topic apache/pulsar/test-topic -p 4

```
