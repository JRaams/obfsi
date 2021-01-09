const memoize = require('fast-memoize');
const db = require('../../dbc/db');
const { assetGroupService } = require('../assetgroups');

const getAll = () => {
  return new Promise((resolve, reject) => {
    const assetsSQL = 'select * from assets;';
    db.query(assetsSQL, (err, res) => {
      if (err) {
        reject(new Error('Error fetching assets...'));
      } else {
        const assets = res.rows;
        assets.forEach((asset) => {
          asset.assetGroup = assetGroupService.getById(asset.assetgroup_id);
        });
        resolve(assets);
      }
    });
  });
};

module.exports = {
  getAll: memoize(getAll),
};
