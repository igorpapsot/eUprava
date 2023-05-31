import { Router } from "express";
import { PoternicaService } from "../services/poternica.service";

export const PoternicaController = Router();

const poternicaService = new PoternicaService()

PoternicaController.get("/poternica", async (req, res) => {
    const poternica = await poternicaService.getPoternicaById(req.query.id as string)
    res.json(poternica)
    return
})

PoternicaController.get("/poternica/gradjanin", async (req, res) => {
    const poternica = await poternicaService.getPoternicaByGradjaninId(req.query.gradjaninId as string)
    res.json(poternica)
    return
})

PoternicaController.get("/poternica/sudija", async (req, res) => {
    const poternica = await poternicaService.getPoternicaBySudijaId(req.query.sudijaId as string)
    res.json(poternica)
    return
})

PoternicaController.post("/poternica", async (req, res) => {
    const poternica = await poternicaService.createPoternica(req.body.sudijaId as string, req.body.gradjaninId as string, req.body.naslov as string, req.body.opis as string)
    res.json(poternica)
    return
})