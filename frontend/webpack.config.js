
const { ElementPlusResolver } = require('unplugin-vue-components/resolvers');

module.exports = {
  // ... other configurations
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: 'vue-loader',s
      },
    ],
  },
  plugins: [
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
};