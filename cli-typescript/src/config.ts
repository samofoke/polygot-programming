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
  const pullOperations = getOperation(op);

  if (pullOperations === Operation.Print) {
    if (op.args.length > 1) {
      throw new Error(`expected 0 or 1 but got ${op.args.length}`);
    }
    return op.args;
  }

  if (pullOperations === Operation.Add) {
    if (op.args.length !== 3) {
      throw new Error(`expected 2 arguments but got ${op.args.length - 1}`);
    }
    return op.args.slice(1);
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
