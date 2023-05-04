import mongoose from "mongoose";
import { DokumentTip, ReqSchema, ZahtevTip } from "../schemas/req.schema";
import * as crypto from "crypto"

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

    async createRequest(zahtevTip: ZahtevTip, dokumentTip: DokumentTip, korisnikId: string) {
        const uuid = crypto.randomUUID()
        const now = new Date()
        const req = await this.reqModel.create({
            id: uuid,
            datum: now.toLocaleDateString("en-GB", {
                day: "2-digit",
                month: "2-digit",
                year: "numeric"
            }),
            zahtevTip: zahtevTip,
            dokumentTip: dokumentTip,
            korisnikId: korisnikId
        })
        return req
    }
    
}