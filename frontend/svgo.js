module.exports = {
  js2svg: {
    indent: 2,
    pretty: true,
  },
  multipass: true,
  plugins: [
    {
      name: 'preset-default',
      params: {
        overrides: {
          minifyStyles: false, // Removes 100% step in @keyframe
          removeHiddenElems: false, // Removes svg sprite content
        },
      },
    },
    'removeDimensions',
    'sortAttrs',
    'convertStyleToAttrs',
    'removeTitle',
    {
      name: 'removeAttrs',
      params: {
        attrs: 'data.*',
      },
    },
  ],
};
