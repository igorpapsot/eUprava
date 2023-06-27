import mongoose from "mongoose";
import { DokumentTip, ReqSchema, ZahtevTip } from "../schemas/req.schema";
import * as crypto from "crypto"
import { formatDate, formatDateTime } from "../utilities/date.util";

export class ReqService {
    readonly reqModel = mongoose.model("req", ReqSchema, "req")

    async getReqById(id: string) {
        const req = await this.reqModel.findOne({id: id})
        return req
    }

    async getReqByUserId(userId: string) {
        const req = await this.reqModel.find({userId: userId})
        return req
    }

    async createRequest(zahtevTip: ZahtevTip, dokumentTip: DokumentTip, korisnikId: string, zakazanDatumVreme: string|undefined, datumIsticanja: string|undefined, jmbgDeteta: string|undefined) {
        const uuid = crypto.randomUUID()
        const now = new Date()
        const req = await this.reqModel.create({
            id: uuid,
            datum: formatDate(now),
            zahtevTip: zahtevTip,
            dokumentTip: dokumentTip,
            korisnikId: korisnikId,
            zakazanDatumVreme: formatDateTime(new Date(zakazanDatumVreme)),
            datumIsticanja: formatDate(new Date(datumIsticanja)),
            jmbgDeteta: jmbgDeteta
        })
        return req
    }
    
}