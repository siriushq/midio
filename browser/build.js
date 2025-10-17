var async = require('async')
var exec = require('child_process').exec
var fs = require('fs')

var isProduction = process.env.NODE_ENV == 'production' ? true : false

rmDir = function(dirPath) {
  try { var files = fs.readdirSync(dirPath); }
  catch(e) { return; }
  if (files.length > 0)
    for (var i = 0; i < files.length; i++) {
      var filePath = dirPath + '/' + files[i];
      if (fs.statSync(filePath).isFile())
        fs.unlinkSync(filePath);
      else
        rmDir(filePath);
    }
  fs.rmdirSync(dirPath);
};

async.waterfall([
    function(cb) {
      rmDir('production');
      rmDir('dev');
      var cmd = 'webpack -p --config webpack.production.config.js'
      if (!isProduction) {
        cmd = 'webpack';
      }
      console.log('Running', cmd)
      exec(cmd, cb)
    },
  ], function(err) {
    if (err) return console.log(err)
  })
