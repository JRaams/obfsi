const memoize = require('fast-memoize');
const db = require('../../dbc/db');
const { getAssetGroupByID } = require('../assetgroups/assetGroupService');

const getAssets = () => {
  return new Promise((resolve, reject) => {
    const assetsSQL = 'select * from assets;';
    db.query(assetsSQL, (err, res) => {
      if (err) {
        reject(new Error('Error fetching assets...'));
      } else {
        const assets = res.rows;
        assets.forEach((asset) => {
          asset.assetGroup = getAssetGroupByID(asset.assetgroup_id);
        });
        resolve(assets);
      }
    });
  });
};

module.exports = {
  getAssets: memoize(getAssets),
};
