# Kafka

Required to install Java, Zookeeper, and Kafka. You need to start Zookeeper first to use Kafka.

Start Zookeeper
```bash
/opt/zookeeper/bin/zkServer.sh start
```

Start Kafka
```bash
nohup bin/kafka-server-start.sh config/server.properties &
```

Create a Kafka Topic
```bash
bin/kafka-topics.sh --create --topic test --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

List all Topic in Kafka
```bash
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```