/* eslint-env es2017 */
module.exports = {
    root: true,
    extends: [
        "plugin:svelte/base"
    ],
    parser: "@typescript-eslint/parser",
    plugins: ["@typescript-eslint"],
    parserOptions: {
        extraFileExtensions: [".svelte"],
        tsconfigRootDir: __dirname,
    },
    overrides: [
        {
            files: ["**/*.svelte"],
            parser: "svelte-eslint-parser",
            parserOptions: {
                parser: "@typescript-eslint/parser"
            }
        }
    ],
    ignorePatterns: [
        "*.config.js",
    ],
};
