import { DokumentTip } from "./dokumentTipEnum";
import { ZahtevTip } from "./zahtevTipEnum";

export interface Zahtev {
    id: string,
    datum: string,
    zahtevTip: ZahtevTip,
    dokumentTip: DokumentTip,
    korisnikId: string
}