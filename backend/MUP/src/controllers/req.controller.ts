import { Router } from "express";
import { ReqService as ReqService } from "../services/req.service";

export const ReqController = Router();

const reqService = new ReqService();

ReqController.get("/req", async (req, res) => {
    const request = await reqService.getReqById(req.query.id as string)
    res.json(request)
    return
})

ReqController.get("/req/user", async (req, res) => {
    const request = await reqService.getReqByUserId(req.query.userId as string)
    res.json(request)
    return
})

ReqController.post("/req", async (req, res) => {
    const request = await reqService.createRequest(req.body.zahtevTip, req.body.dokumentTip, req.body.korisnikId)
    res.json(request)
    return
})
