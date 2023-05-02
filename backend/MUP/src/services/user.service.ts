import mongoose from "mongoose";
import { UserSchema } from "../schemas/user.schema";

export class UserService {

    readonly userModel = mongoose.model("user", UserSchema, "user");

    async getUserById(id: string) {
        const user = await this.userModel.findOne({id: id})
        return user
    }

    async getUserByIdAndSifra(jmbg: string, sifra: string) {
        const user = await this.userModel.findOne({jmbg : jmbg, sifra : sifra})
        return user
    }

}