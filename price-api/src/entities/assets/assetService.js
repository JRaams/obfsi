const memoize = require('fast-memoize');
const db = require('../../dbc/db');
const { assetGroupService } = require('../assetgroups');

const getAll = () => {
  return new Promise((resolve, reject) => {
    const assetsSQL = 'select * from assets;';
    db.query(assetsSQL, (err, res) => {
      if (err) {
        return reject(new Error('Error fetching assets...'));
      } else {
        const assets = res.rows;
        assets.forEach((asset) => {
          asset.assetGroup = assetGroupService.getById(asset.assetgroup_id);
        });
        return resolve(assets);
      }
    });
  });
};

const getById = (id) => {
  return new Promise((resolve, reject) => {
    const assetsSQL = 'select * from assets where id = $1;';
    const values = [id];
    db.query(assetsSQL, values, (err, res) => {
      if (err) {
        return reject(new Error(`Error fetching asset with id: ${id}...`));
      } else {
        if (res.rowCount === 0) {
          return reject(new Error(`No asset found with id ${id}...`));
        }
        const asset = res.rows[0];
        asset.assetGroup = assetGroupService.getById(asset.assetgroup_id);
        return resolve(asset);
      }
    });
  });
};

const getByName = (name) => {
  return new Promise((resolve, reject) => {
    const assetsSQL = 'select * from assets where name = $1;';
    const values = [name];
    db.query(assetsSQL, values, (err, res) => {
      if (err) {
        return reject(new Error(`Error fetching asset with name: ${name}...`));
      } else {
        if (res.rowCount === 0) {
          return reject(new Error(`No asset found with name ${name}...`));
        }
        const asset = res.rows[0];
        asset.assetGroup = assetGroupService.getById(asset.assetgroup_id);
        return resolve(asset);
      }
    });
  });
};

module.exports = {
  getAll: memoize(getAll),
  getById: memoize(getById),
  getByName: memoize(getByName),
};
