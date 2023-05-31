import { Schema } from "mongoose"

export interface Poternica {
    id: string,
    sudijaId: string,
    gradjaninId: string,
    naslov: string,
    opis: string
}

export const PoternicaSchema = new Schema <Poternica> ({
    id: {type: String, required: true},
    sudijaId: {type: String, required: true},
    gradjaninId: {type: String, required: true},
    naslov: {type: String, required: true},
    opis: {type: String},
}, {
    versionKey: false
})