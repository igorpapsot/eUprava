import mongoose from "mongoose";
import { PoternicaSchema } from "../schemas/poternica.schema";

export class PoternicaService {

    readonly poternicaModel = mongoose.model("poternica", PoternicaSchema, "poternica");

    async getPoternicaById(id: string) {
        const poternica = await this.poternicaModel.findOne({id: id})
        return poternica
    }

    async getPoternicaByGradjaninId(gradjaninId: string) {
        const poternica = await this.poternicaModel.findOne({gradjaninId: gradjaninId})
        return poternica
    }

    async getPoternicaBySudijaId(sudijaId: string) {
        const poternica = await this.poternicaModel.find({sudijaId: sudijaId})
        return poternica
    }

    async createPoternica(sudijaId: string, gradjaninId: string, naslov: string, opis: string|undefined) {
        const uuid = crypto.randomUUID()
        const poternica = await this.poternicaModel.create({
            id: uuid,
            sudijaId: sudijaId,
            gradjaninId: gradjaninId,
            naslov: naslov,
            opis: opis
        })
        return poternica
    }
}