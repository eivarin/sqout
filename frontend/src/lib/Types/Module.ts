export interface Flag {
    Description: string;
    Type: string;
    Required: boolean;
    Prefix: string;
    IsEmpty: boolean;
    Default: string;
};

export interface Exe {
    CommandName: string;
    Description: string;
    KeepAlive: boolean;
    FlagsOrder: string[];
    Flags: { [key: string]: Flag };
};

export interface GitInfo {
    Branch: string;
    Commit: string;
};

export interface Module {
    Id: string;
    Path: string;
    IsRepo: boolean;
    GitInfo: GitInfo;
    Exe: Exe;
};