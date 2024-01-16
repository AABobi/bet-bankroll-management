module.exports = {
  publicPath: '/',
  devServer: {
    port: 3000,
    proxy: {
      '^/api': {
        target: 'http://localhost:3000/',
        ws: true,
        changeOrigin: true
      }
    }
  }
}