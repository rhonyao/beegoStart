/**
 * Webpack Config Development
 */

'use strict';

var webpack = require('webpack');
var path = require('path');
var ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = {
  devtool: 'source-map',
  entry: [
    'webpack/hot/only-dev-server',
    __dirname + '/src/app.js'
  ],
  output: {
    path: __dirname + '/web/dist',
    filename: 'build.js'
  },
  devServer: {
    host: '0.0.0.0',
    port: 1998,
    inline: true,
    publicPath: "/web/dist/",
    disableHostCheck: true,
    historyApiFallback: true,
    headers: { "Access-Control-Allow-Origin": "*" }
  },
  babel: {
    presets: ['es2015', 'stage-2'],
    plugins: ['transform-runtime']
  },
  module: {
    loaders: [{
      test: /\.js?$/,
      loader: 'babel',
      exclude: /node_modules/
    }, 
    {
      test: /\.css?$/,
      loader:  ExtractTextPlugin.extract("style-loader", "css-loader") 
    }, 
    {
      test: /\.styl?$/,
      loader:  ExtractTextPlugin.extract("style-loader", "css-loader!stylus-loader") 
    }, 
    {
      test: /\.pug?$/,
      loader:  ExtractTextPlugin.extract("pug-loader") 
    }, 
    {
      test: /\.less?$/,
      loader:  ExtractTextPlugin.extract("style-loader", "css-loader!less-loader"),
    }, 
    {
      test: /\.json$/,
      loader: 'json-loader'
    },
    { test: /\.(eot|svg|ttf|woff|woff2)(\?\S*)?$/, loader: "file-loader" },
    { test: /\.(jpg|png|bmp|gif)(\?\S*)?$/, loader: "url-loader" },
    {
      test: /\.vue?$/,
      loader: 'vue'
    }],
    vue: {
      loaders: {
        js: 'babel'
      }
    }
  },
  resolve: {
    extensions: ['', '.js', '.vue'],
    alias: {
      vue: 'vue/dist/vue.js'
    }
  },
  plugins: [
    new webpack.ProvidePlugin({
        $: "jquery",
        jQuery: "jquery",
        "window.jQuery": "jquery"
    }),
    new ExtractTextPlugin("app.css")
  ]
};
