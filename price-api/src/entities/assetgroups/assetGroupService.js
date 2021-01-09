const memoize = require('fast-memoize');
const db = require('../../dbc/db');

const getAssetGroups = () => {
  return new Promise((resolve, reject) => {
    const sql = 'select * from assetgroups;';
    db.query(sql, (err, res) => {
      if (err) {
        reject(new Error('Error fetching assetgroups...'));
      } else {
        resolve(res.rows);
      }
    });
  });
};

const getAssetGroupByID = (id) => {
  return new Promise((resolve, reject) => {
    const sql = 'select * from assetgroups where id = $1;';
    const values = [id];
    db.query(sql, values, (err, res) => {
      if (err) {
        reject(new Error(`Error fetching assetgroup with id ${id}...`));
      } else {
        if (res.rowCount === 0) {
          reject(new Error(`No assetgroup found with id ${id}...`));
        }
        resolve(res.rows[0]);
      }
    });
  });
};

const getAssetGroupByName = (name) => {
  return new Promise((resolve, reject) => {
    const sql = 'select * from assetgroups where name = $1;';
    const values = [name];
    db.query(sql, values, (err, res) => {
      if (err) {
        reject(new Error(`Error fetching assetgroup with name ${name}...`));
      } else {
        if (res.rowCount === 0) {
          reject(new Error(`No assetgroup found with name ${name}...`));
        }
        resolve(res.rows[0]);
      }
    });
  });
};

module.exports = {
  getAssetGroups: memoize(getAssetGroups),
  getAssetGroupByID: memoize(getAssetGroupByID),
  getAssetGroupByName: memoize(getAssetGroupByName),
};
