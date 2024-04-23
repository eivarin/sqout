export interface Probe {
    Name: string;
    Description: string;
    Options: { [key: string]: string };
    HeartbitInterval: number;
    Module: string;
    Alive: boolean;
}