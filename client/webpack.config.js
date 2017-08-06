const path = require('path');

module.exports = {
    entry: "./src/main.ts",
    output: {
        filename: "game.js",
        path: path.join(__dirname, "../public/")
    },

    devtool: "source-map",

    resolve: {
        extensions: [".ts", ".tsx", ".js", ".json"]
    },

    module: {
        rules: [
            { test: /\.tsx?$/, loader: "awesome-typescript-loader" },
            { enforce: "pre", test: /\.js$/, loader: "source-map-loader" }
        ]
    }
};