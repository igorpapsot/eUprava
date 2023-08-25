import {GPolicajac} from "./gPolicajac";
import {EStatusProvere} from "./EStatusProvere";
import { Poternica } from "./poternica";
import { Gradjanin } from "./gradjanin";

export class ProveraGradjanina {

  constructor() {
  }

  Id: string
  Policajac: GPolicajac
  Gradjanin: Gradjanin
  Vreme: string
  Status: EStatusProvere
  Poternica: Poternica
}
