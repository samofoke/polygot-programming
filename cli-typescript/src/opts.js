"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const cli = require("command-line-args");
function getOpts() {
    return cli([
        {
            name: "args",
            defaultOption: true,
            type: String,
        },
        {
            name: "config",
            alias: "p",
            type: String,
        },
        {
            name: "pwd",
            alias: "c",
            type: String,
        },
    ]);
}
module.exports = getOpts;
