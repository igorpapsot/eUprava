export enum Tip {
    LICNA, PASOS 
}

export interface Doc {
    id: string,
    korisnikId: string,
    tip: Tip,
    istice: Date
}