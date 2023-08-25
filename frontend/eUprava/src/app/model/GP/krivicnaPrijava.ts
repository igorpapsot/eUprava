import {EStatusPrijave} from "./EStatusPrijave";
import { Optuzeni } from "./optuzeni";

export class KrivicnaPrijava {

  constructor() {
  }

  Id: string
  Privatnost: boolean
  ClanZakonika: string
  Datum: string
  MestoPrijave: string
  TuzilastvoId: string
  Obrazlozenje: string
  Status: EStatusPrijave
  Optuzeni: Optuzeni
}
