import cli from "command-line-args";

export type Opts = {
  args?: string[];
  pwd?: string;
  config?: string;
};

export default function getOpts(): Opts {
  const options = cli([
    {
      name: "args",
      defaultOption: true,
      multiple: true,
      type: String,
    },
    {
      name: "pwd",
      alias: "p",
      type: String,
    },
    {
      name: "config",
      alias: "c",
      type: String,
    },
  ]);
  return options as Opts;
}
