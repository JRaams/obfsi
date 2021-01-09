const { Client } = require('pg');

const connectionString = process.env.DBSOURCE || 'postgres://obfsi@0.0.0.0:5432/obfsi?sslmode=disable';

const client = new Client({
  connectionString,
});
client.connect();

module.exports = client;
