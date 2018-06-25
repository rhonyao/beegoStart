/**
 * Webpack Config Production
 */

'use strict';

var webpack = require('webpack');
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var _ = require('lodash');

var prodPlugins = [
  new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/),
  new webpack.optimize.OccurrenceOrderPlugin(),
  new webpack.optimize.DedupePlugin(),
  new webpack.optimize.CommonsChunkPlugin({
      minChunks: 2,
      name: "vendor"
    }),
  new ExtractTextPlugin("[name].css"),
  new webpack.optimize.UglifyJsPlugin({
    compressor: {
      warnings: false,
      screw_ie8: true
    }
  }),
  new webpack.DefinePlugin({
    'process.env': {
      NODE_ENV: JSON.stringify('production')
    }
  })
];

var webpackDevConfig = require('./webpack.dev.config.js');

var webpackProdConfig = _.extend(webpackDevConfig, {
  devtool: false,
  entry: {
    app: [ __dirname + '/client/app.js'],
    vendor: ['jquery','vue','vuex','vue-router','vue-resource', 
      'bootstrap-daterangepicker',
      'datatables.net',
      'datatables.net-bs',
      'datatables.net-fixedcolumns',
      'bootstrap','moment','numeral',
      'echarts/lib/echarts',
      'echarts/lib/chart/bar',
       'echarts/lib/chart/map',
       'echarts/lib/chart/line',
       'echarts/lib/chart/scatter',
       'echarts/lib/chart/funnel',
       'echarts/lib/chart/pie',
       'echarts/lib/component/tooltip',
       'echarts/lib/component/title',
       'echarts/lib/component/legend',
       'echarts/lib/component/visualMap',
       'echarts/lib/component/markPoint'
    ]
  },
  
  output: {
    path: __dirname + '/www/magicbean/static/dist',
    publicPath: 'http://magicbean-dev.yidian-inc.com/magicbean/static/dist/',
    filename: '[name].js'
  },
  plugins: webpackDevConfig.plugins.concat(prodPlugins)
});

module.exports = webpackProdConfig;
