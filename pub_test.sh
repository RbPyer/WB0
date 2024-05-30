test_case=$(< ./model.json)
nats_url="nats://0.0.0.0:4222"

nats pub Order "$test_case" -s $nats_url