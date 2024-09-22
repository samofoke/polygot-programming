const cli = require("command-line-args");

export type Opts = {
  args?: string[];
  pwd?: string;
  config?: string;
};

function getOpts(): Opts {
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
  ]) as Opts;
}

module.exports = getOpts;
