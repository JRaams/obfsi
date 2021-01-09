const db = require('../../dbc/db');

module.exports.getAssets = () => {
  return new Promise((resolve, reject) => {
    const assetsSQL = 'select * from assets;';
    db.query(assetsSQL, (err, res) => {
      if (err) {
        reject(new Error('Error fetching assets...'));
      } else {
        const assets = res.rows;

        const assetGroupsSQL =
          'select * from assetgroups ag where ag.id = ANY($1::int[]);';
        const assetGroupIds = new Array(...new Set(res.rows.map((r) => r.id)));

        db.query(assetGroupsSQL, [assetGroupIds], (err, res) => {
          if (err) {
            console.info(err);
            reject(new Error('Error fetching assetgroups...'));
          } else {
            const assetGroups = new Map();
            res.rows.forEach((row) => {
              assetGroups[row.id] = row;
            });

            assets.forEach((asset) => {
              asset.assetGroup = assetGroups[asset.assetgroup_id];
            });

            resolve(assets);
          }
        });
      }
    });
  });
};
