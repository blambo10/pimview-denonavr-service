const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');
const path = require('path');
const dotenv = require('dotenv');
const webpack = require('webpack');

// call dotenv and it will return an Object with a parsed key 
const env = dotenv.config().parsed;

// reduce it to a nice object, the same as before
const envKeys = Object.keys(env).reduce((prev, next) => {
prev[`process.env.${next}`] = JSON.stringify(env[next]);
    return prev;
}, {});


console.log("my object: %o", envKeys)

module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map',
  devServer: {
    historyApiFallback: true,
    static: {
      directory: path.join(__dirname, 'dist'),
    },
    port: 3001,
  },
  plugins: [
    new webpack.DefinePlugin(envKeys)
  ],
});
console.log("my object: %o", common)