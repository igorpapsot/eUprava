import {GPolicajac} from "./gPolicajac";
import {EStatusProvere} from "./EStatusProvere";
import { Poternica } from "./poternica";
import { Gradjanin } from "./gradjanin";

export class ProveraGradjanina {

  constructor() {
  }

  id: string
  policajac: GPolicajac
  gradjanin: Gradjanin
  vreme: string
  status: EStatusProvere
  poternica: Poternica
}
