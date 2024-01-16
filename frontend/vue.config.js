module.exports = {
    devServer: {
      port: 3000,
      proxy: {
        '^/users': {
          target: 'http://localhost:3000/',
          ws: true,
          changeOrigin: true
        },
    }
}}