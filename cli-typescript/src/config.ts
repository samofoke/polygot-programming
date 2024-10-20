import getOpts, { Opts } from "./opts";

export enum Operation {
  Print,
  Add,
  Remove,
}

export type Config = {
  args: string[];
  operation: Operation;
  config: string;
  pwd: string;
};

const getPwd = (op: Opts): string => {
  if (op.pwd) {
    return op.pwd;
  }
  return process.cwd();
};

const getConfig = (op: Opts): string => {
  if (op.config) {
    return op.config;
  }

  const home = process.env["HOME"];
  const envLocation = process.env["XDG_CONFIG_HOME"] || home;
  if (!envLocation) {
    throw new Error("unable to determine config location");
  }
  if (envLocation === home) {
    return path.join(envLocation, ".projector.json");
  }

  return path.join(envLocation, "projector", "projector.json");
};

const getArgs = (op: Opts): string[] => {
  if (!op.args || op.args.length === 0) {
    return [];
  }
};

const getOperation = (op: Opts): string => {
  if (!op.args || op.args.length === 0) {
    return Operation.Print;
  }

  if (op[0] === "add") {
    return Operation.Add;
  }
  return Operation.Print;
};

export default function getConfigFile(opts: Opts): Config {
  return {
    pwd: getPwd(opts),
    config: getConfig(opts),
    args: getArgs(opts),
    operation: getOperation(opts),
  };
}
