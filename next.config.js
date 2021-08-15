// eslint-disable-next-line @typescript-eslint/no-var-requires
const path = require('path')

module.exports = {
  reactStrictMode: true,
  sassOptions: {
    includePaths: [path.join(__dirname, 'styles')],
  },
  images: {
    domains: ['picsum.photos', 'unsplash.com'],
  },
}
