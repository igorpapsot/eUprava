import { Pol } from "./polEnum";

export interface Korisnik {
    id: string,
    jmbg: string,
    ime: string,
    prezime: string,
    pol: Pol,
    sifra: string
}