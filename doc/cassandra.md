# Cassandra

Create User `eiffel` `tower`
```sql
CREATE USER user_name WITH PASSWORD 'password';
```

Create Keyspace `eiffel_keyspace`
```sql
CREATE KEYSPACE keyspace_name WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'datacenter1' : 1 };
```
