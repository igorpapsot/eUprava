import mongoose from "mongoose";
import { Pol, UserSchema } from "../schemas/user.schema";

export class UserService {

    readonly userModel = mongoose.model("user", UserSchema, "user");

    async getUserById(id: string) {
        const user = await this.userModel.findOne({id: id})
        return user
    }

    async getUserByJmbgAndSifra(jmbg: string, sifra: string) {
        const user = await this.userModel.findOne({jmbg: jmbg, sifra : sifra})
        return user
    }

    async getUserByJmbg(jmbg: string) {
        const user = await this.userModel.findOne({jmbg: jmbg})
        return user
    }

    async createUser(jmbg: string, ime: string, prezime: string, pol: number, sifra: string) {
        const uuid = crypto.randomUUID()
        const req = await this.userModel.create({
            id: uuid,
            jmbg: jmbg,
            ime: ime,
            prezime: prezime,
            pol: pol,
            sifra: sifra
        })
        return req
    }

}