import { Schema } from "mongoose";

export enum Pol {
    M,Z
}

export interface User {
    id: string,
    jmbg: string,
    ime: string,
    prezime: string,
    pol: Pol,
    sifra: string
}

export const UserSchema = new Schema <User> ({
    id: {type: String, required: true},
    jmbg: {type: String, required: true},
    ime: {type: String, required: true},
    prezime: {type: String, required: true},
    pol: {type: Number, enum: Pol},
    sifra: {type: String, required: true},
})