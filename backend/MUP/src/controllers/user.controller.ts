import { Router, response } from "express";
import { UserService } from "../services/user.service";

export const UserController = Router();

const userService = new UserService();

UserController.get("/user", async (req, res) => {
    const user = await userService.getUserById(req.query.id as string)
    res.json(user)
    return
})

UserController.post("/user/login", async (req, res) => {
    const user = await userService.getUserByJmbgAndSifra(req.body.jmbg as string, req.body.sifra as string)
    res.json(user)
    return 
})