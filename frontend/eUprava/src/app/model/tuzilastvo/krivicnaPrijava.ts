import { Optuzeni } from "./optuzeni"
import { Status } from "./statusEnum"

export class KrivicnaPrijava {

    constructor(){}

    id: string
    privatnost: boolean
    clanZakonika: string
    datum: string | null
    mestoPrijave: string
    tuzilastvoId: string
    obrazlozenje: string
    status: Status
    optuzeni: Optuzeni
    gradjaninId: string
}